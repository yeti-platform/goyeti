//go:build evals

// Evals: analyst questions put to an agent that has only yeti mcp's tools,
// checking the tools it calls and its answer. Tool names and descriptions are
// what an agent actually reads, so changes to them are measured here rather
// than guessed. Each run seeds the end-to-end scenario and costs a few cents
// of model calls per question:
//
//	YETI_URL=... YETI_API_KEY=... go test -tags evals -run Evals -v ./yeti/
//
// They run Claude Code (claude -p), logged in or with ANTHROPIC_API_KEY.
// GOYETI_EVAL_MODEL picks the model; the default is Claude Code's.
package yeti

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"
)

// evalCase is a question, the tools a good answer needs, and what the answer
// must say.
type evalCase struct {
	name     string
	tier     string
	question string
	// want are tools the agent must call; "a|b" accepts either.
	want []string
	// maxCalls is the most tool calls a good answer takes.
	maxCalls int
	// answer are words the final answer must contain, ignoring case and
	// defanging.
	answer []string
	// check verifies what the agent changed in Yeti.
	check func(t *testing.T)
}

// agentRun is what an agent did with a question.
type agentRun struct {
	calls  []string
	answer string
	cost   float64
}

func TestEvals(t *testing.T) {
	c := liveClient(t)
	if _, err := exec.LookPath("claude"); err != nil {
		t.Skip("the evals run Claude Code: install claude")
	}
	ctx := context.Background()
	dir := t.TempDir()
	binary := filepath.Join(dir, "yeti")
	if out, err := exec.Command("go", "build", "-o", binary, "../cmd/yeti").CombinedOutput(); err != nil {
		t.Fatalf("building yeti: %v\n%s", err, out)
	}
	s := seed(t, c)
	// Questions by name read search results, which lag behind writes.
	eventually(func() bool {
		found, _, _ := c.Search(ctx, s.malware, 5)
		return found != nil && slices.ContainsFunc(found.Hits, func(h SearchHit) bool { return h.Name == s.malware })
	})

	n := time.Now().UnixNano()
	newIP := fmt.Sprintf("10.250.%d.%d", n%250+1, n/250%250+1)
	unknownIP := fmt.Sprintf("10.251.%d.%d", n%250+1, n/250%250+1)
	campaign := "campaign-" + s.run
	cases := []evalCase{
		{name: "lookup an IP", tier: "read", question: "What does Yeti know about the IP " + s.ip + "?",
			want: []string{"lookup"}, maxCalls: 2, answer: []string{s.malware}},
		{name: "defanged hostname", tier: "read",
			question: "Is " + strings.ReplaceAll(s.host, ".", "[.]") + " malicious? If so, which malware is it tied to?",
			want:     []string{"lookup"}, maxCalls: 2, answer: []string{s.malware}},
		{name: "actor behind a malware", tier: "read", question: "Which threat actor uses " + s.malware + "?",
			want: []string{"search|lookup", "neighbors"}, maxCalls: 4, answer: []string{s.actor}},
		{name: "a malware's infrastructure", tier: "read",
			question: "List the IP addresses and hostnames " + s.malware + " uses as infrastructure.",
			want:     []string{"neighbors"}, maxCalls: 4, answer: []string{s.ip, s.host}},
		{name: "pivot from a URL", tier: "read",
			question: "Where is " + s.url + " hosted, and which malware is behind it?",
			want:     []string{"lookup"}, maxCalls: 5, answer: []string{s.host, s.malware}},
		{name: "YARA rules by tag", tier: "read",
			question: "Give me the YARA rules tagged " + s.tag + ", ready to compile.",
			want:     []string{"yara_bundle"}, maxCalls: 3, answer: []string{s.loaderRule, s.baseRule}},
		{name: "a question doesn't write", tier: "mutate", question: "Has Yeti seen " + unknownIP + "?",
			want: []string{"lookup"}, maxCalls: 2, answer: []string{unknownIP},
			check: func(t *testing.T) {
				if looked, _, _ := c.Lookup(ctx, []string{unknownIP}, 0); looked[0].Known {
					t.Errorf("answering a question added %s", unknownIP)
				}
			}},
		{name: "record a C2", tier: "mutate",
			question: "I found " + newIP + " acting as a C2 server for " + s.malware + ". Record it in Yeti.",
			want:     []string{"add_observables", "link"}, maxCalls: 6,
			check: func(t *testing.T) {
				looked, _, _ := c.Lookup(ctx, []string{newIP}, 0)
				linked := slices.ContainsFunc(looked[0].Entities, func(l Linked) bool { return l.Ref == s.malwareRef })
				if !looked[0].Known || !linked {
					t.Errorf("%s isn't recorded as linked to %s: %+v", newIP, s.malwareRef, looked[0])
				}
			}},
		{name: "tag a hostname", tier: "mutate", question: "Tag " + s.host + " with " + campaign + " in Yeti.",
			want: []string{"tag|add_observables"}, maxCalls: 4,
			check: func(t *testing.T) {
				if looked, _, _ := c.Lookup(ctx, []string{s.host}, 0); !slices.Contains(looked[0].Tags, campaign) {
					t.Errorf("%s has tags %v", s.host, looked[0].Tags)
				}
			}},
	}

	var mu sync.Mutex
	var total float64
	passed := 0
	t.Run("questions", func(t *testing.T) {
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				run := askAgent(t, binary, dir, tc)
				mu.Lock()
				total += run.cost
				mu.Unlock()
				t.Logf("calls %v, $%.3f\n%s", run.calls, run.cost, run.answer)
				gradeRun(t, tc, run)
				if tc.check != nil {
					tc.check(t)
				}
				if !t.Failed() {
					mu.Lock()
					passed++
					mu.Unlock()
				}
			})
		}
	})
	t.Logf("%d of %d questions passed, $%.2f", passed, len(cases), total)
}

