# Engineering Kiosk: Web page

Built with [Astro](https://astro.build/). 

Check it out at [engineeringkiosk.dev](https://www.engineeringkiosk.dev/)

## Getting started

```sh
$ make init

$ make run
```

## 🧞 Commands

All commands are run from the root of the project, from a terminal.


Available `Makefile` commands:

```
build                          Compiles the application into static content
clean                          Deletes the generated content and node_modules
episode-check                  Checks all Podcast Episodes if all player links (Spotify, etc.) are set
eslint                         Statically analyzes of basic JavaScript scripts
find-tags-that-need-descriptions-content-files-dump Find all used tags in content files (blog posts and podcasts) that need SEO descriptions and dump it down to disk into tag-file
find-tags-that-need-descriptions-content-files Checks all used tags in content files (blog posts and podcasts) that need SEO descriptions and output them on stdout
find-tags-that-need-descriptions-german-tech-podcast-file-dump Find all used tags in the german tech podcasts that need SEO descriptions and dump it down to disk into tag-file
find-tags-that-need-descriptions-german-tech-podcast-file Checks all used tags in the german tech podcasts that need SEO descriptions and output them on stdout
help                           Outputs the help
init                           Installs dependencies
prettier                       Run prettier (file formatting)
run                            Starts the development server
sync-german-tech-podcasts      Syncs German Tech Podcasts data from https://github.com/EngineeringKiosk/GermanTechPodcasts
update-content                 Pulls the latest Podcast RSS feed and updates the content
update-redirects               Writes all short url redirects for Podcast episodes to netlify.toml```

Native npm commands:

| Command           | Action                                       |
|:----------------  |:-------------------------------------------- |
| `npm install`     | Installs dependencies                        |
| `npm run dev`     | Starts local dev server at `localhost:3000`  |
| `npm run build`   | Build your production site to `./dist/`      |
| `npm run preview` | Preview your build locally, before deploying |

Native python script commands:

```
$ python3 scripts/podcast-feed-to-content.py -h
usage: podcast-feed-to-content.py [-h] [mode]

Automate new Podcast Episide parsing

positional arguments:
  mode        Mode to execute. Supported: sync, redirect (default: sync)

options:
  -h, --help  show this help message and exit
```

## Checklist: Releasing a new Podcast episode

1. Run GitHub Actions [Publish Podcast Episodes](https://github.com/EngineeringKiosk/webpage/actions/workflows/publish-podcast-episodes.yml)
2. Copy Episode Single View URL from [Engineering Kiosk at Amazon Music](https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/engineering-kiosk)
3. Add the Episode Single View URL into the Frontmatter of the newly released Podcast episode
4. Commit, push and wait until Netlify is building the new website version.

## Netlify

This page is hosted on Netlify.

Docs for the file based configuration of Netlify are available at https://docs.netlify.com/configure-builds/file-based-configuration/

## Blog posts: Image sizes

### Thumbnail (blog article preview)

| Scale | Width (px) | Height (px) |
| ----- | ---------- | ----------- |
| x1    | 358        | 272         |
| x1.5  | 537        | 408         |
| x2    | 716        | 544         |
| x2.5  | 895        | 680         |
| x3    | 1074       | 816         |
| x3.5  | 1253       | 952         |
| x4    | 1432       | 1088        |
| x4.5  | 1611       | 1224        |
| x5    | 1790       | 1360        |

### Header image (inside a blog post)

| Scale | Width (px) | Height (px) |
| ----- | ---------- | ----------- |
| x1    | 1440       | 641         |
| x1.5  | 2160       | 961,5       |
| x2    | 2880       | 1282        |
| x2.5  | 3600       | 1602,5      |
| x3    | 4320       | 1923        |

### Content images

| Scale | Width (px) | Height (px) |
| ----- | ---------- | ----------- |
| x1    | 750        | 422         |
| x1.5  | 1125       | 633         |
| x2    | 1500       | 844         |
| x2.5  | 1875       | 1055        |
| x3    | 2250       | 1266        |
| x3.5  | 2625       | 1477        |
| x4    | 3000       | 1688        |

## Development note: Podcast player

We are using our own fork of [podigee/podigee-podcast-player](https://github.com/podigee/podigee-podcast-player) which can be found at [EngineeringKiosk/podigee-podcast-player](https://github.com/EngineeringKiosk/podigee-podcast-player).
The main differences of the fork in comparision to its origin can be found in the forks README.

In combination, we downloaded a minified version of [https://github.com/embedly/player.js](https://github.com/embedly/player.js) into `public/js`.
If we aim to upgrade *player.js*, download and replace the minified javascript file.