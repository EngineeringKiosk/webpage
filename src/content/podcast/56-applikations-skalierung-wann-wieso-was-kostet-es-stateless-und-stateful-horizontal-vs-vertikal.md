---
advertiser: ""
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/3dbb8bc3-645b-4b36-9e8b-67c617d69375/engineering-kiosk-56-applikations-skalierung-wann-wieso-was-kostet-es-stateless-und-stateful-horizontal-vs-vertikal
apple_podcasts: https://podcasts.apple.com/us/podcast/56-applikations-skalierung-wann-wieso-was-kostet-es/id1603082924?i=1000597357438&uo=4
audio: https://audio1.redcircle.com/episodes/c3af9548-59fe-4cf7-934a-878aa505b686/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:00:35
  title: Das Märchen der Skalierung und meine Datenbank skaliert nicht
- start: 00:02:55
  title: Was ist Skalierung?
- start: 00:06:45
  title: Braucht man Skalierung überhaupt? Wer muss skalieren?
- start: 00:12:41
  title: Es ist cool auf Scale zu arbeiten
- start: 00:16:23
  title: Wenn wir skalieren können, sparen wir Geld
- start: 00:20:50
  title: Stateless vs. Stateful-Systeme
- start: 00:31:43
  title: Horizontaler vs. Vertikaler skalierung
- start: 00:35:38
  title: Ab wann skaliere ich die Hardware oder optimiere die Applikation?
- start: 00:39:24
  title: Gesteigerte Komplexität durch Skalierung
- start: 00:42:42
  title: Was braucht ihr, um skalieren zu können bzw. damit anzufangen?
- start: 00:48:49
  title: Outro
deezer: https://www.deezer.com/episode/474777487
description: "Die App muss skalieren. Das kann doch nicht so schwer sein, oder? Sekundenschnelles und automatisches Hochskalieren bei einem erhöhten Traffic-Aufkommen. So oder so ähnlich versprechen es die Cloud-Hyperscaler in ihren Marketing-Texten. Das erweckt oft den Anschein, dass das Ganze gar nicht so schwer sein kann. Doch ist dies auch in der Realität so? Eine Applikation skalierbar zu gestalten, ist bei weitem nicht einfach. Stichworte wie Ausfallsicherheit, vertikale- oder horizontale Skalierung, Stateless- oder Stateful-Applications, Loadbalancer und Auto-Discovery, Kubernetes und zusätzliche Code-Komplexität, finanzieller Impact, Load-Tests, Request-Deadlines, Chaos Monkey und Down-Scaling. Alles Begriffe, die damit in Verbindung stehen und einen wichtigen Bestandteil ausmachen. In dieser Episode geben wir einen Überblick über das Thema Application-Skalierung: Was ist das? Wer braucht es? Was sind die Grundbegriffe und welche Grundlagen müssen erfüllt werden, damit das ganze erfolgreich wird? Bonus: Warum Andy eine Märchenstimme hat und warum wir Milchmädchenrechnung nicht mehr sagen sollten."
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./56-applikations-skalierung-wann-wieso-was-kostet-es-stateless-und-stateful-horizontal-vs-vertikal.jpg
length_second: 3026
pubDate: 2023-01-31 05:00:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: B
- name: Wolfi Gassler
  transcriptLetter: A
