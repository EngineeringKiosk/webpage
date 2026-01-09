---
advertiser: trivago
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/b5fcebfa-b7b8-4558-892f-7ab0f0947bbe/engineering-kiosk-91-konsistent-verf%C3%BCgbar-ausfalltolerant-oder-performant-das-cap--und-pacelc-theorem-in-verteilten-systemen
apple_podcasts: https://podcasts.apple.com/us/podcast/91-konsistent-verf%C3%BCgbar-ausfalltolerant-oder-performant/id1603082924?i=1000629992728&uo=4
audio: https://audio1.redcircle.com/episodes/2f6b4d3d-af38-4f97-ad65-5cc65eaaf79a/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:01:16
  title: Das CAP-Theorem
- start: 00:03:37
  title: Verteilte Systeme, Datenhaltung und State
- start: 00:08:17
  title: Was ist das CAP-Theorem?
- start: 00:09:52
  title: Info/Werbung
- start: 00:11:15
  title: Was ist das CAP-Theorem?
- start: 00:13:10
  title: "Das C in CAP: Consistency"
- start: 00:15:12
  title: "Das A in CAP: Availability"
- start: 00:16:16
  title: "Das P in CAP: Partition Tolerance"
- start: 00:18:55
  title: Der Problemfall in verteilten Systemen
- start: 00:20:41
  title: Was bedeutet es, wenn ein verteiltes System alle drei Eigenschaften erfüllen würde
- start: 00:22:56
  title: Wo kommt das CAP-Theorem her? Das Brewer Theorem, ACID und BASE
- start: 00:31:50
  title: "Ausprägung CP: Consistency und Partition Tolerance"
- start: 00:33:31
  title: "Ausprägung AP: Availability und Partition Tolerance"
- start: 00:36:34
  title: "Ausprägung AC: Availability und Consistency"
- start: 00:37:16
  title: "Das CAP-Theorem im Fehlerfall: PACELC"
- start: 00:41:07
  title: Das CAP-Theorem und die Praxis
- start: 00:43:00
  title: MongoDB
- start: 00:45:13
  title: Cassandra
- start: 00:46:36
  title: Kubernetes
- start: 00:49:38
  title: Kafka
deezer: https://www.deezer.com/episode/562049411
description: "Konsistent, Verfügbar und Ausfalltolerant: Wähle zwei - Das CAP-Theorem Stellt euch vor, ein Handwerker könnte die Dienstleistung schnell, günstig und in hoher Qualität leisten. Wäre dies nicht ein Traum? Leider sind alle drei Eigenschaften in der Realität nicht möglich. Und genau so geht es uns mit dem CAP-Theorem in verteilten Systemen mit Datenhaltung. Speziell im aktuellen Zeitalter mit Cloud Computing, horizontaler Skalierung, weltweiter Verfügbarkeit spielt das CAP-Theorem eine essentielle Rolle. Wie soll sich dein System verhalten, wenn die Netzwerk-Verbindung zwischen deinen Compute-Knoten ausfällt? Muss die Datenhaltung konsistent bleiben? Oder sind Inkonsistenzen für eine gewisse Zeit OK, dafür hat die Verfügbarkeit eine höhere Priorität? Um diesen Konflikt geht es in dieser Episode. Bonus: Auf GCP kannst du deine Compute-Instanz auf maximal 12 TB SSD Disk (ohne Netzwerk-Storage) aufstocken.  Das schnelle Feedback zur Episode: 👍 (top)\u00a0👎 (geht so)"
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./91-konsistent-verfügbar-ausfalltolerant-oder-performant-das-cap-und-pacelc-theorem-in-verteilten-systemen.jpg
pubDate: 2023-10-03 04:00:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: B
- name: Wolfi Gassler
  transcriptLetter: A
