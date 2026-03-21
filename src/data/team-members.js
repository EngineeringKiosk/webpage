/**
 * Page-level data for each team member's individual profile page.
 *
 * The slug must match the old static filename (without .astro) so that
 * existing URLs are preserved.  memberKey is the key used in
 * team-members.json (and by getTeamMember()).
 */
const teamMemberPages = [
	{
		slug: "andy-grunwald",
		memberKey: "andy-grunwald",
		title: "Andy Grunwald, Podcast Host & Meetup Organizer - Engineering Kiosk",
		description:
			"Engineering Manager im Bereich Site Reliability Engineering und Software Engineer mit einem Fokus auf Backend-Systeme, Infrastruktur und Engineering Kultur aus Duisburg.",
		ogImageAlt: "Andy Grunwald",
		profileFirstName: "Andy",
		profileLastName: "Grunwald",
		jsonLdName: null, // uses member.name
		h1: "Andy Grunwald",
		imagePosition: "right",
		bioHtml: `<p class="mb-6 text-lg md:text-xl text-coolGray-500 font-medium">
								Engineering Manager im Bereich <i>Site Reliability Engineering</i> und Software Engineer mit einem Fokus auf Backend-Systeme, Infrastruktur und Engineering Kultur. Zieht seine Motivation aus der Team-Arbeit mit Menschen, die motiviert sind und etwas nach vorne bringen wollen (<i
									>Do awesome things with awesome people</i
								>). Wohnt mittem im Ruhrgebiet, in Duisburg (und will das auch) 👷
							</p>
							<p class="mb-6 text-lg md:text-xl text-coolGray-500 font-medium">Langjähriger Open-Source Contributor und Hobby- (bzw. eher Möchtegern)-Landwirt 🧑‍🌾 Mag es, seine Grenzen beim Functional Fitness auszuloten und die Freizeit mit seinem Hund zu verbringen.</p>
							<p class="text-lg md:text-xl text-coolGray-500 font-medium">
								Neben dem <a href="/">Engineering Kiosk</a> betreibt Andy mit <a href="https://app.sourcectl.dev/" class="underline hover:no-underline">sourcectl</a> eine Plattform die Open Source-Kultur sowie Inner Source innerhalb von Firmen zu pushen.
							</p>`,
	},
	{
		slug: "wolfgang-gassler",
		memberKey: "wolfgang-gassler",
		title: "Wolfgang Gassler, Podcast Host & Meetup Organizer - Engineering Kiosk",
		description:
			"Freelancer und Consultant mit einem Interesse in Menschen, Tech-Kultur, Datenbank und Skalierung, Product Entwicklung und Usability sowie Recommender Systems aus Innsbruck.",
		ogImageAlt: "Wolfgang Gassler",
		profileFirstName: "Wolfgang",
		profileLastName: "Gassler",
		jsonLdName: null,
		h1: "Wolfgang Gassler",
		imagePosition: "left",
		bioHtml: `<p class="mb-6 text-lg md:text-xl text-coolGray-500 font-medium">
								<a href="https://gassler.org/consulting" class="underline hover:no-underline">Freelancer und Consultant</a> mit Interesse an Menschen, Tech &amp; Team-Kultur, Datenbank und Skalierung, Produkt Entwicklung und Usability sowie Recommender Systeme. Ehemals Engineering Manager bei <a
									href="https://company.trivago.com/"
									class="underline hover:no-underline">trivago</a
								> und auch viele Jahre in Lehre &amp; Forschung in der <a href="https://dbis-informatik.uibk.ac.at/" class="underline hover:no-underline">Forschungsgruppe Datenbanken und Informationssysteme an der Universität Innsbruck (Österreich)</a>. Gerne unterwegs (auch weit weg von den Bergen
								🌊 ). Ursprünglich aus Innsbruck (ganz nah an den Bergen ⛰️ ).
							</p>
							<p class="mb-6 text-lg md:text-xl text-coolGray-500 font-medium">Freiwilliger Feuerwehrmann und leidenschaftlicher Kletterer. Verbringt einen Großteil seiner Freizeit in der Natur und in Zügen, um die Welt zu sehen.</p>
							<p class="text-lg md:text-xl text-coolGray-500 font-medium">
								Neben dem <a href="/">Engineering Kiosk</a> betreibt Wolfgang mit Freunden als Side Project <a href="https://www.f-online.app/" class="underline hover:no-underline">F-Online.at</a>, die größte E-Learning Plattform für die österreichische Führerscheinprüfung.
							</p>`,
	},
	{
		slug: "christian-schepp-schaefer",
		memberKey: "schepp",
		title: 'Christian "Schepp" Schaefer, Meetup Organizer - Engineering Kiosk',
		description:
			"Freiberuflicher Frontend-Entwickler aus Düsseldorf mit Fokus auf modernem CSS, Barrierefreiheit sowie optimaler Lade- und Laufzeitperformance.",
		ogImageAlt: "Christian Schepp Schaefer",
		profileFirstName: "Christian",
		profileLastName: "Schaefer",
		jsonLdName: "Christian Schaefer",
		h1: 'Christian &bdquo;Schepp&ldquo; Schaefer',
		imagePosition: "right",
		bioHtml: `<p class="mb-6 text-lg md:text-xl text-coolGray-500 font-medium">
								Christian Schaefer, genannt &bdquo;Schepp&ldquo;, ist ein freiberuflicher Frontend-Entwickler aus Düsseldorf. Statt sich ausschließlich auf JavaScript-Frameworks zu verlassen, arbeitet er bevorzugt mit klassischen, serverseitig gerenderten, komponentenbasierten Systemen. Sein Fokus liegt auf modernem CSS, Barrierefreiheit sowie optimaler Lade- und Laufzeitperformance.
							</p>
							<p class="text-lg md:text-xl text-coolGray-500 font-medium">
								Gleichzeitig hat er einen eigenen clientseitigen Renderer entwickelt, der auf Twig.js, Web Workers, Redux und MorphDOM basiert. Neben seiner Arbeit für Kunden organisiert er Meetups wie das <a href="https://www.meetup.com/de-de/css-cafe/" class="underline hover:no-underline">CSS Café</a> und den Engineering Kiosk, richtet Konferenzen wie <a href="https://www.fronteers.nl/" class="underline hover:no-underline">Fronteers</a> aus und ist Co-Host des <a href="https://workingdraft.de/" class="underline hover:no-underline">Working Draft</a> Podcasts.
							</p>`,
	},
	{
		slug: "christoph-stanger",
		memberKey: "christoph-stanger",
		title: "Christoph Stanger, Meetup Organizer - Engineering Kiosk",
		description:
			"Software Entwickler bei Google im Cloud Run Team, begeisterter Vortragender, ehrenamtlicher Jugendleiter beim Alpenverein und leidenschaftlicher Bergsteiger.",
		ogImageAlt: "Christoph Stanger",
		profileFirstName: "Christoph",
		profileLastName: "Stanger",
		jsonLdName: null,
		h1: "Christoph Stanger",
		imagePosition: "right",
		bioHtml: `<p class="mb-6 text-lg md:text-xl text-coolGray-500 font-medium">
								Software Entwickler bei Google im Cloud Run Team und begeisterter Vortragender auf Konferenzen und Universit&auml;ten.
							</p>
							<p class="text-lg md:text-xl text-coolGray-500 font-medium">
								Ehrenamtlicher Jugendleiter beim Alpenverein und leidenschaftlicher Bergsteiger. Liebt es mit dem Camper unterwegs zu sein.
							</p>`,
	},
	{
		slug: "dario-tilgner",
		memberKey: "dario-tilgner",
		title: "Dario Tilgner, Meetup Organizer - Engineering Kiosk",
		description:
			"Head of Department IT, der bestehende Systeme kritisch hinterfragt und neben dem Job Meetups organisiert und Talks hält.",
		ogImageAlt: "Dario Tilgner",
		profileFirstName: "Dario",
		profileLastName: "Tilgner",
		jsonLdName: null,
		h1: "Dario Tilgner",
		imagePosition: "left",
		bioHtml: `<p class="mb-6 text-lg md:text-xl text-coolGray-500 font-medium">
								Im beruflichen Alltag ist Dario Head of Department IT, verantwortet mehrere IT Teams und besch&auml;ftigt sich viel damit, bestehende Systeme kritisch anzuschauen und zu entscheiden, was davon wirklich Sinn ergibt. Technologien und Architekturen werden gern hinterfragt, manchmal direkt, manchmal radikal, aber nie ohne Grund. Jeder mag coole Technologien, Dario mag die die funktionieren. &bdquo;Get the job done&ldquo; und &bdquo;you build it, you run it&ldquo; sind hier die Mottos. Dabei aber essenziell: Technik ist wichtig, Humor auch.
							</p>
							<p class="text-lg md:text-xl text-coolGray-500 font-medium">
								Neben dem Job organisiert Dario gerne Meetups, h&auml;lt selber Talks und tauscht sich in jedem Fall gerne mit Menschen aus, die neugierig sind und Dinge nicht einfach hinnehmen, nur weil sie &bdquo;schon immer so waren&ldquo;.
							</p>`,
	},
	{
		slug: "tim-hannemann",
		memberKey: "tim-hannemann",
		title: "Tim Hannemann, Meetup Organizer - Engineering Kiosk",
		description:
			"Android-Developer der ersten Stunde, Startup-Erfahrung in Berlin, verantwortet den Standortaufbau der Open Reply GmbH in Innsbruck.",
		ogImageAlt: "Tim Hannemann",
		profileFirstName: "Tim",
		profileLastName: "Hannemann",
		jsonLdName: null,
		h1: "Tim Hannemann",
		imagePosition: "left",
		bioHtml: `<p class="mb-6 text-lg md:text-xl text-coolGray-500 font-medium">
								Android-Developer der ersten Stunde lernte Tim noch w&auml;hrend seines Masterstudiums die Startup-Szene in Berlin kennen und lieben. Der Wunsch, das Internationale mit den Wurzeln zu verbinden, begleitet Tim seitdem. Bei der <a href="https://www.reply.com/open-reply-germany/de" class="underline hover:no-underline">Open Reply GmbH</a> verantwortet er unter anderem den Standortaufbau in Innsbruck.
							</p>
							<p class="text-lg md:text-xl text-coolGray-500 font-medium">
								Gerne in den Bergen, renoviert mit Freunden eine Bergh&uuml;tte, betreibt eine <a href="https://www.f-online.at" class="underline hover:no-underline">Lernplattform f&uuml;r die &ouml;sterreichische F&uuml;hrerscheinpr&uuml;fung</a> &ndash; mit Familie daheim in der Rushhour des Lebens kommt Tim aber aktuell gef&uuml;hlt zu &uuml;berhaupt nix mehr. Liebt sein Oldtimer-Rennradl und trinkt keinen Kaffee.
							</p>`,
	},
];

export default teamMemberPages;
