---
advertiser: ''
amazon_music: ''
apple_podcasts: https://podcasts.apple.com/us/podcast/194-was-wurde-aus-mapreduce-und-der-funktionalen-eleganz/id1603082924?i=1000706472441&uo=4
audio: https://audio1.redcircle.com/episodes/4b20fcd3-75e4-4826-b348-b50379d7605e/stream.mp3
chapter:
- start: 00:00:00
  title: 'MapReduce: Ein Deep Dive'
- start: 00:04:32
  title: Info/Werbung
- start: 00:05:32
  title: 'MapReduce: Ein Deep Dive'
- start: 00:15:05
  title: 'Storage: Google File System (GFS) und Hadoop Distributed File System (HDFS)'
- start: 00:21:27
  title: Wie funktioniert MapReduce?
- start: 00:38:10
  title: Seiteneffekte, Determinismus und Reproduzierbarkeit
- start: 00:40:42
  title: "Produktanforderung: Welche Seiten sind in welcher Altersgruppe popul\xE4\
    r?"
- start: 00:47:48
  title: Batch vs. Streaming
- start: 00:50:23
  title: Heutige Relevanz von MapReduce
deezer: https://www.deezer.com/episode/746340731
description: "MapReduce: Ein Deep Dive Im Jahr 2004 war die Verarbeitung von gro\xDF\
  en Datenmengen eine richtige Herausforderung. Einige Firmen hatten daf\xFCr sogenannte\
  \ Supercomputer. Andere haben nur mit der Schulter gezuckt und auf das Ende ihrer\
  \ Berechnung gewartet. Google war einer der Player, der zwar gro\xDFe Datenmengen\
  \ hatte und diese auch verarbeiten wollte, jedoch keine Supercomputer zur Verf\xFC\
  gung hatte. Oder besser gesagt: Nicht das Geld in die Hand nehmen wollte. Was macht\
  \ man also, wenn man ein Problem hat? Eine L\xF6sung suchen. Das hat Jeffrey Dean\
  \ und sein Team getan. Das Ergebnis? Ein revolution\xE4res Paper, wie man mittels\
  \ MapReduce gro\xDFe Datenmengen verteilt auf einfacher Commodity-Hardware verarbeiten\
  \ kann. In dieser Podcast-Episode schauen wir uns das mal genauer an. Wir kl\xE4\
  ren, was MapReduce ist, wie es funktioniert, warum MapReduce so revolution\xE4r\
  \ war, wie es mit Hardware-Ausf\xE4llen umgegangen ist, welche Herausforderungen\
  \ in der Praxis hatte bzw. immer noch hat, was das Google File System, Hadoop und\
  \ HDFS damit zu tun haben und ordnen MapReduce im Kontext der heutigen Technologien\
  \ mit Cloud und Co ein. Eine weitere Episode \u201CPapers We Love\u201D. Bonus:\
  \ Hadoop ist wohl der Elefant im Raum.  Unsere aktuellen Werbepartner findest du\
  \ auf https://engineeringkiosk.dev/partners  Das schnelle Feedback zur Episode:\
  \ \U0001F44D (top)\_\U0001F44E (geht so)  Anregungen, Gedanken, Themen und W\xFC\
  nscheDein Feedback z\xE4hlt! Erreiche uns \xFCber einen der folgenden Kan\xE4le\
  \ \u2026 EngKiosk Community: https://engineeringkiosk.dev/join-discord\_LinkedIn:\
  \ https://www.linkedin.com/company/engineering-kiosk/Email: stehtisch@engineeringkiosk.devMastodon:\
  \ https://podcasts.social/@engkioskBluesky: https://bsky.app/profile/engineeringkiosk.bsky.socialInstagram:\
  \ https://www.instagram.com/engineeringkiosk/ Unterst\xFCtze den Engineering KioskWenn\
  \ du uns etwas Gutes tun m\xF6chtest \u2026 Kaffee schmeckt uns immer\_ Buy us a\
  \ coffee: https://engineeringkiosk.dev/kaffee LinksMapReduce: Simplified Data Processing\
  \ on Large Clusters: https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdfApache\
  \ Hadoop: https://hadoop.apache.org/HDFS Architecture Guide: https://hadoop.apache.org/docs/r1.2.1/hdfs_design.htmlEngineering\
  \ Kiosk Episode #180 Skalierung, aber zu welchem Preis? (Papers We Love): https://engineeringkiosk.dev/podcast/episode/180-skalierung-aber-zu-welchem-preis-papers-we-love/\
  \ Sprungmarken(00:00:00) MapReduce: Ein Deep Dive (00:04:32) Info/Werbung (00:05:32)\
  \ MapReduce: Ein Deep Dive (00:15:05) Storage: Google File System (GFS) und Hadoop\
  \ Distributed File System (HDFS) (00:21:27) Wie funktioniert MapReduce? (00:38:10)\
  \ Seiteneffekte, Determinismus und Reproduzierbarkeit (00:40:42) Produktanforderung:\
  \ Welche Seiten sind in welcher Altersgruppe popul\xE4r? (00:47:48) Batch vs. Streaming\
  \ (00:50:23) Heutige Relevanz von MapReduce  HostsWolfgang Gassler (https://gassler.dev)\_\
  Andy Grunwald (https://andygrunwald.com/)\uFEFF CommunityDiskutiere mit uns und\
  \ vielen anderen Tech-Spezialist\u22C5innen in unserer Engineering Kiosk Community\
  \ unter https://engineeringkiosk.dev/join-discord"
headlines: "anregungen-gedanken-themen-und-wunsche::Anregungen, Gedanken, Themen und\
  \ W\xFCnsche||unterstutze-den-engineering-kiosk::Unterst\xFCtze den Engineering\
  \ Kiosk||links::Links||sprungmarken::Sprungmarken||hosts::Hosts||community::Community"
