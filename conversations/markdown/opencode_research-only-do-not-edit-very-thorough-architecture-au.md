# Audit feature MFEs (@explore subagent)

- **ID**: `ses_016c95ac0ffeBaXtVzxujZB77R`
- **Source Tool**: `opencode`
- **Date**: `2026-08-10 08:08:20`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Research only, do not edit. Very thorough architecture audit of all ~/Development/qcash-ui-* repositories EXCEPT qcash-ui itself, landing-page, header-footer, header-footer-bri, dashboard, main-page, design-system, base. Group repositories by architecture pattern and business domain. Inspect package.json, next/webpack/module federation configs, src/pages/services/store patterns, test setup, README, and dev branch where locally available. Identify concrete systemic and repository-specific technical debt, documentation gaps, and coupling. Return exact file:line references for representative/high-risk evidence and a complete repo grouping suitable for adding to root ARCHITECTURE.md. Ignore node_modules, .next, generated @mf-types.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **glob**: Tool: glob

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (10)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **grep**: Tool: grep
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (9)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **grep**: Tool: grep

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 31: Assistant

# Architecture audit

Research-only audit completed; no files were modified. Scope: **71 repositories** matching `qcash-ui-*`, excluding the seven requested repositories. Ignored `node_modules`, `.next`, and generated `@mf-types`.

## 1. Executive summary

The estate is effectively one large, loosely governed micro-frontend platform:

- **69/71 repositories** use the legacy architecture: Next.js Pages Router, mostly Next `13.5.x`, webpack 5, and `@module-federation/nextjs-mf`.
- **2/71 repositories** are a newer architecture using Next 16, `next-rspack`, Tailwind 4, Biome, and `@module-federation/enhanced`:
  - `qcash-ui-dashboard-auto-collection`
  - `qcash-ui-forex`
- **62 repositories** hard-code the development OpenShift domain in federation configuration.
- **67 repositories** directly depend on `@ordentco/addons-auth-provider`, across many incompatible versions.
- Only **34/71 current worktrees have a test script**; 37 have no executable test contract.
- Local or remote `dev` references exist for **46 repositories**, but many checked-out `master` trees are far behind or divergent from `dev`.
- MFEs are not isolated business modules in practice. They share:
  - token storage keys,
  - login and shell routes,
  - browser events such as `showDialog`,
  - remote component URLs,
  - generated type layouts,
  - API environment-selection code,
  - response/error semantics,
  - CSS exposure conventions.

The biggest architectural risk is therefore not Module Federation itself, but the absence of centrally versioned contracts around federation, authentication, routing, API clients, runtime configuration, and testing.

---

# 2. Complete repository grouping

Unless marked **modern**, every repository below belongs to the legacy Next.js/webpack/`nextjs-mf` pattern.

## Platform administration, IAM and configuration — 19

- `qcash-ui-announcement-management`
- `qcash-ui-company-management`
- `qcash-ui-content-management-system`
- `qcash-ui-cut-off-management`
- `qcash-ui-employee-data`
- `qcash-ui-group-management`
- `qcash-ui-holiday-management`
- `qcash-ui-language-management`
- `qcash-ui-menu-management`
- `qcash-ui-onboarding-tour`
- `qcash-ui-partners-integrations`
- `qcash-ui-rma-management`
- `qcash-ui-role-management`
- `qcash-ui-sistem-management`
- `qcash-ui-subscription-management`
- `qcash-ui-update-user`
- `qcash-ui-user`
- `qcash-ui-user-basic-management`
- `qcash-ui-user-manual`

## Workflow and approval — 2

- `qcash-ui-approval-signature`
- `qcash-ui-workflow`

## Customer support — 2

- `qcash-ui-complaint`
- `qcash-ui-helpdesk-portal`

## Registration and onboarding — 5

- `qcash-ui-account-registration`
- `qcash-ui-account-validation`
- `qcash-ui-registration-giro`
- `qcash-ui-registration-online`
- `qcash-ui-user-registration-online`

## Accounts, receivables, statements and monitoring — 11

