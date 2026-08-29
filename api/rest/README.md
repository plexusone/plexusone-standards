# plexusone-rest

PlexusOne's REST API style guide — a
[api-style-spec](https://github.com/plexusone/api-style-spec) profile
instance defining how Forge product APIs (DashForge, ActionForge,
AgentForge, SystemForge, OmniNotebook) and the
[crossforge](https://github.com/plexusone/crossforge) interop contract
present themselves over HTTP: naming, errors, pagination, authorization
scopes, PRN cross-product references, run resources, streaming, and
versioning.

## Structure

```
api/rest/
└── plexusone-rest.api-style.json   Single source of truth — 17 PX-* rules,
                                     3 conformance levels (minimum/standard/exemplary)
```

Generated artifacts (Spectral ruleset, LLM evaluation rubric, human-readable
Markdown guide) are produced from this file via `api-style generate`
rather than hand-maintained — see
[`RMI-PLEXUSONESTANDARDS-004`](../../docs/specs/initiatives/INIT-PLEXUSONESTANDARDS-001/ROADMAP.md).

## Using this profile

Reference it by file path — it isn't a built-in api-style-spec profile:

```bash
api-style lint openapi.yaml --profile /path/to/plexusone-standards/api/rest/plexusone-rest.api-style.json
api-style generate guide -f /path/to/plexusone-standards/api/rest/plexusone-rest.api-style.json
```

Or via `.api-style.yaml` in the consuming repo:

```yaml
profile: ../plexusone-standards/api/rest/plexusone-rest.api-style.json
level: standard
```

## Design basis

Restated from api-style-spec's `default` profile (citing `default PO-xxx`
per rule) rather than relying on `extends` resolution — the api-style-spec
loader treats `extends` as schema metadata only, not inherited rules, so
this profile is self-contained. Adds PlexusOne-specific rules for
SystemForge authentication/scopes, `prn://` cross-product references
(validated by [resourceref](https://github.com/plexusone/resourceref)),
crossforge run resources and status vocabulary, and SSE streaming
conventions.
