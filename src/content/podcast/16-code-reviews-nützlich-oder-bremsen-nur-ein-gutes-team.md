---
advertiser: ""
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/a5a91d87-c04e-43f1-a2dd-9f137a4fa0ef/engineering-kiosk-16-code-reviews-n%C3%BCtzlich-oder-bremsen-nur-ein-gutes-team
apple_podcasts: https://podcasts.apple.com/us/podcast/16-code-reviews-n%C3%BCtzlich-oder-bremsen-nur-ein-gutes-team/id1603082924?i=1000558699631
audio: https://audio1.redcircle.com/episodes/7c477eeb-e8a1-43e3-9663-8c6d9118645f/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:01:23
  title: Amsterdam, Hipster und Craft-Bier
- start: 00:02:04
  title: git rebase, history rewrite, force push und Kollaboration
- start: 00:04:54
  title: Saubere git history und Merge-Commits
- start: 00:05:17
  title: Code Reviews
- start: 00:02:05
  title: Was ist denn ein Code Review?
- start: 00:07:02
  title: Muss der Code Reviewer immer mehr Erfahrung haben?
- start: 00:08:10
  title: Was wir von Juniors im Job und bei Code Reviews lernen können
- start: 00:09:37
  title: Was ist denn ein Google Fellow?
- start: 00:10:37
  title: Code Reviews als Mittel fürs Onboarding
- start: 00:11:35
  title: Frühes pushen der Arbeit ist wichtig für frühes Feedback
- start: 00:13:15
  title: Wo sollen die Informationen zu meiner Änderung stehen? Im Pull Request oder in der Git Commit message?
- start: 00:15:04
  title: Wen lade ich zu meinem Code Review ein? Code Owners file und Software Repository Mining
- start: 00:17:34
  title: Worauf soll ich als Code Reviewer acht geben und konzentrieren?
- start: 00:20:09
  title: Pro-Aktivität in Code Reviews mit aktiven Vorschlägen und Fragen stellen
- start: 00:23:33
  title: Muss man jedes Feedback aus dem Code Review annehmen? Lass das Ego vor der Tür und Gate-Keeper
- start: 00:26:37
  title: "Kommunikationskonflikte in Code Reviews: Video Calls vs. geschriebene Kommunikation"
- start: 00:28:14
  title: Nitpicking, Code-Formatting, Styleguides, Linter und Automatisierung
- start: 00:32:11
  title: Mindern Pull Requests und Code Reviews die Performance von Teams?
- start: 00:36:56
  title: "Schnelligkeit bei Facebook und Booking: Schnelles Mergen und Code Reviews danach"
- start: 00:39:01
  title: Auslagerung der Verantwortung bei Code Reviews und Blameless Kultur innerhalb der Firma
- start: 00:42:38
  title: "Technische Mittel für sichere Deploys: Feature Flags, Canary Deployments, Dark Launches und Blue-Green-Deployments"
- start: 00:46:08
  title: Wie schnell sollte ein Code Review vom Team gereviewed werden?
- start: 00:53:00
  title: Code Reviews in der Firma vs. Code Reviews in Open Source und die Motivation
- start: 00:54:34
  title: Konsequenzen bei nicht eingehaltenen Service Level Agreements (SLA)
- start: 00:57:28
  title: Outro
deezer: https://www.deezer.com/de/episode/380202837
description: "Code Reviews: Jeder will schnelles Feedback, doch niemand hat Zeit dafür - Eine Hassliebe. Eine Komponente im Alltag jedes Software Engineers. Egal ob Junior, Senior oder Staff-Engineer. Jeder erstellt Code Reviews und kommentiert die Arbeit von den Kollegen. Doch wie sehen gute Code Reviews aus? Was gehört hinein, was bleibt besser draußen? Wie viel Reviewer machen Sinn? Wie geht man mit Nitpicking-Kommentaren und Gatekeepern um? Und allgemein: Zieht dieser zusätzliche Schritt nicht die Performance des Teams runter und ist sowieso Overhead? All das und noch viel mehr in dieser Episode zum Thema Code Reviews. Bonus: Was Faultiere und Markus Söder mit Code Reviews zu tun haben und warum Blubberwasser den Charakter verdirbt."
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts
image: ./16-code-reviews-nützlich-oder-bremsen-nur-ein-gutes-team.jpg
pubDate: 2022-04-26 04:30:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: A
- name: Wolfi Gassler
  transcriptLetter: B
