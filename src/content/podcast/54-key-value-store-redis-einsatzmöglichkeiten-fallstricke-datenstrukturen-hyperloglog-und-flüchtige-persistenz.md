---
advertiser: ""
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/a8eb4417-1d48-4eaa-ae6f-7566cade756e/engineering-kiosk-54-key-value-store-redis-einsatzm%C3%B6glichkeiten-fallstricke-datenstrukturen-hyperloglog-und-fl%C3%BCchtige-persistenz
apple_podcasts: https://podcasts.apple.com/us/podcast/54-key-value-store-redis-einsatzm%C3%B6glichkeiten-fallstricke/id1603082924?i=1000594826183&uo=4
audio: https://audio1.redcircle.com/episodes/c55253d1-4bcb-4ccc-8d3c-5042ceef6064/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:00:49
  title: Die Datenbank Redis, wie wir zu dem Thema kommen und warum andere oft Redis wollen
- start: 00:03:32
  title: Was ist Redis?
- start: 00:06:51
  title: Wie sicher sind meine Daten im In-Memory-Store Redis?
- start: 00:08:07
  title: Was ist Geospatial im Kontext von Redis?
- start: 00:09:34
  title: Was ist HyperLogLog?
- start: 00:12:13
  title: Was sind BitMaps und BitFields?
- start: 00:14:19
  title: "In Redis fängt alles mit einem Key an: Key-Value"
- start: 00:17:07
  title: Warum sollte ich Redis nehmen und nicht eine MySQL-Datenbank?
- start: 00:19:21
  title: "Persistenz-Modelle in Redis: RDB und AOF"
- start: 00:25:49
  title: Redis Cluster und Leader-Follower Architektur
- start: 00:26:55
  title: Wie Andy zu Redis kam
- start: 00:31:36
  title: "Use-Cases von Redis: Caching, Delayed writes"
- start: 00:34:26
  title: Ist Redis eine Datenbank?
- start: 00:36:08
  title: Salvatore Sanfilippo, antirez, der Autor von Redis
- start: 00:40:09
  title: "Fallstricke von Redis: Persistenz, Key-Struktur, Verschachtelte Datenstrukturen"
- start: 00:43:16
  title: Lua Scripting in Redis
- start: 00:45:26
  title: Neue Daten-Strukturen durch Redis Modules
- start: 00:48:00
  title: Redis Protokoll und Kommunikation mit Redis
- start: 00:49:52
  title: Kommerzieller Support von Redis und Firmen-Konstrukt hinter Redis mit RedisLabs
- start: 00:55:27
  title: "Redis in der Cloud-Welt: GCP Memory Store, AWS ElastiCache, Azure Cache for Redis"
- start: 00:59:36
  title: "Andys Lieblings-Datenbank und nicht erwähnte Themen: Cluster, Streams, twemproxy, Redis Sentinel"
deezer: https://www.deezer.com/episode/470664857
description: "Redis - Der open source, in-memory data structure server Viele Software-Engineers haben bereits von Redis gelesen. Primär im Anwendungsfall eines Caches. Doch das ist bei weitem nicht alles, was Redis unter der Haube hat. In dieser Episode schauen wir uns den Data Structure Server mal genauer an. Was ist Redis? Welche Datentypen unterstützt dieser? Was ist Geospatial und HyperLogLog? Kann Redis meine Daten auch persistieren? Welche Use-Cases gibt es neben dem Caching? Wer ist eigentlich der Kopf hinter Redis? Und wie kann ich Redis erweitern, falls ich noch mehr Funktionalität brauche? All das und noch viel mehr Hintergrundwissen zu Redis in dieser Episode. Bonus: Wann unser Co-Host Andy und wann Andreas genannt wird und was Clippy von Word mit Redis zu tun hat."
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./54-key-value-store-redis-einsatzmöglichkeiten-fallstricke-datenstrukturen-hyperloglog-und-flüchtige-persistenz.jpg
pubDate: 2023-01-17 05:00:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: A
- name: Wolfi Gassler
  transcriptLetter: B
