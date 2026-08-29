# PlexusOne Standards

PlexusOne's own house standards — the single front door for "what does
PlexusOne do here," for humans and agents alike.

This repo holds *instances* of PlexusOne's standards specs, organized by
domain. Each spec product stays org-neutral and reusable by anyone; this
repo is where PlexusOne's own configuration of each one lives.

| Domain | Spec product | Status |
|--------|--------------|--------|
| [`design-system/`](design-system/) | [design-system-spec](https://github.com/plexusone/design-system-spec) | Active — design tokens for marketing site, web apps, presentations, doc sites |
| `api/` | [api-style-spec](https://github.com/plexusone/api-style-spec) | Planned — the `plexusone-rest` REST API style profile |

Future domains (module/product naming conventions, PRN scheme
cross-references) land here as they're specced; see
[`docs/specs/initiatives/INIT-PLEXUSONESTANDARDS-001`](docs/specs/initiatives/INIT-PLEXUSONESTANDARDS-001)
for the current roadmap.

## Conventions

- One root `CHANGELOG.json`/`CHANGELOG.md` for the whole repo, using
  conventional-commit scopes per domain (`feat(design-system): ...`,
  `feat(api): ...`) rather than a changelog per domain.
- Each domain directory is self-contained with its own README describing
  its structure and how to validate/consume it.

## License

MIT
