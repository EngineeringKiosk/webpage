---
advertiser: Media Lab Bayern
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/1f4e3a28-fb66-48c5-a59b-55aa5b4528fa/engineering-kiosk-52-asynchrone-verarbeitung-durch-message-queues---vor--und-nachteile
apple_podcasts: https://podcasts.apple.com/de/podcast/52-asynchrone-verarbeitung-durch-message-queues-vor/id1603082924?i=1000591978596
audio: https://audio1.redcircle.com/episodes/5f100418-9dc0-463b-9033-bf99e1b6da3c/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:00:41
  title: Post-Retouren an der Haustür, Warteschlangen und Reservierungs-Scheduling mit Process Mining
- start: 00:05:49
  title: "Das heutige Thema: Message Queuing (mit RabbitMQ)"
- start: 00:07:23
  title: Was sind Message Queues und was sind typische Anwendungsfälle?
- start: 00:10:48
  title: Komponenten einer Message Queues und deine Datenbank als Message Queues
- start: 00:12:30
  title: "Herausforderungen beim Message Queuing: Exactly once delivery"
- start: 00:15:15
  title: "Möglichkeiten durch Message Queuing: Granulare Skalierbarkeit, Micro-Service Kommunikation und Serverless"
- start: 00:16:56
  title: Was ist RabbitMQ? (Erlang, AMQP)
- start: 00:21:53
  title: "Advanced Features von Message Queuing-Systemen: Exchanges, Routing, Priority Queues, Time to live (TTL)"
- start: 00:27:01
  title: Message Acknowledgement + Rejetion und Dead Letter Queues
- start: 00:31:11
  title: Ist Amazon SQS oder Google PubSub eine gute Alternative?
- start: 00:34:22
  title: "Alternative mit Redis: PubSub, Listen und Streams"
- start: 00:36:18
  title: Wo ist der Unterschied zwischen einer Message Queue und einem Stream?
- start: 00:38:51
  title: Kann ich Apache Kafka als Message Queuing System verwenden?
- start: 00:40:29
  title: Ist RabbitMQ oder Apache Kafka einfacher zu installieren und zu betreiben?
- start: 00:42:56
  title: ZeroMQ
- start: 00:44:53
  title: Was spricht gegen RabbitMQ? Operations, idempotente Consumer
- start: 00:48:34
  title: "Andere Protokolle fürs Message Queuing: MQTT, HTTP und WebSockets"
- start: 00:49:46
  title: Erfahrung durch die Nutzung in Side Projects und Zusammenfassung
deezer: https://www.deezer.com/episode/466631357
description: "Asynchrone Verarbeitung durch Message Queues: Was ist das und wofür ist das gut? In vielen Applikationen gibt es Bereiche, die einfach etwas Zeit für die Verarbeitung brauchen, aber das klassische Anfrage/Antwort (Request/Response) Verhalten nicht blockieren sollen. Oft werden dafür asynchrone Prozesse verwendet. Durch den Einsatz von Message Queues ergeben sich weitere Vorteile wie u.a. granulare Skalierbarkeit und Unabhängigkeit von einzelnen Programmiersprachen. RabbitMQ ist einer der Platzhirsche im Bereich Open-Source-Message-Broker. In dieser Episode klären wir wofür Message Queues gut sind, bei welchen klassischen Anwendungsfällen diese helfen können, welche Herausforderungen diese Darstellen, wo der Unterschied zu Pub/Sub oder Streams ist und was Redis, Kafka und ZeroMQ damit zu tun hat. Bonus: Warum Software rostet."
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./52-asynchrone-verarbeitung-durch-message-queues-vor-und-nachteile.jpg
pubDate: 2023-01-03 05:00:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: A
- name: Wolfi Gassler
  transcriptLetter: B
