export function URLify(element) {
	// Replace whitespace with -
	let e = element.trim().replace(/\s/g, '-');

	// Replace dots (.) with nothing.
	e = e.replace(/\./g, '');

	return {
		name: element,
		url: e.toLowerCase(),
	};
}