- `qcash-ui-account-management`
- `qcash-ui-account-receivable`
- `qcash-ui-account-statement`
- `qcash-ui-account-summary`
- `qcash-ui-bank-monitoring-transaction`
- `qcash-ui-beneficiary-account-management`
- `qcash-ui-global-account-statement`
- `qcash-ui-global-account-summary`
- `qcash-ui-rdn-summary`
- `qcash-ui-rdn-transaction`
- `qcash-ui-report-statement`

## Payments, cards and transfers — 12

- `qcash-ui-bifast`
- `qcash-ui-cashcard`
- `qcash-ui-credit-card`
- `qcash-ui-ewallet-topup`
- `qcash-ui-fund-transfer`
- `qcash-ui-global-external-transfer`
- `qcash-ui-global-internal-transfer`
- `qcash-ui-mass-brizzi`
- `qcash-ui-mass-transfer`
- `qcash-ui-payroll-transfer`
- `qcash-ui-rtgs`
- `qcash-ui-swift`

## Trade finance, treasury and liquidity — 5

- `qcash-ui-amend-cancel-lc`
- `qcash-ui-digi-guarantee`
- `qcash-ui-forex` — **modern**
- `qcash-ui-issuance-lc`
- `qcash-ui-liquidities`

## Government, tax and regulatory payments — 4

- `qcash-ui-local-tax-dki-jakarta`
- `qcash-ui-mpn`
- `qcash-ui-regol`
- `qcash-ui-span`

## Industry, SCM and specialized corporate products — 11

- `qcash-ui-dashboard-auto-collection` — **modern**
- `qcash-ui-dashboard-dhe`
- `qcash-ui-df-pupuk-indonesia`
- `qcash-ui-do-pertamina`
- `qcash-ui-do-pupuk-indonesia`
- `qcash-ui-do-semen-indonesia`
- `qcash-ui-mitra-asuransi`
- `qcash-ui-pencairan-bulog`
- `qcash-ui-plafon-pegadaian`
- `qcash-ui-scm`
- `qcash-ui-scm-cu`

---

# 3. Architecture patterns

## Pattern A: Legacy federated Next.js Pages Router — 69 repositories

Typical structure:

```text
pages/           route entries and federation exposure targets
containers/      stateful business screens
components/      local and remote-component wrappers
services/        duplicated Axios clients and API modules
styles/          generated Tailwind output exposed through federation
config.ts        environment-by-hostname and feature-toggle selection
```

Representative configuration:

- Federation directly exposes page files and CSS:
  - `~/Development/qcash-ui-account-management/next.config.js:11-20`
  - `~/Development/qcash-ui-workflow/federation.config.js:9-15`
- Standalone deployment and business-specific `basePath`:
  - `~/Development/qcash-ui-account-management/next.config.js:56-57`
  - `~/Development/qcash-ui-workflow/next.config.js:12-13`
- Every MFE alters Next's webpack pipeline to add `style-loader`, force a browser target, and install federation:
  - `~/Development/qcash-ui-account-management/next.config.js:58-69`
  - `~/Development/qcash-ui-global-account-summary/next.config.js:47-51`

This template is consistent enough to automate, but it has been copied rather than maintained as a shared preset.

## Pattern B: Modern Next 16/Rspack federation — 2 repositories

### `qcash-ui-dashboard-auto-collection`

- Next 16.2.9 and `@module-federation/enhanced`:
  - `~/Development/qcash-ui-dashboard-auto-collection/package.json:21-31`
- Rspack federation runs only on client builds:
  - `~/Development/qcash-ui-dashboard-auto-collection/next.config.ts:12-31`
- Exposes pages from `src/pages` and a generated stylesheet module:
  - `~/Development/qcash-ui-dashboard-auto-collection/next.config.ts:47-52`
- Has Vitest, Biome, Tailwind 4 and a host-compatibility development proxy:
  - `~/Development/qcash-ui-dashboard-auto-collection/package.json:11-19`
  - `~/Development/qcash-ui-dashboard-auto-collection/README.md:19-41`

### `qcash-ui-forex`

- Next 16.2.4, Rspack and enhanced federation:
  - `~/Development/qcash-ui-forex/package.json:20-38`
  - `~/Development/qcash-ui-forex/next.config.ts:1-5`
- Explicitly shares React/React DOM as eager singletons:
  - `~/Development/qcash-ui-forex/next.config.ts:32-35`
- Consumes design-system type artifacts from a cluster URL:
  - `~/Development/qcash-ui-forex/next.config.ts:41-54`

