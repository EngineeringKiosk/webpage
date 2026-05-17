// Newsletter signup constants used across the site.
//
// `NEWSLETTER_SOURCES` — analytics tags identifying *where* a signup
// originated. Free-form strings on the wire; the registry exists so call
// sites share canonical values.
//
// `NEWSLETTER_LISTS` — Plunk list identifiers consumed by the Cloudflare
// Worker. These MUST match the lists configured in the worker; renaming a
// value here without updating the worker silently breaks signups.
//
// `NEWSLETTER_LOCALES` — BCP 47 / ISO 639 language codes passed through to
// Plunk's `data.locale` to localize the unsubscribe footer and hosted
// preferences pages. See https://docs.useplunk.com/guides/localization for
// the full list of languages Plunk supports.
export const NEWSLETTER_SOURCES = {
	HOME_CTA: 'website-home-cta',
	NEWSLETTER_PAGE: 'website-newsletter-page',
	MEETUP: 'website-meetup',
};

export const NEWSLETTER_LISTS = {
	MEETUP_RHINE_RUHR: 'meetup_rhine_ruhr',
	MEETUP_ALPS: 'meetup_alps',
};

export const NEWSLETTER_LOCALES = {
	DE: 'de',
	EN: 'en',
};
