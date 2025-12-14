package utils

import (
	"testing"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"#-1 Wrap Up 2022 und 1. Geburtstag: Learnings, Statistiken und was 2023 geplant ist",
			"-1-wrap-up-2022-und-1-geburtstag-learnings-statistiken-und-was-2023-geplant-ist",
		},
		{
			"#00 Developer fangen bei 0 an zu zählen",
			"00-developer-fangen-bei-0-an-zu-zählen",
		},
		{
			"#01 Side Projects - Fluch oder Segen für die Karriere?",
			"01-side-projects-fluch-oder-segen-für-die-karriere",
		},
		{
			"#02 Technologienzoo Side Projects",
			"02-technologienzoo-side-projects",
		},
		{
			"#03 Over-Engineering, das Werkzeug des Teufels?",
			"03-over-engineering-das-werkzeug-des-teufels",
		},
		{
			"#04 Lohnt der Einstieg in Open Source?",
			"04-lohnt-der-einstieg-in-open-source",
		},
		{
			"#05 Team Lead - der einzige Ausweg",
			"05-team-lead-der-einzige-ausweg",
		},
		{
			"#06 Hype oder Hope: Job-Titel und Beförderungen",
			"06-hype-oder-hope-job-titel-und-beförderungen",
		},
		{
			"#07 Die Freelance Freiheit",
			"07-die-freelance-freiheit",
		},
		{
			"#08 Vergiss doch Datenbanken!",
			"08-vergiss-doch-datenbanken",
		},
		{
			"#09 Ukraine",
			"09-ukraine",
		},
		{
			"#10 Das Karriere Booster Meeting 1:1s",
			"10-das-karriere-booster-meeting-11s",
		},
		{
			"#11 Die Suche nach dem IT Traumjob",
			"11-die-suche-nach-dem-it-traumjob",
		},
		{
			"#12 Make oder Buy",
			"12-make-oder-buy",
		},
		{
			"#13 Produktivität",
			"13-produktivität",
		},
		{
			"#14 async und await: asynchrones Arbeiten im Alltag",
			"14-async-und-await-asynchrones-arbeiten-im-alltag",
		},
		{
			"#15 Source Code Kommentare, Git Commits Messages, Merge Commits und Branch-Visualisierungs-Kunst",
			"15-source-code-kommentare-git-commits-messages-merge-commits-und-branch-visualisierungs-kunst",
		},
		{
			"#16 Code Reviews: Nützlich oder bremsen nur ein gutes Team?",
			"16-code-reviews-nützlich-oder-bremsen-nur-ein-gutes-team",
		},
		{
			"#17 Was können wir beim Incident Management von der Feuerwehr lernen?",
			"17-was-können-wir-beim-incident-management-von-der-feuerwehr-lernen",
		},
		{
			"#18 Ziele und Performance-Metriken für Teams und mich selbst",
			"18-ziele-und-performance-metriken-für-teams-und-mich-selbst",
		},
		{
			"#19 Datenbank-Deepdive (oder das Ende einer Ära): von Redis bis ClickHouse",
			"19-datenbank-deepdive-oder-das-ende-einer-ära-von-redis-bis-clickhouse",
		},
		{
			"#20 Off-Boarding und On-Boarding: Wie verlasse ich eine Firma richtig?",
			"20-off-boarding-und-on-boarding-wie-verlasse-ich-eine-firma-richtig",
		},
		{
			"#21 Static Site Generators & DIE Webseite",
			"21-static-site-generators-die-webseite",
		},
		{
			"#22 NoSQL: ACID, BASE, Ende einer Ära Teil 2",
			"22-nosql-acid-base-ende-einer-ära-teil-2",
		},
		{
			"#23 Schaltest du noch oder automatisiert du schon: Home Automation",
			"23-schaltest-du-noch-oder-automatisiert-du-schon-home-automation",
		},
		{
			"#24 Infrastructure as Code oder old man yells at cloud",
			"24-infrastructure-as-code-oder-old-man-yells-at-cloud",
		},
		{
			"#25 Tech-Entlassungswellen & Job-Interview Skills",
			"25-tech-entlassungswellen-job-interview-skills",
		},
		{
			"#26 My English is not the yellow from the egg - Arbeiten in internationalen Teams",
			"26-my-english-is-not-the-yellow-from-the-egg-arbeiten-in-internationalen-teams",
		},
		{
			"#27 Sicherheit in der Dependency Hölle",
			"27-sicherheit-in-der-dependency-hölle",
		},
		{
			"#28 O(1), O(log n), O(n^2) - Ist die Komplexität von Algorithmen im Entwickler-Alltag relevant?",
			"28-o1-olog-n-on2-ist-die-komplexität-von-algorithmen-im-entwickler-alltag-relevant",
		},
		{
			"#29 Die andere Seite: Meetups & Konferenzen organisieren",
			"29-die-andere-seite-meetups-konferenzen-organisieren",
		},
		{
			"#30 Ist ein Informatikstudium sinnvoll? Welche Ausbildung für Devs?",
			"30-ist-ein-informatikstudium-sinnvoll-welche-ausbildung-für-devs",
		},
		{
			"#31 Ich automatisiere mir die Welt wie sie mir gefällt (mit GitHub Actions)",
			"31-ich-automatisiere-mir-die-welt-wie-sie-mir-gefällt-mit-github-actions",
		},
		{
			"#32 Die richtigen Leute anstellen: Die Recruiting Pipeline",
			"32-die-richtigen-leute-anstellen-die-recruiting-pipeline",
		},
		{
			"#33 Andy im Team Lead Bewerbungsgespräch",
			"33-andy-im-team-lead-bewerbungsgespräch",
		},
		{
			"#34 Wie cloudy bist du?",
			"34-wie-cloudy-bist-du",
		},
		{
			"#35 Knowledge Sharing oder die Person, die nie \"gehen\" sollte...",
			"35-knowledge-sharing-oder-die-person-die-nie-gehen-sollte",
		},
		{
			"#36 Sechs-stellige IT-Gehälter? Wie? Was? Wo? Fair?",
			"36-sechs-stellige-it-gehälter-wie-was-wo-fair",
		},
		{
			"#37 Mit IT-Büchern Geld verdienen? Wer liest überhaupt noch Bücher?",
			"37-mit-it-büchern-geld-verdienen-wer-liest-überhaupt-noch-bücher",
		},
		{
			"#38 Monitoring, Metriken, Tracing, Alerting, Observability",
			"38-monitoring-metriken-tracing-alerting-observability",
		},
		{
			"#39 Gemischte Tüte: Software Engineer, Github, OpenSource, Git und HomeOffice",
			"39-gemischte-tüte-software-engineer-github-opensource-git-und-homeoffice",
		},
		{
			"#40 Wie wird man und Frau zum Senior Dev?",
			"40-wie-wird-man-und-frau-zum-senior-dev",
		},
		{
			"#41 SQL Injections - Ein unterschätztes Risiko",
			"41-sql-injections-ein-unterschätztes-risiko",
		},
		{
			"#42 Lexer, Parser und Open Source in Counterstrike",
			"42-lexer-parser-und-open-source-in-counterstrike",
		},
		{
			"#43 Cloud vs. On-Premise: Die Entscheidungshilfe",
			"43-cloud-vs-on-premise-die-entscheidungshilfe",
		},
		{
			"#44 Der Weg zum hochperformanten Team",
			"44-der-weg-zum-hochperformanten-team",
		},
		{
			"#45 Datengetriebene Entscheidungen und der perfekte Dashboard Stack",
			"45-datengetriebene-entscheidungen-und-der-perfekte-dashboard-stack",
		},
		{
			"#46 Welches Problem löst Docker?",
			"46-welches-problem-löst-docker",
		},
		{
			"#47 Wer Visionen hat, soll zum Arzt!?",
			"47-wer-visionen-hat-soll-zum-arzt",
		},
		{
			"#48 Der Layer unter Docker: containerd, Kubernetes, Container Runtime Interface, CRI-O und Open Container Initiative (OCI)",
			"48-der-layer-unter-docker-containerd-kubernetes-container-runtime-interface-cri-o-und-open-container-initiative-oci",
		},
		{
			"#49 Die Zukunft: Software Repository Mining",
			"49-die-zukunft-software-repository-mining",
		},
		{
			"#50 DEI erhöht Innovation und den finanziellen Erfolg",
			"50-dei-erhöht-innovation-und-den-finanziellen-erfolg",
		},
		{
			"#51 Was ist das Staff (Engineer) Level?",
			"51-was-ist-das-staff-engineer-level",
		},
		{
			"#52 Asynchrone Verarbeitung durch Message Queues - Vor- und Nachteile",
			"52-asynchrone-verarbeitung-durch-message-queues-vor-und-nachteile",
		},
		{
			"#53 Cloud / NoCode/ AI / ChatGPT ersetzen unsere Jobs?",
			"53-cloud-nocode-ai-chatgpt-ersetzen-unsere-jobs",
		},
		{
			"#54 Key Value Store Redis: Einsatzmöglichkeiten, Fallstricke, Datenstrukturen, HyperLogLog und (flüchtige) Persistenz",
			"54-key-value-store-redis-einsatzmöglichkeiten-fallstricke-datenstrukturen-hyperloglog-und-flüchtige-persistenz",
		},
		{
			"#55 Weiterbildung: Zertifizierung, Newsletter, Konferenzen, ... Wie? Warum? Und wer zahlt das Ganze?",
			"55-weiterbildung-zertifizierung-newsletter-konferenzen-wie-warum-und-wer-zahlt-das-ganze",
		},
		{
			"#56 Applikations-Skalierung: Wann, wieso, was kostet es? Stateless und Stateful, Horizontal vs. Vertikal",
			"56-applikations-skalierung-wann-wieso-was-kostet-es-stateless-und-stateful-horizontal-vs-vertikal",
		},
		{
			"#57 Servant Leadership: Führungsstil der neuen Generation?",
			"57-servant-leadership-führungsstil-der-neuen-generation",
		},
		{
			"#58 Software-Updates, alte Software, Long Term Support und End of Life-Dates",
			"58-software-updates-alte-software-long-term-support-und-end-of-life-dates",
		},
		{
			"#59 Kann man mit Open Source Geld verdienen?",
			"59-kann-man-mit-open-source-geld-verdienen",
		},
		{
			"#60 On-Call: Warum auch Software-Engineers auf Rufbereitschaft sein sollten",
			"60-on-call-warum-auch-software-engineers-auf-rufbereitschaft-sein-sollten",
		},
		{
			"#61 Schwierige 1-on-1 Situationen und Lösungsvorschläge",
			"61-schwierige-1-on-1-situationen-und-lösungsvorschläge",
		},
		{
			"#62 Technologien konsolidieren, oder wie Startups sammeln?",
			"62-technologien-konsolidieren-oder-wie-startups-sammeln",
		},
		{
			"#63 Spaß mit Zahlen: Under- und Overflows, Rückwärtslaufende Zeit, Negative Modulos und Währungsbeträge",
			"63-spaß-mit-zahlen-under-und-overflows-rückwärtslaufende-zeit-negative-modulos-und-währungsbeträge",
		},
		{
			"#64 Infrastruktur-Bingo: Forward-, Reverse-, SOCKS-Proxy, Load Balancing und gibt es einen Unterschied zwischen Load-Balancer und Reverse-Proxy?",
			"64-infrastruktur-bingo-forward-reverse-socks-proxy-load-balancing-und-gibt-es-einen-unterschied-zwischen-load-balancer-und-reverse-proxy",
		},
		{
			"#65 Clean Code macht Software langsam",
			"65-clean-code-macht-software-langsam",
		},
		{
			"#66 Stressfreie Produktivität in der hektischen Welt mit Getting Things Done",
			"66-stressfreie-produktivität-in-der-hektischen-welt-mit-getting-things-done",
		},
		{
			"#67 Die Netz-Entlastung des Internets: Content Delivery Networks (CDNs)",
			"67-die-netz-entlastung-des-internets-content-delivery-networks-cdns",
		},
		{
			"#68 Im \"Flow\" und Deepwork mit Kirill Sivy",
			"68-im-flow-und-deepwork-mit-kirill-sivy",
		},
		{
			"#69 MySQL vs. MariaDB",
			"69-mysql-vs-mariadb",
		},
		{
			"#70 Alan Turing: Der Vater der heutigen Informatik (Turing-Complete, Turing-Test, Halting-Problem, Turing-Maschine, Captcha)",
			"70-alan-turing-der-vater-der-heutigen-informatik-turing-complete-turing-test-halting-problem-turing-maschine-captcha",
		},
		{
			"#71 Tim Berners-Lee: Was ist das World Wide Web und was ist seine Zukunft?",
			"71-tim-berners-lee-was-ist-das-world-wide-web-und-was-ist-seine-zukunft",
		},
		{
			"#72 Meetings: Jeder hat sie, keiner will sie",
			"72-meetings-jeder-hat-sie-keiner-will-sie",
		},
		{
			"#73 Cache-freundliches Programmieren, CPU-Caches, Ersetzungsstrategien und Cache-Invalidierung",
			"73-cache-freundliches-programmieren-cpu-caches-ersetzungsstrategien-und-cache-invalidierung",
		},
		{
			"#74 REST: Das oft falsch verstandene Architektur-Paradigma",
			"74-rest-das-oft-falsch-verstandene-architektur-paradigma",
		},
		{
			"#75 Evaluierung deiner Job-Performance, Team-Feedback und LNO Framework",
			"75-evaluierung-deiner-job-performance-team-feedback-und-lno-framework",
		},
		{
			"#76 Mit Open Source 100.000$ verdienen, Sponsorware und Plattform-Risiken bei GitHub mit Martin Donath",
			"76-mit-open-source-100000-verdienen-sponsorware-und-plattform-risiken-bei-github-mit-martin-donath",
		},
		{
			"#77 Kinder, Coding und AI: Die Zukunft der Informatik-Bildung mit Diana Knodel",
			"77-kinder-coding-und-ai-die-zukunft-der-informatik-bildung-mit-diana-knodel",
		},
		{
			"#78 Microservice & Monolith: Was die Industrie in den letzten 9 Jahren gelernt hat",
			"78-microservice-monolith-was-die-industrie-in-den-letzten-9-jahren-gelernt-hat",
		},
		{
			"#79 Font-Engineering und Schriftarten fürs Programmieren mit Philipp Acsany",
			"79-font-engineering-und-schriftarten-fürs-programmieren-mit-philipp-acsany",
		},
		{
			"#80 Junior Devs: Steckt das wahre Potential in unerfahrenen Talenten?",
			"80-junior-devs-steckt-das-wahre-potential-in-unerfahrenen-talenten",
		},
		{
			"#81 Copilot & AI im Dev-Test: Produktivitäts-Boost oder nur Hype?",
			"81-copilot-ai-im-dev-test-produktivitäts-boost-oder-nur-hype",
		},
		{
			"#82 Hinter den Kulissen: Die Informatik-Doktorarbeit und ist der Dr. Titel in der heutigen IT-Welt noch relevant?",
			"82-hinter-den-kulissen-die-informatik-doktorarbeit-und-ist-der-dr-titel-in-der-heutigen-it-welt-noch-relevant",
		},
		{
			"#83 Transparenz im Tech-Leadership & Fehlerkultur: Wie weit kann ich gehen?",
			"83-transparenz-im-tech-leadership-fehlerkultur-wie-weit-kann-ich-gehen",
		},
		{
			"#84 Die Evolution von JavaScript: Vom Ducktyping zum Monopol im Browser mit Peter Kröner",
			"84-die-evolution-von-javascript-vom-ducktyping-zum-monopol-im-browser-mit-peter-kröner",
		},
		{
			"#85 Von Entwicklerin zur Engineering Managerin: Erfahrungen und Learnings mit Isabelle Glasmacher",
			"85-von-entwicklerin-zur-engineering-managerin-erfahrungen-und-learnings-mit-isabelle-glasmacher",
		},
		{
			"#86 Open Source als Herz einer Firma mit Nextcloud Gründer Frank Karlitschek",
			"86-open-source-als-herz-einer-firma-mit-nextcloud-gründer-frank-karlitschek",
		},
		{
			"#87 Die DORA-Metriken: Ist Software-Entwicklungs-Performance messbar?",
			"87-die-dora-metriken-ist-software-entwicklungs-performance-messbar",
		},
		{
			"#88 Die Personalabteilung: Freund oder Feind? mit Patrick Kuster",
			"88-die-personalabteilung-freund-oder-feind-mit-patrick-kuster",
		},
		{
			"#89 Die Klimakrise und Green IT: unser Einfluss über Hardware, Farben, Web-Performance und Green-Hosting mit Christian Schäfer",
			"89-die-klimakrise-und-green-it-unser-einfluss-über-hardware-farben-web-performance-und-green-hosting-mit-christian-schäfer",
		},
		{
			"#90 Inner Source: Open Source Best Practices zur besseren Zusammenarbeit zwischen Teams mit Sebastian Spier",
			"90-inner-source-open-source-best-practices-zur-besseren-zusammenarbeit-zwischen-teams-mit-sebastian-spier",
		},
		{
			"#91 Konsistent, Verfügbar, Ausfalltolerant oder Performant: Das CAP- und PACELC-Theorem in verteilten Systemen",
			"91-konsistent-verfügbar-ausfalltolerant-oder-performant-das-cap-und-pacelc-theorem-in-verteilten-systemen",
		},
		{
			"#92 Technologie trifft Deutsche Ausbildungskultur: Die moderne IT-Berufsausbildung mit Stefan Macke",
			"92-technologie-trifft-deutsche-ausbildungskultur-die-moderne-it-berufsausbildung-mit-stefan-macke",
		},
		{
			"#93 Barbara Liskov - Das L in SOLID (Liskovsches Substitutionsprinzip & Abstraktion)",
			"93-barbara-liskov-das-l-in-solid-liskovsches-substitutionsprinzip-abstraktion",
		},
		{
			"#94 Die Realität des Freelancings: Zwischen Selbstbestimmung und Unsicherheit mit Index out of bounds",
			"94-die-realität-des-freelancings-zwischen-selbstbestimmung-und-unsicherheit-mit-index-out-of-bounds",
		},
		{
			"#95 Effiziente Knowledge Sharing Formate: Wissen teilen und begeistern",
			"95-effiziente-knowledge-sharing-formate-wissen-teilen-und-begeistern",
		},
		{
			"#96 Selbstgemacht vs. Fertigprodukt: Ein Blick auf das Not-Invented-Here-Phänomen",
			"96-selbstgemacht-vs-fertigprodukt-ein-blick-auf-das-not-invented-here-phänomen",
		},
		{
			"#97 Metriken, Hypothesen und Fehler: A/B-Testing in der Praxis mit Philipp Monreal",
			"97-metriken-hypothesen-und-fehler-ab-testing-in-der-praxis-mit-philipp-monreal",
		},
		{
			"#98 Der Hype um Rust mit Matthias Endler",
			"98-der-hype-um-rust-mit-matthias-endler",
		},
		{
			"#99 Modernes SQL ist mehr als SELECT * FROM - mit Markus Winand",
			"99-modernes-sql-ist-mehr-als-select-from-mit-markus-winand",
		},
		{
			"#100 Episoden: ein Tech Rückblick auf 2022/23, Predictions 2024 und viel Tech Trivia",
			"100-episoden-ein-tech-rückblick-auf-202223-predictions-2024-und-viel-tech-trivia",
		},
		{
			"#101 Observability und OpenTelemetry mit Severin Neumann",
			"101-observability-und-opentelemetry-mit-severin-neumann",
		},
		{
			"#102 Quereinstieg in die Software-Entwicklung mit Melanie Patrick",
			"102-quereinstieg-in-die-software-entwicklung-mit-melanie-patrick",
		},
		{
			"#103 Plattform Engineering und Interne Developer Plattformen mit Puja Abbassi",
			"103-plattform-engineering-und-interne-developer-plattformen-mit-puja-abbassi",
		},
		{
			"#104 Präsentieren mit Wirkung: Public Speaking und Storytelling mit Anna Momber",
			"104-präsentieren-mit-wirkung-public-speaking-und-storytelling-mit-anna-momber",
		},
		{
			"#105 Cloud-Ausfallsicherheit: Die Realität von Regionen und Availability Zones",
			"105-cloud-ausfallsicherheit-die-realität-von-regionen-und-availability-zones",
		},
		{
			"#106 CI - Continuous Integration in der Praxis mit Michael Lihs von Thoughtworks",
			"106-ci-continuous-integration-in-der-praxis-mit-michael-lihs-von-thoughtworks",
		},
		{
			"#107 Entwickler-Alltag: Die \"bösen\" Ablenkungen und das ewige Leiden mit dem Fokus",
			"107-entwickler-alltag-die-bösen-ablenkungen-und-das-ewige-leiden-mit-dem-fokus",
		},
		{
			"#108 Agile Multi-Team Projekte: Die Kunst, hunderte Leute effektiv zu koordinieren mit Stephan Strack",
			"108-agile-multi-team-projekte-die-kunst-hunderte-leute-effektiv-zu-koordinieren-mit-stephan-strack",
		},
		{
			"#109 Freeze! Warum dein Code manchmal eine Pause braucht",
			"109-freeze-warum-dein-code-manchmal-eine-pause-braucht",
		},
		{
			"#110 OKRs und Beyond: Agile Unternehmensführung mit Marco Alberti von Murakamy",
			"110-okrs-und-beyond-agile-unternehmensführung-mit-marco-alberti-von-murakamy",
		},
		{
			"#111 Side-Projects: Zwei Entwickler overengineeren einen Podcast",
			"111-side-projects-zwei-entwickler-overengineeren-einen-podcast",
		},
		{
			"#112 Das Engineering Manager Pendulum: Zwischen Coding und Leadership mit Tom Bartel",
			"112-das-engineering-manager-pendulum-zwischen-coding-und-leadership-mit-tom-bartel",
		},
		{
			"#113 Selbstmarketing ohne Bullshit: Brag Documents",
			"113-selbstmarketing-ohne-bullshit-brag-documents",
		},
		{
			"#114 Sales Engineers: Engineering und Sales in einer Person vereint mit Patrick Pissang",
			"114-sales-engineers-engineering-und-sales-in-einer-person-vereint-mit-patrick-pissang",
		},
		{
			"#115 Die Shift Left Philosophie: Mehr Verantwortung für Devs",
			"115-die-shift-left-philosophie-mehr-verantwortung-für-devs",
		},
		{
			"#116 KI unterstützte Software Entwicklung: Ein Reality Check mit Birgitta Böckeler von Thoughtworks",
			"116-ki-unterstützte-software-entwicklung-ein-reality-check-mit-birgitta-böckeler-von-thoughtworks",
		},
		{
			"#117 Vanilla Web: Niedrige Kopplung & hohe Kohäsion mit Golo Roden von the native web",
			"117-vanilla-web-niedrige-kopplung-hohe-kohäsion-mit-golo-roden-von-the-native-web",
		},
		{
			"#118 Wie funktioniert eine moderne Suche? Von Indexierung bis Ranking",
			"118-wie-funktioniert-eine-moderne-suche-von-indexierung-bis-ranking",
		},
		{
			"#119 Der Jobwechsel: Einblick und Erfahrungsaustausch mit UNMUTE IT",
			"119-der-jobwechsel-einblick-und-erfahrungsaustausch-mit-unmute-it",
		},
		{
			"#120 No-Code ist technische Schuld!",
			"120-no-code-ist-technische-schuld",
		},
		{
			"#121 YAML: Mehr als Konfiguration! Aliases, Tags und YAMLScript mit Tina Müller",
			"121-yaml-mehr-als-konfiguration-aliases-tags-und-yamlscript-mit-tina-müller",
		},
		{
			"#122 Ich hasse Re-Orgs",
			"122-ich-hasse-re-orgs",
		},
		{
			"#123 The Bread Code: vom Entwickler zum Brot-Influencer mit Hendrik Kleinwächter",
			"123-the-bread-code-vom-entwickler-zum-brot-influencer-mit-hendrik-kleinwächter",
		},
		{
			"#124 Technische Glaubwürdigkeit bewahren: Müssen Leads den Code kennen?",
			"124-technische-glaubwürdigkeit-bewahren-müssen-leads-den-code-kennen",
		},
		{
			"#125 Die Kunst der technischen Dokumentation mit Jana Aydinbas",
			"125-die-kunst-der-technischen-dokumentation-mit-jana-aydinbas",
		},
		{
			"#126 Killing the Mutant: Teststrategien mit Sebastian Bergmann",
			"126-killing-the-mutant-teststrategien-mit-sebastian-bergmann",
		},
		{
			"#127 Imposter-Syndrom & Peter-Prinzip mit Dr. Fanny Jimenez",
			"127-imposter-syndrom-peter-prinzip-mit-dr-fanny-jimenez",
		},
		{
			"#128 Devs müssen wissenschaftliche Papers lesen!?",
			"128-devs-müssen-wissenschaftliche-papers-lesen",
		},
		{
			"#129 Simplify Your Stack: Files statt Datenbanken!",
			"129-simplify-your-stack-files-statt-datenbanken",
		},
		{
			"#130 Wie gutes UX-Design entsteht mit Robin Titus",
			"130-wie-gutes-ux-design-entsteht-mit-robin-titus",
		},
		{
			"#131 Equity in Tech-Startups: Mehr als nur Gehalt mit Philipp \"Pip\" Klöckner",
			"131-equity-in-tech-startups-mehr-als-nur-gehalt-mit-philipp-pip-klöckner",
		},
		{
			"#132 Prometheus: Revolution im Monitoring mit Mitbegründer Julius Volz",
			"132-prometheus-revolution-im-monitoring-mit-mitbegründer-julius-volz",
		},
		{
			"#133 Die wichtige Rolle von 1on1s in Zeiten der Arbeiterlosigkeit",
			"133-die-wichtige-rolle-von-1on1s-in-zeiten-der-arbeiterlosigkeit",
		},
		{
			"#134 Wir profitieren, sie leiden: Die Schattenseiten von Open Source",
			"134-wir-profitieren-sie-leiden-die-schattenseiten-von-open-source",
		},
		{
			"#135 Design Documents & RFCs: Der Weg zu besserer Software-Architektur",
			"135-design-documents-rfcs-der-weg-zu-besserer-software-architektur",
		},
		{
			"#136 Als Knowledge Worker fit und gesund bleiben mit Patrick Cole",
			"136-als-knowledge-worker-fit-und-gesund-bleiben-mit-patrick-cole",
		},
		{
			"#137 Die Schaltsekunde und ihre IT-Folgen: Ein Sekundenbruchteil mit Impact",
			"137-die-schaltsekunde-und-ihre-it-folgen-ein-sekundenbruchteil-mit-impact",
		},
		{
			"#138 Gemeinsam stark: Jobsharing und Tandems in der modernen Arbeitswelt mit Anna Drüing-Schlüter",
			"138-gemeinsam-stark-jobsharing-und-tandems-in-der-modernen-arbeitswelt-mit-anna-drüing-schlüter",
		},
		{
			"#139 Security Engineering und Hacking-Wettbewerbe mit Frederik Braun von Mozilla",
			"139-security-engineering-und-hacking-wettbewerbe-mit-frederik-braun-von-mozilla",
		},
		{
			"#140 Tech-Leadership: Die technische Vision als Leitfaden für Teams",
			"140-tech-leadership-die-technische-vision-als-leitfaden-für-teams",
		},
		{
			"#141 Datenjournalismus - zwischen Grafik und Fakten mit Michael Kreil",
			"141-datenjournalismus-zwischen-grafik-und-fakten-mit-michael-kreil",
		},
		{
			"#142 Ist Return to Office die Zukunft? Was die Wissenschaft sagt - mit Jean-Victor Alipour vom IFO",
			"142-ist-return-to-office-die-zukunft-was-die-wissenschaft-sagt-mit-jean-victor-alipour-vom-ifo",
		},
		{
			"#143 Ship It! Deployment-Strategien und Anti-Patterns auf der letzten Meile",
			"143-ship-it-deployment-strategien-und-anti-patterns-auf-der-letzten-meile",
		},
		{
			"#144 Die unterschätzte Macht der Zeit: Wie NTP und PTP die Welt synchronisieren mit Daniel Boldt und Thomas Behn von Meinberg",
			"144-die-unterschätzte-macht-der-zeit-wie-ntp-und-ptp-die-welt-synchronisieren-mit-daniel-boldt-und-thomas-behn-von-meinberg",
		},
		{
			"#145 Große Open Source Projekte managen: 20 Jahre im TYPO3 Projekt mit Benni Mack",
			"145-große-open-source-projekte-managen-20-jahre-im-typo3-projekt-mit-benni-mack",
		},
		{
			"#146 Warum ist Doom so faszinierend für die Software-Entwicklung?",
			"146-warum-ist-doom-so-faszinierend-für-die-software-entwicklung",
		},
		{
			"#147 Mechanische Tastaturen: Vom Klick zum Kult mit Philipp Hoeler-Lutz von Click! Clack! Hack!",
			"147-mechanische-tastaturen-vom-klick-zum-kult-mit-philipp-hoeler-lutz-von-click-clack-hack",
		},
		{
			"#148 Wenn Open Source eigene Wege geht: Forking und seine Folgen",
			"148-wenn-open-source-eigene-wege-geht-forking-und-seine-folgen",
		},
		{
			"#149 Recommender Systems: Funktionsweise und Forschungstrends mit Eva Zangerle",
			"149-recommender-systems-funktionsweise-und-forschungstrends-mit-eva-zangerle",
		},
		{
			"#150 Die ThinkPad-Faszination: Technik, Design und Nostalgie mit Christian Stankowic vom ThinkPad Museum",
			"150-die-thinkpad-faszination-technik-design-und-nostalgie-mit-christian-stankowic-vom-thinkpad-museum",
		},
		{
			"#151 Räumliche Indexstrukturen: Grundpfeiler in Geo-Systemen, Games und Machine Learning",
			"151-räumliche-indexstrukturen-grundpfeiler-in-geo-systemen-games-und-machine-learning",
		},
		{
			"#152 Warum i und j als Zählvariablen genutzt werden",
			"152-warum-i-und-j-als-zählvariablen-genutzt-werden",
		},
		{
			"#153 Wie hoste ich ein Large Language Modell (LLM) mit Kubernetes in 5 Minuten mit Data Science Deep Dive",
			"153-wie-hoste-ich-ein-large-language-modell-llm-mit-kubernetes-in-5-minuten-mit-data-science-deep-dive",
		},
		{
			"#154 Architektur-Diskussion: Design eines einfachen und robusten Preis-Scrapers",
			"154-architektur-diskussion-design-eines-einfachen-und-robusten-preis-scrapers",
		},
		{
			"#155 Cursor AI mit der programmier.bar",
			"155-cursor-ai-mit-der-programmierbar",
		},
		{
			"#156 Inbox Zero: der Pro-Tipp für deine Produktivität",
			"156-inbox-zero-der-pro-tipp-für-deine-produktivität",
		},
		{
			"#157 Agile Arbeitsmethoden - Extreme Programming mit den Coding Buddies",
			"157-agile-arbeitsmethoden-extreme-programming-mit-den-coding-buddies",
		},
		{
			"#158 Zykel-Erkennung in einer Linked List",
			"158-zykel-erkennung-in-einer-linked-list",
		},
		{
			"#159 Verhaltensbezogene Interview-Fragen und STAR-Methode",
			"159-verhaltensbezogene-interview-fragen-und-star-methode",
		},
		{
			"#160 Grace Hopper mit UNMUTE IT",
			"160-grace-hopper-mit-unmute-it",
		},
		{
			"#161 Sichere Daten trotz physischem Zugriff: Disk Encryption und Integritätsschutz von Laptops bis IoT-Devices mit David Gstir von sigma star",
			"161-sichere-daten-trotz-physischem-zugriff-disk-encryption-und-integritätsschutz-von-laptops-bis-iot-devices-mit-david-gstir-von-sigma-star",
		},
		{
			"#162 Event Sourcing & Märchen mit Golo Roden von the native web",
			"162-event-sourcing-märchen-mit-golo-roden-von-the-native-web",
		},
		{
			"#163 Benevolent Dictator for Life (BDFL)",
			"163-benevolent-dictator-for-life-bdfl",
		},
		{
			"#164 Suchalgorithmen: Lineare- und Binäre Suche & Indizes mit Stefan Macke vom IT Berufe Podcast",
			"164-suchalgorithmen-lineare-und-binäre-suche-indizes-mit-stefan-macke-vom-it-berufe-podcast",
		},
		{
			"#165 Pessimistisches und Optimistisches Sperren in Datenbanken: Eine Geschichte",
			"165-pessimistisches-und-optimistisches-sperren-in-datenbanken-eine-geschichte",
		},
		{
			"#166 Das Fediverse mit Christian Stankowic vom Focus on Linux Podcast",
			"166-das-fediverse-mit-christian-stankowic-vom-focus-on-linux-podcast",
		},
		{
			"#167 Tabs vs. Spaces mit dem Index out of bounds Podcast",
			"167-tabs-vs-spaces-mit-dem-index-out-of-bounds-podcast",
		},
		{
			"#168 Beyond Learning Budgets: Was Team-Entwicklung wirklich braucht",
			"168-beyond-learning-budgets-was-team-entwicklung-wirklich-braucht",
		},
		{
			"#169 Deno, die Alternative zu Node.js mit der programmier.bar",
			"169-deno-die-alternative-zu-nodejs-mit-der-programmierbar",
		},
		{
			"#170 - 404 Not Found!",
			"170-404-not-found",
		},
		{
			"#171 Vergiss deine Maus mit Philipp Hoeler-Lutz von Click! Clack! Hack!",
			"171-vergiss-deine-maus-mit-philipp-hoeler-lutz-von-click-clack-hack",
		},
		{
			"#172 Die kuriosesten Versionsnummern bekannter Software mit Matthias Endler von Rust in Production",
			"172-die-kuriosesten-versionsnummern-bekannter-software-mit-matthias-endler-von-rust-in-production",
		},
		{
			"#173 Rekursion: Das Ende ist nah!",
			"173-rekursion-das-ende-ist-nah",
		},
		{
			"#174 Frontend-Engineering Metriken im Team einführen mit dem Working Draft Podcast",
			"174-frontend-engineering-metriken-im-team-einführen-mit-dem-working-draft-podcast",
		},
		{
			"#175 Von Lustig bis Traurig: Wenn Open Source Geschichten schreibt",
			"175-von-lustig-bis-traurig-wenn-open-source-geschichten-schreibt",
		},
		{
			"#176 Der Engineering Kiosk wird 3 Jahre alt!",
			"176-der-engineering-kiosk-wird-3-jahre-alt",
		},
		{
			"#177 Stream Processing & Kafka: Die Basis moderner Datenpipelines mit Stefan Sprenger",
			"177-stream-processing-kafka-die-basis-moderner-datenpipelines-mit-stefan-sprenger",
		},
		{
			"#178 Code der bewegt: Infotainmentsysteme auf Kreuzfahrtschiffen mit Sebastian Hammerl",
			"178-code-der-bewegt-infotainmentsysteme-auf-kreuzfahrtschiffen-mit-sebastian-hammerl",
		},
		{
			"#179 MLOps: Machine Learning in die Produktion bringen mit Michelle Golchert und Sebastian Warnholz",
			"179-mlops-machine-learning-in-die-produktion-bringen-mit-michelle-golchert-und-sebastian-warnholz",
		},
		{
			"#180 Skalierung, aber zu welchem Preis? (Papers We Love)",
			"180-skalierung-aber-zu-welchem-preis-papers-we-love",
		},
		{
			"#181 Von Code zu Value: Wie Entwickler·innen Business-Mehrwert schaffen",
			"181-von-code-zu-value-wie-entwicklerinnen-business-mehrwert-schaffen",
		},
		{
			"#182 Happy Birthday SQL: 50 Jahre Abfragesprache",
			"182-happy-birthday-sql-50-jahre-abfragesprache",
		},
		{
			"#183 Event-Sourcing: Die intelligente Datenarchitektur mit semantischer Historie – mit Golo Roden",
			"183-event-sourcing-die-intelligente-datenarchitektur-mit-semantischer-historie-mit-golo-roden",
		},
		{
			"#184 GPU Programmierung - von CUDA bis OpenMP mit Peter Thoman",
			"184-gpu-programmierung-von-cuda-bis-openmp-mit-peter-thoman",
		},
		{
			"#185 Der Mainframe ist tot, lang lebe der Mainframe! Von COBOL bis JavaScript am Mainframe mit Tobias Leicher von IBM",
			"185-der-mainframe-ist-tot-lang-lebe-der-mainframe-von-cobol-bis-javascript-am-mainframe-mit-tobias-leicher-von-ibm",
		},
		{
			"#186 Von CNC-Fräse bis RFID-Tag: Wenn Informatik zur Kunst wird mit Sabine Wieluch aka Bleeptrack",
			"186-von-cnc-fräse-bis-rfid-tag-wenn-informatik-zur-kunst-wird-mit-sabine-wieluch-aka-bleeptrack",
		},
		{
			"#187 Meeresschutz mit Code – Sea Shepherds Tech-Einsatz mit Florian Stadler",
			"187-meeresschutz-mit-code-sea-shepherds-tech-einsatz-mit-florian-stadler",
		},
		{
			"#188 Spieleentwicklung: Die Königsdisziplin der Informatik mit Dominic Pacher",
			"188-spieleentwicklung-die-königsdisziplin-der-informatik-mit-dominic-pacher",
		},
		{
			"#189 Fuzzing: Wenn der Zufall dein bester Tester ist mit Prof. Andreas Zeller",
			"189-fuzzing-wenn-der-zufall-dein-bester-tester-ist-mit-prof-andreas-zeller",
		},
		{
			"#190 Mehr Meetings, mehr Macht? Der Weg zur Tech-Führungskraft",
			"190-mehr-meetings-mehr-macht-der-weg-zur-tech-führungskraft",
		},
		{
			"#191 Graphdatenbanken: von GraphRAG bis Cypher mit Michael Hunger von Neo4j",
			"191-graphdatenbanken-von-graphrag-bis-cypher-mit-michael-hunger-von-neo4j",
		},
		{
			"#196 Star Wars auf GitHub: 4,5 Mio. Fake-Sterne entdeckt",
			"196-star-wars-auf-github-45-mio-fake-sterne-entdeckt",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := Slugify(test.input)

			if result != test.expected {
				t.Errorf("Expected %q, but got %q", test.expected, result)
			}
		})
	}
}
