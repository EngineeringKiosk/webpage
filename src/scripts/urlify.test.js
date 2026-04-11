import { describe, it, expect } from 'vitest';
import { URLify } from './urlify.js';

describe('URLify', () => {
	it('converts spaces to hyphens', () => {
		expect(URLify('Hello World')).toEqual({ name: 'Hello World', url: 'hello-world' });
	});

	it('removes dots', () => {
		expect(URLify('Node.js')).toEqual({ name: 'Node.js', url: 'nodejs' });
	});

	it('lowercases the url', () => {
		expect(URLify('TypeScript')).toEqual({ name: 'TypeScript', url: 'typescript' });
	});

	it('preserves original name', () => {
		const result = URLify('Infrastructure as Code');
		expect(result.name).toBe('Infrastructure as Code');
		expect(result.url).toBe('infrastructure-as-code');
	});

	it('trims whitespace', () => {
		expect(URLify('  padded  ')).toEqual({ name: '  padded  ', url: 'padded' });
	});

	it('handles multiple spaces', () => {
		const result = URLify('a b');
		expect(result.url).toBe('a-b');
	});

	it('handles combined dots and spaces', () => {
		expect(URLify('C.I. / C.D.')).toEqual({ name: 'C.I. / C.D.', url: 'ci-/-cd' });
	});
});
