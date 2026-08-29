# PlexusOne Design System - LLM Context

> This document is generated from `design-system.json` and provides structured context for AI agents implementing PlexusOne interfaces.

## Design System Overview

**Name**: PlexusOne
**Version**: 0.1.0
**Description**: Design system for PlexusOne - AI Agent Development Platform

---

## Color Tokens

Use these tokens instead of hardcoded color values. Reference by CSS variable: `var(--plexus-{id})`.

### Background Colors

| Token | Value | Semantic | Usage |
|-------|-------|----------|-------|
| `dark` | `#0a0e1a` | background | Primary dark background |
| `navy` | `#0f172a` | background-secondary | Card backgrounds, sections |
| `slate` | `#1e293b` | background-tertiary | Elevated surfaces, inputs |

### Brand Colors

| Token | Value | Semantic | Usage |
|-------|-------|----------|-------|
| `cyan` | `#06b6d4` | primary | Primary accent, links, active states |
| `cyan-light` | `#22d3ee` | primary-light | Hover states, highlights |
| `cyan-dark` | `#0891b2` | primary-dark | Pressed states, borders |
| `purple` | `#8b5cf6` | accent | Secondary accent, focus rings |
| `purple-light` | `#a78bfa` | accent-light | Hover states |
| `purple-dark` | `#7c3aed` | accent-dark | Pressed states |
| `pink` | `#ec4899` | tertiary | Tertiary accent, gradients |
| `pink-light` | `#f472b6` | tertiary-light | Hover states |
| `pink-dark` | `#db2777` | tertiary-dark | Pressed states |
| `violet` | `#a855f7` | highlight | Gradient accent |

### Text Colors

| Token | Value | Semantic | Usage |
|-------|-------|----------|-------|
| `text` | `#f1f5f9` | text | Primary text on dark backgrounds |
| `text-muted` | `#94a3b8` | text-muted | Secondary text, descriptions |

---

## Typography

### Font Families

| Token | Value | Usage |
|-------|-------|-------|
| `sans` | `Inter, system-ui, -apple-system, sans-serif` | Body text, UI elements |
| `mono` | `ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace` | Code blocks, technical content |

**CSS Variable**: `var(--plexus-font-{id})`

### Font Sizes

| Token | Value |
|-------|-------|
| `xs` | `0.75rem` (12px) |
| `sm` | `0.875rem` (14px) |
| `base` | `1rem` (16px) |
| `lg` | `1.125rem` (18px) |
| `xl` | `1.25rem` (20px) |
| `2xl` | `1.5rem` (24px) |
| `3xl` | `1.875rem` (30px) |
| `4xl` | `2.25rem` (36px) |

### Font Weights

| Token | Value |
|-------|-------|
| `normal` | 400 |
| `medium` | 500 |
| `semibold` | 600 |
| `bold` | 700 |

### Line Heights

| Token | Value |
|-------|-------|
| `tight` | 1.25 |
| `normal` | 1.5 |
| `relaxed` | 1.625 |

---

## Spacing

**Base Unit**: `0.25rem` (4px)

**CSS Variable**: `var(--plexus-space-{id})`

| Token | Value |
|-------|-------|
| `0` | `0` |
| `1` | `0.25rem` (4px) |
| `2` | `0.5rem` (8px) |
| `4` | `1rem` (16px) |
| `6` | `1.5rem` (24px) |
| `8` | `2rem` (32px) |
| `12` | `3rem` (48px) |
| `16` | `4rem` (64px) |

---

## Border Radius

**CSS Variable**: `var(--plexus-radius-{id})`

| Token | Value |
|-------|-------|
| `none` | `0` |
| `sm` | `0.125rem` (2px) |
| `default` | `0.25rem` (4px) |
| `md` | `0.375rem` (6px) |
| `lg` | `0.5rem` (8px) |
| `xl` | `0.75rem` (12px) |
| `full` | `9999px` (pill shape) |

---

## Color Combinations

Use these approved color combinations to ensure consistency and accessibility.

### Dark Canvas (Primary)

- **Background**: `dark` (`#0a0e1a`)
- **Text**: `text` (`#f1f5f9`)
- **Accent**: `cyan` (`#06b6d4`)
- **Usage**: Primary dark background for main content

### Card Surface

- **Background**: `navy` (`#0f172a`)
- **Text**: `text` (`#f1f5f9`)
- **Accent**: `purple` (`#8b5cf6`)
- **Usage**: Elevated card surfaces, modal backgrounds

### Gradient Accent

- **Background**: `dark` (`#0a0e1a`)
- **Text**: `text` (`#f1f5f9`)
- **Accent**: Gradient from `cyan` → `purple` → `pink`
- **Usage**: Hero sections, CTAs, brand emphasis

---

## Anti-Patterns

**DO NOT:**

1. **Hardcode color values** - Always use CSS variables (`var(--plexus-{token})`)
2. **Use cyan for large background areas** - Reserve for accents only
3. **Use text colors for backgrounds** - `text` and `text-muted` are for text only
4. **Mix color semantics** - Don't use `primary` colors where `accent` is expected
5. **Skip hover/pressed states** - Each color has `-light` (hover) and `-dark` (pressed) variants
6. **Use arbitrary spacing** - Always use spacing scale tokens
7. **Mix font families** - Use `sans` for UI, `mono` for code only

---

## Accessibility Requirements

### Color Contrast

- **Normal text**: Minimum 4.5:1 contrast ratio (WCAG AA)
- **Large text**: Minimum 3:1 contrast ratio
- **UI components**: Minimum 3:1 against adjacent colors

### Focus States

- All interactive elements must have visible focus indicators
- Use `purple` for focus rings: `box-shadow: 0 0 0 2px var(--plexus-purple)`
- Focus must be visible in both light and dark contexts

### Motion

- Respect `prefers-reduced-motion` media query
- Keep animations under 200ms for micro-interactions
- Provide alternative to motion-dependent information

---

## Profile-Specific Overrides

### Marp Presentations

For projected presentations, use brighter values:

| Token | Default | Marp Override |
|-------|---------|---------------|
| `cyan` | `#06b6d4` | `#00d4ff` |
| `pink` | `#ec4899` | `#e040fb` |
| `text` | `#f1f5f9` | `#e8eaf6` |

### PDF/Print

For print output, invert to light background:

| Token | Default | Print Override |
|-------|---------|----------------|
| `dark` | `#0a0e1a` | `#ffffff` |
| `navy` | `#0f172a` | `#f8fafc` |
| `text` | `#f1f5f9` | `#1e293b` |

---

## Implementation Checklist

When implementing PlexusOne UI:

- [ ] Import CSS variables from design system
- [ ] Use semantic color tokens (not hardcoded hex values)
- [ ] Apply spacing scale for margins/padding
- [ ] Use approved font families
- [ ] Implement all interactive states (default, hover, pressed, focus, disabled)
- [ ] Verify color contrast meets WCAG AA
- [ ] Test with keyboard navigation
- [ ] Respect reduced motion preference
