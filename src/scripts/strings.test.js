import { describe, it, expect } from 'vitest';
import { cutText, cleanDescription } from './strings.js';

describe('cutText', () => {
	it('returns text unchanged if shorter than limit', () => {
		expect(cutText('Hello', 10)).toBe('Hello');
	});

	it('cuts at last space before limit and adds ellipsis', () => {
		expect(cutText('Hello World Foo Bar', 11)).toBe('Hello World...');
	});

	it('returns full text if no space found before limit', () => {
		expect(cutText('Superlongwordwithoutspaces', 10)).toBe('Superlongwordwithoutspaces');
	});

	it('returns text unchanged if exactly at limit', () => {
		expect(cutText('Hello', 5)).toBe('Hello');
	});

	it('handles empty string', () => {
		expect(cutText('', 10)).toBe('');
	});

	it('cuts at word boundary for longer text', () => {
		const text = 'This is a longer sentence that should be cut at a word boundary';
		const result = cutText(text, 30);
		expect(result).toBe('This is a longer sentence that...');
		expect(result.length).toBeLessThanOrEqual(33); // 30 + '...'
	});
});

describe('cleanDescription', () => {
	it('removes text from marker onwards', () => {
		expect(cleanDescription('Good content. Unsere aktuellen Werbepartner: Sponsor1')).toBe('Good content.');
	});

	it('returns unchanged if no marker present', () => {
		expect(cleanDescription('Just a normal description')).toBe('Just a normal description');
	});

	it('trims whitespace before marker', () => {
		expect(cleanDescription('Content   Unsere aktuellen Werbepartner')).toBe('Content');
	});

	it('handles marker at the very start', () => {
		expect(cleanDescription('Unsere aktuellen Werbepartner and more')).toBe('');
	});
});