These two repositories are an emerging replacement architecture, but they do not yet share one canonical implementation: one uses Rspack's `ModuleFederationPluginV1`, the other `@module-federation/enhanced/rspack`.

---

# 4. Systemic technical debt and coupling

## Critical: hard-coded development federation topology

At least 62 repositories embed `internal-service.example.com` in federation configuration. This makes independent environments and local composition fragile and turns a development cluster into a build-time dependency.

Representative evidence:

- `~/Development/qcash-ui-account-management/next.config.js:7-10`
- `~/Development/qcash-ui-workflow/federation.config.js:5-8`
- `~/Development/qcash-ui-dashboard-auto-collection/next.config.ts:35-38`
- `~/Development/qcash-ui-forex/next.config.ts:29-31`

`qcash-ui-global-account-summary` nominally supports environment variables but then unconditionally overwrites `ENV_DOMAIN`:

- `~/Development/qcash-ui-global-account-summary/next.config.js:4-7`

A runtime remote registry or shared federation preset should own remote names, paths and environment resolution.

## Critical: authentication and shell behavior duplicated inside feature MFEs

Feature-level Axios clients directly know:

- `access-token`, `refresh-token`, and `login` storage keys;
- `/homepage` and `/main-page`;
- the `"Another Login Detected"` backend string;
- the `showDialog` DOM event contract;
- retry behavior for status 401/408/499.

Representative evidence:

- `~/Development/qcash-ui-workflow/services/custom-axios.ts:45-68`
- `~/Development/qcash-ui-workflow/services/custom-axios.ts:107-160`
- `~/Development/qcash-ui-account-management/services/custom-axios.ts:28-48`
- `~/Development/qcash-ui-account-management/services/custom-axios.ts:181-232`

This couples each MFE to shell internals and preserves the obsolete `/main-page` login route even though the root architecture describes `/landing-page` as current.

The account-management client is 257 lines and includes authentication, retry queues, WAF HTML parsing, business-error normalization, routing and dialog rendering:

- `~/Development/qcash-ui-account-management/services/custom-axios.ts:51-162`
- `~/Development/qcash-ui-account-management/services/custom-axios.ts:164-255`

These responsibilities need a shared API/auth adapter with a typed host interface.

## High: dependency and singleton drift

There are 67 direct Ordentco auth-provider dependencies, ranging from `0.9.70` through `0.9.142-mfe`, with both exact and range constraints.

Examples:

- `~/Development/qcash-ui-beneficiary-account-management/package.json:23`
- `~/Development/qcash-ui-menu-management/package.json:20`
- `~/Development/qcash-ui-dashboard-auto-collection/package.json:23`

Federation commonly disables version enforcement:

- `~/Development/qcash-ui-account-management/next.config.js:23-40`
- `~/Development/qcash-ui-workflow/federation.config.js:20-42`

`requiredVersion: false` hides incompatibility instead of managing it. Shared libraries, especially auth and i18n contexts, need an estate-wide compatibility matrix and enforced versions.

## High: build checks are bypassed

At least nine repositories suppress lint and/or TypeScript failures during builds. The strongest example disables both:

- `~/Development/qcash-ui-user-basic-management/next.config.js:49-56`

Other examples include:

- `~/Development/qcash-ui-workflow/next.config.js:14-19`
- `~/Development/qcash-ui-account-management/next.config.js:49-55`

This allows incompatible remote interfaces and broken generated-type imports to reach deployment.

## High: federation exposes implementation-level page modules

Most MFEs expose route files directly. That couples hosts to:

- repository directory structure,
- Next router semantics,
- dynamic route parameter names,
- page-level providers,
- CSS side effects.

Examples:

- `~/Development/qcash-ui-account-management/next.config.js:12-20`
- `~/Development/qcash-ui-workflow/federation.config.js:10-15`

Stable adapters such as `src/exposes/account-list.tsx` would provide a safer public API.

## High: CSS and webpack ownership are unclear

Legacy projects add a second CSS/Sass loader chain to Next's webpack config:

- `~/Development/qcash-ui-global-account-summary/next.config.js:47-50`
- `~/Development/qcash-ui-menu-management/next.config.js:49-57`

