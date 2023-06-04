// Open Close JS for burger menu

const openMenu = () => {
	const menus = Array.from(document.querySelectorAll('.navbar-menu'));
	menus.map((m) => m.classList.remove('hidden'));
};

const closeMenu = () => {
	const menus = Array.from(document.querySelectorAll('.navbar-menu'));
	menus.map((m) => m.classList.add('hidden'));
};

const scrollToElement = (el) => {
	const offsetTop = document.querySelector(el).offsetTop;
	window.scrollTo({
		top: offsetTop,
		behavior: 'smooth',
	});
}