image: ./194-was-wurde-aus-mapreduce-und-der-funktionalen-eleganz-in-verteilten-systemen.jpg
length_second: 3672
pubDate: 2025-05-06 04:00:00+00:00
rtlplus: ''
six_user_needs: []
speaker:
- name: Andy Grunwald
  transcriptLetter: A
- name: Wolfi Gassler
  transcriptLetter: B
spotify: https://open.spotify.com/episode/1Zn2WYPFQPxKhIOeZdeype
tags: []
title: '#194 Was wurde aus MapReduce und der funktionalen Eleganz in verteilten Systemen?'
transcript_raw: ''
transcript_slim: ''
youtube: https://www.youtube.com/watch?v=AtQdtUZVYbw

---
<p>MapReduce: Ein Deep Dive</p><p>Im Jahr 2004 war die Verarbeitung von großen Datenmengen eine richtige Herausforderung. Einige Firmen hatten dafür sogenannte Supercomputer. Andere haben nur mit der Schulter gezuckt und auf das Ende ihrer Berechnung gewartet. Google war einer der Player, der zwar große Datenmengen hatte und diese auch verarbeiten wollte, jedoch keine Supercomputer zur Verfügung hatte. Oder besser gesagt: Nicht das Geld in die Hand nehmen wollte.</p><p>Was macht man also, wenn man ein Problem hat? Eine Lösung suchen. Das hat Jeffrey Dean und sein Team getan. Das Ergebnis? Ein revolutionäres Paper, wie man mittels MapReduce große Datenmengen verteilt auf einfacher Commodity-Hardware verarbeiten kann.</p><p>In dieser Podcast-Episode schauen wir uns das mal genauer an. Wir klären, was MapReduce ist, wie es funktioniert, warum MapReduce so revolutionär war, wie es mit Hardware-Ausfällen umgegangen ist, welche Herausforderungen in der Praxis hatte bzw. immer noch hat, was das Google File System, Hadoop und HDFS damit zu tun haben und ordnen MapReduce im Kontext der heutigen Technologien mit Cloud und Co ein.</p><p>Eine weitere Episode “Papers We Love”.</p><p>Bonus: Hadoop ist wohl der Elefant im Raum.</p><p><br></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p><br></p><p><strong>Das schnelle Feedback zur Episode:</strong></p><p><a href="https://api.openpodcast.dev/feedback/194/upvote" rel="nofollow"><strong>👍 (top)</strong></a><strong> </strong><a href="https://api.openpodcast.dev/feedback/194/downvote" rel="nofollow"><strong>👎 (geht so)</strong></a></p><p><br></p><h3 id="anregungen-gedanken-themen-und-wunsche">Anregungen, Gedanken, Themen und Wünsche</h3><p>Dein Feedback zählt! Erreiche uns über einen der folgenden Kanäle …</p><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>LinkedIn: <a href="https://www.linkedin.com/company/engineering-kiosk/" rel="nofollow">https://www.linkedin.com/company/engineering-kiosk/</a></li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Bluesky: <a href="https://bsky.app/profile/engineeringkiosk.bsky.social" rel="nofollow">https://bsky.app/profile/engineeringkiosk.bsky.social</a></li><li>Instagram: <a href="https://www.instagram.com/engineeringkiosk/" rel="nofollow">https://www.instagram.com/engineeringkiosk/</a></li></ul><p><br></p><h3 id="unterstutze-den-engineering-kiosk">Unterstütze den Engineering Kiosk</h3><p>Wenn du uns etwas Gutes tun möchtest … Kaffee schmeckt uns immer </p><ul><li>Buy us a coffee: <a href="https://engineeringkiosk.dev/kaffee">https://engineeringkiosk.dev/kaffee</a></li></ul><p><br></p><h3 id="links">Links</h3><ul><li>MapReduce: Simplified Data Processing on Large Clusters: <a href="https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf" rel="nofollow">https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf</a></li><li>Apache Hadoop: <a href="https://hadoop.apache.org/" rel="nofollow">https://hadoop.apache.org/</a></li><li>HDFS Architecture Guide: <a href="https://hadoop.apache.org/docs/r1.2.1/hdfs_design.html" rel="nofollow">https://hadoop.apache.org/docs/r1.2.1/hdfs_design.html</a></li><li>Engineering Kiosk Episode #180 Skalierung, aber zu welchem Preis? (Papers We Love): <a href="https://engineeringkiosk.dev/podcast/episode/180-skalierung-aber-zu-welchem-preis-papers-we-love/">https://engineeringkiosk.dev/podcast/episode/180-skalierung-aber-zu-welchem-preis-papers-we-love/</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) MapReduce: Ein Deep Dive</p><p>(00:04:32) Info/Werbung</p><p>(00:05:32) MapReduce: Ein Deep Dive</p><p>(00:15:05) Storage: Google File System (GFS) und Hadoop Distributed File System (HDFS)</p><p>(00:21:27) Wie funktioniert MapReduce?</p><p>(00:38:10) Seiteneffekte, Determinismus und Reproduzierbarkeit</p><p>(00:40:42) Produktanforderung: Welche Seiten sind in welcher Altersgruppe populär?</p><p>(00:47:48) Batch vs. Streaming</p><p>(00:50:23) Heutige Relevanz von MapReduce</p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://gassler.dev" rel="nofollow">https://gassler.dev</a>) </li><li>Andy Grunwald (<a href="https://andygrunwald.com/" rel="nofollow">https://andygrunwald.com/</a>)</li></ul><p>﻿</p><h3 id="community">Community</h3><p>Diskutiere mit uns und vielen anderen Tech-Spezialist⋅innen in unserer Engineering Kiosk Community unter <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </p>