---
name: yeti
description: Query the Yeti threat intelligence platform with the yeti command. Use it to check whether IPs, domains, URLs or hashes are known threats, find malware, actors and campaigns by name, pivot through their links, and fetch YARA rules. Needs YETI_URL and YETI_API_KEY.
---

# Yeti

`yeti` queries a Yeti server. The first word after `yeti` is the tier:
`yeti read …` never changes anything in Yeti.

## Investigate

1. **Observables first.** For IPs, hostnames, URLs or hashes, defanged or not,
   look them all up in one call:

   ```
   yeti read lookup 203.0.113.7 evil[.]example[.]com 44d88612fea8a8f36de82e1278abb02f
   ```

   Each value says whether Yeti knows it (`known`), its `tags` and
   `context_sources`, linked `entities` (malware, actors, campaigns…) and
   `observables`, and `indicators` that match it. Unknown values are reported,
   never added. Pipe long lists in: `cut -f1 iocs.tsv | yeti read lookup`.

2. **Names and topics.** For "Emotet", "TA505" or "credential dumping":

   ```
   yeti read search emotet
   ```

   Hits found by meaning come first, with a 0-1 `score`.

3. **Pivot.** Every object has a reference such as `entities/45`. Follow it:

   ```
   yeti read neighbors entities/45                         # direct links
   yeti read neighbors entities/45 --target-type observable --hops 2
   yeti read get entities/45                              # every field
   ```

4. **Detection content.**

   ```
   yeti read yara-bundle --tag emotet --raw > emotet.yar
   yeti read yara-bundle indicators/67 indicators/68
   ```

`yeti read types` lists the object types, for `--target-type`.

## Reading the output

Output is one line of JSON. Add `--pretty` to read it yourself.

- Success: `{"data": …}`. When results were cut short, `"truncated":
  {"omitted": N, "hint": …}` says how many and how to get them; inside a
  lookup result, `"omitted"` counts what each list left out. Raise `--limit`
  when you need more.
- `"warnings"` lists problems that didn't stop the operation, e.g. semantic
  search being unavailable.
- Failure: `{"error": {"code", "message", "hint", "retryable"}}`. Follow the
  `hint`. Retry only when `retryable` is true.

Exit codes: 0 success, 1 failed operation, 2 bad usage.

## Setup

Set `YETI_URL` (e.g. `https://yeti.example.com`) and `YETI_API_KEY`.
Optionally set `YETI_AGENT_NAME` so Yeti can tell your requests apart.
