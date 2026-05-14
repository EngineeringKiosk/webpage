import { describe, it, expect } from 'vitest';
import { isValidEmail, buildNewsletterPayload, interpretNewsletterResponse } from './newsletter.js';

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
	it('returns the canonical worker payload shape with all fields', () => {
		expect(
			buildNewsletterPayload({
				email: 'andy@example.com',
				newsletters: ['general'],
				source: 'home-cta',
				locale: 'de',
				honeypot: '',
			})
		).toEqual({
			email: 'andy@example.com',
			newsletters: ['general'],
			honeypot: '',
			source: 'home-cta',
			locale: 'de',
		});
	});

	it('omits source when not provided', () => {
		expect(
			buildNewsletterPayload({
				email: 'andy@example.com',
				newsletters: ['general'],
				honeypot: '',
			})
		).toEqual({
			email: 'andy@example.com',
			newsletters: ['general'],
			honeypot: '',
		});
	});

	it('omits source when empty string', () => {
		const payload = buildNewsletterPayload({
			email: 'andy@example.com',
			newsletters: ['general'],
			source: '',
			honeypot: '',
		});
		expect(payload).not.toHaveProperty('source');
	});

	it('omits locale when not provided', () => {
		const payload = buildNewsletterPayload({
			email: 'andy@example.com',
			newsletters: ['general'],
			honeypot: '',
		});
		expect(payload).not.toHaveProperty('locale');
	});

	it('omits locale when empty string', () => {
		const payload = buildNewsletterPayload({
			email: 'andy@example.com',
			newsletters: ['general'],
			locale: '',
			honeypot: '',
		});
		expect(payload).not.toHaveProperty('locale');
	});

	it('defaults honeypot to empty string when undefined', () => {
		const payload = buildNewsletterPayload({
			email: 'andy@example.com',
			newsletters: ['general'],
		});
		expect(payload.honeypot).toBe('');
	});

	it('preserves a non-empty honeypot value (worker decides what to do)', () => {
		const payload = buildNewsletterPayload({
			email: 'andy@example.com',
			newsletters: ['general'],
			honeypot: 'gotcha',
		});
		expect(payload.honeypot).toBe('gotcha');
	});

	it('passes the newsletters array through verbatim', () => {
		expect(
			buildNewsletterPayload({
				email: 'andy@example.com',
				newsletters: ['general', 'meetup_alps'],
			}).newsletters
		).toEqual(['general', 'meetup_alps']);
	});
});

describe('interpretNewsletterResponse', () => {
	it('treats 200 as success', () => {
		expect(interpretNewsletterResponse({ status: 200, messages })).toEqual({
			kind: 'success',
			message: messages.success,
		});
	});

	it('treats 201 as success', () => {
		expect(interpretNewsletterResponse({ status: 201, messages })).toEqual({
			kind: 'success',
			message: messages.success,
		});
	});

	it('maps 400 to validation + invalidEmail', () => {
		expect(interpretNewsletterResponse({ status: 400, messages })).toEqual({
			kind: 'validation',
			message: messages.invalidEmail,
		});
	});

	it('maps 405 to server + generic', () => {
		expect(interpretNewsletterResponse({ status: 405, messages })).toEqual({
			kind: 'server',
			message: messages.generic,
		});
	});

	it('maps 500 to server + generic', () => {
		expect(interpretNewsletterResponse({ status: 500, messages })).toEqual({
			kind: 'server',
			message: messages.generic,
		});
	});

	it('maps 502 to server + generic', () => {
		expect(interpretNewsletterResponse({ status: 502, messages })).toEqual({
			kind: 'server',
			message: messages.generic,
		});
	});
});
