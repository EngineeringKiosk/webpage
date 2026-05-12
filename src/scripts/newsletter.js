// Newsletter signup helpers.
// Backend: the Engineering Kiosk newsletter worker
// (https://api.engineeringkiosk.dev/newsletter/subscribe). Plunk integration,
// double-opt-in orchestration, and bot protection now live server-side.

const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

export function isValidEmail(value) {
	if (typeof value !== 'string') {
		return false;
	}
	return EMAIL_PATTERN.test(value.trim());
}

export function buildNewsletterPayload({ email, newsletters, source, honeypot } = {}) {
	const payload = {
		email,
		newsletters,
		honeypot: typeof honeypot === 'string' ? honeypot : '',
	};
	if (typeof source === 'string' && source.length > 0) {
		payload.source = source;
	}
	return payload;
}

export function interpretNewsletterResponse({ status, messages }) {
	if (status >= 200 && status < 300) {
		return { kind: 'success', message: messages.success };
	}
	if (status === 400) {
		return { kind: 'validation', message: messages.invalidEmail };
	}
	return { kind: 'server', message: messages.generic };
}