They also expose generated global CSS. This creates ordering, duplication, hydration and class-collision risks. Styling prefixes exist per repository, but no documented host-level ordering or lifecycle contract was found.

## Medium: state and remote types cross boundaries

State management is inconsistent: most repositories use local React state, selected repositories use TanStack Query, and a few use Zustand.

Examples:

- Remote-generated UI types are embedded in local global state:
  - `~/Development/qcash-ui-scm/store/globalStore.ts:1-16`
- Persisted browser state directly invokes services:
  - `~/Development/qcash-ui-fund-transfer/store/useMasterDataStore.ts:11-37`

The latter persists environment-derived master data without a version, user, tenant or environment key, creating stale/cross-session data risk.

---

# 5. Testing audit

## Coverage of executable test contracts

- **34/71 repositories** currently define a `test` script.
- Four use Vitest:
  - `qcash-ui-dashboard-auto-collection`
  - `qcash-ui-helpdesk-portal`
  - `qcash-ui-plafon-pegadaian`
  - `qcash-ui-sistem-management`
- The remainder use Jest.
- No consistent estate-wide integration, contract, browser or host-composition test layer was found.

### Concrete debt

- `qcash-ui-menu-management` permits an empty suite:
  - `~/Development/qcash-ui-menu-management/package.json:12-13`
- It combines Jest 30 with `ts-jest` 29:
  - `~/Development/qcash-ui-menu-management/package.json:53-67`
- `qcash-ui-beneficiary-account-management` also combines Jest 30 with `ts-jest` 29:
  - `~/Development/qcash-ui-beneficiary-account-management/package.json:55-72`
- That repository has both `jest.config.js` and `jest.config.ts`, making configuration ownership ambiguous.
- Tests generally mock federation or generated remote types; therefore they validate local rendering but not host/remote compatibility.
- `qcash-ui-forex` has no test script despite introducing the new Rspack architecture:
  - `~/Development/qcash-ui-forex/package.json:5-18`

Recommended minimum contract per MFE:

1. unit tests,
2. exposed-module import smoke test,
3. generated/public type validation,
4. remoteEntry availability test,
5. host-mounted route test,
6. API/auth adapter contract test.

---

# 6. Documentation gaps

All repositories generally contain a README, but most are cloned boilerplate rather than operational architecture documentation.

Typical generic text only says “Next.js + Module Federation” and `yarn dev`; it omits remote names, exposed modules, base paths, environment variables, API ownership and host requirements.

Examples:

- `~/Development/qcash-ui-account-management/README.md:1-15`
- `~/Development/qcash-ui-sistem-management/README.md:1-40`

The latter still discusses nonexistent sample API routes and Vercel deployment.

`qcash-ui-forex` is especially stale:

- It describes `pages/index.tsx`, but source is under `src/pages`:
  - `~/Development/qcash-ui-forex/README.md:17-23`
- It documents four package managers while package scripts mix Bun and npm:
  - `~/Development/qcash-ui-forex/README.md:5-15`
  - `~/Development/qcash-ui-forex/package.json:6-18`

Missing from most READMEs:

- business owner and lifecycle status;
- shell route and `basePath`;
- federation remote name and expose map;
- required host-provided providers;
- supported auth-provider version;
- API/environment-variable matrix;
- local shell integration instructions;
- state persistence and storage keys;
- deployment port and health checks;
- test/coverage expectations;
- deprecation or replacement status.

---

# 7. Repository-specific high-risk findings

## `qcash-ui-user-basic-management`

Production builds explicitly ignore both lint and TypeScript errors:

- `~/Development/qcash-ui-user-basic-management/next.config.js:49-56`

Federated type generation is also commented out:

- `~/Development/qcash-ui-user-basic-management/next.config.js:85-87`

## `qcash-ui-workflow`

The development `FederatedTypesPlugin` remote URL ends at `/static/chunks/` rather than `remoteEntry.js` or a type artifact:

- `~/Development/qcash-ui-workflow/next.config.js:37-47`

It also disables version requirements across a very large shared dependency surface:

- `~/Development/qcash-ui-workflow/federation.config.js:20-42`

## `qcash-ui-global-account-summary`

It mutates `process.env.ENV_DOMAIN` to the dev cluster regardless of the earlier environment fallback:

- `~/Development/qcash-ui-global-account-summary/next.config.js:4-7`

