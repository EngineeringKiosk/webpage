import { defineConfig } from 'astro/config';
import tailwind from '@astrojs/tailwind';
import sitemap from '@astrojs/sitemap';
import mdx from '@astrojs/mdx';

// do not add to sitemap if specified string is contained in path
const exludeFromSitemap = ['meetup/alps/promote/', 'meetup/alps/qr/'];

// https://astro.build/config
export default defineConfig({
	site: 'https://engineeringkiosk.dev/',
	trailingSlash: 'always',
	// redirects: {
	// 	'/meetup/alps/qr/': 'https://google.at'
	//   },
	integrations: [
		tailwind(),
		sitemap({
			filter: (page) => !exludeFromSitemap.some((path) => page.includes(path)),
		}),
		mdx(),
	],
});
