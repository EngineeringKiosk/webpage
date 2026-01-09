import { z, defineCollection } from 'astro:content';

// Schema for Podcast Episodes
const podcastEpisodeCollection = defineCollection({
	type: 'content',
	schema: ({ image }) =>
		z.object({
			advertiser: z.string(),
			amazon_music: z.string(),
			apple_podcasts: z.string(),
			audio: z.string(),
			chapter: z.array(
				z.object({
					start: z.string(),
					title: z.string(),
				})
			),
			deezer: z.string().optional(),
			description: z.string(),
			headlines: z.string(),
			image: image(),
			pubDate: z.date(),
			speaker: z.array(
				z.object({
					name: z.string(),
					transcriptLetter: z.string().optional(),
				})
			),
			spotify: z.string().optional(),
			tags: z.array(z.string()),
			title: z.string(),
			transcript_slim: z.string(),
			youtube: z.string().optional(),
		}),
});

// Schema for Blog Entries
const blogEntryCollection = defineCollection({
	type: 'content',
	schema: ({ image }) =>
		z.object({
			title: z.string(),
			subtitle: z.string(),
			description: z.string(),
			tags: z.array(z.string()),
			pubDate: z.date(),
			thumbnail: image(),
			headerimage: image(),
		}),
});

// Schema for Meetups
const meetupCollection = defineCollection({
	type: 'content',
	schema: ({ image }) =>
		z.object({
			date: z.date(),
			eventId: z.string().optional(),
			registrationClosed: z.boolean().optional(),
			location: z.object({
				name: z.string(),
				address: z.string(),
				url: z.string().optional(),
				logo: image().optional(),
				note: z.string().optional(),
			}),
			talks: z.array(
				z.object({
					avatar: image().optional(),
					name: z.string(),
					title: z.string(),
					description: z.string(),
					github: z.string().optional(),
					twitter: z.string().optional(),
					// string or array of strings
					linkedin: z.union([z.string(), z.array(z.string())]).optional(),
					mastodon: z.string().optional(),
					website: z.string().optional(),
					bio: z.string().optional(),
					slides: z.string().optional(),
				})
			),
			participants: z
				.object({
					registered: z.number(),
					present: z
						.object({
							total: z.number().optional(),
							male: z.number().optional(),
							female: z.number().optional(),
						})
						.optional(),
					newParticipants: z.number().optional(),
				})
				.optional(),
			speakers: z
				.object({
					female: z.number().optional(),
					male: z.number().optional(),
					nonbinary: z.number().optional(),
				})
				.optional(),
		}),
});

// Schema for German Tech Podcasts
const germantechpodcastsCollection = defineCollection({
	type: 'data',
	schema: ({ image }) =>
		z.object({
			name: z.string(),
			slug: z.string(),
			website: z.string(),
			podcastIndexID: z.number(),
			rssFeed: z.string(),
			spotify: z.string(),
			description: z.string(),
			tags: z.array(z.string()).nullable(),
			weekly_downloads_avg: z.object({
				value: z.number(),
				updated: z.string(),
			}),
			episodeCount: z.number(),
			latestEpisodePublished: z.number(),
			archive: z.boolean(),
			itunesID: z.number(),
			image: image(),
		}),
});

// Schema for awesome-software-engineering-games
const awesomeSoftwareEngineeringGamesCollection = defineCollection({
	type: 'data',
	schema: ({ image }) =>
		z.object({
			name: z.string(),
			slug: z.string(),
			website: z.string(),
			repository: z.string().optional(),
			programmable: z.boolean().optional(),
			steamID: z.number(),
			platforms: z.object({
				windows: z.boolean(),
				mac: z.boolean(),
				linux: z.boolean(),
			}),
			release_date: z.object({
				coming_soon: z.boolean(),
				date: z.string(),
			}),
			image: image(),
			german_content: z.object({
				short_description: z.string(),
				categories: z.array(z.string()).optional(),
				genres: z.array(z.string()),
			}),
			english_content: z.object({
				short_description: z.string(),
				categories: z.array(z.string()).optional(),
				genres: z.array(z.string()),
			}),
		}),
});

// Export a single `collections` object to register your collection(s)
export const collections = {
	podcast: podcastEpisodeCollection,
	blog: blogEntryCollection,
	meetup: meetupCollection,
	germantechpodcasts: germantechpodcastsCollection,
	'awesome-software-engineering-games': awesomeSoftwareEngineeringGamesCollection,
};
