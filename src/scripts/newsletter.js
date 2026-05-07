// Newsletter signup helpers.
// Backend: Plunk (https://docs.useplunk.com/). Sends `POST /v1/track` with
// subscribed:false so Plunk's double-opt-in workflow handles confirmation.
// Custom field `data.newsletter = true` powers a dynamic segment in Plunk.

export const PLUNK_API_BASE = 'https://next-api.useplunk.com';
export const NEWSLETTER_EVENT = 'newsletter-signup';

const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

export function isValidEmail(value) {
	if (typeof value !== 'string') {
		return false;
	}
	return EMAIL_PATTERN.test(value.trim());
}

export function buildNewsletterPayload({ email, source } = {}) {
	const data = { newsletter: true };
	if (typeof source === 'string' && source.length > 0) {
		data.source = source;
	}
	return {
		email,
		event: NEWSLETTER_EVENT,
		subscribed: false,
		data,
	};
}

export function interpretNewsletterResponse({ status, body, messages }) {
	const safeBody = body && typeof body === 'object' ? body : {};
	const bodyMessage = typeof safeBody.message === 'string' ? safeBody.message : '';

	if (status >= 200 && status < 300) {
		return { kind: 'success', message: messages.success };
	}

	if (status === 400 || status === 422) {
		if (mentionsEmailIssue(bodyMessage)) {
			return { kind: 'validation', message: messages.invalidEmail };
		}
		return { kind: 'validation', message: bodyMessage || messages.generic };
	}

	if (status === 401 || status === 403) {
		return { kind: 'server', message: messages.generic };
	}

	if (status === 429) {
		return { kind: 'rateLimit', message: messages.generic };
	}

	return { kind: 'server', message: messages.generic };
}

function mentionsEmailIssue(message) {
	if (!message) {
		return false;
	}
	const lowered = message.toLowerCase();
	return lowered.includes('email') || lowered.includes('e-mail') || lowered.includes('format');
}
