# Frontend — Storybook note

Storybook and its related devDependencies were removed from this frontend package to allow a clean upgrade to Svelte 5 and `@sveltejs/vite-plugin-svelte@^6`.

Why:
- Several Storybook addons in this repository depend on Svelte 3 / older `@sveltejs/vite-plugin-svelte` versions. Those peer dependency conflicts prevented a normal `npm install` for the frontend.

What changed:
- Storybook-related devDependencies and `sb` / `build-storybook` scripts were removed from `package.json`.

How to restore Storybook (quick, temporary):

1. Re-add the previous Storybook devDependencies (example versions used previously):

```bash
cd application/frontend
npm install --save-dev \
  @storybook/addon-actions@^6.5.12 \
  @storybook/addon-essentials@^6.5.12 \
  @storybook/addon-interactions@^6.5.12 \
  @storybook/addon-links@^6.5.12 \
  @storybook/builder-vite@^0.2.3 \
  @storybook/svelte@^6.5.12 \
  @storybook/testing-library@^0.0.13 \
  --legacy-peer-deps
```

2. Re-add the scripts you need in `package.json`:

```json
"scripts": {
  "sb": "start-storybook -p 6006",
  "build-storybook": "build-storybook"
}
```

Notes:
- The `--legacy-peer-deps` approach is a temporary workaround. For a long-term fix, migrate Storybook to a version compatible with Svelte 5 (a dedicated migration is recommended, ideally in a separate branch/PR).
- Alternatively, reintroduce Storybook in a separate branch where you can upgrade Storybook and its addons together and resolve any breaking changes.

If you want, I can create a migration branch and attempt the full Storybook v10 upgrade and compatibility fixes.