It is coupled to both global components and the excluded main-page remote:

- `~/Development/qcash-ui-global-account-summary/next.config.js:12-18`

## `qcash-ui-regol`, `qcash-ui-rma-management`, `qcash-ui-dashboard-dhe`

Current checked-out configurations create remote entries but expose no modules:

- `~/Development/qcash-ui-regol/next.config.js:12-15`
- `~/Development/qcash-ui-rma-management/next.config.js:12-17`
- `~/Development/qcash-ui-dashboard-dhe/next.config.js:12-17`

These appear to be skeleton, dormant, or stale stable branches and should be explicitly classified.

## `qcash-ui-registration-giro`

The current branch also has no federation exposures:

- `~/Development/qcash-ui-registration-giro/next.config.js:12-17`

However, `origin/dev` is 30 commits ahead and changes approximately 885 architecture/test paths, including renaming `next.config.js` to `next.config.mjs`. The current branch is therefore not a reliable picture of active development.

## `qcash-ui-approval-signature`

`images.domains` contains full URLs rather than hostnames, which does not match Next's domain-list contract:

- `~/Development/qcash-ui-approval-signature/next.config.js:38-43`

## `qcash-ui-dashboard-auto-collection`

This is the most complete modern implementation, but federation still disables webpack caching:

- `~/Development/qcash-ui-dashboard-auto-collection/next.config.ts:12-14`

It mixes React 18 with `react-is` 19:

- `~/Development/qcash-ui-dashboard-auto-collection/package.json:27-34`

## `qcash-ui-forex`

Package-manager ownership is unclear: scripts invoke Bun, npm and bare Next, while both `yarn.lock` and `bun.lock` exist:

- `~/Development/qcash-ui-forex/package.json:5-18`

Remote type acquisition is tied directly to the development cluster:

- `~/Development/qcash-ui-forex/next.config.ts:41-49`

There are no tests or API/service implementation yet, so this is closer to an architectural scaffold than a production-complete feature.

---

# 8. `dev` branch observations

A local or remote `dev` ref was available for 46 repositories.

Important branch-state risks:

- Many repositories are checked out on `master` while `dev` contains major architecture and test additions.
- Particularly large dev-side changes exist in:
  - `qcash-ui-registration-giro`
  - `qcash-ui-digi-guarantee`
  - `qcash-ui-scm-cu`
  - `qcash-ui-rma-management`
  - `qcash-ui-payroll-transfer`
  - `qcash-ui-language-management`
  - `qcash-ui-role-management`
  - `qcash-ui-dashboard-dhe`
- Conversely, several `master` branches are hundreds of commits ahead of an obsolete `origin/dev`, including:
  - `qcash-ui-ewallet-topup`
  - `qcash-ui-global-internal-transfer`
- `qcash-ui-account-management` is on a release branch and has diverged 19 commits in each direction from local `dev`.

Therefore, root architecture documentation should record **lifecycle branch per repository**, not state globally that `dev` is always active and `master` is always stable.

---

# 9. Recommended root `ARCHITECTURE.md` additions

Add:

1. The complete domain grouping above.
2. An architecture-generation field per repository:
   - `legacy-nextjs-mf`
   - `modern-next-rspack-mf`
   - `skeleton/dormant`
3. A registry table with:
   - owner/domain,
   - active branch,
   - Next/federation generation,
   - base path,
   - remote name,
   - exposed modules,
   - consumed remotes,
   - auth-provider version,
   - state/data strategy,
   - test runner,
   - lifecycle status.
4. A formal host/MFE contract covering:
   - routing,
   - auth context,
   - storage keys,
   - session expiry,
   - dialogs/events,
   - i18n,
   - error shape,
   - process IDs,
   - CSS loading,
   - remote type publication.
5. A single environment-aware federation manifest instead of repository-local dev URLs.
6. A shared API client package instead of copied `custom-axios.ts`.
7. A migration plan from the 69 legacy MFEs to the modern Rspack pattern.
8. CI gates prohibiting:
   - ignored TypeScript/lint failures,
   - empty exposure maps for active MFEs,
   - `passWithNoTests`,
   - hard-coded cluster remotes,
   - incompatible Jest/transformer versions,
   - undocumented public exposures.

---

