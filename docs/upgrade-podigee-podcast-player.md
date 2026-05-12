# Upgrading the Podigee Podcast Player

## Scope

This document describes how to refresh the vendored copy of the
[EngineeringKiosk/podigee-podcast-player](https://github.com/EngineeringKiosk/podigee-podcast-player)
fork that lives under `public/podcast-player/`.

It does **not** cover `public/js/player-0.1.0.min.js`. That is a
separate dependency from [embedly/player.js](https://github.com/embedly/player.js)
with its own upgrade procedure (see the "Podcast player" section in
the top-level [`README.md`](../README.md)).

## How the player is integrated

The episode pages use the player in **iframe-script mode**.

1. `src/layouts/podcast-episode.astro` builds a `window.playerConfiguration`
   object (theme, chapter marks, transcript URL, ...) and includes:

   ```html
   <script class="podigee-podcast-player"
           src="/podcast-player/javascripts/podigee-podcast-player.js"
           data-configuration="playerConfiguration"></script>
   ```

2. That loader script scans the page for `script.podigee-podcast-player`,
   builds an `<iframe>` pointing at
   `/podcast-player/podigee-podcast-player.html?id=…&iframeMode=script`,
   and posts the configuration JSON into it via `postMessage`.

3. The iframe document is **self-contained**: the fork's Gulp build
   inlines the theme CSS and the embed JS bundle into
   `podigee-podcast-player.html` at build time (`<!-- inject:head:css -->`
   and inline `<style>` / `<script>` blocks). The browser never
   requests an external `*-embed.js` or theme CSS file.

A separate small script, `public/js/podcast-player-tracking.js`, hooks
the iframe's `playerjs` events (via Embedly's `player.js`) to push
play/pause/ended events into Matomo.

## Files we vendor

Only the files actually requested by the browser at runtime:

| Path under `public/podcast-player/` | Role |
| --- | --- |
| `podigee-podcast-player.html` | Iframe document (CSS and embed JS inlined at build time) |
| `javascripts/podigee-podcast-player.js` | Loader served by the script tag |
| `fonts/podigee-podcast-player.ttf` | Icon font (modern browsers) |
| `fonts/podigee-podcast-player.woff` | Icon font (modern browsers) |
| `fonts/podigee-podcast-player.svg` | Icon font (legacy fallback) |
| `fonts/podigee-podcast-player.eot` | Icon font (legacy IE) — kept for `@font-face` parity with upstream |
| `images/chromcast.png` | Chromecast UI icon, referenced from inlined CSS |
| `README.md` | Pointer to the fork commit this snapshot was built from |

## Files we do NOT vendor (and why)

The fork's `build/` directory contains additional artifacts that we
intentionally skip — none of them are loaded in iframe-script mode:

- `podigee-podcast-player-direct.html`,
  `javascripts/podigee-podcast-player-embed.js`,
  `javascripts/podigee-podcast-player-direct.js` — direct-mode bundles
  for non-iframe embeds.
- `stylesheets/app.css`, `stylesheets/app-direct.css` — already inlined
  into `podigee-podcast-player.html` by the build.
- `themes/default/index.css`, `themes/default/index.html`,
  `themes/default-dark/`, `themes/legacy/`, `themes/minimal/` —
  default theme is inlined; other themes unused.
- `fonts/podigee-podcast-player.json` — icomoon glyph metadata
  (~1.9 MB); not referenced by any `@font-face` rule.
- All `*.gz` siblings — Netlify performs its own compression.

If a future change to the fork starts referencing one of these from
the inlined HTML, add it to the vendor list above and re-sync.

## Upgrade procedure

### 1. Refresh the fork's build

In the fork checkout
(`/Users/andygrunwald/Development/EngineeringKiosk/podigee-podcast-player`):

```bash
git status                  # working tree must be clean
git pull                    # or check out the commit you want to ship
make init                   # install deps (Yarn 4)
make clean && make build    # regenerate build/ from scratch
```

`make clean && make build` is recommended because partially regenerated
`build/` directories (mixed file timestamps) have shipped slightly
inconsistent snapshots in the past.

Note the commit you're pinning to:

```bash
git rev-parse HEAD
git branch --show-current
```

### 2. Copy the runtime files into the webpage

From the webpage repo root, on a fresh feature branch
(`andygrunwald/upgrade-podigee-podcast-player` or similar):

| Source (`<fork>/build/…`) | Destination (`public/podcast-player/…`) |
| --- | --- |
| `podigee-podcast-player.html` | `podigee-podcast-player.html` |
| `javascripts/podigee-podcast-player.js` | `javascripts/podigee-podcast-player.js` |
| `fonts/podigee-podcast-player.ttf` | `fonts/podigee-podcast-player.ttf` |
| `fonts/podigee-podcast-player.woff` | `fonts/podigee-podcast-player.woff` |
| `fonts/podigee-podcast-player.svg` | `fonts/podigee-podcast-player.svg` |
| `fonts/podigee-podcast-player.eot` | `fonts/podigee-podcast-player.eot` |
| `images/chromcast.png` | `images/chromcast.png` |

Plain `cp` is sufficient. Do not copy anything not listed above.

### 3. Update the vendor pointer

Edit `public/podcast-player/README.md` and replace the `git/sha:…`
link and the branch reference with the commit hash and branch you
captured in step 1.

### 4. Validate

From the webpage repo root:

```bash
make build              # Astro build must succeed
make test-javascript    # Vitest sanity check
make prettier           # should be a no-op; vendored files are in .prettierignore
```

Then run a manual smoke test:

```bash
make run
```

Open an episode page (e.g. `http://localhost:4321/podcast/episode/…/`)
and confirm:

- The player iframe renders with the default theme.
- Play / pause / chapter marks work.
- Clicking a transcript line jumps the audio (`data-trans-start`
  handler in `public/js/podcast-player-tracking.js`).
- DevTools **Network** tab requests, from `/podcast-player/...`, only
  `podigee-podcast-player.html`, the loader JS, the font files, and
  `chromcast.png`. No 404s, no orphan requests to removed files.
- DevTools **Console** is clean and `window._paq` entries fire on
  play / pause / ended.

### 5. Open the PR

Split the work into reviewable, atomic commits — at minimum:

1. (optional) Cleanup of unused vendored files, if any new dead weight
   has accumulated since the last sync.
2. The player bump itself (binaries + `public/podcast-player/README.md`
   pointer update).
3. Docs, if this guide needs adjusting.

## When something inside the iframe changes

Behavioural regressions from a fork bump almost always surface in
`public/js/podcast-player-tracking.js`, which bridges the iframe's
`playerjs` events into Matomo. If the fork changes the event names or
the iframe message protocol, that file is what needs updating — the
loader and the iframe document themselves are vendored as-is.

## Related

- Fork repo: <https://github.com/EngineeringKiosk/podigee-podcast-player>
- Upstream:  <https://github.com/podigee/podigee-podcast-player>
- Embedly `player.js` (separate upgrade): see the top-level
  [`README.md`](../README.md) "Podcast player" section.
