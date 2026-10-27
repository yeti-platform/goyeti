# goyeti

A client for [Yeti](https://github.com/yeti-platform/yeti) built for AI
agents: one `yeti` binary with a command line and an
[MCP](https://modelcontextprotocol.io) server.

```
yeti read lookup evil.example.com 203.0.113.7
yeti read search "cobalt strike beacons"
yeti mcp
```

The first word of every command is its tier — `read`, `mutate` or `admin` —
so what a command can change is visible before it runs.

**Status:** being rewritten from scratch. See [docs/DESIGN.md](docs/DESIGN.md)
for what we are building. The previous generated client is available as
`github.com/yeti-platform/goyeti@v0.1.0`.
