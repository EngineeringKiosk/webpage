import teamMembers from '../data/team-members.json';

/**
 * Returns the raw team member object for the given key.
 * Throws if the key is not found.
 */
export function getTeamMember(memberKey) {
  const member = teamMembers[memberKey];
  if (!member) {
    throw new Error(`Team member "${memberKey}" not found in team-members.json`);
  }
  return member;
}

/**
 * Returns an array of social profile URLs for a team member,
 * suitable for use as schema.org sameAs values.
 */
export function getSocialUrls(member) {
  const urls = [];
  if (member.social.github) urls.push(`https://github.com/${member.social.github}`);
  if (member.social.x) urls.push(`https://x.com/${member.social.x}`);
  if (member.social.linkedin) urls.push(`https://linkedin.com/in/${member.social.linkedin}`);
  if (member.social.bluesky) urls.push(member.social.bluesky);
  if (member.social.mastodon) urls.push(member.social.mastodon);
  if (member.social.website) urls.push(member.social.website);
  return urls;
}

/**
 * Returns a flat, props-compatible object for a team member.
 * Social fields are spread to the top level and any overrides are merged on top.
 */
export function resolveTeamMember(memberKey, overrides = {}) {
  const member = getTeamMember(memberKey);
  const { social, badges = [], ...rest } = member;
  return { ...rest, ...social, badges, ...overrides };
}
