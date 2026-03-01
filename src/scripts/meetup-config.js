export const meetupConfigs = {
	alps: {
		collectionName: 'meetup-alps',
		meetupName: 'Engineering Kiosk Alps',
		cityName: 'Innsbruck',
		regionName: 'Alps',
		urlSlug: 'alps',
		shortUrl: 'engineeringkiosk.dev/alps',
		registrationUrl: 'https://engineeringkiosk.dev/alps',
		backgroundImage: '/meetup/alps/images/ibk.jpg',
		signOffNamesNewsletter: 'Wolfi, Tim, Romedius, Christoph, Mirjam',
		signOffNamesAnnounce: 'Wolfgang, Tim, Romedius, Christoph and Mirjam',
		socialPostTitle: 'Social Media Post Meetup Alps - Engineering Kiosk - Innsbruck',
		socialPostDescription: 'Tech Meetup in Innsbruck, brings together tech enthusiasts, engineers, and professionals from various fields to explore the intersection of engineering culture, open source, people and technology.',
		newsletterTitle: 'Social Media Post Meetup Alps - Engineering Kiosk - Innsbruck',
		announceNewsletterTitle: 'Newsletter Template for Engineering Kiosk Alps',
		announceNewsletterDescription: 'Tech Meetup in Innsbruck, brings together tech enthusiasts, engineers, and professionals from various fields to explore the intersection of engineering culture, open source, people and technology.',
	},
	'rhine-ruhr': {
		collectionName: 'meetup-rhine-ruhr',
		meetupName: 'Engineering Kiosk Rhine-Ruhr',
		cityName: 'the Rhine-Ruhr region',
		regionName: 'Rhine-Ruhr',
		urlSlug: 'rhine-ruhr',
		shortUrl: 'engineeringkiosk.dev/rhine-ruhr',
		registrationUrl: 'https://engineeringkiosk.dev/rhine-ruhr',
		backgroundImage: '/meetup/rhine-ruhr/images/rr.jpg',
		signOffNamesNewsletter: 'Andy, Schepp, Dario',
		signOffNamesAnnounce: 'Andy, Schepp and Dario',
		socialPostTitle: 'Social Media Post Meetup Rhine-Ruhr - Engineering Kiosk',
		socialPostDescription: 'Tech Meetup in the Rhine-Ruhr region, brings together tech enthusiasts, engineers, and professionals from various fields to explore the intersection of engineering culture, open source, people and technology.',
		newsletterTitle: 'Social Media Post Meetup Rhine-Ruhr - Engineering Kiosk',
		announceNewsletterTitle: 'Newsletter Template for Engineering Kiosk Rhine-Ruhr',
		announceNewsletterDescription: 'Tech Meetup in the Rhine-Ruhr region, brings together tech enthusiasts, engineers, and professionals from various fields to explore the intersection of engineering culture, open source, people and technology.',
	},
};

export function getMeetupConfig(meetupId) {
	const config = meetupConfigs[meetupId];
	if (!config) {
		throw new Error(`Unknown meetup ID: ${meetupId}. Available: ${Object.keys(meetupConfigs).join(', ')}`);
	}
	return config;
}