spotify: https://open.spotify.com/episode/12JJPgrlzpQmrLWEiqT1mR
tags:
- Team
- Software Qualität
- Frontend
- Backend
- Software Engineering
title: "#16 Code Reviews: Nützlich oder bremsen nur ein gutes Team?"
transcript_slim: src/data/transcripts/16-transcript-slim.json
youtube: https://www.youtube.com/watch?v=uvPmOmb-HsE
---
<p>Code Reviews: Jeder will schnelles Feedback, doch niemand hat Zeit dafür - Eine Hassliebe.</p><p>Eine Komponente im Alltag jedes Software Engineers. Egal ob Junior, Senior oder Staff-Engineer. Jeder erstellt Code Reviews und kommentiert die Arbeit von den Kollegen. Doch wie sehen gute Code Reviews aus? Was gehört hinein, was bleibt besser draußen? Wie viel Reviewer machen Sinn? Wie geht man mit Nitpicking-Kommentaren und Gatekeepern um? Und allgemein: Zieht dieser zusätzliche Schritt nicht die Performance des Teams runter und ist sowieso Overhead?</p><p>All das und noch viel mehr in dieser Episode zum Thema Code Reviews.</p><p>Bonus: Was Faultiere und Markus Söder mit Code Reviews zu tun haben und warum Blubberwasser den Charakter verdirbt.</p><p><br></p><p>Feedback an <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a> oder via Twitter an <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p> </p><h3 id="links">Links</h3><ul><li>Google Fellows: <a href="https://www.newyorker.com/magazine/2018/12/10/the-friendship-that-made-google-huge" rel="nofollow">https://www.newyorker.com/magazine/2018/12/10/the-friendship-that-made-google-huge</a></li><li>Code Owners File: <a href="https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners" rel="nofollow">https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners</a></li><li>Software Repository Mining: <a href="https://en.wikipedia.org/wiki/Mining_software_repositories" rel="nofollow">https://en.wikipedia.org/wiki/Mining_software_repositories</a></li><li>Round Robin Code Review Assignment: <a href="https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team" rel="nofollow">https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team</a></li><li>Inner Source: <a href="https://de.wikipedia.org/wiki/Inner_Source" rel="nofollow">https://de.wikipedia.org/wiki/Inner_Source</a></li><li>Suggesting Changes on GitHub: <a href="http://haacked.com/archive/2019/06/03/suggested-changes/" rel="nofollow">http://haacked.com/archive/2019/06/03/suggested-changes/</a></li><li>Are Pull Requests Holding Back Your Team?: <a href="https://betterprogramming.pub/are-pull-requests-holding-back-your-team-e8aec48986c2" rel="nofollow">https://betterprogramming.pub/are-pull-requests-holding-back-your-team-e8aec48986c2</a></li><li>Extrem Programming: <a href="https://de.wikipedia.org/wiki/Extreme_Programming" rel="nofollow">https://de.wikipedia.org/wiki/Extreme_Programming</a></li></ul><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00) Intro</p><p>(01:23) Amsterdam, Hipster und Craft-Bier</p><p>(02:04) git rebase, history rewrite, force push und Kollaboration</p><p>(04:54) Saubere git history und Merge-Commits</p><p>(05:17) Code Reviews</p><p>(02:05) Was ist denn ein Code Review?</p><p>(07:02) Muss der Code Reviewer immer mehr Erfahrung haben?</p><p>(08:10) Was wir von Juniors im Job und bei Code Reviews lernen können</p><p>(09:37) Was ist denn ein Google Fellow?</p><p>(10:37) Code Reviews als Mittel fürs Onboarding</p><p>(11:35) Frühes pushen der Arbeit ist wichtig für frühes Feedback</p><p>(13:15) Wo sollen die Informationen zu meiner Änderung stehen? Im Pull Request oder in der Git Commit message?</p><p>(15:04) Wen lade ich zu meinem Code Review ein? Code Owners file und Software Repository Mining</p><p>(17:34) Worauf soll ich als Code Reviewer acht geben und konzentrieren?</p><p>(20:09) Pro-Aktivität in Code Reviews mit aktiven Vorschlägen und Fragen stellen</p><p>(23:33) Muss man jedes Feedback aus dem Code Review annehmen? Lass das Ego vor der Tür und Gate-Keeper</p><p>(26:37) Kommunikationskonflikte in Code Reviews: Video Calls vs. geschriebene Kommunikation</p><p>(28:14) Nitpicking, Code-Formatting, Styleguides, Linter und Automatisierung</p><p>(32:11) Mindern Pull Requests und Code Reviews die Performance von Teams?</p><p>(36:56) Schnelligkeit bei Facebook und Booking: Schnelles Mergen und Code Reviews danach</p><p>(39:01) Auslagerung der Verantwortung bei Code Reviews und Blameless Kultur innerhalb der Firma</p><p>(42:38) Technische Mittel für sichere Deploys: Feature Flags, Canary Deployments, Dark Launches und Blue-Green-Deployments</p><p>(46:08) Wie schnell sollte ein Code Review vom Team gereviewed werden?</p><p>(53:00) Code Reviews in der Firma vs. Code Reviews in Open Source und die Motivation</p><p>(54:34) Konsequenzen bei nicht eingehaltenen Service Level Agreements (SLA)</p><p>(57:28) Outro</p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://twitter.com/schafele" rel="nofollow">https://twitter.com/schafele</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><p>Engineering Kiosk Podcast: Anfragen an <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a> oder via Twitter an <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></p>