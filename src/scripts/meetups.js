import { getCollection } from 'astro:content';

export async function createMeetupHelpers(collectionName) {
	let allMeetups = await getCollection(collectionName);
	allMeetups = allMeetups.map((meetup) => meetup.data);
	allMeetups = allMeetups.map((meetup) => ({
		...meetup,
		talks: meetup.talks.map((talk) => ({
			...talk,
			title: talk.title || 'To be announced',
			name: talk.name || 'Here could be your name',
			description: talk.description || 'The description will be available soon',
			bio: talk.bio || 'How awesome would it be to read your profile here?',
		})),
	}));
	const todayEndOfDay = new Date(new Date().setHours(23, 59, 59, 999)).getTime();

	function getNextMeetups(limit = undefined, timeDivider = undefined) {
		if (!timeDivider) {
			timeDivider = todayEndOfDay;
		}
		const meetups = allMeetups.filter((meetup) => new Date(meetup.date).getTime() > timeDivider);
		meetups.sort((a, b) => new Date(a.date).valueOf() - new Date(b.date).valueOf());

		return limit ? meetups.slice(0, limit) : meetups;
	}

	// to not immediately show the next meetup as soon as it's over
	// add a buffer of X days
	function getNextMeetup(bufferDays = 0) {
		const timeDivider = new Date(todayEndOfDay - bufferDays * 24 * 60 * 60 * 1000);
		return getNextMeetups(1, timeDivider)[0];
	}

	function getPastMeetups(limit = undefined, timeDivider = undefined) {
		if (!timeDivider) {
			timeDivider = todayEndOfDay;
		}
		const meetups = allMeetups.filter((meetup) => new Date(meetup.date).getTime() <= timeDivider);
		meetups.sort((a, b) => new Date(b.date).valueOf() - new Date(a.date).valueOf());

		return limit ? meetups.slice(0, limit) : meetups;
	}

	function getAllMeetups() {
		return allMeetups;
	}

	return { getNextMeetups, getNextMeetup, getPastMeetups, getAllMeetups };
}
