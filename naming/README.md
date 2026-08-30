# Module & Product Naming

PlexusOne's naming conventions for Go modules and products, and how they
connect to the two grammars that actually run through the ecosystem's
code: PRN resource references and OAuth scopes. Unlike `design-system/`
and `api/rest/`, this domain isn't an instance of an external spec
product — it's original PlexusOne content, decided through direct
practice rather than specced up front.

## The four naming patterns

### `omni*` — abstraction layer, or an earned flagship runtime

Two distinct meanings share the prefix:

1. **Abstraction layer** — a stable interface over multiple interchangeable
   provider/backend implementations. `omnillm` (LLM provider abstraction:
   `omnillm-core` types + `omnillm` + thick providers `omni-openai`,
   `omni-anthropic`, ...), `omniworkflow` (orchestration abstraction: in-memory,
   durable, Temporal, Step Functions providers behind one interface),
   `omnistorage`, `omniobserve`.
2. **Flagship runtime** — a single-user, local-first application that has
   earned the prefix by being genuinely useful on its own, independent of
   any `*forge` product orchestrating it. `omniagent` (an OpenClaw-derived
   agent runtime, runnable locally, hosted solo, or orchestrated by
   AgentForge in the cloud) is the first of these.

**The exception-count rule**: two flagship exceptions to the
abstraction-layer default is precedent; a third ends the "exception"
framing entirely and the flagship sense becomes an equally established
meaning of the prefix, not a special case. `omninotebook` is reserved as
the second flagship exception — earned only once it's genuinely
multi-kernel/multi-format (an abstraction in its own right) **or** reaches
independent flagship status with users outside DashForge, following the
same local → hosted-solo → forge-orchestrated arc as `omniagent`.
`omnikernel` (a hypothetical kernel-protocol abstraction extracted from
`omninotebook`) is a reserved name, not yet a module — extract it only
when a second kernel backend actually justifies the abstraction.

### `*forge` — a multi-tenant, cloud-hosted product application

`DashForge`, `ActionForge`, `AgentForge`. Each is composed on
`SystemForge` for identity, tenancy, and authorization, and each is
something an end user directly interacts with — never the name for a
library or engine. The engine *behind* a forge is typically `omni*` or
plain-named: `omniworkflow` behind ActionForge, `omniagent` behind
AgentForge.

`SystemForge` itself carries the suffix despite being the suite's
foundation layer, not a vertical product — a deliberate, single exception:
it ships its own IAM console and UI shell (a real application, not just a
library), so it still qualifies as a product app under the rule. It's the
one every other forge depends on, not a peer vertical.

### `*-core` — the deterministic types/IR layer of its sibling module

The pattern is `{name}-core` + `{name}`: `omnillm-core`/`omnillm`,
`omninotebook-core`/`omninotebook`, `omniroadmap-core`/`omniroadmap`,
`omnidevx-core`/`omnidevx`. The `-core` module holds contracts, types, and
an IR — deterministic, infrastructure-independent, and lean on
dependencies (often near-stdlib). It **never** imports its own sibling
implementation module; the dependency runs one direction only, so a
consumer that needs only the types never pulls in the runtime.

### Plain functional names — everything else

Used when a module is a leaf library, a spec-instance mechanism, or a
tool that doesn't fit the abstraction-engine or product-app shape:
`resourceref`, `crossforge`, `deppolicy`, `guardsql`, `schemakit`,
`agentauth`.

**`crossforge` is worth calling out explicitly** — its name ends in
"-forge" but it is *not* a `*forge` product: it's the interop contract
(PRN invocation, run vocabulary, the resolver interface) that every forge
and every `omni*` engine can build on, deliberately kept dependency-free
and permissively licensed. The name is a nod to being the suite's lingua
franca, not a claim to be a tenant-facing application — the one name in
the ecosystem that could be mistaken for breaking its own rule if you
don't know the story.

## Choosing a name for a new module

1. **Does it ship a UI/console end users interact with directly, composed
   on SystemForge for tenancy?** → `*forge`.
2. **Is it a stable interface over multiple interchangeable
   provider/backend implementations?** → `omni*` (abstraction sense).
3. **Is it a local-first runtime that is genuinely, independently useful
   before any forge orchestrates it — and are you prepared to defend it
   as only the second or third such exception?** → `omni*` (flagship
   sense). If you're past the third, this is no longer an exception to
   justify — treat it as an established sense of the prefix.
4. **Is it the types/IR layer of an existing sibling module?** →
   `{sibling}-core`, one direction of dependency only.
5. **None of the above?** → a plain, functional name. Don't reach for a
   prefix or suffix to signal importance; `resourceref` and `crossforge`
   are foundational precisely because their names don't claim to be.

## Why this matters beyond aesthetics: naming feeds the grammars that run through the code

Module and product names aren't just labels — the `*forge` product name
is literally the `{app}` value that appears in two grammars everywhere in
the ecosystem:

```text
OAuth scope:   {app}:{resource}:{verb}         e.g. dashforge:notebook:run
PRN:           prn://{authority}/{app}/{org}/{resourceType}/{id}
               e.g. prn://plexusone/dashforge/org_7f3/notebook/nb_churn@prod
```

- **PRN grammar** — the mechanism (`ResourceRef`, the `scheme.Profile`
  engine) lives in [resourceref](https://github.com/plexusone/resourceref);
  PlexusOne's own `prn` scheme instance — the real authorities, shapes,
  and resource-type registry — lives in this repo at
  [`resourceref/prn-profile.json`](../resourceref/prn-profile.json).
- **Scope grammar** — defined as rule `PX-011` in
  [`api/rest/plexusone-rest.api-style.json`](../api/rest/plexusone-rest.api-style.json),
  mapping one-to-one onto SystemForge's authz resource types and the PRN
  `app`/`resourceType` segments — the same string means the same thing in
  an OAuth scope and in a resource reference.
- **Invocation** — [crossforge](https://github.com/plexusone/crossforge)
  builds the invocation contract, run vocabulary, and resolver on top of
  resourceref's PRN grammar; it does not define naming or scoping itself.

A forge product's name isn't renamed lightly — it propagates into every
scope string and every PRN minted for that product's resources for as
long as any of them exist.
