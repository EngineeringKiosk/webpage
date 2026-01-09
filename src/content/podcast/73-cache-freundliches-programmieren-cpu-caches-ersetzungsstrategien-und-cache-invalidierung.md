---
advertiser: ""
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/32005152-5853-4d9e-a37f-565996c54508/engineering-kiosk-73-cache-freundliches-programmieren-cpu-caches-ersetzungsstrategien-und-cache-invalidierung
apple_podcasts: https://podcasts.apple.com/us/podcast/73-cache-freundliches-programmieren-cpu-caches-ersetzungsstrategien/id1603082924?i=1000614941494&uo=4
audio: https://audio1.redcircle.com/episodes/c4c3001c-2a6c-4384-9409-474583b5fb53/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:01:02
  title: Es gibt nur zwei harte Probleme in der Informatik
- start: 00:06:56
  title: Caching und Cache Invalidierung
- start: 00:07:25
  title: Was ist ein Cache?
- start: 00:10:36
  title: "Jeder nutzt Caches jeden Tag: Cache-Layer"
- start: 00:12:34
  title: Cache-Hierarchien und wie relevant sind CPU Caches (L1 bis L4) heute noch?
- start: 00:16:39
  title: Wie funktioniert eigentlich ein Cache? Cache Hit, Cache Miss, Cache Ratio
- start: 00:21:27
  title: "Ersetzungsstrategien / Eviction-Policies: FIFO, LIFO, LRU, TTL, MRU, LFU"
- start: 00:31:57
  title: Wie kann ich Cache-Freundlich programmieren um den L1 bis L4 Cache richtig zu nutzen und Monitoring von L1 bis L4 Caches
- start: 00:48:39
  title: Ist die Funktionsweise von L1 bis L4 Caches valide für andere Arten von Caches?
- start: 00:51:12
  title: Warum ist Cache Invalidierung ein hartes Problem?
deezer: https://www.deezer.com/episode/512905195
description: "There are only two hard things in Computer Science: cache invalidation and naming things. Caches sind einfach überall. Jede Aktion auf einem Computer nutzt eine Vielzahl an Caches. Sei es der Browser Cache, DNS-Cache, In-Memory Cache auf dem Server oder dein lokaler CPU Cache L1-L4. Doch was sind Caches eigentlich? Welche Cache-Layer und Cache-Hierarchien gibt es? Wie funktionieren Caches? Wie kann ich Cache-freundlich programmieren? Was passiert, wenn der Cache voll ist und was sind Eviction-Policies? Wie relevant sind heutzutage eigentlich die CPU-Caches L1 bis L4 für die normalen Software-Entwickler*innen? Wie kann ich verifizieren, ob mein Code Cache-freundlich ist? Und warum ist Cache Invalidation eigentlich ein hartes Problem? Bonus: Was Bandlaufwerke mit Caching und niederländisches Hähnchen mit Queues zu tun haben.  Unsere aktuellen Werbepartner findest du auf https://engineeringkiosk.dev/partners \u00a0 Das schnelle Feedback zur Episode: 👍 (top)  👎 (geht so)"
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./73-cache-freundliches-programmieren-cpu-caches-ersetzungsstrategien-und-cache-invalidierung.jpg
pubDate: 2023-05-30 04:00:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: A
- name: Wolfi Gassler
  transcriptLetter: B
