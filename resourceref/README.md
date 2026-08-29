# resourceref (prn scheme instance)

PlexusOne's own `prn` scheme — the `resourceref` [scheme.Profile](https://github.com/plexusone/resourceref/blob/main/scheme/profile.go)
instance that defines what a valid `prn://` reference looks like across
the Forge suite. This is the *instance*; the mechanism (the `ResourceRef`
type, the profile engine) lives in the org-neutral
[resourceref](https://github.com/plexusone/resourceref) module — the same
relationship this repo has with design-system-spec and api-style-spec.

## What's here

```
resourceref/
├── prn-profile.json        The instance: authorities, shapes, resource
│                           types, version grammar — data, not code
├── go.mod                  A nested Go module, scoped to this domain,
│                           for the one thing that needs Go: validation
└── prn_profile_test.go     Loads the profile via scheme.Load and checks
                            it against real prn:// references used across
                            the ecosystem's design docs
```

## The scheme

```text
prn://{authority}/{app}/{org}/{resourceType}/{id}[@{version|alias}]   tenant-scoped
prn://{authority}/{resourceType}/{path...}                            ecosystem artifact

prn://plexusone/dashforge/org_7f3/notebook/nb_churn@prod
prn://plexusone/actionforge/org_1/workflow-run/run_abc
prn://productbuildershq/roadmap/RMI-142
```

- **Authorities**: `plexusone` (the Forge suite and crossforge/omni*
  engines), `productbuildershq` (roadmap/initiative artifacts),
  `grokify` (shared foundation modules).
- **Shapes**: `tenant-scoped` (app/org/resourceType/id — most Forge
  resources) and `ecosystem-artifact` (resourceType/id... — org-less
  artifacts like roadmap items).
- **Resource types**: registered per (authority, app) only where the
  vocabulary is settled today (dashforge, actionforge). Every other app
  is intentionally unrestricted — the registry is additive, not a gate
  new apps have to clear before they can mint PRNs. Add entries as each
  app's resource-type vocabulary stabilizes.
- **Versions**: `@14` (immutable version) or `@prod` (mutable alias),
  optional at the grammar level. Whether a given caller *requires* a
  pinned reference (crossforge invocation targets do) is enforced by that
  caller, not by this profile.

## Validating changes

```bash
cd resourceref && go test ./...
```

Any edit to `prn-profile.json` that breaks structural validity
(`scheme.Load`) or any of the golden reference examples fails here — this
is the gate the roadmap calls "schema-validated in CI," run via
`plexusone/.github`'s shared `go-ci`/`go-lint` workflows scoped to this
subdirectory.

**CI caveat**: `go.mod` currently carries a local `replace` for
`github.com/plexusone/resourceref`, since that module has no tagged
release or pushed remote yet (same situation actionforge and omniworkflow
are already in with their own temporary replaces). The workflows will run
red on GitHub — no sibling checkout to resolve the replace against — until
resourceref is published and the replace is dropped
(`RMI-RESOURCEREF-006`). `go test ./...` here, locally, is the real
correctness signal in the meantime; the CI files exist now so nothing
else has to change when that dependency clears.

## Consuming this profile

```go
import (
    "embed"
    "github.com/plexusone/resourceref/scheme"
)

//go:embed prn-profile.json
var prnProfileJSON []byte

profile, err := scheme.Load(prnProfileJSON)
```

crossforge's PRN target layer (`RMI-CROSSFORGE-002`) is the first real
consumer — see [resourceref's embedding guide](https://github.com/plexusone/resourceref/blob/main/docs/embedding.md)
for the general pattern of adding a PRN-typed field to your own IR.
