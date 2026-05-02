import { describe, it, expect } from 'vitest';
import { millisecondsToHumanTimestamp, humanTimestampToSecondsTo, monthSuffixedDay, formatDateWithWeekdayAndOrdinal, year } from './date.js';

describe('millisecondsToHumanTimestamp', () => {
	it.each([
		[0, '00:00:00'],
		[1000, '00:00:01'],
		[60000, '00:01:00'],
		[3600000, '01:00:00'],
		[3661000, '01:01:01'],
		[500, '00:00:00'], // sub-second rounds down
		[86399000, '23:59:59'], // just under 24 hours
		[5430000, '01:30:30'],
	])('millisecondsToHumanTimestamp(%i) = %s', (ms, expected) => {
		expect(millisecondsToHumanTimestamp(ms)).toBe(expected);
	});
});

describe('humanTimestampToSecondsTo', () => {
	it.each([
		['00:00:00', 0],
		['01:30:45', 5445],
		['00:05:30', 330],
		['(01:30:45)', 5445], // parentheses stripped
		['(00:05:30)', 330],
		['00:00:01', 1],
		['23:59:59', 86399],
	])('humanTimestampToSecondsTo(%s) = %i', (ts, expected) => {
		expect(humanTimestampToSecondsTo(ts)).toBe(expected);
	});
});

describe('monthSuffixedDay', () => {
	it.each([
		// 1st, 2nd, 3rd
		['2026-01-01', 'en-US', 'January 1st'],
		['2026-01-02', 'en-US', 'January 2nd'],
		['2026-01-03', 'en-US', 'January 3rd'],
		// 4th through 20th all get "th"
		['2026-01-04', 'en-US', 'January 4th'],
		['2026-01-11', 'en-US', 'January 11th'],
		['2026-01-12', 'en-US', 'January 12th'],
		['2026-01-13', 'en-US', 'January 13th'],
		// 21st, 22nd, 23rd
		['2026-01-21', 'en-US', 'January 21st'],
		['2026-01-22', 'en-US', 'January 22nd'],
		['2026-01-23', 'en-US', 'January 23rd'],
		['2026-01-30', 'en-US', 'January 30th'],
	])('monthSuffixedDay(%s, %s) = %s', (date, locale, expected) => {
		expect(monthSuffixedDay(date, locale)).toBe(expected);
	});
});

describe('formatDateWithWeekdayAndOrdinal', () => {
	it('formats a date with weekday and ordinal suffix', () => {
		// 2026-02-26 is a Thursday
		expect(formatDateWithWeekdayAndOrdinal('2026-02-26', 'en-US')).toBe('Thursday, 26th of February 2026');
	});

	it('handles 1st suffix', () => {
		// 2026-03-01 is a Sunday
		expect(formatDateWithWeekdayAndOrdinal('2026-03-01', 'en-US')).toBe('Sunday, 1st of March 2026');
	});

	it('handles 2nd suffix', () => {
		// 2026-03-02 is a Monday
		expect(formatDateWithWeekdayAndOrdinal('2026-03-02', 'en-US')).toBe('Monday, 2nd of March 2026');
	});

	it('handles 3rd suffix', () => {
		// 2026-03-03 is a Tuesday
		expect(formatDateWithWeekdayAndOrdinal('2026-03-03', 'en-US')).toBe('Tuesday, 3rd of March 2026');
	});

	it('handles 11th (not 11st)', () => {
		// 2026-03-11 is a Wednesday
		expect(formatDateWithWeekdayAndOrdinal('2026-03-11', 'en-US')).toBe('Wednesday, 11th of March 2026');
	});
});

describe('year', () => {
	it('extracts year from date string', () => {
		expect(year('2026-04-10', 'en-US')).toBe('2026');
	});

	it('extracts year from Date object', () => {
		expect(year(new Date(2025, 0, 1), 'en-US')).toBe('2025');
	});
});
