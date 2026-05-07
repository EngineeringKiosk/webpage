import { describe, it, expect } from 'vitest';
import { isValidEmail, buildNewsletterPayload, interpretNewsletterResponse, NEWSLETTER_EVENT, PLUNK_API_BASE } from './newsletter.js';

const messages = {
	loading: 'Sende …',
	success: 'Danke! Bitte bestätige deine Anmeldung über den Link in deiner Inbox.',
	generic: 'Etwas ist schiefgelaufen. Bitte versuche es später erneut.',
	invalidEmail: 'Bitte gib eine gültige E-Mail-Adresse ein.',
};

describe('isValidEmail', () => {
	it.each([
		['a@b.co', true],
		['andy@example.com', true],
		['first.last+tag@sub.example.de', true],
		['', false],
		['foo', false],
		['foo@', false],
		['foo@bar', false],
		['a b@c.de', false],
		['@example.com', false],
		['no-at-sign.com', false],
	])('isValidEmail(%j) = %s', (input, expected) => {
		expect(isValidEmail(input)).toBe(expected);
	});

	it('returns false for non-string values', () => {
		expect(isValidEmail(undefined)).toBe(false);
		expect(isValidEmail(null)).toBe(false);
		expect(isValidEmail(42)).toBe(false);
	});

	it('trims surrounding whitespace before validating', () => {
		expect(isValidEmail('  andy@example.com  ')).toBe(true);
	});
});

describe('buildNewsletterPayload', () => {
	it('returns the canonical Plunk payload shape', () => {
		expect(buildNewsletterPayload({ email: 'andy@example.com', source: 'home-cta' })).toEqual({
			email: 'andy@example.com',
			event: 'newsletter-signup',
			subscribed: false,
			data: { newsletter: true, source: 'home-cta' },
		});
	});

	it('omits data.source when source is not provided', () => {
		expect(buildNewsletterPayload({ email: 'andy@example.com' })).toEqual({
			email: 'andy@example.com',
			event: 'newsletter-signup',
			subscribed: false,
			data: { newsletter: true },
		});
	});

	it('omits data.source when source is empty string', () => {
		const payload = buildNewsletterPayload({ email: 'andy@example.com', source: '' });
		expect(payload.data).toEqual({ newsletter: true });
	});

	it('uses the exported event constant', () => {
		const payload = buildNewsletterPayload({ email: 'andy@example.com' });
		expect(payload.event).toBe(NEWSLETTER_EVENT);
	});
});

describe('interpretNewsletterResponse', () => {
	it('treats 200 as success', () => {
		expect(interpretNewsletterResponse({ status: 200, body: { success: true }, messages })).toEqual({
			kind: 'success',
			message: messages.success,
		});
	});

	it('treats 201 as success', () => {
		expect(interpretNewsletterResponse({ status: 201, body: {}, messages })).toEqual({
			kind: 'success',
			message: messages.success,
		});
	});

	it('maps 400 with email-related message to validation + invalidEmail', () => {
		expect(
			interpretNewsletterResponse({
				status: 400,
				body: { message: 'Invalid email format' },
				messages,
			})
		).toEqual({ kind: 'validation', message: messages.invalidEmail });
	});

	it('maps 422 with email-related message to validation + invalidEmail', () => {
		expect(
			interpretNewsletterResponse({
				status: 422,
				body: { message: 'email is not valid' },
				messages,
			})
		).toEqual({ kind: 'validation', message: messages.invalidEmail });
	});

	it('falls back to body.message for non-email validation errors', () => {
		expect(
			interpretNewsletterResponse({
				status: 400,
				body: { message: 'Some other validation error' },
				messages,
			})
		).toEqual({ kind: 'validation', message: 'Some other validation error' });
	});

	it('uses generic fallback when 400 has no body message', () => {
		expect(interpretNewsletterResponse({ status: 400, body: {}, messages })).toEqual({
			kind: 'validation',
			message: messages.generic,
		});
	});

	it('maps 401 to server + generic', () => {
		expect(interpretNewsletterResponse({ status: 401, body: { message: 'Unauthorized' }, messages })).toEqual({
			kind: 'server',
			message: messages.generic,
		});
	});

	it('maps 403 to server + generic', () => {
		expect(interpretNewsletterResponse({ status: 403, body: {}, messages })).toEqual({
			kind: 'server',
			message: messages.generic,
		});
	});

	it('maps 429 to rateLimit + generic', () => {
		expect(interpretNewsletterResponse({ status: 429, body: {}, messages })).toEqual({
			kind: 'rateLimit',
			message: messages.generic,
		});
	});

	it('maps 500 to server + generic', () => {
		expect(interpretNewsletterResponse({ status: 500, body: {}, messages })).toEqual({
			kind: 'server',
			message: messages.generic,
		});
	});

	it('handles a non-object body without throwing', () => {
		expect(() => interpretNewsletterResponse({ status: 500, body: null, messages })).not.toThrow();
		expect(() => interpretNewsletterResponse({ status: 500, body: 'oops', messages })).not.toThrow();
	});
});

describe('module constants', () => {
	it('event name avoids reserved Plunk prefixes', () => {
		expect(NEWSLETTER_EVENT.startsWith('email.')).toBe(false);
		expect(NEWSLETTER_EVENT.startsWith('contact.')).toBe(false);
		expect(NEWSLETTER_EVENT.startsWith('segment.')).toBe(false);
	});

	it('Plunk API base points at api.useplunk.com', () => {
		expect(PLUNK_API_BASE).toBe('https://api.useplunk.com');
	});
});
