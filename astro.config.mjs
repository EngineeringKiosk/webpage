import { defineConfig } from 'astro/config';
import tailwindcss from '@tailwindcss/vite';
import sitemap from '@astrojs/sitemap';
import mdx from '@astrojs/mdx';

// do not add to sitemap if specified string is contained in path
const exludeFromSitemap = ['meetup/alps/promote/'];

// https://astro.build/config
export default defineConfig({
	site: 'https://engineeringkiosk.dev/',
	trailingSlash: 'always',

	integrations: [
		sitemap({
			filter: (page) => !exludeFromSitemap.some((path) => page.includes(path)),
		}),
		mdx(),
	],

	vite: {
		plugins: [tailwindcss()],
	},
});
