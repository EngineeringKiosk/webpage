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
 * Returns a flat, props-compatible object for a team member.
 * Social fields are spread to the top level and any overrides are merged on top.
 */
export function resolveTeamMember(memberKey, overrides = {}) {
  const member = getTeamMember(memberKey);
  const { social, badges = [], ...rest } = member;
  return { ...rest, ...social, badges, ...overrides };
}