spotify: https://open.spotify.com/episode/2dEBTNiNu3C6kDenwK9eFf
tags:
- Backend
- Software Engineering
- Datenbanken
- Cloud
title: "#91 Konsistent, Verfügbar, Ausfalltolerant oder Performant: Das CAP- und PACELC-Theorem in verteilten Systemen"
transcript_slim: src/data/transcripts/91-transcript-slim.json
youtube: https://youtu.be/Uk4PzoF5sAg
---
<p>Konsistent, Verfügbar und Ausfalltolerant: Wähle zwei - Das CAP-Theorem</p><p>Stellt euch vor, ein Handwerker könnte die Dienstleistung schnell, günstig und in hoher Qualität leisten. Wäre dies nicht ein Traum? Leider sind alle drei Eigenschaften in der Realität nicht möglich. Und genau so geht es uns mit dem CAP-Theorem in verteilten Systemen mit Datenhaltung. Speziell im aktuellen Zeitalter mit Cloud Computing, horizontaler Skalierung, weltweiter Verfügbarkeit spielt das CAP-Theorem eine essentielle Rolle.</p><p>Wie soll sich dein System verhalten, wenn die Netzwerk-Verbindung zwischen deinen Compute-Knoten ausfällt? Muss die Datenhaltung konsistent bleiben? Oder sind Inkonsistenzen für eine gewisse Zeit OK, dafür hat die Verfügbarkeit eine höhere Priorität? Um diesen Konflikt geht es in dieser Episode.</p><p>Bonus: Auf GCP kannst du deine Compute-Instanz auf maximal 12 TB SSD Disk (ohne Netzwerk-Storage) aufstocken.</p><p><br></p><p><strong>Das schnelle Feedback zur Episode:</strong></p><p><a href="https://api.openpodcast.dev/feedback/91/upvote" rel="nofollow"><strong>👍 (top)</strong></a><strong> </strong><a href="https://api.openpodcast.dev/feedback/91/downvote" rel="nofollow"><strong>👎 (geht so)</strong></a></p><p><br></p><p>Feedback (gerne auch als Voice Message)</p><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per <a href="https://engineeringkiosk.dev/kontakt/">Email</a> oder WhatsApp Voice Message an +49 15678 136776</p><p><br></p><h3 id="links">Links</h3><ul><li>CAP-Theorem: <a href="https://de.wikipedia.org/wiki/CAP-Theorem" rel="nofollow">https://de.wikipedia.org/wiki/CAP-Theorem</a></li><li>Eric Brewer: <a href="https://de.wikipedia.org/wiki/Eric_Brewer_(Informatiker)" rel="nofollow">https://de.wikipedia.org/wiki/Eric_Brewer_(Informatiker)</a></li><li>Slides zum CAP Theorem (Towards robust distributed systems): <a href="https://people.eecs.berkeley.edu/~brewer/cs262b-2004/PODC-keynote.pdf" rel="nofollow">https://people.eecs.berkeley.edu/~brewer/cs262b-2004/PODC-keynote.pdf</a></li><li>PACELC theorem: <a href="https://en.wikipedia.org/wiki/PACELC_theorem" rel="nofollow">https://en.wikipedia.org/wiki/PACELC_theorem</a></li><li>CAP Theorem Explained: <a href="https://learncloudnative.com/blog/2019-11-22-basics_cap_theorem" rel="nofollow">https://learncloudnative.com/blog/2019-11-22-basics_cap_theorem</a></li><li>Jepsen: Elasticsearch 1.5.0: <a href="https://aphyr.com/posts/323-jepsen-elasticsearch-1-5-0" rel="nofollow">https://aphyr.com/posts/323-jepsen-elasticsearch-1-5-0</a></li><li>Jepsen: MongoDB stale reads: <a href="https://aphyr.com/posts/322-jepsen-mongodb-stale-reads" rel="nofollow">https://aphyr.com/posts/322-jepsen-mongodb-stale-reads</a></li><li>Jepsen: Cassandra: <a href="https://aphyr.com/posts/294-jepsen-cassandra" rel="nofollow">https://aphyr.com/posts/294-jepsen-cassandra</a></li><li>Jepsen: Kafka: <a href="https://aphyr.com/posts/293-jepsen-kafka" rel="nofollow">https://aphyr.com/posts/293-jepsen-kafka</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) Intro</p><p>(00:01:16) Das CAP-Theorem</p><p>(00:03:37) Verteilte Systeme, Datenhaltung und State</p><p>(00:08:17) Was ist das CAP-Theorem?</p><p>(00:09:52) Info/Werbung</p><p>(00:11:15) Was ist das CAP-Theorem?</p><p>(00:13:10) Das C in CAP: Consistency</p><p>(00:15:12) Das A in CAP: Availability</p><p>(00:16:16) Das P in CAP: Partition Tolerance</p><p>(00:18:55) Der Problemfall in verteilten Systemen</p><p>(00:20:41) Was bedeutet es, wenn ein verteiltes System alle drei Eigenschaften erfüllen würde</p><p>(00:22:56) Wo kommt das CAP-Theorem her? Das Brewer Theorem, ACID und BASE</p><p>(00:31:50) Ausprägung CP: Consistency und Partition Tolerance</p><p>(00:33:31) Ausprägung AP: Availability und Partition Tolerance</p><p>(00:36:34) Ausprägung AC: Availability und Consistency</p><p>(00:37:16) Das CAP-Theorem im Fehlerfall: PACELC</p><p>(00:41:07) Das CAP-Theorem und die Praxis</p><p>(00:43:00) MongoDB</p><p>(00:45:13) Cassandra</p><p>(00:46:36) Kubernetes</p><p>(00:49:38) Kafka</p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://mastodon.social/@woolf" rel="nofollow">https://mastodon.social/@woolf</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul>