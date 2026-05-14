// Newsletter signup constants used across the site.
//
// `NEWSLETTER_SOURCES` — analytics tags identifying *where* a signup
// originated. Free-form strings on the wire; the registry exists so call
// sites share canonical values.
//
// `NEWSLETTER_LISTS` — Plunk list identifiers consumed by the Cloudflare
// Worker. These MUST match the lists configured in the worker; renaming a
// value here without updating the worker silently breaks signups.
export const NEWSLETTER_SOURCES = {
	HOME_CTA: 'website-home-cta',
	NEWSLETTER_PAGE: 'website-newsletter-page',
	MEETUP: 'website-meetup',
};

export const NEWSLETTER_LISTS = {
	MEETUP_RHINE_RUHR: 'meetup_rhine_ruhr',
};