spotify: https://open.spotify.com/episode/0LmFRcBMaZGonYPSE1x15j
tags:
- Backend
- Software Engineering
title: "#52 Asynchrone Verarbeitung durch Message Queues - Vor- und Nachteile"
transcript_slim: src/data/transcripts/52-transcript-slim.json
youtube: https://www.youtube.com/watch?v=ew9SFEmpbOg
---
<p>Asynchrone Verarbeitung durch Message Queues: Was ist das und wofür ist das gut?</p><p>In vielen Applikationen gibt es Bereiche, die einfach etwas Zeit für die Verarbeitung brauchen, aber das klassische Anfrage/Antwort (Request/Response) Verhalten nicht blockieren sollen. Oft werden dafür asynchrone Prozesse verwendet. Durch den Einsatz von Message Queues ergeben sich weitere Vorteile wie u.a. granulare Skalierbarkeit und Unabhängigkeit von einzelnen Programmiersprachen. RabbitMQ ist einer der Platzhirsche im Bereich Open-Source-Message-Broker.</p><p>In dieser Episode klären wir wofür Message Queues gut sind, bei welchen klassischen Anwendungsfällen diese helfen können, welche Herausforderungen diese Darstellen, wo der Unterschied zu Pub/Sub oder Streams ist und was Redis, Kafka und ZeroMQ damit zu tun hat.</p><p>Bonus: Warum Software rostet.</p><p><br></p><p>Feedback (gerne auch als Voice Message)</p><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per <a href="https://engineeringkiosk.dev/kontakt/">Email</a> oder WhatsApp Voice Message an +49 15678 136776</p><p><br></p><h3 id="links">Links</h3><ul><li>RabbitMQ: <a href="https://www.rabbitmq.com/" rel="nofollow">https://www.rabbitmq.com/</a></li><li>ActiveMQ: <a href="https://activemq.apache.org/" rel="nofollow">https://activemq.apache.org/</a></li><li>AMQP: <a href="https://www.amqp.org/" rel="nofollow">https://www.amqp.org/</a></li><li>Jakarta Messaging: <a href="https://de.wikipedia.org/wiki/Jakarta_Messaging" rel="nofollow">https://de.wikipedia.org/wiki/Jakarta_Messaging</a></li><li>Red Hat / JBoss AMQ: <a href="https://www.redhat.com/de/technologies/jboss-middleware/amq" rel="nofollow">https://www.redhat.com/de/technologies/jboss-middleware/amq</a></li><li>Apache Kafka: <a href="https://kafka.apache.org/" rel="nofollow">https://kafka.apache.org/</a></li><li>ZeroMQ: <a href="https://zeromq.org/" rel="nofollow">https://zeromq.org/</a></li><li>Erlang mnesia: <a href="https://www.erlang.org/doc/man/mnesia.html" rel="nofollow">https://www.erlang.org/doc/man/mnesia.html</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p><span>(00:00:00) Intro</span></p><p><span>(00:00:41) Post-Retouren an der Haustür, Warteschlangen und Reservierungs-Scheduling mit Process Mining</span></p><p><span>(00:05:49) Das heutige Thema: Message Queuing (mit RabbitMQ)</span></p><p><span>(00:07:23) Was sind Message Queues und was sind typische Anwendungsfälle?</span></p><p><span>(00:10:48) Komponenten einer Message Queues und deine Datenbank als Message Queues</span></p><p><span>(00:12:30) Herausforderungen beim Message Queuing: Exactly once delivery</span></p><p><span>(00:15:15) Möglichkeiten durch Message Queuing: Granulare Skalierbarkeit, Micro-Service Kommunikation und Serverless</span></p><p><span>(00:16:56) Was ist RabbitMQ? (Erlang, AMQP)</span></p><p><span>(00:21:53) Advanced Features von Message Queuing-Systemen: Exchanges, Routing, Priority Queues, Time to live (TTL)</span></p><p><span>(00:27:01) Message Acknowledgement + Rejetion und Dead Letter Queues</span></p><p><span>(00:31:11) Ist Amazon SQS oder Google PubSub eine gute Alternative?</span></p><p><span>(00:34:22) Alternative mit Redis: PubSub, Listen und Streams</span></p><p><span>(00:36:18) Wo ist der Unterschied zwischen einer Message Queue und einem Stream?</span></p><p><span>(00:38:51) Kann ich Apache Kafka als Message Queuing System verwenden?</span></p><p><span>(00:40:29) Ist RabbitMQ oder Apache Kafka einfacher zu installieren und zu betreiben?</span></p><p><span>(00:42:56) ZeroMQ</span></p><p><span>(00:44:53) Was spricht gegen RabbitMQ? Operations, idempotente Consumer</span></p><p><span>(00:48:34) Andere Protokolle fürs Message Queuing: MQTT, HTTP und WebSockets</span></p><p><span>(00:49:46) Erfahrung durch die Nutzung in Side Projects und Zusammenfassung</span></p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://mastodon.social/@woolf" rel="nofollow">https://mastodon.social/@woolf</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul>