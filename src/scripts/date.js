// returns the date in the format of Monday, October 12, 2023
export function formatDate(date, locale = 'de-DE') {
	const options = {
		weekday: 'long',
		year: 'numeric',
		month: 'long',
		day: 'numeric',
	};
	return new Date(date).toLocaleDateString(locale, options);
}

export function formatUnixTimestampToDate(timestamp, locale = 'de-DE') {
	const options = {
		weekday: 'long',
		year: 'numeric',
		month: 'long',
		day: 'numeric',
	};
	return new Date(timestamp * 1000).toLocaleDateString(locale, options);
}

// returns the date in the format of October 12, 2023
export function formatDateWithoutWeekday(date, locale = 'de-DE') {
	const options = {
		year: 'numeric',
		month: 'long',
		day: 'numeric',
	};
	return new Date(date).toLocaleDateString(locale, options);
}

// returns the time in the format HH:MM
export function formatTime(date, locale = 'de-DE', timeZone = 'Europe/Berlin') {
	const options = {
		hour: 'numeric',
		minute: 'numeric',
		timeZone,
	};
	return new Date(date).toLocaleTimeString(locale, options);
}

/**
 * Transforms a number of milliseconds (e.g. 1240) to
 * a human readable timestamp in the format hours:minutes:seconds.
 *
 * @param int ms
 * @returns string
 */
export function millisecondsToHumanTimestamp(ms) {
	const daysms = ms % (24*60*60*1000);
	const hours = Math.floor(daysms / (60*60*1000));
	const hoursms = ms % (60*60*1000);
	const minutes = Math.floor(hoursms / (60*1000));
	const minutesms = ms % (60*1000);
	const sec = Math.floor(minutesms / 1000);

	return String(hours).padStart(2, '0') + ":" + String(minutes).padStart(2, '0') + ":" + String(sec).padStart(2, '0');
}

/**
 * Transforms a a human readable timestamp in the format hours:minutes:seconds
 * into number of seconds.
 *
 * @param string ts
 * @returns int
 */
export function humanTimestampToSecondsTo(ts) {
	const humanTimestamp = ts.replaceAll(/[)(]/g, "");
	const timestampParts = humanTimestamp.split(":");
	const hourSeconds = parseInt(timestampParts[0]) * 60 * 60;
	const minuteSeconds = parseInt(timestampParts[1]) * 60;
	const seconds = parseInt(timestampParts[2]);

	return hourSeconds + minuteSeconds + seconds;
}