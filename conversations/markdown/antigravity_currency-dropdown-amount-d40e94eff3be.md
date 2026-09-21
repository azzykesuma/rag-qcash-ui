# Rewrite `~/Development/qcash-ui-issuance-lc/components/u...

- **ID**: `antigravity-f01730ea5e0be789a30e5b9fc84acffd79246489bb4863312d95d40e94eff3be`
- **Source Tool**: `antigravity`
- **Date**: `2026-08-06 08:33:07`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

Rewrite `~/Development/qcash-ui-issuance-lc/components/ui/qui-currency-amount.tsx` to use `MoneyInput` from `@/components/global` instead of `NumberField` and `QuiLabel` from QUI.

The global `MoneyInput` component (from `qcash-global-component/money-input`) wraps number inputs with currency selection. Check its type definition at `~/Development/qcash-ui-issuance-lc/@mf-types/qcash-global-component/_types/components/form/money-input/money-input.d.ts`.

The current `QuiCurrencyAmount` component provides:
- Amount input with locale-aware number formatting
- Currency dropdown (prefix autocomplete)
- Props: `label`, `amount`, `currency`, `currencyOptions`, `variant`, `name`, `currencyName`, `placeholder`, `disabled`, `readOnly`, `amountReadOnly`, `required`, `minValue`, `maxValue`, `step`, `maximumIntegerDigits`, `maximumFractionDigits`, `locale`, `onAmountChange`, `onCurrencyChange`, `isInvalid`

If `MoneyInput` from global provides similar functionality, use it. If not, implement a plain HTML + React approach:

1. Check the MoneyInput type definition file first
2. If MoneyInput API is compatible, use it
3. If not, implement a plain approach:
   - Use `Select` from `@/components/global` for currency dropdown
   - Use `Input` from `@/components/global` for amount input
   - Handle number formatting manually

The key is: keep the same external API (`QuiCurrencyAmountProps`) so callers don't break.

Also preserve these exported utilities:
- `resolveNumberFieldLocale(language?: string): QuiCurrencyAmountLocale`
- `export type QuiCurrencyAmountOption`
- `export type QuiCurrencyAmountLocale`

Read the full current file first, read the MoneyInput type definition, then rewrite the component.

Write the updated file to `~/Development/qcash-ui-issuance-lc/components/ui/qui-currency-amount.tsx`.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **view_file**: 
- **view_file**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **run_command**: 
- **run_command**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 8: Assistant

I have rewritten `~/Development/qcash-ui-issuance-lc/components/ui/qui-currency-amount.tsx` according to your instructions.

Since the global `MoneyInput` API does not natively support a fully customizable currency dropdown (like `currencyOptions`), I used the fallback plain approach:
- Replaced `NumberField` and `QuiLabel` from QUI with `Input` and `Select` from `@/components/global`.
- Integrated `Select` for the currency dropdown with `Input` for the amount input into a layout side-by-side.
- Manually handled parsing and formatting number inputs using the specified locale, including clamping maximum integer and fraction digits, checking valid inputs per locale separators, and formatting upon blur. 
- Maintained all exported types (`QuiCurrencyAmountOption`, `QuiCurrencyAmountLocale`) and functions (`resolveNumberFieldLocale`), and completely preserved the previous `QuiCurrencyAmountProps` interface so all callers continue to work seamlessly.

---

