# Explore project structure (@explore subagent)

- **Author**: `azzy`
- **ID**: `ses_0ba5a6990ffe66SuhFy7lgc6Xc`
- **Source Tool**: `opencode`
- **Date**: `2026-07-09 13:51:47`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Explore the project at D:\code\mobile-fitflow. I need to understand:
1. The full project structure (top-level files and directories)
2. The main entry point (App.tsx, app/ directory, or similar)
3. How the splash screen and app initialization work
4. The package.json contents
5. The app.json or app.config.js contents
6. Any error boundary or root component setup

Return ALL file paths and key code snippets related to app initialization, splash screen handling, and the root component.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>

I'll start by exploring the top-level structure and key configuration files in parallel.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (6)</summary>

- **read**: Tool: read
- **glob**: Tool: glob
- **grep**: Tool: grep
- **grep**: Tool: grep
- **grep**: Tool: grep
- **read**: Tool: read

</details>

Now let me explore the `src/` directory, the `app/` directory (expo-router convention), and search for splash screen and error boundary references.

---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (8)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>

The `app/` directory is inside `src/`. Let me explore it fully and search for initialization files.

---

## Turn 5: Assistant

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

Excellent! Now let me read all the key initialization files: root layout, auth context, and other important setup files.

---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

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
- **read**: Tool: read
- **read**: Tool: read

</details>

Now let me read the remaining initialization-related files: auth token init, supabase, global CSS, theme, screen wrapper, and other config files.

---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (12)</summary>

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
- **read**: Tool: read
- **read**: Tool: read

</details>

Let me also grab the remaining layout files, the `.env`, the scripts, and the API client for completeness.

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **read**: Tool: read
- **read**: Tool: read
- **read**: Tool: read

</details>

Let me grab the assets structure and the reset-project script for completeness.

---

## Turn 9: Assistant

Here is the complete analysis of the project at `D:\code\mobile-fitflow`.

---

## 1. FULL PROJECT STRUCTURE (Top-Level)

```
D:\code\mobile-fitflow/
├── .aipulse/                  # AI pulse config directory
├── .claude/                   # Claude AI config directory
├── .env                       # Environment variables (Supabase, API, encryption keys)
├── .expo/                     # Expo build/cache artifacts
├── .git/
├── .gitignore
├── .playwright-mcp/
├── .vscode/
├── AGENTS.md                  # Agent instructions file
├── android/                   # Android native project
├── app.json                   # Expo app configuration
├── assets/                    # Static assets
│   ├── expo.icon/
│   └── images/                # App icons, splash, logos, tutorial images
├── babel.config.js            # Babel config with NativeWind
├── CLAUDE.md
├── dist/                      # Build output
├── eas.json                   # EAS Build config
├── expo-env.d.ts              # Expo type references
├── LICENSE
├── metro.config.js            # Metro bundler with NativeWind
├── nativewind-env.d.ts        # NativeWind types
├── node_modules/
├── package-lock.json
├── package.json
├── README.md
├── scripts/
│   └── reset-project.js       # Utility to reset app to blank state
├── src/                       # ** MAIN SOURCE CODE **
│   ├── app/                   # Expo Router file-based routes
│   ├── components/            # Reusable UI components
│   ├── constants/             # Theme constants
│   ├── features/              # Feature modules (auth, workouts, meals, progress, habits)
│   ├── global.css             # Tailwind base + custom font variables
│   ├── hooks/                 # Custom hooks (theme, color scheme)
│   └── lib/                   # Core libraries (auth, API client, encryption, supabase)
├── tailwind.config.js
└── tsconfig.json              # TypeScript config with @/ path alias to ./src
```

---

## 2. MAIN ENTRY POINT

**File:** `D:\code\mobile-fitflow\package.json` (line 3)

```json
"main": "expo-router/entry"
```

There is **no** traditional `App.tsx` or `App.js`. The project uses **Expo Router** (file-based routing), so the entry point is `expo-router/entry`, which automatically loads routes from `src/app/`. The root layout is the first file loaded:

**File:** `D:\code\mobile-fitflow\src\app\_layout.tsx`

```typescript
import { AuthProvider } from "@/lib/auth-context";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { DarkTheme, DefaultTheme, ThemeProvider, Stack } from "expo-router";
import * as SplashScreen from "expo-splash-screen";
import { useState } from "react";
import { useColorScheme } from "react-native";
import "../global.css";

// Prevent splash screen from auto-hiding
SplashScreen.preventAutoHideAsync().catch(() => {});

export default function RootLayout() {
  const colorScheme = useColorScheme();
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            retry: 1,
            staleTime: 30_000,
          },
        },
      })
  );

  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <ThemeProvider value={colorScheme === "dark" ? DarkTheme : DefaultTheme}>
          <Stack screenOptions={{ headerShown: false }}>
            <Stack.Screen name="index" />
            <Stack.Screen name="(auth)/login" />
            <Stack.Screen name="(auth)/register" />
            <Stack.Screen name="(app)" />
          </Stack>
        </ThemeProvider>
      </AuthProvider>
    </QueryClientProvider>
  );
}
```

**Key observations about the root component:**
- It serves as the **RootLayout** (the only root-level component).
- It wraps everything in three providers: `QueryClientProvider` (React Query) > `AuthProvider` > `ThemeProvider`.
- It defines four route groups in a `Stack` navigator: `index` (landing), `(auth)/login`, `(auth)/register`, and `(app)` (authenticated tab area).
- There is **no explicit Error Boundary** anywhere in the project.

---

## 3. SPLASH SCREEN & APP INITIALIZATION

### 3a. Static Configuration (`app.json`)

**File:** `D:\code\mobile-fitflow\app.json`

```json
{
  "expo": {
    "plugins": [
      "expo-router",
      [
        "expo-splash-screen",
        {
          "backgroundColor": "#208AEF",
          "image": "./assets/images/splash-icon.png",
          "imageWidth": 76
        }
      ]
    ]
  }
}
```

The splash screen is configured via the `expo-splash-screen` plugin with:
- Background color: `#208AEF` (blue)
- Image: `./assets/images/splash-icon.png`
- Image width: `76` dp

### 3b. Runtime Splash Screen Logic (`_layout.tsx`)

**File:** `D:\code\mobile-fitflow\src\app\_layout.tsx` (line 10)

```typescript
// Prevent splash screen from auto-hiding
SplashScreen.preventAutoHideAsync().catch(() => {});
```

This is called at module import time (outside the component), which means it runs **before any component renders**. However, note that `SplashScreen.hideAsync()` is **never explicitly called** anywhere in the codebase -- the splash screen will be auto-hidden by Expo Router once the first route is fully rendered (Expo Router manages this automatically when using `preventAutoHideAsync`).

### 3c. Auth Initialization Flow

**File:** `D:\code\mobile-fitflow\src\lib\auth-context.tsx`

```typescript
useEffect(() => {
  const initializeAuth = async () => {
    await initAuthTokens();
    await checkUser();
  };
  initializeAuth();
}, []);
```

The `AuthProvider` performs these steps on mount:
1. **`initAuthTokens()`** (from `src/lib/auth-token.ts`) -- Loads access/refresh tokens from `expo-secure-store` (native) or `localStorage` (web) into memory:
   ```typescript
   export async function initAuthTokens() {
     if (isInitialized) return;
     if (Platform.OS === 'web') {
       if (typeof window !== 'undefined') {
         memoryAccessToken = window.localStorage.getItem(AUTH_TOKEN_KEY) ?? undefined;
         memoryRefreshToken = window.localStorage.getItem(REFRESH_TOKEN_KEY) ?? undefined;
       }
     } else {
       try {
         const access = await SecureStore.getItemAsync(AUTH_TOKEN_KEY);
         const refresh = await SecureStore.getItemAsync(REFRESH_TOKEN_KEY);
         memoryAccessToken = access ?? undefined;
         memoryRefreshToken = refresh ?? undefined;
       } catch (e) {
         console.warn("Failed to load secure tokens", e);
       }
     }
     isInitialized = true;
   }
   ```
2. **`checkUser()`** -- If a token exists, calls `getCurrentUser()` to verify the session and set the user state. Otherwise marks as unauthenticated.
3. While loading, `isLoading` is `true` and the landing page (`src/app/index.tsx`) shows a **spinning loader**:
   ```tsx
   if (isLoading) {
     return (
       <View className="flex-1 items-center justify-center bg-slate-950">
         <View className="size-8 animate-spin rounded-full border-2 border-lime-300 border-t-transparent" />
       </View>
     );
   }
   ```
4. After loading, if the user **is authenticated**, the landing page immediately redirects to `/(app)/dashboard`.

### 3d. App (Authenticated) Layout Guard

**File:** `D:\code\mobile-fitflow\src\app\(app)\_layout.tsx`

```typescript
useEffect(() => {
  if (!isLoading && !isAuthenticated) {
    router.replace("/");
  }
}, [isAuthenticated, isLoading]);
```

The `(app)` group layout also checks auth status. If not authenticated, it redirects back to the landing page. While loading, it shows an `ActivityIndicator`.

---

## 4. PACKAGE.JSON CONTENTS

**File:** `D:\code\mobile-fitflow\package.json` (52 lines)

```json
{
  "name": "mobile-fitflow",
  "main": "expo-router/entry",
  "version": "1.0.0",
  "dependencies": {
    "@expo/ui": "~57.0.3",
    "@expo/vector-icons": "^15.1.1",
    "@supabase/supabase-js": "^2.110.0",
    "@tanstack/react-query": "^5.101.2",
    "expo": "~57.0.2",
    "expo-constants": "~57.0.3",
    "expo-device": "~57.0.0",
    "expo-font": "~57.0.0",
    "expo-glass-effect": "~57.0.0",
    "expo-image": "~57.0.0",
    "expo-image-picker": "^57.0.2",
    "expo-linking": "~57.0.1",
    "expo-router": "~57.0.3",
    "expo-secure-store": "~57.0.0",
    "expo-splash-screen": "~57.0.2",
    "expo-status-bar": "~57.0.0",
    "expo-symbols": "~57.0.0",
    "expo-system-ui": "~57.0.0",
    "expo-web-browser": "~57.0.0",
    "formik": "^2.4.9",
    "nativewind": "^4.2.6",
    "node-forge": "^1.4.0",
    "react": "19.2.3",
    "react-dom": "19.2.3",
    "react-native": "0.86.0",
    "react-native-gesture-handler": "~2.32.0",
    "react-native-reanimated": "4.5.0",
    "react-native-safe-area-context": "~5.7.0",
    "react-native-screens": "4.25.2",
    "react-native-web": "~0.21.0",
    "react-native-worklets": "0.10.0"
  },
  "devDependencies": {
    "@types/react": "~19.2.2",
    "tailwindcss": "^3.3.2",
    "typescript": "~6.0.3"
  },
  "scripts": {
    "start": "expo start --port 8085",
    "reset-project": "node ./scripts/reset-project.js",
    "android": "expo run:android",
    "ios": "expo run:ios",
    "web": "expo start --web",
    "lint": "expo lint"
  },
  "private": true
}
```

Notable: Uses **Expo SDK 57**, **React 19.2.3**, **React Native 0.86.0**, **NativeWind 4.2.6**, **TanStack React Query**, **Formik**, and **node-forge** for RSA password encryption.

---

## 5. APP.JSON / APP CONFIG CONTENTS

**File:** `D:\code\mobile-fitflow\app.json` (49 lines)

```json
{
  "expo": {
    "name": "mobile-fitflow",
    "slug": "mobile-fitflow",
    "version": "1.0.0",
    "orientation": "portrait",
    "icon": "./assets/images/icon.png",
    "scheme": "mobilefitflow",
    "userInterfaceStyle": "automatic",
    "ios": {
      "icon": "./assets/expo.icon"
    },
    "android": {
      "adaptiveIcon": {
        "backgroundColor": "#E6F4FE",
        "foregroundImage": "./assets/images/android-icon-foreground.png",
        "backgroundImage": "./assets/images/android-icon-background.png",
        "monochromeImage": "./assets/images/android-icon-monochrome.png"
      },
      "predictiveBackGestureEnabled": false,
      "package": "com.anonymous.mobilefitflow"
    },
    "web": {
      "output": "single",
      "favicon": "./assets/images/favicon.png"
    },
    "plugins": [
      "expo-router",
      [
        "expo-splash-screen",
        {
          "backgroundColor": "#208AEF",
          "image": "./assets/images/splash-icon.png",
          "imageWidth": 76
        }
      ]
    ],
    "experiments": {
      "typedRoutes": true,
      "reactCompiler": true
    },
    "extra": {
      "router": {},
      "eas": {
        "projectId": "63c1d5d7-dd45-4071-9b57-5239cd4c5ee9"
      }
    }
  }
}
```

No `app.config.js` file exists -- only `app.json`.

---

## 6. ERROR BOUNDARY & ROOT COMPONENT SETUP

### Error Boundary

There is **no custom Error Boundary** implemented in this project. The app relies on:
- **React Native's built-in error handling** (redbox in dev, crash in production).
- **Expo's default error overlay** during development.
- The `try/catch` blocks scattered throughout the code (e.g., in `auth-context.tsx` `checkUser`, in `api-client.ts` for network/refresh errors).

### Root Component Tree (visual hierarchy)

```
RootLayout (src/app/_layout.tsx)
  ├── QueryClientProvider (React Query)
  │   └── AuthProvider (src/lib/auth-context.tsx)
  │       │   ┌─ initAuthTokens()  → loads tokens from SecureStore/localStorage
  │       │   └─ checkUser()       → validates token against /api/auth/me
  │       └── ThemeProvider (expo-router: DarkTheme | DefaultTheme)
  │           └── Stack Navigator (expo-router)
  │               ├── index                          → LandingPage (src/app/index.tsx)
  │               ├── (auth)/login                   → LoginScreen (src/app/(auth)/login.tsx)
  │               ├── (auth)/register                → RegisterScreen (src/app/(auth)/register.tsx)
  │               └── (app)                          → AppLayout (src/app/(app)/_layout.tsx)
  │                   │  (guarded by isAuthenticated check)
  │                   └── Tabs Navigator
  │                       ├── dashboard
  │                       ├── meals (nested Stack)
  │                       ├── workouts (nested Stack)
  │                       ├── progress (nested Stack)
  │                       └── settings
```

### Initialization Sequence

1. **Module load**: `_layout.tsx` is parsed -> `global.css` is imported (Tailwind/NativeWind) -> `SplashScreen.preventAutoHideAsync()` is called.
2. **RootLayout renders**: Creates `QueryClient`, wraps app in providers.
3. **AuthProvider mounts**: Calls `initAuthTokens()` (loads tokens from storage), then `checkUser()` (validates session with backend). `isLoading` is `true` during this process.
4. **Landing page renders** (`index.tsx`): Shows a spinner while `isLoading` is true. Once loading finishes:
   - If authenticated: redirects to `/(app)/dashboard`
   - If not authenticated: shows the landing page with "Get Started" / "Log In" buttons
5. **AppLayout renders** (when `/(app)` route is accessed): Checks auth again; redirects to `/` if not authenticated.
6. **Splash screen**: Auto-hidden by Expo Router once the first route screen is committed (no explicit `SplashScreen.hideAsync()` call in user code).

---

## COMPLETE FILE PATH INDEX

### Configuration Files
- `D:\code\mobile-fitflow\package.json`
- `D:\code\mobile-fitflow\app.json`
- `D:\code\mobile-fitflow\tsconfig.json`
- `D:\code\mobile-fitflow\babel.config.js`
- `D:\code\mobile-fitflow\metro.config.js`
- `D:\code\mobile-fitflow\tailwind.config.js`
- `D:\code\mobile-fitflow\eas.json`
- `D:\code\mobile-fitflow\.env`

### Entry Points & Layouts
- `D:\code\mobile-fitflow\src\app\_layout.tsx` -- **Root layout (splash, providers, stack)**
- `D:\code\mobile-fitflow\src\app\index.tsx` -- **Landing page (auth check + redirect)**
- `D:\code\mobile-fitflow\src\app\(app)\_layout.tsx` -- **Authenticated tab layout**
- `D:\code\mobile-fitflow\src\app\(app)\workouts\_layout.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\progress\_layout.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\meals\_layout.tsx`

### Auth & Initialization Core
- `D:\code\mobile-fitflow\src\lib\auth-context.tsx` -- **AuthProvider + useAuth hook**
- `D:\code\mobile-fitflow\src\lib\auth-token.ts` -- **Token storage/retrieval (SecureStore/localStorage)**
- `D:\code\mobile-fitflow\src\lib\api-client.ts` -- **HTTP client with auto-refresh**
- `D:\code\mobile-fitflow\src\lib\supabase.ts` -- **Supabase client**
- `D:\code\mobile-fitflow\src\lib\password-encryption.ts` -- **RSA-OAEP password encryption**
- `D:\code\mobile-fitflow\src\features\auth\api.ts`
- `D:\code\mobile-fitflow\src\features\auth\hooks.ts`
- `D:\code\mobile-fitflow\src\features\auth\types.ts`
- `D:\code\mobile-fitflow\src\features\auth\validation.ts`

### Screens
- `D:\code\mobile-fitflow\src\app\(auth)\login.tsx`
- `D:\code\mobile-fitflow\src\app\(auth)\register.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\dashboard.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\settings.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\workouts\index.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\workouts\new.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\workouts\[id].tsx`
- `D:\code\mobile-fitflow\src\app\(app)\workouts\session\[id].tsx`
- `D:\code\mobile-fitflow\src\app\(app)\meals\index.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\meals\list.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\meals\search.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\meals\[id].tsx`
- `D:\code\mobile-fitflow\src\app\(app)\progress\index.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\progress\body-measurement.tsx`
- `D:\code\mobile-fitflow\src\app\(app)\progress\photos.tsx`
- `D:\code\mobile-fitflow\src\app\explore.tsx`

### Components
- `D:\code\mobile-fitflow\src\components\layout\screen-wrapper.tsx`
- `D:\code\mobile-fitflow\src\components\ui\button.tsx`
- `D:\code\mobile-fitflow\src\components\ui\field.tsx`
- `D:\code\mobile-fitflow\src\components\ui\collapsible.tsx`
- `D:\code\mobile-fitflow\src\components\app-tabs.tsx`
- `D:\code\mobile-fitflow\src\components\app-tabs.web.tsx`
- `D:\code\mobile-fitflow\src\components\themed-text.tsx`
- `D:\code\mobile-fitflow\src\components\themed-view.tsx`
- `D:\code\mobile-fitflow\src\components\animated-icon.tsx`
- `D:\code\mobile-fitflow\src\components\animated-icon.web.tsx`
- `D:\code\mobile-fitflow\src\components\external-link.tsx`
- `D:\code\mobile-fitflow\src\components\hint-row.tsx`
- `D:\code\mobile-fitflow\src\components\web-badge.tsx`
- `D:\code\mobile-fitflow\src\global.css`

### Hooks & Constants
- `D:\code\mobile-fitflow\src\hooks\use-theme.ts`
- `D:\code\mobile-fitflow\src\hooks\use-color-scheme.ts`
- `D:\code\mobile-fitflow\src\hooks\use-color-scheme.web.ts`
- `D:\code\mobile-fitflow\src\constants\theme.ts`

### Feature Modules
- `D:\code\mobile-fitflow\src\features\workouts\api.ts`
- `D:\code\mobile-fitflow\src\features\workouts\hooks.ts`
- `D:\code\mobile-fitflow\src\features\workouts\types.ts`
- `D:\code\mobile-fitflow\src\features\meals\api.ts`
- `D:\code\mobile-fitflow\src\features\meals\hooks.ts`
- `D:\code\mobile-fitflow\src\features\meals\types.ts`
- `D:\code\mobile-fitflow\src\features\progress\api.ts`
- `D:\code\mobile-fitflow\src\features\progress\hooks.ts`
- `D:\code\mobile-fitflow\src\features\progress\types.ts`

### Assets
- `D:\code\mobile-fitflow\assets\images\splash-icon.png`
- `D:\code\mobile-fitflow\assets\images\icon.png`
- `D:\code\mobile-fitflow\assets\images\favicon.png`
- `D:\code\mobile-fitflow\assets\images\android-icon-foreground.png`
- `D:\code\mobile-fitflow\assets\images\android-icon-background.png`
- `D:\code\mobile-fitflow\assets\images\android-icon-monochrome.png`

### Other
- `D:\code\mobile-fitflow\scripts\reset-project.js`
- `D:\code\mobile-fitflow\nativewind-env.d.ts`
- `D:\code\mobile-fitflow\expo-env.d.ts`

---

