# Engineering Kiosk Website

Available at [engineeringkiosk.dev](https://www.engineeringkiosk.dev/).

## Getting started

```sh
$ make init
$ make run
```

## 🧞 Commands

All commands are run from the root of the project, from a terminal.

Available `Makefile` commands:

```
build                                               Compiles the application into static content
check-episode-player-urls                           Checks all podcast episodes if all player links (Spotify, etc.) are set
clean                                               Deletes all generated items (like node_modules, build output, caches)
find-missing-tag-descriptions-content-files         Finds all used tags in content files that need SEO descriptions and output them on stdout
find-missing-tag-descriptions-german-tech-podcast   Finds all used tags in the german tech podcasts that need SEO descriptions and output them on stdout
help                                                Outputs the help
init-javascript                                     Installs JavaScript dependencies
init-python                                         Installs python dependencies and creates a virtualenv
init                                                Installs all dependencies (JavaScript and python)
optimize-episode-transcriptions                     Trims and shortens the episode text transcriptions to only the data we actually need
prettier                                            Run code formatter prettier (for JavaScript)
run                                                 Starts the development server
update-episode-content-no-api                       Pulls the latest Podcast RSS feed from RedCircle and updates the content (without doing external API calls)
update-episode-content                              Pulls the latest Podcast RSS feed from RedCircle and updates the content
update-episode-redirects                            Writes all short url redirects for Podcast episodes to netlify.toml
update-german-tech-podcasts                         Updates the German Tech Podcasts data from https://github.com/EngineeringKiosk/GermanTechPodcasts
update-missing-tag-descriptions-content-files       Finds all used tags in content files that need SEO descriptions and updates the tag-file
update-missing-tag-descriptions-german-tech-podcast Find all used tags in the german tech podcasts that need SEO descriptions and updates the tag-file
update-podcast-statistics                           Calculates and updates podcast statistics
```

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

## Development notes

### Podcast player

We are using our own fork of [podigee/podigee-podcast-player](https://github.com/podigee/podigee-podcast-player) which can be found at [EngineeringKiosk/podigee-podcast-player](https://github.com/EngineeringKiosk/podigee-podcast-player).
The main differences of the fork in comparision to its origin can be found in the forks README.

In combination, we downloaded a minified version of [https://github.com/embedly/player.js](https://github.com/embedly/player.js) into `public/js`.
If we aim to upgrade _player.js_, download and replace the minified javascript file.
