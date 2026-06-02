// Display labels and small helpers for the awesome-software-engineering-movies
// content collection. The upstream JSON keeps category/type/platform values in
// English; this module is the single place that converts them to the German
// strings shown in the UI and used to build URL slugs.

export const CATEGORY_LABELS = {
	'Programming Languages': 'Programmiersprachen',
	'Applications / Frameworks / Systems': 'Anwendungen / Frameworks / Systeme',
	'Culture / Society': 'Kultur / Gesellschaft',
	'Culture / People': 'Kultur / Personen',
};

export const TYPE_LABELS = {
	Documentary: 'Dokumentation',
	Movie: 'Film',
	'TV Series': 'TV-Serie',
};

export const LINK_PLATFORM_LABELS = {
	youtube: 'YouTube',
	netflix: 'Netflix',
	amazon_prime_video: 'Amazon Prime',
	hbo_max: 'HBO Max',
	apple_tv: 'Apple TV',
	rtl_plus: 'RTL+',
	bpb: 'Bundeszentrale für politische Bildung',
};

// Brand SVGs that exist under public/images/brands/. Anything not in this set
// falls back to the generic website icon.
const KNOWN_BRAND_ICONS = new Set(['youtube', 'netflix', 'amazon_prime_video', 'hbo_max', 'apple_tv', 'rtl_plus']);

export function iconFor(platform) {
	if (KNOWN_BRAND_ICONS.has(platform)) {
		return `/images/brands/${platform}.svg`;
	}
	return '/images/elements/icon-website-2.svg';
}

// Returns either the mapped value or the raw key, so unknown/new upstream
// values still render rather than disappearing silently.
export function labelFor(map, key) {
	return map[key] ?? key;
}

// Slugifies a (German) display label into a URL-safe segment:
//   "Kultur / Gesellschaft"               -> "kultur-gesellschaft"
//   "Anwendungen / Frameworks / Systeme"  -> "anwendungen-frameworks-systeme"
//   "TV-Serie"                            -> "tv-serie"
export function slugLabel(label) {
	return label
		.toLowerCase()
		.replace(/[\s/.]+/g, '-')
		.replace(/-+/g, '-')
		.replace(/^-|-$/g, '');
}

// URL slug for a category, derived from the German label.
export function categorySlug(category) {
	return slugLabel(labelFor(CATEGORY_LABELS, category));
}

// URL slug for a type, derived from the German label.
export function typeSlug(type) {
	return slugLabel(labelFor(TYPE_LABELS, type));
}

// Headline used on per-type listing pages. Each type has its own canonical
// phrasing (not a literal label substitution).
export function typeHeadline(type) {
	switch (type) {
		case 'Documentary':
			return 'Dokumentationen für Softwareentwickler:innen';
		case 'Movie':
			return 'Filme für Softwareentwickler:innen';
		case 'TV Series':
			return 'TV-Serien für Softwareentwickler:innen';
		default:
			return `${labelFor(TYPE_LABELS, type)} für Softwareentwickler:innen`;
	}
}

// Title to render on a movie card. Prefer the German override when present,
// otherwise the canonical short `name`. The verbose `title` field (often
// includes subtitles like "| An origin story") is the last resort for the
// rare case where both `localized.de.title` and `name` are missing.
export function localizedTitle(movie) {
	return movie.data.localized?.de?.title || movie.data.name || movie.data.title;
}

export function localizedDescription(movie) {
	return movie.data.localized?.de?.description || movie.data.description;
}
