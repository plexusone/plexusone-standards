# ROADMAP — PlexusOne Standards — Repo Restructure & REST API Style Guide

**Initiative:** `INIT-PLEXUSONESTANDARDS-001`
**Repository:** `github.com/plexusone/plexusone-standards`

The repo (renamed from plexusone-design-system) becomes the single front door
for PlexusOne standards, organized by domain: `design-system/` (existing
content, instance of design-system-spec) and `api/` (new, instance of
api-style-spec — starting with `api/rest/`). Spec products (api-style-spec,
design-system-spec) stay org-neutral; house content lives here. Note for the
profile work: api-style-spec's profile `extends` is schema-only (loader does
not resolve inheritance), so the plexusone-rest profile is self-contained
with `source:` citations to default PO-* rules.

Full PX-rule inventory for `RMI-PLEXUSONESTANDARDS-003`: OpenAPI 3.1 generated
from Go types (Huma); camelCase properties with documented telemetry and
OpenAI-compat exceptions; kebab-case /api/v1 paths; RFC 9457 errors; cursor
pagination; prefixed string IDs and ISO 8601 UTC timestamps; SystemForge
bearer/BFF auth; {app}:{resource}:{verb} scopes; PRN cross-product references;
run resources with normalized status vocabulary; SSE for server push;
additive-only versioning with CI diff; verbResource operationIds; conformance
levels minimum/standard/exemplary.

## Phase 1 — Repository Restructure

**Theme:** One standards repo, domain directories, nothing breaks

- [x] `RMI-PLEXUSONESTANDARDS-001` Move design-system content into design-system/ (design-system.json, agents/, loops/, examples/, llm-context.md), update package.json name and internal paths, root README as domain index
- [ ] `RMI-PLEXUSONESTANDARDS-002` Update external references to the old repo name/paths (website build, any CI consumers); adopt conventional-commit scopes (api, design-system) in CHANGELOG.json workflow
  - Depends on: `RMI-PLEXUSONESTANDARDS-001`

## Phase 2 — REST API Style Guide (api/rest)

**Theme:** The plexusone-rest profile — machine-readable house rules, generated artifacts

- [x] `RMI-PLEXUSONESTANDARDS-003` Author api/rest/plexusone-rest.api-style.json profile: PX rules for Go-first OpenAPI 3.1, naming, errors, pagination, SystemForge auth and scopes, PRN references, run resources, SSE, versioning; conformance levels
  - Depends on: `RMI-PLEXUSONESTANDARDS-001`
- [ ] `RMI-PLEXUSONESTANDARDS-004` Generated artifacts and validation: spectral ruleset + LLM rubric + Markdown guide via api-style generate, schema-validate the profile, lint an exemplar spec; resolve omniagent-rest's dangling basedOn pointer to this profile's location
  - Depends on: `RMI-PLEXUSONESTANDARDS-003`
- [x] `RMI-PLEXUSONESTANDARDS-005` Naming-conventions standard: omni*/forge/*-core module naming rules, plain functional names otherwise; PRN and scope grammar cross-references to resourceref and crossforge
  - Depends on: `RMI-PLEXUSONESTANDARDS-001`
- [x] `RMI-PLEXUSONESTANDARDS-006` PRN instance profile (resourceref/ domain): prn scheme definition against the resourceref profile format — authorities, path shapes, resource-type registries, version/alias rules; schema-validated in CI
  - Depends on: `RMI-PLEXUSONESTANDARDS-001`

## Phase 3 — Adoption in Forge Repos

**Theme:** Dogfood the profile — .api-style.yaml + CI lint where the APIs live

- [ ] `RMI-DASHFORGE-018` dashforge: .api-style.yaml pointing at plexusone-rest, api-style lint of the Huma-generated OpenAPI in CI, exceptions documented for legacy pre-Huma endpoints
  - Depends on: `RMI-PLEXUSONESTANDARDS-004`
- [ ] `RMI-ACTIONFORGE-040` actionforge: .api-style.yaml + CI lint for the server API surface
  - Depends on: `RMI-PLEXUSONESTANDARDS-004`
- [ ] `RMI-AGENTFORGE-003` agentforge: .api-style.yaml + CI lint for the BFF API surface (OpenAI-compatible endpoints continue under the omniagent-rest profile)
  - Depends on: `RMI-PLEXUSONESTANDARDS-004`

## Phase 4 — One Guide for Every Product

**Theme:** plexusone-rest covers every product built on SystemForge, including APIs published to customers

- [x] `RMI-PLEXUSONESTANDARDS-007` plexusone-rest v1.1.0: the `/api/{api}/v{N}/` path form beside the `/api/v{N}/` default, camelCase path parameters, organization selectors in paths authorized against the principal, slugs as alternate keys, and new rules for operation audiences, preconditions, dry runs, content-addressed uploads and unauthenticated endpoints
  - Depends on: `RMI-PLEXUSONESTANDARDS-004`
