!(function t(e, i, r) {
	function n(a, o) {
		if (!i[a]) {
			if (!e[a]) {
				var c = 'function' == typeof require && require;
				if (!o && c) return c(a, !0);
				if (s) return s(a, !0);
				throw new Error("Cannot find module '" + a + "'");
			}
			var u = (i[a] = { exports: {} });
			e[a][0].call(
				u.exports,
				function (t) {
					var i = e[a][1][t];
					return n(i || t);
				},
				u,
				u.exports,
				t,
				e,
				i,
				r
			);
		}
		return i[a].exports;
	}
	for (var s = 'function' == typeof require && require, a = 0; a < r.length; a++) n(r[a]);
	return n;
})(
	{
		1: [
			function (t, e, i) {
				var r, n, s;
				(n = t('./iframe_resizer.coffee')),
					(s = t('./subscribe_button_trigger.coffee')),
					(r = class {
						constructor(t) {
							var e;
							(this.elem = t),
								(e = this.elem.getAttribute('data-configuration').replace(/(^\s+|\s+$)/g, '')),
								(this.id = this.randomId(e)),
								(this.configuration = 'string' == typeof e ? (e.match(/^{/) ? JSON.parse(e) : this.getInSiteConfig(e) || { json_config: e }) : e),
								(this.configuration.parentLocationHash = window.location.hash),
								(this.configuration.embedCode = this.elem.outerHTML);
							try {
								this.configuration.customOptions = JSON.parse(this.elem.getAttribute('data-options'));
							} catch (t) {
								console.debug('[Podigee Podcast Player] data-options has invalid JSON');
							}
							(this.url = `${this.origin()}/podigee-podcast-player.html?id=${this.id}&iframeMode=script`), this.buildIframe(), this.setupListeners(), this.replaceElem(), this.configuration && this.injectConfiguration(), this.setupSubscribeButton();
						}
						getInSiteConfig(t) {
							var e, i;
							return 0 !== t.indexOf('http') && t.match(/\./) && !t.match(/^\//)
								? ((e = t.split('.')),
								  (i = null),
								  e.forEach(function (t) {
										return (i = null === i ? window[t] : i[t]);
								  }),
								  i)
								: window[t];
						}
						randomId(t) {
							var e, i, r, n, s;
							if (((e = 0), 0 === t.length)) return e;
							for (
								i = (t) => {
									if (!isNaN(t)) return (e = (e << 5) - e + t), (e &= e);
								},
									r = n = 0,
									s = t.length;
								0 <= s ? n <= s : n >= s;
								r = 0 <= s ? ++n : --n
							)
								i(t.charCodeAt(r));
							return `pdg-${e.toString(16).substring(1)}`;
						}
						origin() {
							return (this.elem.src || this.elem.getAttribute('src'))
								.match(/(^.*\/)/)[0]
								.replace(/javascripts\/$/, '')
								.replace(/\/$/, '');
						}
						buildIframe() {
							return (
								(this.iframe = document.createElement('iframe')),
								(this.iframe.id = this.id),
								(this.iframe.scrolling = 'no'),
								(this.iframe.src = this.url),
								(this.iframe.style.border = '0'),
								(this.iframe.style.overflowY = 'hidden'),
								(this.iframe.style.transition = 'height 100ms linear'),
								(this.iframe.style.minWidth = '100%'),
								(this.iframe.width = '1px'),
								(this.iframe.title = 'Podcast'),
								this.iframe.setAttribute('aria-label', 'Podcast'),
								this.iframe
							);
						}
						setupListeners() {
							return n.listen('resizePlayer', this.iframe);
						}
						setupSubscribeButton() {
							return window.addEventListener(
								'message',
								(t) => {
									var e;
									try {
										e = JSON.parse(t.data || t.originalEvent.data);
									} catch (t) {
										return;
									}
									if (e.id === this.iframe.id && 'loadSubscribeButton' === e.listenTo) return new s(this.iframe).listen();
								},
								!1
							);
						}
						replaceElem() {
							return (this.iframe.className += this.elem.className), this.elem.parentNode.replaceChild(this.iframe, this.elem);
						}
						injectConfiguration() {
							return window.addEventListener(
								'message',
								(t) => {
									var e, i;
									try {
										i = JSON.parse(t.data || t.originalEvent.data);
									} catch (t) {
										return;
									}
									if (i.id === this.iframe.id && 'sendConfig' === i.listenTo) return (e = this.configuration.constructor === String ? this.configuration : JSON.stringify(this.configuration)), this.iframe.contentWindow.postMessage(e, '*');
								},
								!1
							);
						}
					}),
					new (class {
						constructor() {
							var t, e, i, n, s;
							if (((s = []), 0 !== (e = document.querySelectorAll('script.podigee-podcast-player, div.podigee-podcast-player')).length)) {
								for (i = 0, n = e.length; i < n; i++) (t = e[i]), s.push(new r(t));
								window.podigeePodcastPlayers = s;
							}
						}
					})();
			},
			{ './iframe_resizer.coffee': 2, './subscribe_button_trigger.coffee': 4 },
		],
		2: [
			function (t, e, i) {
				var r;
				(r = class {
					static listen(t, e, i = {}, r) {
						return window.addEventListener(
							'message',
							(n) => {
								var s, a, o;
								try {
									a = JSON.parse(n.data || n.originalEvent.data);
								} catch (t) {
									return;
								}
								if (a.id === e.id && a.listenTo === t)
									return (s = a.height + (i.height || 0)), (o = /%$/.test(a.width) ? a.width : a.width + (i.width || 0)), (e.style.height = `${s}px`), (e.style.maxHeight = `${s}px`), (e.style.width = `${o}px`), (e.style.maxWidth = `${o}px`), null != r ? r(e) : void 0;
							},
							!1
						);
					}
				}),
					(e.exports = r);
			},
			{},
		],
		3: [
			function (t, e, i) {
				!(function () {
					var t;
					if (!('function' == typeof window.CustomEvent || this.CustomEvent.toString().indexOf('CustomEventConstructor') > -1))
						((t = function (t, e) {
							var i;
							return (e = e || { bubbles: !1, cancelable: !1, detail: void 0 }), (i = document.createEvent('CustomEvent')).initCustomEvent(t, e.bubbles, e.cancelable, e.detail), i;
						}).prototype = window.Event.prototype),
							(window.CustomEvent = t);
				})();
			},
			{},
		],
		4: [
			function (t, e, i) {
				var r;
				t('./polyfills/custom_event.coffee'),
					(r = class {
						constructor(t) {
							(this.referenceElement = t), (this.referenceId = this.referenceElement.id), (this.id = this.randomId(this.referenceElement.toString())), this.buildTags(), this.insert();
						}
						buildTags() {
							return (
								(this.scriptTag = document.createElement('script')),
								(this.scriptTag.className = 'podlove-subscribe-button'),
								(this.scriptTag.src = 'https://cdn.podigee.com/subscribe-button/javascripts/app.js'),
								(this.scriptTag.dataset.language = 'en'),
								(this.scriptTag.dataset.size = 'medium'),
								this.scriptTag.setAttribute('data-hide', !0),
								this.scriptTag.setAttribute('data-buttonid', this.id),
								(this.button = document.createElement('button')),
								(this.button.style.display = 'none'),
								(this.button.className = `podlove-subscribe-button-${this.id}`)
							);
						}
						insert() {
							return this.referenceElement.parentNode.insertBefore(this.scriptTag, this.referenceElement.nextSibling), this.referenceElement.parentNode.insertBefore(this.button, this.referenceElement.nextSibling);
						}
						randomId(t) {
							var e, i, r, n, s, a;
							if (
								((s = Math.floor(65536 * (1 + Math.random()))
									.toString(16)
									.substring(1)),
								(e = 0),
								0 === (t += s).length)
							)
								return e;
							for (
								i = (t) => {
									if (!isNaN(t)) return (e = (e << 5) - e + t), (e &= e);
								},
									r = n = 0,
									a = t.length;
								0 <= a ? n <= a : n >= a;
								r = 0 <= a ? ++n : --n
							)
								i(t.charCodeAt(r));
							return e.toString(16).substring(1);
						}
						listen() {
							return window.addEventListener(
								'message',
								(t) => {
									var e, i;
									try {
										e = JSON.parse(t.data || t.originalEvent.data);
									} catch (t) {
										return;
									}
									if ('subscribeButtonTrigger' === e.listenTo && e.id === this.referenceId) return ((i = e.detail).id = this.id), (t = new CustomEvent('openSubscribeButtonPopup', { detail: i })), document.body.dispatchEvent(t);
								},
								!1
							);
						}
					}),
					(e.exports = r);
			},
			{ './polyfills/custom_event.coffee': 3 },
		],
	},
	{},
	[1]
);
