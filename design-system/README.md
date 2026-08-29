# PlexusOne Design System

Design tokens and brand configuration for PlexusOne properties — marketing
website, web apps, presentations, and doc sites.

This is PlexusOne's own *instance* of
[design-system-spec](https://github.com/plexusone/design-system-spec)'s
format, published as the `@plexusone/design-system` npm package from this
directory. It is one domain within the
[plexusone-standards](../README.md) repository.

## Structure

```
design-system/
├── design-system.json    # Single source of truth
├── package.json
├── agents/                # Design-review/fix agent definitions
├── examples/              # Recorded compliance-check examples
├── loops/                 # Automated compliance-check loop config
└── llm-context.md         # Generated LLM-facing context (from design-system.json)
```

## Color Palette

| Token | Value | Role |
|-------|-------|------|
| `dark` | `#0a0e1a` | Primary background |
| `navy` | `#0f172a` | Card backgrounds |
| `cyan` | `#06b6d4` | Primary accent |
| `purple` | `#8b5cf6` | Secondary accent |
| `pink` | `#ec4899` | Tertiary accent |

## Validation

```bash
ajv validate --spec=draft2020 --strict=false \
  -s path/to/design-system.schema.json \
  -d design-system.json
```

## License

MIT