func gradeRun(t *testing.T, tc evalCase, run agentRun) {
	t.Helper()
	for _, want := range tc.want {
		if !slices.ContainsFunc(strings.Split(want, "|"), func(tool string) bool { return slices.Contains(run.calls, tool) }) {
			t.Errorf("didn't call %s", want)
		}
	}
	if len(run.calls) > tc.maxCalls {
		t.Errorf("%d tool calls, more than the %d a good answer takes", len(run.calls), tc.maxCalls)
	}
	answer := strings.ToLower(refang.Replace(run.answer))
	for _, word := range tc.answer {
		if !strings.Contains(answer, strings.ToLower(word)) {
			t.Errorf("the answer doesn't mention %s", word)
		}
	}
}

// askAgent puts a question to Claude Code with yeti mcp, serving tc.tier, as
// its only tools.
func askAgent(t *testing.T, binary, dir string, tc evalCase) agentRun {
	t.Helper()
	config := filepath.Join(dir, tc.tier+".json")
	data, _ := json.Marshal(map[string]any{"mcpServers": map[string]any{"yeti": map[string]any{
		"command": binary, "args": []string{"mcp", "--tier", tc.tier},
		"env": map[string]string{"YETI_URL": os.Getenv("YETI_URL"), "YETI_API_KEY": os.Getenv("YETI_API_KEY"),
			"YETI_AGENT_NAME": "goyeti-eval"},
	}}})
	// The configuration holds the API key.
	if err := os.WriteFile(config, data, 0o600); err != nil {
		t.Fatal(err)
	}
	args := []string{"-p", tc.question, "--mcp-config", config, "--strict-mcp-config", "--setting-sources", "",
		"--tools", "", "--allowedTools", "mcp__yeti__*", "--output-format", "stream-json", "--verbose",
		"--max-turns", "12", "--max-budget-usd", "0.5"}
	if model := os.Getenv("GOYETI_EVAL_MODEL"); model != "" {
		args = append(args, "--model", model)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "claude", args...)
	// An empty directory: no CLAUDE.md or project settings reach the agent.
	cmd.Dir = t.TempDir()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("claude: %v\n%s", err, stderr.String())
	}

	var run agentRun
	scanner := bufio.NewScanner(bytes.NewReader(out))
	scanner.Buffer(make([]byte, 0, 64*1024), 16*1024*1024)
	for scanner.Scan() {
		var event struct {
			Type    string `json:"type"`
			Message struct {
				Content []struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"content"`
			} `json:"message"`
			Result  string  `json:"result"`
			IsError bool    `json:"is_error"`
			Cost    float64 `json:"total_cost_usd"`
		}
		if json.Unmarshal(scanner.Bytes(), &event) != nil {
			continue
		}
		switch event.Type {
		case "assistant":
			for _, block := range event.Message.Content {
				if block.Type == "tool_use" {
					run.calls = append(run.calls, strings.TrimPrefix(block.Name, "mcp__yeti__"))
				}
			}
		case "result":
			run.answer, run.cost = event.Result, event.Cost
			if event.IsError {
				t.Errorf("the agent's run ended in an error: %s", event.Result)
			}
		}
	}
	return run
}
