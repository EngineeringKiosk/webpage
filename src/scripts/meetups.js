import { getCollection } from 'astro:content';

let allMeetups = await getCollection('meetup');
allMeetups = allMeetups.map((meetup) => meetup.data);
const todayEndOfDay = new Date(new Date().setHours(23, 59, 59, 999)).getTime();

export function getNextMeetups(limit = undefined, timeDivider = undefined) {
	if (!timeDivider) {
		timeDivider = todayEndOfDay;
	}
	const meetups = allMeetups.filter((meetup) => new Date(meetup.date).getTime() > timeDivider);
	meetups.sort((a, b) => new Date(a.date).valueOf() - new Date(b.date).valueOf());

	return limit ? meetups.slice(0, limit) : meetups;
}

// to not immediately show the next meetup as soon as it's over
// add a buffer of X days
export function getNextMeetup(bufferDays = 0) {
	const timeDivider = new Date(new Date().getTime() - bufferDays * 24 * 60 * 60 * 1000);
	return getNextMeetups(1, timeDivider)[0];
}

export function getPastMeetups(limit = undefined, timeDivider = undefined) {
	if (!timeDivider) {
		timeDivider = todayEndOfDay;
	}
	const meetups = allMeetups.filter((meetup) => new Date(meetup.date).getTime() <= timeDivider);
	meetups.sort((a, b) => new Date(b.date).valueOf() - new Date(a.date).valueOf());

	return limit ? meetups.slice(0, limit) : meetups;
}
