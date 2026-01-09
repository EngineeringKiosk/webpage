---
advertiser: ABOUT YOU
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/72510b99-a2b0-4fca-8ecd-a3969347c78f/engineering-kiosk-101-observability-und-opentelemetry-mit-severin-neumann
apple_podcasts: https://podcasts.apple.com/us/podcast/101-observability-und-opentelemetry-mit-severin-neumann/id1603082924?i=1000638267638&uo=4
audio: https://audio1.redcircle.com/episodes/627265ce-6d39-4cc5-a9af-c1c3a3b68095/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:01:07
  title: Unser Gast Severin Neumann, Sales Engineering und Demo Monkey
- start: 00:06:50
  title: Was ist Observability und wie unterscheidet es sich von Monitoring?
- start: 00:06:54
  title: Info/Werbung
- start: 00:07:56
  title: Was ist Observability und wie unterscheidet es sich von Monitoring?
- start: 00:15:24
  title: "Signale bei Observability: Metrics, Logs und Traces"
- start: 00:17:25
  title: Was ist OpenTelemetry?
- start: 00:25:59
  title: APM-Anbieter und der Lock-in-Effekt
- start: 00:28:38
  title: OpenTelemetry als offener Standard
- start: 00:35:38
  title: Die Sicht von Dev und Ops auf OpenTelemetry
- start: 00:41:11
  title: Wie binde ich OpenTelemetry in meine App ein?
- start: 00:48:02
  title: Auto-Instrumentation, Microservice-Architektur und Trace-Headers
- start: 00:51:03
  title: Overhead beim Erheben von Daten und eigene Metadaten in Traces
- start: 00:56:12
  title: Speicherung von Observability-Daten
- start: 01:00:55
  title: Pitfalls und die Shortcomings von OpenTelemetry
- start: 01:05:26
  title: OpenTelemetry-Dokumentation
deezer: https://www.deezer.com/episode/584273752
description: "Effektive Observability mit OpenTelemetry Früher waren viele Applikationen eine Black Box, besonders für die Ops aka Betriebsabteilung. Dann fing das Logging an. Apps haben Log-Lines geschrieben, zum Beispiel wann die App fertig hochgefahren ist oder wenn etwas schief gegangen ist. In einer Art und Weise haben durch Logs die Devs angefangen, mit den Ops-Leuten zu kommunizieren. Irgendwann später gab es Metriken. Wie viel RAM verbraucht die App, wie oft wurde der Garbage Collector getriggert oder auch Business-Metriken, wie oft eine Bestellung ausgeführt wurde oder wann eine Geo- anstatt einer Text-Suche gestartet wurde. War das alles? Nein. Der neueste Hype: Traces. Eine genaue Einsicht, welchen Code-Path die App genommen hat und wie lange dieser gedauert hat inkl. aller Metadaten, die wir uns wünschen. Und wenn man dies nun alles in einen Sack packt, es gut durchschüttelt und man ein System hat, das man auf Basis dieser Daten fragen stellen kann, nennt man das Observability. Und genau da setzt das Projekt OpenTelemetry an. In dieser Episode sprechen wir mit dem Experten Severin Neumann über Observability und OpenTelemetry. Bonus: Was ist ein Sales-Engineer?  Unsere aktuellen Werbepartner findest du auf https://engineeringkiosk.dev/partners \u00a0 Das schnelle Feedback zur Episode: 👍 (top)\u00a0👎 (geht so)"
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./101-observability-und-opentelemetry-mit-severin-neumann.jpg
pubDate: 2023-12-12 05:00:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: A
- name: Wolfi Gassler
  transcriptLetter: B
- name: Severin Neumann
  transcriptLetter: C
