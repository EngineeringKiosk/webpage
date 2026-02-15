import rss from '@astrojs/rss';
import { getCollection } from 'astro:content';

// Docs: https://docs.astro.build/en/guides/rss/
export async function GET(context) {
	const blogPosts = await getCollection('blog');
	const podcastEpisodes = await getCollection('podcast');

	let rssItems = blogPosts.map((post) => ({
		title: post.data.title,
		pubDate: post.data.pubDate,
		description: post.data.description,
		link: `/posts/${post.slug}/`,
	}));

	podcastEpisodes.map(function (episode) {
		rssItems.push({
			title: episode.data.title,
			pubDate: episode.data.pubDate,
			description: episode.data.description,
			link: `/podcast/episode/${episode.slug}/`,
		});
	});

	return rss({
		title: 'Engineering Kiosk',
		description: 'Der deutschsprachige Software-Engineering-Podcast, Tech-Meetups und Discord-Community rund um die Themen Engineering-Kultur, Open Source und Technologie.',
		site: context.site,
		items: rssItems,
		customData: `<language>de-de</language>`,
	});
}
