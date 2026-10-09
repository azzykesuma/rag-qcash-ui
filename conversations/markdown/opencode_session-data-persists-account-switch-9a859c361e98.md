# Session data persists after account switch

- **ID**: `opencode-f187d538634b48f4a99f0d13c19f95ebed84e044232e21779c7f9a859c361e98`
- **Source Tool**: `opencode`
- **Date**: `2026-09-28 09:22:50`
- **Project**: `qcash-ui`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

i have an issue with the session management on this repo. basically, when i logged in as ba, and logged out , then login again as cu,the header footer information is still listed as ba, along with their menu. can you investigate this issue?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **read**: Tool: read
- **grep**: Tool: grep
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>

I’ll trace login, logout, persisted session state, and header/footer/menu data flow, then identify where the prior BA identity survives the CU login. I’ll keep this investigation read-only and return a concrete fix and verification plan.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **task**: Tool: task
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read
- **grep**: Tool: grep

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash

</details>

The current branch already contains comments/workarounds aimed at this exact stale-user scenario, so this is not just a generic storage leak. The key risk is that the host and federated header/footer can hold separate auth-provider instances, while the provider’s five-minute validation cache is not scoped to the active token.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **bash**: Tool: bash
- **context7_resolve-library-id**: Tool: context7_resolve-library-id

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **context7_query-docs**: Tool: context7_query-docs
- **bash**: Tool: bash
- **atlassian-bitbucket-dc_bitbucket_getRepository**: Tool: atlassian-bitbucket-dc_bitbucket_getRepository
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode

</details>

The current header/footer repository confirms the failing path: its logout goes to `/main-page` without `?logout=true`, while this host only hard-resets singleton state for the query-param flow and only clears cross-user caches on `/landing-page`. The remote does clear menu storage, but it does not notify `HostAuthGate`, and the provider itself leaves several identity/global-store fields alive.

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_getCommits**: Tool: atlassian-bitbucket-dc_bitbucket_getCommits
- **grep**: Tool: grep
- **bash**: Tool: bash

</details>

A second, more direct inconsistency is present in the header/footer’s current indexed source: its layout, profile dropdown, and menu hook still import `@ordentco/addons-auth-provider`, while this host mounts only the BRI provider. That means the visible username/user type/menu are not guaranteed to come from the host session state at all; the remote’s federation config mentioning BRI does not fix those legacy imports.

---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **atlassian-bitbucket-dc_bitbucket_getPullRequests**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequests
- **atlassian-bitbucket-dc_bitbucket_getPullRequests**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequests
- **atlassian-bitbucket-dc_bitbucket_getPullRequests**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequests
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **task**: Tool: task

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **atlassian-bitbucket-dc_bitbucket_getPullRequest**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequest
- **atlassian-bitbucket-dc_bitbucket_getPullRequestChanges**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestChanges
- **atlassian-bitbucket-dc_bitbucket_getPullRequest**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequest
- **atlassian-bitbucket-dc_bitbucket_getPullRequestChanges**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestChanges
- **atlassian-bitbucket-dc_bitbucket_getPullRequest**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequest
- **atlassian-bitbucket-dc_bitbucket_getPullRequestChanges**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestChanges

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **atlassian-bitbucket-dc_bitbucket_getPullRequestDiff**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestDiff
- **atlassian-bitbucket-dc_bitbucket_getPullRequestDiff**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestDiff
- **atlassian-bitbucket-dc_bitbucket_getPullRequestDiff**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestDiff
- **atlassian-bitbucket-dc_bitbucket_getPullRequestDiff**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestDiff

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **atlassian-bitbucket-dc_bitbucket_getPullRequestDiff**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestDiff
- **atlassian-bitbucket-dc_bitbucket_getPullRequestDiff**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestDiff
- **atlassian-bitbucket-dc_bitbucket_getPullRequestDiff**: Tool: atlassian-bitbucket-dc_bitbucket_getPullRequestDiff

</details>



---

## Turn 14: Assistant

**Investigation Result**
The issue is a cross-user state leak involving the host, auth provider, and federated header/footer.

Primary causes:

1. `@bri/addons-auth-provider` is not configured as a singleton in `next.config.js:16`, despite the architecture requiring one shared React context across host and remotes.
2. The header/footer logout does not emit the `qc-bridge-sync` logout event expected by `HostAuthGate.tsx:131`.
3. The normal header/footer logout redirects to `/main-page` without `?logout=true`. Therefore, the reload workaround in `pages/_app.tsx:94` does not run.
4. Cache cleanup in `pages/_app.tsx:101` only applies to `/landing-page`, not `/main-page`.
5. The installed auth provider’s five-minute freshness cache is not associated with an access token. `auth.js:319` can therefore consider BA data fresh after the token changes to CU.
6. Provider logout clears menus and authorities, but does not reset username, user type, company information, product authorities, readiness, or the global store.
7. The BRI global store does not clear products/system configuration when the token is removed.
8. `HostAuthGate.tsx:244` can publish the auth object captured before `guard()` completed.

This sequence explains the symptom:

```text
BA authenticated
→ BA identity/menu cached in provider and global store
→ logout clears token and some storage, but not all in-memory state
→ CU token stored
→ guard freshness check exits early or a separate provider instance is used
→ host marks CU token validated
→ header/footer still renders BA identity and menu
```

The `qcash-ui-header-footer` prerelease branch was migrated to BRI auth, but it recently reverted from `^0.10.7-mfe` to `^0.10.6-mfe` in PR 701. The host currently uses `0.10.7-mfe`, making the missing singleton contract especially risky.

