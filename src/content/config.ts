import { z, defineCollection } from 'astro:content';

// Schema for Podcast Episodes
const podcastEpisodeCollection = defineCollection({
	type: 'content',
	schema: ({ image }) => z.object({
		amazon_music: z.string(),
		apple_podcasts: z.string(),
		audio: z.string(),
		chapter: z.array(
			z.object({
				start: z.string(),
				title: z.string(),
			})
		),
		deezer: z.string(),
		description: z.string(),
		google_podcasts: z.string(),
		headlines: z.string(),
		image: image(),
		length_second: z.number(),
		pubDate: z.date(),
		rtlplus: z.string(),
		speaker: z.array(
			z.object({
				name: z.string(),
				transcriptLetter: z.string(),
				website: z.string(),
			})
		),
		spotify: z.string(),
		tags: z.array(z.string()),
		title: z.string(),
		youtube: z.string(),
	}),
});

// Schema for Blog Entries
const blogEntryCollection = defineCollection({
	type: 'content',
	schema: ({ image }) => z.object({
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
  schema: z.object({
    date: z.date(),
    location: z.object({
      name: z.string(),
      address: z.string(),
      url: z.string().optional(),
      logo: z.string().optional()
    }),
    talks: z.array(z.object({
        avatar: z.string().optional(),
        name: z.string(),
        title: z.string(),
        description: z.string(),
        github: z.string().optional(),
        twitter: z.string().optional(),
        linkedin: z.string().optional(),
        website: z.string().optional(),
        bio: z.string().optional(),
        slides: z.string().optional()
      })
    )
  })
});

// Export a single `collections` object to register your collection(s)
export const collections = {
  podcast: podcastEpisodeCollection,
  blog: blogEntryCollection,
  meetup: meetupCollection
};
