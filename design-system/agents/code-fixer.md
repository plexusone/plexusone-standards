---
name: code-fixer
description: Fixes PlexusOne design system violations
extends: design-system-spec/design-fixer
model: sonnet
tools: [Read, Grep, Glob, Edit, Write, Bash]
allowedTools: [Read, Grep, Glob, Edit, Write, "Bash(dss *)"]
context:
  - design-system.json
  - llm-context.md
---

# PlexusOne Code Fixer

This agent extends the generic `design-fixer` agent from design-system-spec with PlexusOne-specific token mappings.

## Configuration

- **Design System**: PlexusOne v0.1.0
- **Base Agent**: `design-system-spec/agents/specs/design-fixer.md`
- **Context**: `llm-context.md` (generated from design-system.json)

## PlexusOne Token Mappings

### Colors

| Hardcoded | PlexusOne Token |
|-----------|-----------------|
| `#0a0e1a` | `var(--plexus-dark)` |
| `#0f172a` | `var(--plexus-navy)` |
| `#1e293b` | `var(--plexus-slate)` |
| `#06b6d4` | `var(--plexus-cyan)` |
| `#22d3ee` | `var(--plexus-cyan-light)` |
| `#0891b2` | `var(--plexus-cyan-dark)` |
| `#8b5cf6` | `var(--plexus-purple)` |
| `#a78bfa` | `var(--plexus-purple-light)` |
| `#7c3aed` | `var(--plexus-purple-dark)` |
| `#ec4899` | `var(--plexus-pink)` |
| `#f472b6` | `var(--plexus-pink-light)` |
| `#db2777` | `var(--plexus-pink-dark)` |
| `#a855f7` | `var(--plexus-violet)` |
| `#f1f5f9` | `var(--plexus-text)` |
| `#94a3b8` | `var(--plexus-text-muted)` |

### Spacing

| Hardcoded | PlexusOne Token |
|-----------|-----------------|
| `4px` | `var(--plexus-space-1)` |
| `8px` | `var(--plexus-space-2)` |
| `16px` | `var(--plexus-space-4)` |
| `24px` | `var(--plexus-space-6)` |
| `32px` | `var(--plexus-space-8)` |
| `48px` | `var(--plexus-space-12)` |
| `64px` | `var(--plexus-space-16)` |

### Typography

| Hardcoded | PlexusOne Token |
|-----------|-----------------|
| `Inter, system-ui, ...` | `var(--plexus-font-sans)` |
| `ui-monospace, ...` | `var(--plexus-font-mono)` |

### Border Radius

| Hardcoded | PlexusOne Token |
|-----------|-----------------|
| `2px` | `var(--plexus-radius-sm)` |
| `4px` | `var(--plexus-radius-default)` |
| `6px` | `var(--plexus-radius-md)` |
| `8px` | `var(--plexus-radius-lg)` |
| `12px` | `var(--plexus-radius-xl)` |
| `9999px` | `var(--plexus-radius-full)` |

## PlexusOne-Specific Fixes

### Gradient Syntax

```css
/* PlexusOne brand gradient */
background: linear-gradient(
  135deg,
  var(--plexus-cyan),
  var(--plexus-purple),
  var(--plexus-pink)
);
```

### Focus Ring

```css
/* PlexusOne focus style */
.interactive:focus-visible {
  outline: none;
  box-shadow: 0 0 0 2px var(--plexus-purple);
}
```

### Dark Background Pattern

```css
/* Primary dark canvas */
.container {
  background: var(--plexus-dark);
  color: var(--plexus-text);
}

/* Elevated surface */
.card {
  background: var(--plexus-navy);
}

/* Tertiary surface */
.input {
  background: var(--plexus-slate);
}
```

## Invocation

```bash
# Run via AssistantKit
assistantkit run code-fixer --issues validation-output.md

# The fixer reads design-system.json for token values
```

## See Also

- Generic agent: `github.com/plexusone/design-system-spec/agents/specs/design-fixer.md`
- PlexusOne tokens: `llm-context.md`
- Violation examples: `examples/violations.tsx`
