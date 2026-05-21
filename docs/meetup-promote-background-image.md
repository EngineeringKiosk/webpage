# Meetup promote page: background image

This document explains how to change the background image of the social-share / promote page that lives at `/meetup/<region>/promote/` (e.g. [engineeringkiosk.dev/meetup/rhine-ruhr/promote/](https://engineeringkiosk.dev/meetup/rhine-ruhr/promote/)).

The same asset is also used as the page's Open Graph preview (`<meta property="og:image">`), i.e. what social platforms render when the page URL is shared.

## How the image is rendered

The promote page paints a 700×700 px canvas with the image as a CSS background and the next-meetup details (date, city, host, talks) overlaid on top.

CSS used:

```css
background-size: cover;
background-position: center;
background-repeat: no-repeat;
```

- Square assets fill the canvas with no cropping.
- Off-aspect assets (e.g. 16:9) are scaled to cover the square and **center-cropped** on the long sides.

## Recommended dimensions

| Use         | Width × Height | Aspect | Notes                                                                                |
| ----------- | -------------- | ------ | ------------------------------------------------------------------------------------ |
| Recommended | 1200×1200      | 1:1    | Crisp on retina, square OG preview accepted by Facebook, X/Twitter and LinkedIn      |
| Minimum     | 700×700        | 1:1    | Exact canvas size — may look soft on retina displays                                 |
| Avoid       | non-square     | ≠ 1:1  | Will be center-cropped by `background-size: cover` (top/bottom or sides get clipped) |

Format: JPG or PNG. Keep file size reasonable (< 500 kB) — the asset is loaded both for the visible page and for social preview crawlers.

## Where to put the file

Drop the image into the region's public images folder:

```
public/meetup/<region>/images/<your-file>.jpg
```

Existing examples:

- `public/meetup/rhine-ruhr/images/rr.jpg` — Rhine-Ruhr default
- `public/meetup/alps/images/ibk.jpg` — Alps default

A naming convention like `<event-id>-<short-slug>.jpg` (e.g. `2606-duisburg.jpg`) keeps per-event assets easy to find.

## Setting it per meetup (recommended)

Add the optional `promoteBackgroundImage` field to the meetup's frontmatter, referencing the asset by its absolute public path:

```yaml
---
date: 2026-06-10T18:00:00+02:00
promoteBackgroundImage: '/meetup/rhine-ruhr/images/2606-duisburg.jpg'
location:
  name: 'Krankikom GmbH'
  address: 'Calaisplatz 5, 47051 Duisburg'
  # ...
talks:
  # ...
---
```

The schema field is defined on `meetupSchema` in [`src/content.config.ts`](../src/content.config.ts) and shared by both `meetup-alps` and `meetup-rhine-ruhr` collections.

## Changing the region-wide default

If you want to swap the default that's used **when no per-meetup override is set**, replace the file the page currently points to. The default path is wired in the region's promote page:

- Rhine-Ruhr: [`src/pages/meetup/rhine-ruhr/promote/index.astro`](../src/pages/meetup/rhine-ruhr/promote/index.astro) → `ogImage="/meetup/rhine-ruhr/images/rr.jpg"`
- Alps: [`src/pages/meetup/alps/promote/index.astro`](../src/pages/meetup/alps/promote/index.astro) → `ogImage="/meetup/alps/images/ibk.jpg"`

Either overwrite the file at that path with a new square asset, or change the path in the `.astro` file to point at a new asset.

## Fallback behavior

The component (`src/components/meetup/PromoteSocialImage.astro`) picks the effective image like this:

1. If there is an upcoming meetup **and** it has `promoteBackgroundImage` set → that wins.
2. Otherwise → the region's default `ogImage` prop.

The chosen image drives both the visible background **and** the `og:image` meta tag, so social previews stay in sync with what visitors see on the page.