spotify: https://open.spotify.com/episode/1RkWRmsfnRMBfDrYKNGDRf
tags:
- DevOps
- Software Engineering
- Interview
title: "#101 Observability und OpenTelemetry mit Severin Neumann"
transcript_slim: src/data/transcripts/101-transcript-slim.json
youtube: https://www.youtube.com/watch?v=h8NdY4zBJIE
---
<p>Effektive Observability mit OpenTelemetry</p><p>Früher waren viele Applikationen eine Black Box, besonders für die Ops aka Betriebsabteilung. Dann fing das Logging an. Apps haben Log-Lines geschrieben, zum Beispiel wann die App fertig hochgefahren ist oder wenn etwas schief gegangen ist. In einer Art und Weise haben durch Logs die Devs angefangen, mit den Ops-Leuten zu kommunizieren.</p><p>Irgendwann später gab es Metriken. Wie viel RAM verbraucht die App, wie oft wurde der Garbage Collector getriggert oder auch Business-Metriken, wie oft eine Bestellung ausgeführt wurde oder wann eine Geo- anstatt einer Text-Suche gestartet wurde.</p><p>War das alles? Nein. Der neueste Hype: Traces. Eine genaue Einsicht, welchen Code-Path die App genommen hat und wie lange dieser gedauert hat inkl. aller Metadaten, die wir uns wünschen.</p><p>Und wenn man dies nun alles in einen Sack packt, es gut durchschüttelt und man ein System hat, das man auf Basis dieser Daten fragen stellen kann, nennt man das Observability.</p><p>Und genau da setzt das Projekt OpenTelemetry an.</p><p>In dieser Episode sprechen wir mit dem Experten Severin Neumann über Observability und OpenTelemetry.</p><p>Bonus: Was ist ein Sales-Engineer?</p><p><br></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p> </p><p><strong>Das schnelle Feedback zur Episode:</strong></p><p><a href="https://api.openpodcast.dev/feedback/101/upvote" rel="nofollow"><strong>👍 (top)</strong></a><strong> </strong><a href="https://api.openpodcast.dev/feedback/101/downvote" rel="nofollow"><strong>👎 (geht so)</strong></a></p><p><br></p><p>Feedback (gerne auch als Voice Message)</p><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per <a href="https://engineeringkiosk.dev/kontakt/">Email</a> oder WhatsApp Voice Message an +49 15678 136776</p><p><br></p><h3 id="links">Links</h3><ul><li>Severin Neumann: <a href="https://www.linkedin.com/in/severinneumann/" rel="nofollow">https://www.linkedin.com/in/severinneumann/</a></li><li>DemoMonkey: <a href="https://github.com/svrnm/DemoMonkey" rel="nofollow">https://github.com/svrnm/DemoMonkey</a></li><li>OpenTelemetry: <a href="https://opentelemetry.io/" rel="nofollow">https://opentelemetry.io/</a></li><li>OpenTracing: <a href="https://opentracing.io/" rel="nofollow">https://opentracing.io/</a></li><li>W3C Trace Context: <a href="https://www.w3.org/TR/trace-context/" rel="nofollow">https://www.w3.org/TR/trace-context/</a></li><li>W3C Trace Context: AMQP protocol: <a href="https://w3c.github.io/trace-context-amqp/" rel="nofollow">https://w3c.github.io/trace-context-amqp/</a></li><li>OpenTelemetry Vendors: <a href="https://opentelemetry.io/ecosystem/vendors/" rel="nofollow">https://opentelemetry.io/ecosystem/vendors/</a></li><li>OpenTelemetry Instrumentation: <a href="https://opentelemetry.io/docs/instrumentation/" rel="nofollow">https://opentelemetry.io/docs/instrumentation/</a></li><li>Dynatrace: <a href="https://www.dynatrace.com/de/" rel="nofollow">https://www.dynatrace.com/de/</a></li><li>AppDynamics: <a href="https://www.appdynamics.com/" rel="nofollow">https://www.appdynamics.com/</a></li><li>OpenTelemetry Tracing for Monoliths - Phillip Carter, Honeycomb: <a href="https://www.youtube.com/watch?v=kzXT0WlTBpw" rel="nofollow">https://www.youtube.com/watch?v=kzXT0WlTBpw</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) Intro</p><p>(00:01:07) Unser Gast Severin Neumann, Sales Engineering und Demo Monkey</p><p>(00:06:50) Was ist Observability und wie unterscheidet es sich von Monitoring?</p><p>(00:06:54) Info/Werbung</p><p>(00:07:56) Was ist Observability und wie unterscheidet es sich von Monitoring?</p><p>(00:15:24) Signale bei Observability: Metrics, Logs und Traces</p><p>(00:17:25) Was ist OpenTelemetry?</p><p>(00:25:59) APM-Anbieter und der Lock-in-Effekt</p><p>(00:28:38) OpenTelemetry als offener Standard</p><p>(00:35:38) Die Sicht von Dev und Ops auf OpenTelemetry</p><p>(00:41:11) Wie binde ich OpenTelemetry in meine App ein?</p><p>(00:48:02) Auto-Instrumentation, Microservice-Architektur und Trace-Headers</p><p>(00:51:03) Overhead beim Erheben von Daten und eigene Metadaten in Traces</p><p>(00:56:12) Speicherung von Observability-Daten</p><p>(01:00:55) Pitfalls und die Shortcomings von OpenTelemetry</p><p>(01:05:26) OpenTelemetry-Dokumentation</p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://mastodon.social/@woolf" rel="nofollow">https://mastodon.social/@woolf</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul>