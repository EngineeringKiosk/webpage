import rss from '@astrojs/rss';

// Docs: https://docs.astro.build/en/guides/rss/
export const get = () => rss({
    title: 'Engineering Kiosk',
    description: 'Der deutschsprachige Software-Engineering-Podcast mit Wolfgang Gassler und Andy Grunwald rund um die Themen Engineering-Kultur, Open Source und Technologie.',
    site: import.meta.env.SITE,
    // Currently hitting a bug here
    // See https://github.com/withastro/astro/issues/3946
    items: import.meta.glob([
        './blog/post/*.md', 
        './podcast/episode/*.md'
    ])
  });