(function () {
	let maxTriesToSetPlayerCommands = 100;
	var _paq = window._paq || [];
	const setPlayerCommands = function() {
		const iframe = document.querySelector('iframe.podigee-podcast-player');
		if (iframe && playerjs && playerjs.Player) {
			// Tracking docs at https://developer.matomo.org/api-reference/tracking-javascript
			const player = new playerjs.Player(iframe);
			player.on('ready', function () {
				player.on('play', function (data) {
					_paq.push([
						'trackEvent',
						'PodcastPlayer', // Category
						'play', // Action
						window.playerConfiguration.episode.title, // Event name
						Math.round(data.seconds), // Numeric value
					]);
				});
				player.on('pause', function (data) {
					_paq.push([
						'trackEvent',
						'PodcastPlayer', // Category
						'pause', // Action
						window.playerConfiguration.episode.title, // Event name
						Math.round(data.seconds), // Numeric value
					]);
				});
				player.on('ended', function () {
					_paq.push([
						'trackEvent',
						'PodcastPlayer', // Category
						'end', // Action
						window.playerConfiguration.episode.title, // Event name
					]);
				});
			});

			document.addEventListener('DOMContentLoaded', function() {
				// Click events for the transcriptions
				const transcriptsElements = document.querySelectorAll('[data-trans-start]');
				for (let item of transcriptsElements) {
					item.classList.add('cursor-pointer');
					item.addEventListener('click', (e) => {
						player.setCurrentTime(item.dataset.transStart);
						player.play()
					});
				}

				// Click events for the chapters in the shownotes
				const chapterElements = document.querySelectorAll('[data-chapter-start]');
				for (let item of chapterElements) {
					item.classList.add('cursor-pointer');
					item.addEventListener('click', (e) => {
						player.setCurrentTime(item.dataset.chapterStart);
						player.play()
					});
				}
			})
		} else {
			// as longs as the player and iframe wasn't loaded yet, try again
			if (maxTriesToSetPlayerCommands > 0) {
				maxTriesToSetPlayerCommands--;
				setTimeout(setPlayerCommands, 300);
			}
		}
	}
	setPlayerCommands();
})();
