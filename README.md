# PlexusOne Standards

PlexusOne's own house standards — the single front door for "what does
PlexusOne do here," for humans and agents alike.

This repo holds *instances* of PlexusOne's standards specs, organized by
domain. Each spec product stays org-neutral and reusable by anyone; this
repo is where PlexusOne's own configuration of each one lives.

| Domain | Spec product | Status |
|--------|--------------|--------|
| [`design-system/`](design-system/) | [design-system-spec](https://github.com/plexusone/design-system-spec) | Active — design tokens for marketing site, web apps, presentations, doc sites |
| [`api/rest/`](api/rest/) | [api-style-spec](https://github.com/plexusone/api-style-spec) | Active — the `plexusone-rest` REST API style profile |
| [`resourceref/`](resourceref/) | [resourceref](https://github.com/plexusone/resourceref) | Active — the `prn` scheme instance (authorities, shapes, resource types) |
| [`naming/`](naming/) | — (original PlexusOne content) | Active — module/product naming conventions and their tie to PRN/scope grammar |

New domains land here as they're specced; see
[`docs/specs/initiatives/INIT-PLEXUSONESTANDARDS-001`](docs/specs/initiatives/INIT-PLEXUSONESTANDARDS-001)
for the current roadmap.

## Conventions

- **One root `CHANGELOG.json`/`CHANGELOG.md`** for the whole repo (via
  [structured-changelog](https://github.com/grokify/structured-changelog)),
  not a changelog per domain.
- **Scope commits by domain**: `feat(design-system): ...`,
  `fix(api): ...`. Domain scopes match directory names one-to-one; add a
  new scope only when a new domain directory is added. Omit the scope for
  changes to the repo's own structure (this README, the domain index
  table, restructuring) rather than reaching for one domain's name.
  `schangelog parse-commits` reads the scope directly off each commit
  when triaging changes into `CHANGELOG.json` — there's no separate
  "scope" field in the changelog IR itself, so the commit message is the
  only place this information has to be right.
- Each domain directory is self-contained with its own README describing
  its structure and how to validate/consume it.

## License

MIT