spotify: https://open.spotify.com/episode/6ZAM0fZpQAA6WZqcz5MiH7
tags:
- Software Engineering
- Backend
title: "#73 Cache-freundliches Programmieren, CPU-Caches, Ersetzungsstrategien und Cache-Invalidierung"
transcript_slim: src/data/transcripts/73-transcript-slim.json
youtube: https://www.youtube.com/watch?v=6vHAhiVTj10
---
<p>There are only two hard things in Computer Science: cache invalidation and naming things.</p><p>Caches sind einfach überall. Jede Aktion auf einem Computer nutzt eine Vielzahl an Caches. Sei es der Browser Cache, DNS-Cache, In-Memory Cache auf dem Server oder dein lokaler CPU Cache L1-L4. Doch was sind Caches eigentlich? Welche Cache-Layer und Cache-Hierarchien gibt es? Wie funktionieren Caches? Wie kann ich Cache-freundlich programmieren? Was passiert, wenn der Cache voll ist und was sind Eviction-Policies? Wie relevant sind heutzutage eigentlich die CPU-Caches L1 bis L4 für die normalen Software-Entwickler*innen? Wie kann ich verifizieren, ob mein Code Cache-freundlich ist? Und warum ist Cache Invalidation eigentlich ein hartes Problem?</p><p>Bonus: Was Bandlaufwerke mit Caching und niederländisches Hähnchen mit Queues zu tun haben.</p><p><br></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p> </p><p><strong>Das schnelle Feedback zur Episode:</strong></p><p><a href="https://api.openpodcast.dev/feedback/73/upvote" rel="nofollow"><strong>👍 (top)</strong></a><strong>  </strong><a href="https://api.openpodcast.dev/feedback/73/downvote" rel="nofollow"><strong>👎 (geht so)</strong></a></p><p><br></p><p>Feedback (gerne auch als Voice Message)</p><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per <a href="https://engineeringkiosk.dev/kontakt/">Email</a> oder WhatsApp Voice Message an +49 15678 136776</p><p><br></p><h3 id="links">Links</h3><ul><li>KIP-932: Queues for Kafka: <a href="https://cwiki.apache.org/confluence/display/KAFKA/KIP-932%3A+Queues+for+Kafka" rel="nofollow">https://cwiki.apache.org/confluence/display/KAFKA/KIP-932%3A+Queues+for+Kafka</a></li><li>Engineering Kiosk Episode #65 Clean Code macht Software langsam: <a href="https://engineeringkiosk.dev/podcast/episode/65-clean-code-macht-software-langsam/">https://engineeringkiosk.dev/podcast/episode/65-clean-code-macht-software-langsam/</a></li><li>Phil Karlton&#39;s original home page from Netscape: <a href="https://www.karlton.org/karlton/" rel="nofollow">https://www.karlton.org/karlton/</a></li><li>Engineering Kiosk Episode #67 Die Netz-Entlastung des Internets: Content Delivery Networks (CDNs): <a href="https://engineeringkiosk.dev/podcast/episode/67-die-netz-entlastung-des-internets-content-delivery-networks-cdns/">https://engineeringkiosk.dev/podcast/episode/67-die-netz-entlastung-des-internets-content-delivery-networks-cdns/</a></li><li>Redis Key eviction: https://redis.io/docs/reference/eviction/</li><li>Engineering Kiosk Episode #54 Key Value Store Redis: Einsatzmöglichkeiten, Fallstricke, Datenstrukturen, HyperLogLog und (flüchtige) Persistenz: <a href="https://engineeringkiosk.dev/podcast/episode/54-key-value-store-redis-einsatzm%C3%B6glichkeiten-fallstricke-datenstrukturen-hyperloglog-und-fl%C3%BCchtige-persistenz/">https://engineeringkiosk.dev/podcast/episode/54-key-value-store-redis-einsatzm%C3%B6glichkeiten-fallstricke-datenstrukturen-hyperloglog-und-fl%C3%BCchtige-persistenz/</a></li><li>Cache replacement policies auf Wikipedia: <a href="https://en.wikipedia.org/wiki/Cache_replacement_policies" rel="nofollow">https://en.wikipedia.org/wiki/Cache_replacement_policies</a></li><li>Why Cache Invalidation is Hard: <a href="https://levelup.gitconnected.com/why-cache-invalidation-is-hard-and-how-to-solve-it-7921ffe4f6eb" rel="nofollow">https://levelup.gitconnected.com/why-cache-invalidation-is-hard-and-how-to-solve-it-7921ffe4f6eb</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) Intro</p><p>(00:01:02) Es gibt nur zwei harte Probleme in der Informatik</p><p>(00:06:56) Caching und Cache Invalidierung</p><p>(00:07:25) Was ist ein Cache?</p><p>(00:10:36) Jeder nutzt Caches jeden Tag: Cache-Layer</p><p>(00:12:34) Cache-Hierarchien und wie relevant sind CPU Caches (L1 bis L4) heute noch?</p><p>(00:16:39) Wie funktioniert eigentlich ein Cache? Cache Hit, Cache Miss, Cache Ratio</p><p>(00:21:27) Ersetzungsstrategien / Eviction-Policies: FIFO, LIFO, LRU, TTL, MRU, LFU</p><p>(00:31:57) Wie kann ich Cache-Freundlich programmieren um den L1 bis L4 Cache richtig zu nutzen und Monitoring von L1 bis L4 Caches</p><p>(00:48:39) Ist die Funktionsweise von L1 bis L4 Caches valide für andere Arten von Caches?</p><p>(00:51:12) Warum ist Cache Invalidierung ein hartes Problem?</p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://mastodon.social/@woolf" rel="nofollow">https://mastodon.social/@woolf</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul>