// Build a URL -> lastmod ISO-string map for use in the @astrojs/sitemap
// `serialize` hook. Runs once at config-load time and reads the relevant
// content directories synchronously so the sitemap step can do plain Map
// lookups.
//
// Coverage:
//   - Podcast episodes:  pubDate from frontmatter
//   - Blog posts:        pubDate from frontmatter
//   - Meetup events:     date from frontmatter (Alps + Rhine-Ruhr)
//   - Aggregate index pages (/podcast/, /blog/, /meetup/alps/, /meetup/rhine-ruhr/,
//     /deutsche-tech-podcasts/, /filme-fuer-softwareentwickler/, /, ...): max of
//     their underlying entries' timestamps.
//
// Out of scope (kept on default = build time): tag pages, per-genre game
// pages, per-category/per-type movie pages. Those refresh on every build anyway.

import fs from 'node:fs';
import path from 'node:path';

const ROOT = path.resolve(new URL('..', import.meta.url).pathname);

function readFrontmatterValue(filePath, key) {
	let text;
	try {
		text = fs.readFileSync(filePath, 'utf8');
	} catch {
		return null;
	}
	const fmMatch = text.match(/^---\s*\n([\s\S]*?)\n---/);
	if (!fmMatch) return null;
	const re = new RegExp(`^${key}:\\s*['"]?([^'"\\n]+?)['"]?\\s*$`, 'm');
	const m = fmMatch[1].match(re);
	return m ? m[1] : null;
}

function toIsoOrNull(value) {
	if (!value) return null;
	const d = new Date(value);
	return Number.isFinite(d.getTime()) ? d.toISOString() : null;
}

function listMatchingFiles(dir, regex) {
	const abs = path.join(ROOT, dir);
	if (!fs.existsSync(abs)) return [];
	return fs
		.readdirSync(abs)
		.filter((f) => regex.test(f))
		.map((f) => ({ name: f, full: path.join(abs, f) }));
}

function maxIso(values) {
	let max = null;
	for (const v of values) {
		if (!v) continue;
		if (!max || v > max) max = v;
	}
	return max;
}

function readJson(filePath) {
	try {
		return JSON.parse(fs.readFileSync(filePath, 'utf8'));
	} catch {
		return null;
	}
}

function stripExtension(name) {
	return name.replace(/\.(md|mdx|json)$/i, '');
}

export function buildLastmodMap() {
	const map = new Map();
	const set = (urlPath, iso) => {
		if (iso) map.set(urlPath, iso);
	};

	// --- Podcast episodes ----------------------------------------------------
	const episodeIsos = [];
	for (const { name, full } of listMatchingFiles('src/content/podcast', /\.md$/)) {
		const iso = toIsoOrNull(readFrontmatterValue(full, 'pubDate'));
		if (iso) {
			episodeIsos.push(iso);
			set(`/podcast/episode/${stripExtension(name)}/`, iso);
		}
	}
	const newestEpisode = maxIso(episodeIsos);
	set('/podcast/', newestEpisode);
	// The homepage prominently surfaces the latest episode.
	set('/', newestEpisode);

	// --- Blog posts ----------------------------------------------------------
	const blogIsos = [];
	for (const { name, full } of listMatchingFiles('src/content/blog', /\.(md|mdx)$/)) {
		const iso = toIsoOrNull(readFrontmatterValue(full, 'pubDate'));
		if (iso) {
			blogIsos.push(iso);
			set(`/blog/post/${stripExtension(name)}/`, iso);
		}
	}
	set('/blog/', maxIso(blogIsos));

	// --- Meetup events -------------------------------------------------------
	for (const flavor of ['alps', 'rhine-ruhr']) {
		const isos = [];
		for (const { name, full } of listMatchingFiles(`src/content/meetup-${flavor}`, /\.(md|mdx)$/)) {
			const iso = toIsoOrNull(readFrontmatterValue(full, 'date'));
			if (iso) {
				isos.push(iso);
				set(`/meetup/${flavor}/${stripExtension(name)}/`, iso);
			}
		}
		set(`/meetup/${flavor}/`, maxIso(isos));
	}

	// --- German Tech Podcasts directory -------------------------------------
	const gtpIsos = [];
	for (const { full } of listMatchingFiles('src/content/germantechpodcasts', /\.json$/)) {
		const data = readJson(full);
		const epoch = data?.latestEpisodePublished;
		if (typeof epoch === 'number' && epoch > 0) {
			gtpIsos.push(new Date(epoch * 1000).toISOString());
		}
	}
	set('/deutsche-tech-podcasts/', maxIso(gtpIsos));

	// --- Movies directory ----------------------------------------------------
	const movieIsos = [];
	for (const { full } of listMatchingFiles('src/content/awesome-software-engineering-movies', /\.json$/)) {
		const data = readJson(full);
		const iso = toIsoOrNull(data?.publishedAt);
		if (iso) movieIsos.push(iso);
	}
	set('/filme-fuer-softwareentwickler/', maxIso(movieIsos));

	// --- Games directory -----------------------------------------------------
	const gameIsos = [];
	for (const { full } of listMatchingFiles('src/content/awesome-software-engineering-games', /\.json$/)) {
		const data = readJson(full);
		const iso = toIsoOrNull(data?.release_date?.date);
		if (iso) gameIsos.push(iso);
	}
	set('/spiele-fuer-softwareentwickler/', maxIso(gameIsos));

	return map;
}

export function pathFromSitemapUrl(siteUrl, item) {
	// item.url is the full canonical URL (e.g. https://engineeringkiosk.dev/foo/).
	// Sitemaps emit percent-encoded paths; decode so we can match raw filenames
	// (umlauts in German episode slugs).
	try {
		const u = new URL(item.url);
		return decodeURI(u.pathname);
	} catch {
		return item.url.replace(siteUrl.replace(/\/$/, ''), '');
	}
}

const DIST_DIR = path.join(ROOT, 'dist');

// Reads the rendered HTML for a sitemap URL and returns the og:image absolute
// URL, or null if the page has none. Runs at sitemap-serialise time, after
// Astro has already written every page to disk.
export function extractOgImage(siteUrl, urlPath) {
	const trimmed = urlPath.replace(/^\/+|\/+$/g, '');
	const candidate = trimmed ? path.join(DIST_DIR, trimmed, 'index.html') : path.join(DIST_DIR, 'index.html');
	let html;
	try {
		html = fs.readFileSync(candidate, 'utf8');
	} catch {
		return null;
	}
	const m = html.match(/<meta[^>]+property=["']og:image["'][^>]+content=["']([^"']+)["']/i);
	if (!m) return null;
	try {
		return new URL(m[1], siteUrl).toString();
	} catch {
		return null;
	}
}
