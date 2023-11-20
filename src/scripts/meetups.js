import { getCollection } from "astro:content";

let allMeetups = await getCollection("meetup");
allMeetups = allMeetups.map((meetup) => meetup.data)
const timeDivider = new Date(new Date().setHours(23, 59, 59, 999)).getTime();


export function getNextMeetups(limit = undefined) {
  const meetups = allMeetups.filter((meetup) => new Date(meetup.date).getTime() > timeDivider);
  meetups.sort((a, b) => new Date(a.date).valueOf() - new Date(b.date).valueOf());

  return meetups.slice(0, limit)
}

export function getPastMeetups(limit = undefined) {
  const meetups = allMeetups.filter((meetup) => (new Date(meetup.date).getTime() <= timeDivider));
  meetups.sort((a, b) => new Date(b.date).valueOf() - new Date(a.date).valueOf());

  return meetups.slice(0, limit)
}
