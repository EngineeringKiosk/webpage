import { describe, it, expect } from 'vitest';
import { orderByName, sortByDateDesc } from './sorting.js';

describe('orderByName', () => {
	it('sorts items by data.name case-insensitively', () => {
		const items = [
			{ data: { name: 'Zebra' } },
			{ data: { name: 'apple' } },
			{ data: { name: 'Banana' } },
		];
		const sorted = [...items].sort(orderByName);
		expect(sorted.map((i) => i.data.name)).toEqual(['apple', 'Banana', 'Zebra']);
	});

	it('returns 0 for equal names', () => {
		const a = { data: { name: 'test' } };
		const b = { data: { name: 'Test' } };
		expect(orderByName(a, b)).toBe(0);
	});
});

describe('sortByDateDesc', () => {
	it('sorts items by pubDate descending', () => {
		const items = [
			{ data: { pubDate: '2024-01-01' } },
			{ data: { pubDate: '2024-06-15' } },
			{ data: { pubDate: '2024-03-10' } },
		];
		sortByDateDesc(items);
		expect(items.map((i) => i.data.pubDate)).toEqual(['2024-06-15', '2024-03-10', '2024-01-01']);
	});

	it('supports custom date field name', () => {
		const items = [
			{ data: { date: '2024-01-01' } },
			{ data: { date: '2024-12-25' } },
		];
		sortByDateDesc(items, 'date');
		expect(items.map((i) => i.data.date)).toEqual(['2024-12-25', '2024-01-01']);
	});

	it('handles empty array', () => {
		const items = [];
		sortByDateDesc(items);
		expect(items).toEqual([]);
	});
});
