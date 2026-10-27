# goyeti

A client for [Yeti](https://github.com/yeti-platform/yeti) built for AI
agents: one `yeti` binary with a command line and an
[MCP](https://modelcontextprotocol.io) server.

The first word of every command is its tier, so what a command can change is
visible before it runs. `yeti read …` never changes anything in Yeti.

```
export YETI_URL=https://yeti.example.com YETI_API_KEY=...

yeti read lookup 203.0.113.7 evil[.]example[.]com   # what Yeti knows about these values
yeti read search emotet                              # objects by keyword and by meaning
yeti read neighbors entities/45 --hops 2             # pivot through the graph
yeti read get entities/45                            # one object in full
yeti read yara-bundle --tag emotet --raw > emotet.yar
yeti read types                                      # the object types the server knows
```

Output is one line of JSON: `{"data": …}`, or `{"error": {"code", "message",
"hint", "retryable"}}`. Exit codes: 0 success, 1 failed operation, 2 bad usage.

## For agents

- **Coding agents** (Claude Code and the like): give them
  [skill/yeti/SKILL.md](skill/yeti/SKILL.md), and allow `yeti read` without
  prompting, e.g. `Bash(yeti read:*)` in Claude Code.
- **MCP hosts:** run `yeti mcp`, which serves the read operations as tools over
  stdio:

  ```
  claude mcp add yeti -e YETI_URL=https://yeti.example.com -e YETI_API_KEY=... -- yeti mcp
  ```

## Install

Download a binary from the [releases](https://github.com/yeti-platform/goyeti/releases),
or `go install github.com/yeti-platform/goyeti/cmd/yeti@latest`.

## Status

The read tier is done; changing Yeti (`yeti mutate`) and administering it
(`yeti admin`) come next. See [docs/DESIGN.md](docs/DESIGN.md) for what we're
building and why. The previous generated client is available as
`github.com/yeti-platform/goyeti@v0.1.0`.
