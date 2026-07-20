import { defineCollection } from 'astro:content';
import { z } from 'astro/zod';
import { glob } from 'astro/loaders';

// Schema for Podcast Episodes
const podcastEpisodeCollection = defineCollection({
	loader: glob({ pattern: '**/[^_]*.md', base: './src/content/podcast' }),
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
			hideSpeakerInTranscript: z.boolean().optional(),
		}),
});

// Schema for Blog Entries
const blogEntryCollection = defineCollection({
	loader: glob({ pattern: '**/[^_]*.{md,mdx}', base: './src/content/blog' }),
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

// Schema for Meetups (shared between meetup-alps and meetup-rhine-ruhr)
const meetupSchema = ({ image }) =>
	z.object({
		date: z.date(),
		eventId: z.string().optional(),
		registrationClosed: z.boolean().optional(),
		// Per-event override for the promote-page background image (also
		// used as the page's Open Graph preview). Absolute /public/ path,
		// e.g. "/meetup/rhine-ruhr/images/2606-duisburg.jpg". Square 1200×1200
		// is the recommended target. Omit to fall back to the region default.
		promoteBackgroundImage: z.string().optional(),
		location: z.object({
			name: z.string(),
			address: z.string(),
			url: z.string().optional(),
			logo: image().optional(),
			note: z.string().optional(),
		}),
		talks: z.array(
			z.object({
				avatar: z.union([image(), z.array(image())]).optional(),
				name: z.string(),
				title: z.string(),
				description: z.string(),
				github: z.string().optional(),
				x: z.string().optional(),
				// string or array of strings
				linkedin: z.union([z.string(), z.array(z.string())]).optional(),
				bluesky: z.string().optional(),
				mastodon: z.string().optional(),
				threads: z.string().optional(),
				instagram: z.string().optional(),
				youtube: z.string().optional(),
				website: z.string().optional(),
				bio: z.string().optional(),
				// Either a filename inside `public/meetup/<region>/slides/` (e.g. "2602-roland-golla.pdf")
				// or an absolute http(s) URL to a hosted slide deck. Filenames get prefixed with
				// /meetup/<region>/slides/ at render time (Talk.astro); URLs are used as-is.
				slides: z.string().optional(),
				language: z.enum(['en', 'de']).optional(),
				affiliation: z
					.object({
						role: z.string(),
						organization: z.string(),
					})
					.optional(),
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
				// Three states for newParticipants:
				//   - omitted: not tracked (legacy / not yet considered) — no sub-line rendered
				//   - null:    explicitly "no first-timer data for this event" — renders "(no data for first timers)"
				//   - number:  actual count (0 is a legitimate datum meaning zero first-timers)
				newParticipants: z.number().nullable().optional(),
			})
			.optional(),
		speakers: z
			.object({
				female: z.number().optional(),
				male: z.number().optional(),
				nonbinary: z.number().optional(),
			})
			.optional(),
	});

const meetupAlpsCollection = defineCollection({
	loader: glob({ pattern: '**/[^_]*.{md,mdx}', base: './src/content/meetup-alps' }),
	schema: meetupSchema,
});
const meetupRhineRuhrCollection = defineCollection({
	loader: glob({ pattern: '**/[^_]*.{md,mdx}', base: './src/content/meetup-rhine-ruhr' }),
	schema: meetupSchema,
});

// Schema for German Tech Podcasts
const germantechpodcastsCollection = defineCollection({
	loader: glob({ pattern: '**/[^_]*.json', base: './src/content/germantechpodcasts' }),
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
	loader: glob({ pattern: '**/[^_]*.json', base: './src/content/awesome-software-engineering-games' }),
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
			license: z
				.object({
					name: z.string(),
					spdx_id: z.string(),
					url: z.string().url(),
				})
				.optional(),
		}),
});

// Schema for awesome-software-engineering-movies
const awesomeSoftwareEngineeringMoviesCollection = defineCollection({
	loader: glob({ pattern: '**/[^_]*.json', base: './src/content/awesome-software-engineering-movies' }),
	schema: ({ image }) =>
		z.object({
			name: z.string(),
			slug: z.string(),
			// `title` may be empty for entries that only have a localized title.
			title: z.string(),
			description: z.string(),
			// Open string values; the German display label lives in src/scripts/movie-labels.js.
			category: z.string(),
			type: z.string(),
			tags: z.array(z.string()),
			// `language`/`subtitles` are explicitly null for some upstream entries
			// (no language metadata from the source), not just absent.
			language: z.array(z.string()).nullable(),
			subtitles: z.array(z.string()).nullable(),
			// Open platform map: "youtube", "netflix", "bpb", ... -> URL.
			links: z.record(z.string(), z.string().url()),
			localized: z
				.record(
					z.string(),
					z.object({
						title: z.string().optional(),
						description: z.string().optional(),
						links: z.record(z.string(), z.string().url()).optional(),
					})
				)
				.optional(),
			// ISO-8601 duration ("PT12M49S") or empty for non-YouTube entries.
			duration: z.string(),
			// ISO-8601 timestamp or empty for non-YouTube entries.
			publishedAt: z.string(),
			channel: z.object({
				id: z.string(),
				title: z.string(),
			}),
			ratings: z
				.object({
					youtube: z
						.object({
							likeCount: z.number(),
							refreshedAt: z.string(),
						})
						.optional(),
					imdb: z
						.object({
							averageRating: z.number(),
							numVotes: z.number(),
							refreshedAt: z.string(),
						})
						.optional(),
				})
				.optional(),
			views: z
				.object({
					youtube: z.number(),
				})
				.optional(),
			imdbID: z.string().optional(),
			youtubeTrailerForThumbnail: z.string().url().optional(),
			image: image(),
		}),
});

// Export a single `collections` object to register your collection(s)
export const collections = {
	podcast: podcastEpisodeCollection,
	blog: blogEntryCollection,
	'meetup-alps': meetupAlpsCollection,
	'meetup-rhine-ruhr': meetupRhineRuhrCollection,
	germantechpodcasts: germantechpodcastsCollection,
	'awesome-software-engineering-games': awesomeSoftwareEngineeringGamesCollection,
	'awesome-software-engineering-movies': awesomeSoftwareEngineeringMoviesCollection,
};