**Implementation Plan**
1. Restore `singleton: true` for `@bri/addons-auth-provider` in `next.config.js`.
2. Add a single host session-reset mechanism that:
   - Clears `session-user-data`, `productMenu`, `productRoles`, `validateMenu`, and menu ownership/cache keys.
   - Resets `window.__QCASH_HOST_AUTH_GUARD__`.
   - Replaces `window.__QCASH_AUTH_BRIDGE__` with a guest/not-ready state.
   - Remounts the BRI auth and global-store providers to discard all BA React state.
3. Trigger that reset from every logout path using `qc-bridge-sync` with `source: "logout"`.
4. Update header/footer logout to emit that event and use the same logout contract.
5. Apply fallback cleanup on both `/landing-page` and `/main-page` when no token exists.
6. Change bridge publication so it uses post-guard state rather than the pre-guard auth snapshot.
7. Raise an auth-provider fix so its freshness cache is scoped to the active token and logout resets every auth/global-store field.
8. Add a regression test for:
   `BA login → logout without closing tab → CU login → CU header/profile/menu only`.

No files were changed and no tests were run because this was a read-only investigation. The findings assume the reported issue is on the `prerelease` deployment; the exact environment should be confirmed before implementation.

---

## Turn 15: User

can you check which data is being used to populate the header footer information? what is the source of truth for that component?

---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **grep**: Tool: grep
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode

</details>

I’ll separate the displayed identity fields from the navigation menu, because the header/footer remote obtains them through different paths. I’ll trace each field back through the provider, API calls, global store, and browser-cache fallbacks.

---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (7)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (5)</summary>

- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode
- **atlassian-bitbucket-dc_bitbucket_searchCode**: Tool: atlassian-bitbucket-dc_bitbucket_searchCode

</details>



---

## Turn 19: Assistant

**Source Of Truth**
For the migrated `prerelease` header/footer, the immediate source of truth for user information is the React context returned by:

```ts
useAuth() from "@bri/addons-auth-provider"
```

The backend source behind that context is primarily `GET /auth/me`.

**Data Flow**
```text
localStorage["access-token"]
    ↓
BriProviders
    ↓
HostAuthGate calls briAuth.guard()
    ↓
GET /auth/me
    ↓
BRI AuthProvider React state
    ↓
header/footer useAuth()
```

Relevant host locations:

- Provider mounted in `components/providers/BRIProviders.tsx:19`
- Host reads the token in `components/providers/HostAuthGate.tsx:180`
- Host calls `guard()` in `components/providers/HostAuthGate.tsx:240`
- Provider calls `/auth/me` in `node_modules/@bri/addons-auth-provider/dist/services/auth.js:79`
- Provider populates username/company/user type in `node_modules/@bri/addons-auth-provider/dist/src/auth.js:395`
- Header/footer is loaded as a remote from `constants/features/registry.ts:44`

**Header Fields**
The header/footer layout reads these values directly from `useAuth()`:

```ts
const {
  username,
  companyName,
  companyCode,
  userType,
  isAuthoritiesReady,
} = useBriAuth();
```

They are passed into the header:

```tsx
<Header
  username={username}
  companyName={companyName}
  companyCode={companyCode}
  loading={!isAuthoritiesReady}
/>
```

The profile dropdown separately reads:

```ts
const { logout, username, userType, region } = useAuth();
```

Therefore:

| Displayed information | Runtime source | Backend source |
|---|---|---|
| Username | BRI auth React context | `/auth/me` |
| Company name | BRI auth React context | `/auth/me` |
| Company code | BRI auth React context | `/auth/me` |
| User type, BA/CU | BRI auth React context | `/auth/me` |
| Region | BRI auth React context | `/auth/me` |
| Authorities and roles | BRI auth React context | `/auth/me` |
| Auth-provider menus | BRI auth React context | `/menu/me` |

`window.__QCASH_AUTH_BRIDGE__` is diagnostic/compatibility data published by `HostAuthGate`; the migrated prerelease header/footer no longer uses it as its primary source.

**Navigation Menu**
The visible navigation menu has several sources rather than one strict source of truth.

The migrated `useMenu()` combines:

1. `userType`, `userID`, and token from BRI `useAuth()`.
2. Products from BRI `useGlobalStore()`, originally loaded from `GET /product`.
3. Menu licenses/tree from `POST /menu/me`.
4. Product roles from `GET /auth/me`.
5. An additional product list from `GET /product`.
6. Three-minute `ahooks` request caches.
7. Browser fallbacks:
   - `localStorage.productMenu`
   - `localStorage.validateMenu`
   - `localStorage.productRoles`
   - `localStorage.menuCachedForUserType`

The effective product-menu priority is:

```text
BRI global-store products
→ freshly fetched products
→ localStorage productMenu/productRoles
```

The product list and `/menu/me` result are then combined by `transformMenu()` to produce the rendered navigation.

**Footer**
I found no `useAuth()` or `useGlobalStore()` usage under the footer implementation. It appears to be static/configuration/localization-driven. The stale session symptom is therefore primarily the header/profile and navigation menu, not footer-specific user state.

**Important Finding**
There are two different “truths” depending on the deployed header/footer branch:

- `prerelease`, after header/footer PR 634: directly uses `@bri/addons-auth-provider`.
- Current indexed `master`: still contains `@ordentco/addons-auth-provider` imports.

Consequently, the actual runtime source of truth depends on which `qcash-ui-header-footer` artifact your environment loads. For the migrated prerelease build, it should be the host’s BRI context. If Module Federation resolves a separate provider instance, that separate in-memory context becomes the practical source, which explains why BA data can survive after the host has a CU token.

For this bug specifically:

- Stale header identity most likely comes from old BRI provider React state.
- Stale navigation can additionally come from BRI global-store products, `ahooks` cache, or local-storage menu fallbacks.

---

