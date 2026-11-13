# "qcash-ui-account-receivable": {

- **ID**: `antigravity-b79920edcb198032f2ec9cf62c260a25c15ff81068c50e8248033a2fe29ce53f`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-11 11:50:47`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

"qcash-ui-account-receivable": {
    toggleKey: "toggleAccountReceivable",
    routes: ["/account-receivable"],
    i18n: ["account-receivable"],
    config: {
      independentLocale: true,
    },
  }, can you investigate what is the usage of independent locale?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 17: Assistant

In QCash UI's Micro-Frontend (MFE) architecture, **`independentLocale`** controls the **priority order for loading i18n translation bundles** between the remote module (or translation API) and the host shell application.

---

### Definition & Code Implementation

Defined in [`FeatureOptions`](file:///C:/Work/qcash-ui/services/federation/types.ts#L26-L34):

```typescript
export type FeatureOptions = {
  /**
   * When true, locale files are fetched exclusively from the feature's own
   * remote origin (and the centralised language API in config.languageApiUrl
   * when available), bypassing the shell's /public/locales entirely.
   * Use this for modules that own their translation files independently.
   */
  independentLocale?: boolean;
};
```

Evaluated during locale resolution in [`getLoadPathCandidates`](file:///C:/Work/qcash-ui/services/federation/utils/i18n.ts#L41-L60):

```typescript
function getLoadPathCandidates(lng: string, ns: string, sources: ModuleLocaleSources): string[] {
  const { independentLocale, languageApiUrl, modulePath, moduleOrigin, shellOrigin } = sources;
  const localePath = `${modulePath}/locales/${lng}/${ns}.json`;
  const primary = getLoadPath(lng, ns);

  const remoteCandidates = [
    ...(languageApiUrl ? [joinUrl(languageApiUrl, `?lang=${lng}&module=${ns}`)] : []),
    ...(moduleOrigin && modulePath ? [joinUrl(moduleOrigin, localePath)] : []),
  ];

  const hostCandidates = [
    ...(primary ? [primary, ...(shellOrigin && primary.startsWith("/") ? [joinUrl(shellOrigin, primary)] : [])] : []),
  ];

  const ordered = independentLocale
    ? [...remoteCandidates, ...hostCandidates]
    : [...hostCandidates, ...remoteCandidates];

  return [...new Set(ordered)];
}
```

---

### How It Works

When a federated module is loaded (e.g. by [`useModuleLoading`](file:///C:/Work/qcash-ui/components/federation/hooks/useModuleLoading.ts#L41) calling [`dynamicLoadLocales`](file:///C:/Work/qcash-ui/services/federation/utils/i18n.ts#L155)), the system attempts to fetch translation JSONs for each namespace in order:

| Mode | Candidate Fetch Order | Behavior |
| :--- | :--- | :--- |
| **`independentLocale: true`** | **1.** Language API (`/language/v1/translations`)<br>**2.** Remote MFE origin (`${moduleOrigin}/${modulePath}/locales/${lng}/${ns}.json`)<br>**3.** Host shell (`/v2/locales/${lng}/${ns}.json`) | **Remote-First**: Prioritizes the module's own remote translations or the centralized API. Host shell locales are only used as fallback. |
| **`independentLocale: false`** *(or omitted)* | **1.** Host shell (`/v2/locales/${lng}/${ns}.json`)<br>**2.** Language API<br>**3.** Remote MFE origin | **Host-First**: Prioritizes translation files bundled inside the host shell. |

---

### Why It Is Used for `qcash-ui-account-receivable`

1. **Independent Module Lifecycle**: It allows the `qcash-ui-account-receivable` repository and team to manage, update, and deploy its own translation files independently with its own release cycle.
2. **Prevents Stale Shell Locales**: If the host shell happens to have an older copy of `account-receivable.json`, setting `independentLocale: true` ensures the browser pulls the latest translations from the remote module CDN or Language API rather than using the stale host version.

---

