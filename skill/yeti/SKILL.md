---
name: yeti
description: Query and update the Yeti threat intelligence platform with the yeti command. Use it to check whether IPs, domains, URLs or hashes are known threats, find malware, actors and campaigns by name, pivot through their links, fetch YARA rules, record findings as observables, entities, links, tags and context, delete objects, and run feeds and analytics. Needs YETI_URL and YETI_API_KEY.
---

# Yeti

`yeti` queries and updates a Yeti server. The first word after `yeti` is the
tier: `yeti read …` never changes anything in Yeti; `yeti mutate …` adds to it;
`yeti admin …` deletes from it and runs its tasks.

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

## Record findings

`yeti mutate` creates what's missing and leaves what exists, so running a
command twice changes nothing more. Add `--dry-run` to see what would change
first.

```
yeti mutate add observables 203.0.113.7 evil[.]example[.]com --tag campaign-x
yeti mutate add entity malware Qakbot --description "Banking trojan" --tag loader
yeti mutate add indicator yara qakbot_loader --pattern @qakbot.yar --diamond capability
yeti mutate link observables/123 c2 entities/45
yeti mutate tag observables/123 entities/45 --tag campaign-x
yeti mutate untag observables/123 --tag false-positive
yeti mutate add-context observables/123 note="beaconing every 60s" first_seen=2026-10-01
```

- `add observables` guesses each value's type and reports values it doesn't
  recognize as `failed`; pass `--type` for types it can't guess, like `path`.
- An existing entity or indicator is never changed by `add`: only the tags
  are added.
- Context is recorded under the source `agent:<your name>`
  (`YETI_AGENT_NAME`).

Each item reports its `action`: `created`, `exists`, `linked`, `tagged`,
`untagged`, `unchanged`, `context_added` or `failed` with an `error`; the
admin tier adds `deleted`, `unlinked`, `absent`, `queued`, `enabled` and
`disabled`. When some
items fail, the others still happen, `"failed"` counts the failures and the
exit code is 1.

## Clean up and run tasks

`yeti admin` removes things for good, so look before deleting: `--dry-run`
says what each object is and how many links would go with it. Deleting
needs `--yes`. Only delete what you created or were asked to remove.

```
yeti admin delete observables/123 entities/45 --dry-run
yeti admin delete observables/123 entities/45 --yes
yeti admin unlink observables/123 c2 entities/45 --yes
```

Objects and links that don't exist report `absent`, so deleting twice is
safe.

Tasks are Yeti's feeds, analytics and exports. `yeti read tasks` lists them,
with their status and last run; `run-task` queues a run, which Yeti does in the
background.

```
yeti read tasks abuse
yeti admin enable-task AbuseIPDB
yeti admin run-task AbuseIPDB
yeti admin run-task VTDomainReport --param value=example.com
yeti admin disable-task AbuseIPDB
```

- `run-task`, `enable-task` and `disable-task` take a task's exact name; a
  wrong one suggests close matches.
- A task must be enabled to run. Oneshot analytics act on one observable:
  pass its value as `--param value=…`.
- Follow a run with `yeti read tasks <name>`: `status`, `status_message` and
  `last_run`.

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

Exit codes: 0 success, 1 failed operation or failed items, 2 bad usage.

## Setup

Set `YETI_URL` (e.g. `https://yeti.example.com`) and `YETI_API_KEY`.
Optionally set `YETI_AGENT_NAME` so Yeti can tell your requests apart.