spotify: https://open.spotify.com/episode/2QGybmHnaPnv8LhwMCKmDW
tags:
- Backend
- DevOps
- Cloud
title: "#56 Applikations-Skalierung: Wann, wieso, was kostet es? Stateless und Stateful, Horizontal vs. Vertikal"
transcript_slim: src/data/transcripts/56-transcript-slim.json
youtube: https://www.youtube.com/watch?v=fpWfjz3YjPg
---
<p><span>Die App muss skalieren. Das kann doch nicht so schwer sein, oder?</span></p><p><span>Sekundenschnelles und automatisches Hochskalieren bei einem erhöhten Traffic-Aufkommen. So oder so ähnlich versprechen es die Cloud-Hyperscaler in ihren Marketing-Texten. Das erweckt oft den Anschein, dass das Ganze gar nicht so schwer sein kann. Doch ist dies auch in der Realität so? Eine Applikation skalierbar zu gestalten, ist bei weitem nicht einfach. Stichworte wie Ausfallsicherheit, vertikale- oder horizontale Skalierung, Stateless- oder Stateful-Applications, Loadbalancer und Auto-Discovery, Kubernetes und zusätzliche Code-Komplexität, finanzieller Impact, Load-Tests, Request-Deadlines, Chaos Monkey und Down-Scaling. Alles Begriffe, die damit in Verbindung stehen und einen wichtigen Bestandteil ausmachen.</span></p><p><span>In dieser Episode geben wir einen Überblick über das Thema Application-Skalierung: Was ist das? Wer braucht es? Was sind die Grundbegriffe und welche Grundlagen müssen erfüllt werden, damit das ganze erfolgreich wird?</span></p><p><span>Bonus: Warum Andy eine Märchenstimme hat und warum wir Milchmädchenrechnung nicht mehr sagen sollten.</span></p><p><br></p><p><span>Feedback (gerne auch als Voice Message)</span></p><ul><li><span>Email: </span><a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li><span>Mastodon: </span><a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li><span>Twitter: </span><a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li><span>WhatsApp </span>+49 15678 136776</li></ul><p><br></p><p><span>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per </span><a href="https://engineeringkiosk.dev/kontakt/">Email</a><span> oder WhatsApp Voice Message an </span>+49 15678 136776</p><p><br></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p> </p><h3 id="links">Links</h3><ul><li><span>Rule of 40: </span><a href="https://aktien.guide/blog/rule-of-40-einfach-erklaert" rel="nofollow">https://aktien.guide/blog/rule-of-40-einfach-erklaert</a></li><li><span>Kubernetes: </span><a href="https://kubernetes.io/" rel="nofollow">https://kubernetes.io/</a></li><li><span>Amazon S3: </span><a href="https://aws.amazon.com/de/s3/" rel="nofollow">https://aws.amazon.com/de/s3/</a></li><li><span>Vitess: </span><a href="https://vitess.io/" rel="nofollow">https://vitess.io/</a></li><li><span>Ceph: </span><a href="https://ceph.io/" rel="nofollow">https://ceph.io/</a></li><li><span>Chaos Monkey: </span><a href="https://github.com/Netflix/chaosmonkey/" rel="nofollow">https://github.com/Netflix/chaosmonkey/</a></li><li><span>Zu Besuch bei Hetzner Datacenter: </span><a href="https://www.youtube.com/watch?v=F0KRLaw8Di8" rel="nofollow">https://www.youtube.com/watch?v=F0KRLaw8Di8</a></li><li><span>ProxySQL: </span><a href="https://proxysql.com/" rel="nofollow">https://proxysql.com/</a></li><li><span>PlanetScale: </span><a href="https://planetscale.com/" rel="nofollow">https://planetscale.com/</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p><span>(00:00:00) Intro</span></p><p><span>(00:00:35) Das Märchen der Skalierung und meine Datenbank skaliert nicht</span></p><p><span>(00:02:55) Was ist Skalierung?</span></p><p><span>(00:06:45) Braucht man Skalierung überhaupt? Wer muss skalieren?</span></p><p><span>(00:12:41) Es ist cool auf Scale zu arbeiten</span></p><p><span>(00:16:23) Wenn wir skalieren können, sparen wir Geld</span></p><p><span>(00:20:50) Stateless vs. Stateful-Systeme</span></p><p><span>(00:31:43) Horizontaler vs. Vertikaler skalierung</span></p><p><span>(00:35:38) Ab wann skaliere ich die Hardware oder optimiere die Applikation?</span></p><p><span>(00:39:24) Gesteigerte Komplexität durch Skalierung</span></p><p><span>(00:42:42) Was braucht ihr, um skalieren zu können bzw. damit anzufangen?</span></p><p><span>(00:48:49) Outro</span></p><p><br></p><h3 id="hosts">Hosts</h3><ul><li><span>Wolfgang Gassler (</span><a href="https://mastodon.social/@woolf" rel="nofollow">https://mastodon.social/@woolf</a><span>)</span></li><li><span>Andy Grunwald (</span><a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a><span>)</span></li></ul><p><span><span>﻿</span></span></p><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li><span>Email: </span><a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li><span>Mastodon: </span><a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li><span>Twitter: </span><a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li><span>WhatsApp </span>+49 15678 136776</li></ul><p><br></p>