spotify: https://open.spotify.com/episode/6NiUofGRvIYQV0arsK6QFh
tags:
- DevOps
- Datenbanken
title: "#54 Key Value Store Redis: Einsatzmöglichkeiten, Fallstricke, Datenstrukturen, HyperLogLog und (flüchtige) Persistenz"
transcript_slim: src/data/transcripts/54-transcript-slim.json
youtube: https://www.youtube.com/watch?v=Y6ajwtbTExY
---
<p>Redis - Der open source, in-memory data structure server</p><p>Viele Software-Engineers haben bereits von Redis gelesen. Primär im Anwendungsfall eines Caches. Doch das ist bei weitem nicht alles, was Redis unter der Haube hat. In dieser Episode schauen wir uns den Data Structure Server mal genauer an. Was ist Redis? Welche Datentypen unterstützt dieser? Was ist Geospatial und HyperLogLog? Kann Redis meine Daten auch persistieren? Welche Use-Cases gibt es neben dem Caching? Wer ist eigentlich der Kopf hinter Redis? Und wie kann ich Redis erweitern, falls ich noch mehr Funktionalität brauche? All das und noch viel mehr Hintergrundwissen zu Redis in dieser Episode.</p><p>Bonus: Wann unser Co-Host Andy und wann Andreas genannt wird und was Clippy von Word mit Redis zu tun hat.</p><p><br></p><p>Feedback (gerne auch als Voice Message)</p><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per <a href="https://engineeringkiosk.dev/kontakt/">Email</a> oder WhatsApp Voice Message an +49 15678 136776</p><p><br></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p> </p><h3 id="links">Links</h3><ul><li>Engineering Kiosk Episode #53 Cloud / NoCode/ AI / ChatGPT ersetzen unsere Jobs?: <a href="https://engineeringkiosk.dev/podcast/episode/53-cloud-nocode-ai-chatgpt-ersetzen-unsere-jobs/">https://engineeringkiosk.dev/podcast/episode/53-cloud-nocode-ai-chatgpt-ersetzen-unsere-jobs/</a></li><li>Redis: <a href="https://redis.io/" rel="nofollow">https://redis.io/</a></li><li>Redis Data types: <a href="https://redis.io/docs/data-types/" rel="nofollow">https://redis.io/docs/data-types/</a></li><li>MySQL Spatial Data types: <a href="https://dev.mysql.com/doc/refman/8.0/en/spatial-types.html" rel="nofollow">https://dev.mysql.com/doc/refman/8.0/en/spatial-types.html</a></li><li>Engineering Kiosk #28 O(1), O(log n), O(n^2) - Ist die Komplexität von Algorithmen im Entwickler-Alltag relevant?: <a href="https://engineeringkiosk.dev/podcast/episode/28-o1-olog-n-on2-ist-die-komplexit%C3%A4t-von-algorithmen-im-entwickler-alltag-relevant/">https://engineeringkiosk.dev/podcast/episode/28-o1-olog-n-on2-ist-die-komplexit%C3%A4t-von-algorithmen-im-entwickler-alltag-relevant/</a></li><li>Redis Persistence: <a href="https://redis.io/docs/management/persistence/" rel="nofollow">https://redis.io/docs/management/persistence/</a></li><li>Consistent hashing: <a href="https://en.wikipedia.org/wiki/Consistent_hashing" rel="nofollow">https://en.wikipedia.org/wiki/Consistent_hashing</a></li><li>Redis Cluster: <a href="https://redis.io/docs/management/scaling/" rel="nofollow">https://redis.io/docs/management/scaling/</a></li><li>Learn Redis the hard way (in production): <a href="https://tech.trivago.com/post/learn-redis-the-hard-way/" rel="nofollow">https://tech.trivago.com/post/learn-redis-the-hard-way/</a></li><li>memcached: <a href="https://memcached.org/" rel="nofollow">https://memcached.org/</a></li><li>GitHub Issue #1771: &#34;Software watchdog crashes redis during rdb save point&#34;: <a href="https://github.com/redis/redis/issues/1771" rel="nofollow">https://github.com/redis/redis/issues/1771</a></li><li>Salvatore Sanfilippo (antirez) auf GitHub: <a href="https://github.com/antirez" rel="nofollow">https://github.com/antirez</a></li><li>PL/SQL: <a href="https://de.wikipedia.org/wiki/PL/SQL" rel="nofollow">https://de.wikipedia.org/wiki/PL/SQL</a></li><li>Redis Scripting with Lua: <a href="https://redis.io/docs/manual/programmability/eval-intro/" rel="nofollow">https://redis.io/docs/manual/programmability/eval-intro/</a></li><li>Redis Modules: <a href="https://redis.io/resources/modules/" rel="nofollow">https://redis.io/resources/modules/</a></li><li>Redis Clients: <a href="https://redis.io/resources/clients/" rel="nofollow">https://redis.io/resources/clients/</a></li><li>Redis serialization protocol (RESP) specification: <a href="https://redis.io/docs/reference/protocol-spec/" rel="nofollow">https://redis.io/docs/reference/protocol-spec/</a></li><li>DB Engines: <a href="https://db-engines.com/en/ranking" rel="nofollow">https://db-engines.com/en/ranking</a></li><li>Redis vs. MySQL Benchmarks: <a href="https://dzone.com/articles/redis-vs-mysql-benchmarks" rel="nofollow">https://dzone.com/articles/redis-vs-mysql-benchmarks</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) Intro</p><p>(00:00:49) Die Datenbank Redis, wie wir zu dem Thema kommen und warum andere oft Redis wollen</p><p>(00:03:32) Was ist Redis?</p><p>(00:06:51) Wie sicher sind meine Daten im In-Memory-Store Redis?</p><p>(00:08:07) Was ist Geospatial im Kontext von Redis?</p><p>(00:09:34) Was ist HyperLogLog?</p><p>(00:12:13) Was sind BitMaps und BitFields?</p><p>(00:14:19) In Redis fängt alles mit einem Key an: Key-Value</p><p>(00:17:07) Warum sollte ich Redis nehmen und nicht eine MySQL-Datenbank?</p><p>(00:19:21) Persistenz-Modelle in Redis: RDB und AOF</p><p>(00:25:49) Redis Cluster und Leader-Follower Architektur</p><p>(00:26:55) Wie Andy zu Redis kam</p><p>(00:31:36) Use-Cases von Redis: Caching, Delayed writes</p><p>(00:34:26) Ist Redis eine Datenbank?</p><p>(00:36:08) Salvatore Sanfilippo, antirez, der Autor von Redis</p><p>(00:40:09) Fallstricke von Redis: Persistenz, Key-Struktur, Verschachtelte Datenstrukturen</p><p>(00:43:16) Lua Scripting in Redis</p><p>(00:45:26) Neue Daten-Strukturen durch Redis Modules</p><p>(00:48:00) Redis Protokoll und Kommunikation mit Redis</p><p>(00:49:52) Kommerzieller Support von Redis und Firmen-Konstrukt hinter Redis mit RedisLabs</p><p>(00:55:27) Redis in der Cloud-Welt: GCP Memory Store, AWS ElastiCache, Azure Cache for Redis</p><p>(00:59:36) Andys Lieblings-Datenbank und nicht erwähnte Themen: Cluster, Streams, twemproxy, Redis Sentinel</p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://mastodon.social/@woolf" rel="nofollow">https://mastodon.social/@woolf</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul>