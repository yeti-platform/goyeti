# goyeti design

goyeti is a client for [Yeti](https://github.com/yeti-platform/yeti) built for
AI agents first and people second. It ships as a single `yeti` binary with two
front ends over the same operations: a command line and an
[MCP](https://modelcontextprotocol.io) server.

This document records what we are building and why, so that later changes can
be checked against it. Update it when a decision changes.

## Goals

- **An agent can investigate with Yeti without help.** Given a hostname, a
  hash or a question, an agent finds what Yeti knows, follows it through the
  graph, and gets detection content, in a few calls.
- **Cheap to install anywhere an agent runs.** One static binary: no Python
  environment, instant start-up.
- **Safe by construction.** Reading never changes anything. Changes are
  explicit in the command itself, so that people and agent harnesses can see
  and approve them.
- **Small, stable surface.** About a dozen operations shaped around tasks, not
  the API's ~120 endpoints. The output is the contract, not Yeti's schemas.

## Non-goals

- A complete, typed binding of every Yeti endpoint. yeti-python covers
  Python integrations; this client only exposes what agents need.
- A human-first TUI. People can use the CLI, but its defaults favour agents.
- Compatibility with goyeti v0.1.0, the generated client. It stays available
  at that tag.

## Shape

```
cmd/yeti/      the binary: CLI commands and `yeti mcp`
yeti/          the operations, as a Go package: what both front ends call
internal/api/  HTTP, authentication and the Yeti endpoints the operations use
skill/         SKILL.md teaching agents to use the CLI
docs/          this document
```

`internal/api` is hand-written and decodes Yeti's responses loosely, keeping
the fields the operations need. Generating a client from Yeti's OpenAPI spec
(what v0.1.0 did) gives ~300 types for unions discriminated on `type`, breaks
whenever Yeti adds a type, and exposes none of it usefully to an agent.

## Command grammar: the tier is the first word

```
yeti read   <operation> [args]   never changes anything
yeti mutate <operation> [args]   creates, links, tags, adds context
yeti admin  <operation> [args]   deletes, runs tasks, manages users
yeti mcp    [--tier read|mutate|admin]
```

The tier is part of the command, not a flag, so that:

- a reader, human or agent, sees what a command can do before running it;
- agent harnesses can allow a tier by prefix. In Claude Code,
  `Bash(yeti read:*)` lets every read run unprompted while `yeti mutate` still
  asks;
- an operation can't silently gain write access through a forgotten flag.

`yeti mcp` serves the read tier by default. `--tier mutate` adds the mutate
tools, `--tier admin` all of them. Each MCP tool carries the matching
annotations (`readOnlyHint`, `destructiveHint`, `idempotentHint`) so MCP hosts
can apply their own approval rules too.

`yeti mcp` speaks stdio only. A streamable-HTTP server would answer anyone who
can reach its port with the API key it was started with, so it waits until it
can authenticate each MCP client.

A tier is a ceiling the client enforces, not a permission: Yeti still checks
the user's roles on every request.

## Operations

### read

| Operation | What it answers | Yeti endpoints |
|---|---|---|
| `lookup <value>...` | For each value: is it known, its type, tags, context sources, linked entities and observables, and indicators that match it. | `POST /graph/match` |
| `search <query>` | Objects matching a phrase, by keyword and by meaning. | `POST /search/`, `POST /search/semantic` |
| `get <ref>` | One object in full. | `GET /{observables,entities,indicators,dfiq}/{id}` |
| `neighbors <ref>` | Objects linked to one object, within a bounded number of hops. | `POST /graph/search` |
| `yara-bundle` | A compilable YARA bundle for rules given by reference or tag, dependencies included. | `POST /indicators/yara/bundle` |
| `types` | The object types this server knows, by family. | `GET /system/types` |

`lookup` never adds unknown values, although `/graph/match` can. It refangs
defanged values (`evil[.]example[.]com`, `hxxp://`) and reports the value it
looked up as `normalized`.

### mutate (milestone 2)

`add`, `link`, `tag`, `add-context`: each creates or updates rather than
duplicates, reports per-item results for batches, and supports `--dry-run`.
Writes identify the agent (see Attribution).

### admin (later)

`delete`, `run-task`, user and API key management.

## Output contract

The operations' output is the client's API. It is designed for a reader that
pays for every token.

- **JSON on standard output**, compact by default; `--pretty` indents it.
  Exit code 0 on success, 1 on a failed operation, 2 on bad usage.
- **One envelope.** Success: `{"data": ...}`, plus `"truncated"` when results
  were cut short. Failure: `{"error": {"code", "message", "hint",
  "retryable"}}`.
- **Errors say what to do next.** `hint` names the fix: the valid types for an
  unknown type, the missing configuration, the permission a 403 needs.
  `retryable` is true only when the same call can succeed later (timeouts,
  rate limits, 5xx).
- **Stable references.** Every object is referred to as `<collection>/<id>`,
  e.g. `observables/123`, the form Yeti's graph uses. `get` and `neighbors`
  accept them as they are.
- **Summaries first.** List results carry a summary per object (reference,
  type, name or value, tag names, description truncated to a line). Full
  objects come from `get`.
- **No silent truncation.** Every list has a limit (`--limit`, defaults per
  operation). When results are cut, `"truncated": {"omitted": N, "hint": …}`
  says how many were left out and how to ask for them. Lists inside a result,
  like a looked-up value's entities, report what they cut in that result's
  `"omitted"`.
- **Warnings.** `"warnings"` lists problems that didn't stop the operation,
  e.g. semantic search being unavailable on the server.
- **Deterministic.** The same data gives the same output, in the same order.

## Authentication and configuration

- `YETI_URL` and `YETI_API_KEY` in the environment, or `--url` / `--api-key`.
- The API key is exchanged at `/auth/api-token` for a short-lived session
  token, cached for the process and exchanged again on a 401. The agent never
  sees a token.
- Later: Google access tokens through `/auth/google-access-token`, for agents
  that hold Google credentials but no Yeti key.

## Attribution (milestone 2)

Every request sends `User-Agent: goyeti/<version>`, plus the agent's name when
given (`--agent-name` / `YETI_AGENT_NAME`). Context written by `mutate` uses the
source `agent:<name>`, so Yeti's timeline shows what an agent changed.

## Testing

- **Unit tests** against `net/http/httptest` servers that return recorded Yeti
  responses: decoding, output shape, errors and tiers.
- **End-to-end tests** against a real Yeti, behind the `integration` build
  tag. They seed a small investigation, with names unique to the run, and check
  every operation against it. CI runs them against a yeti-docker production
  deployment of both the last Yeti release and Yeti's `main` (the `dev`
  images), on every pull request and weekly. Locally:
  `YETI_URL=... YETI_API_KEY=... go test -tags integration ./...`.
- **Evals** (later): a short list of analyst questions with the operations we
  expect an agent to call. Tool names and descriptions are the interface an
  agent actually reads, so changes to them are measured, not guessed.

## Server-side changes that would help

- **A version endpoint.** Yeti doesn't report its version (its OpenAPI spec
  says `0.1.0`), so the client can't warn about an incompatible server.
- **A bounded investigation endpoint.** `neighbors` maps onto `/graph/search`
  today; yeti#1338 (draft) is the server half of a richer version.
- **Per-value grouping in `/graph/match`.** It returns flat lists, which the
  client regroups by value. It also lists each entity once across all values,
  so an entity linked to two looked-up values shows up under the first only.

## Milestones

1. **read** (done): `lookup`, `search`, `get`, `neighbors`, `yara-bundle`,
   `types`, through the CLI and `yeti mcp` (stdio); API key authentication;
   SKILL.md; static binaries from goreleaser.
2. **mutate**: `add`, `link`, `tag`, `add-context`, with `--dry-run` and
   attribution.
3. **admin** and evals.

## Decisions

| Decision | Why |
|---|---|
| Go, one static binary | Installs anywhere an agent runs; no runtime to manage. |
| CLI and MCP over one operations package | Coding agents use a CLI and a skill cheaply; other hosts need MCP; neither can drift from the other. |
| The tier is the first word of the command | Visible in the command, allowlistable by prefix, can't be enabled by accident. |
| No generated client | The operations' output is the contract; Yeti's schemas stay internal. |
| `<collection>/<id>` references | Yeti's own graph form; works with `get` and `neighbors` as is. |
| JSON by default, compact | Agents parse it reliably and it costs fewer tokens. |
