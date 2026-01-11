export function cutText(input, len) {
	let text = input;
	if (text.length > len) {
		const cutAt = text.lastIndexOf(' ', len);
		if (cutAt != -1) {
			text = text.substring(0, cutAt) + '...';
		}
	}

	return text;
}

export function cleanDescription(input) {
	const cutMarker = 'Unsere aktuellen Werbepartner';
	const index = input.indexOf(cutMarker);
	if (index !== -1) {
		return input.substring(0, index).trim();
	}
	return input;
}
