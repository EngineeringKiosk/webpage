// Canonical `source` values for the newsletter signup payload. Every call
// site that renders <NewsletterForm /> (directly or via the meetup
// NewsletterWorker) imports from here so analytics can attribute signups to
// a surface without typos drifting between pages.
export const NEWSLETTER_SOURCES = {
	HOME_CTA: 'website-home-cta',
	NEWSLETTER_PAGE: 'website-newsletter-page',
	MEETUP: 'website-meetup',
};
