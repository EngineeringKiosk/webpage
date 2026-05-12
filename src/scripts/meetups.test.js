import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';

vi.mock('astro:content', () => ({
	getCollection: vi.fn(),
}));

import { getCollection } from 'astro:content';
import { createMeetupHelpers } from './meetups.js';

function buildMeetup(date) {
	return {
		data: {
			date,
			location: { name: 'Venue', address: 'Some street 1' },
			talks: [{ name: 'Speaker', title: 'Talk title', description: 'Desc' }],
		},
	};
}

describe('createMeetupHelpers', () => {
	beforeEach(() => {
		vi.useFakeTimers();
		// noon UTC keeps "today" = 2026-05-12 in any reasonable timezone
		vi.setSystemTime(new Date('2026-05-12T12:00:00Z'));
	});

	afterEach(() => {
		vi.useRealTimers();
		vi.clearAllMocks();
	});

	it('returns a meetup later today as next (the 1-day buffer treats meetup-day as upcoming)', async () => {
		getCollection.mockResolvedValue([buildMeetup('2026-05-12T18:30:00+02:00')]);
		const { getNextMeetup } = await createMeetupHelpers('meetup-alps');
		expect(getNextMeetup(1)).toBeDefined();
	});

	it('treats yesterday meetup as past, not next', async () => {
		getCollection.mockResolvedValue([buildMeetup('2026-05-11T08:00:00Z')]);
		const { getNextMeetup, getPastMeetups } = await createMeetupHelpers('meetup-alps');
		expect(getNextMeetup(1)).toBeUndefined();
		expect(getPastMeetups()).toHaveLength(1);
	});

	it('returns tomorrow meetup as next', async () => {
		getCollection.mockResolvedValue([buildMeetup('2026-05-13T20:00:00Z')]);
		const { getNextMeetup } = await createMeetupHelpers('meetup-alps');
		expect(getNextMeetup(1)).toBeDefined();
	});

	it('returns undefined when no upcoming meetups exist', async () => {
		getCollection.mockResolvedValue([
			buildMeetup('2026-04-01T18:30:00Z'),
			buildMeetup('2026-03-01T18:30:00Z'),
		]);
		const { getNextMeetup } = await createMeetupHelpers('meetup-alps');
		expect(getNextMeetup(1)).toBeUndefined();
	});

	it('picks the soonest upcoming meetup when several are scheduled', async () => {
		getCollection.mockResolvedValue([
			buildMeetup('2026-07-01T18:30:00Z'),
			buildMeetup('2026-05-13T18:30:00Z'),
			buildMeetup('2026-06-15T18:30:00Z'),
		]);
		const { getNextMeetup } = await createMeetupHelpers('meetup-alps');
		expect(getNextMeetup(1).date).toBe('2026-05-13T18:30:00Z');
	});
});
