export function orderByName(a, b) {
	if (a.data.name.toLowerCase() < b.data.name.toLowerCase()) {
		return -1;
	}

	if (a.data.name.toLowerCase() > b.data.name.toLowerCase()) {
		return 1;
	}

	return 0;
}

export function sortByDateDesc(items, dateField = 'pubDate') {
	return items.sort((a, b) => new Date(b.data[dateField]).valueOf() - new Date(a.data[dateField]).valueOf());
}
