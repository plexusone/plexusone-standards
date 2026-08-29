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
├── plexusone-rest.api-style.json   Single source of truth — 17 PX-* rules,
│                                    3 conformance levels (minimum/standard/exemplary)
├── plexusone-rest.spectral.yaml    Generated — deterministic Spectral ruleset
├── plexusone-rest.rubric.json      Generated — LLM evaluation rubric
└── plexusone-rest.guide.md         Generated — human-readable Markdown guide
```

The three generated files are build artifacts, never hand-edited — after
changing `plexusone-rest.api-style.json`, regenerate them:

```bash
api-style generate spectral -f plexusone-rest.api-style.json -o plexusone-rest.spectral.yaml
api-style generate rubric   -f plexusone-rest.api-style.json -o plexusone-rest.rubric.json
api-style generate guide    -f plexusone-rest.api-style.json -o plexusone-rest.guide.md
```

Both the schema-validity of the source profile and the spectral ruleset's
actual behavior are proven, not assumed: `ajv` validates the profile
against api-style-spec's own JSON Schema (`schemakit lint` alone only
checks property casing, not structural conformance — insufficient on its
own), and linting an api-style-spec exemplar against this profile
confirms PX-003 genuinely catches an unversioned `/v1/...` path rather
than passing everything through.

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
