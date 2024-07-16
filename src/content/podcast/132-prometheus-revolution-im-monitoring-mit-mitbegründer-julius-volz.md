---
advertiser: ''
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/49dd37ed-872e-421f-b85a-c2f62145e38f/engineering-kiosk-132-prometheus-revolution-im-monitoring-mit-mitbegr%C3%BCnder-julius-volz
apple_podcasts: https://podcasts.apple.com/us/podcast/132-prometheus-revolution-im-monitoring-mit-mitbegr%C3%BCnder/id1603082924?i=1000662365264&uo=4
audio: https://audio1.redcircle.com/episodes/0ddae330-a8a8-4fe6-a334-4a50199531c1/stream.mp3
chapter:
- start: 00:00:00
  title: Prometheus mit Julius Volz
- start: 00:07:58
  title: Was ist Prometheus?
- start: 00:16:24
  title: Observability, Service Discovery
- start: 00:21:04
  title: Selbstentwicklung eines Monitoring-Tools innerhalb einer Audio-Firma
- start: 00:27:33
  title: MVP und Inspiration von Borgmon
- start: 00:34:17
  title: Pull- vs. Push-Modell
- start: 00:53:28
  title: PromQL und der Vergleich zu SQL
- start: 01:01:59
  title: Visualisierung von Metriken
- start: 01:04:48
  title: Flaws in Prometheus
- start: 01:13:30
  title: Wie steige ich ins Thema Prometheus ein?
deezer: https://www.deezer.com/episode/651867812
description: "\xDCberwachen von Applikationen in Zeiten von dynamischer Infrastruktur\
  \ Cloud hier, Serverless da, Container-Scheduler dort. In Zeiten von dynamischen\
  \ Infrastrukturen wei\xDF man gar nicht mehr so genau, auf welchem Server und Port\
  \ deine Applikation eigentlich l\xE4uft. Dies wirft die gro\xDFe Frage auf: Wie\
  \ \xFCberwache ich meine Applikation denn eigentlich so ordentlich, dass ich sicherstellen\
  \ kann, dass diese so funktioniert, wie ich mir das initial gedacht habe? Die Antwort\
  \ dreht sich oft um den de facto Standard im Cloud Native Monitoring-Segment: Prometheus.\
  \ In dieser Episode sprechen wir mit Julius Volz, einem der zwei initialen Autoren\
  \ von Prometheus. Mit ihm sprechen wir \xFCber die Entstehungsgeschichte von Prometheus\
  \ bei SoundCloud, wie sich das System von traditionellen Monitoring-Systemen unterscheidet,\
  \ warum mit PromQL eine eigene Query-Language ins leben gerufen wurde aber auch\
  \ welche Flaws er nach 12 Jahren Entwicklung gerne beheben w\xFCrde. Bonus: Wer\
  \ kennt noch Nagios, Ganglia oder Graphite?  Das schnelle Feedback zur Episode:\
  \ \U0001F44D (top)\_\U0001F44E (geht so)  Feedback EngKiosk Community: https://engineeringkiosk.dev/join-discord\_\
  Buy us a coffee: https://engineeringkiosk.dev/kaffeeEmail: stehtisch@engineeringkiosk.devLinkedIn:\
  \ https://www.linkedin.com/company/engineering-kiosk/Mastodon: https://podcasts.social/@engkioskTwitter:\
  \ https://twitter.com/EngKiosk Gerne behandeln wir auch euer Audio Feedback in einer\
  \ der n\xE4chsten Episoden, einfach die Audiodatei per Email an stehtisch@engineeringkiosk.dev.\
  \  LinksJulius Volz: https://juliusv.com/Prometheus: https://prometheus.io/PromLabs\
  \ YouTube Channel: https://www.youtube.com/@PromLabsPromLabs Trainngs: https://training.promlabs.com/Large-scale\
  \ cluster management at Google with Borg: https://static.googleusercontent.com/media/research.google.com/de//pubs/archive/43438.pdfSoundcloud:\
  \ https://soundcloud.com/Prometheus: The Documentary https://www.youtube.com/watch?v=rT4fJNbfe14\_\
  Demo-Metriken: https://demo.promlabs.com/metricsEngineering Kiosk Episode #101 Observability\
  \ und OpenTelemetry mit Severin Neumann: https://engineeringkiosk.dev/ep101 Sprungmarken(00:00:00)\
  \ Prometheus mit Julius Volz (00:07:58) Was ist Prometheus? (00:16:24) Observability,\
  \ Service Discovery (00:21:04) Selbstentwicklung eines Monitoring-Tools innerhalb\
  \ einer Audio-Firma (00:27:33) MVP und Inspiration von Borgmon (00:34:17) Pull-\
  \ vs. Push-Modell (00:53:28) PromQL und der Vergleich zu SQL (01:01:59) Visualisierung\
  \ von Metriken (01:04:48) Flaws in Prometheus (01:13:30) Wie steige ich ins Thema\
  \ Prometheus ein?  HostsWolfgang Gassler (https://mastodon.social/@woolf)Andy Grunwald\
  \ (https://twitter.com/andygrunwald) FeedbackEngKiosk Community: https://engineeringkiosk.dev/join-discord\_\
  Buy us a coffee: https://engineeringkiosk.dev/kaffeeEmail: stehtisch@engineeringkiosk.devLinkedIn:\
  \ https://www.linkedin.com/company/engineering-kiosk/Mastodon: https://podcasts.social/@engkioskTwitter:\
  \ https://twitter.com/EngKiosk"
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback::Feedback
image: "./132-prometheus-revolution-im-monitoring-mit-mitbegr\xFCnder-julius-volz.jpg"
length_second: 4611
pubDate: 2024-07-16 04:00:00+00:00
rtlplus: ''
six_user_needs:
- Update me
- Educate me
speaker:
- name: Andy Grunwald
  transcriptLetter: A
  website: https://andygrunwald.com/
- name: Wolfi Gassler
  transcriptLetter: B
  website: https://wolfgang.gassler.org/
- name: Julius Volz
  transcriptLetter: C
  website: https://juliusv.com/
spotify: https://open.spotify.com/episode/0RzSEPPwhMfEHjUILDTtyz
tags:
- Cloud
- DevOps
- Backend
- Interview
title: "#132 Prometheus: Revolution im Monitoring mit Mitbegr\xFCnder Julius Volz"
transcript_raw: src/data/transcripts/132-transcript.zip
transcript_slim: src/data/transcripts/132-transcript-slim.json
youtube: https://www.youtube.com/watch?v=J9ugvkek92o

---
<p>Überwachen von Applikationen in Zeiten von dynamischer Infrastruktur</p><p>Cloud hier, Serverless da, Container-Scheduler dort. In Zeiten von dynamischen Infrastrukturen weiß man gar nicht mehr so genau, auf welchem Server und Port deine Applikation eigentlich läuft. Dies wirft die große Frage auf: Wie überwache ich meine Applikation denn eigentlich so ordentlich, dass ich sicherstellen kann, dass diese so funktioniert, wie ich mir das initial gedacht habe?</p><p>Die Antwort dreht sich oft um den de facto Standard im Cloud Native Monitoring-Segment: Prometheus.</p><p>In dieser Episode sprechen wir mit Julius Volz, <span>einem der zwei initialen Autoren</span> von Prometheus.</p><p>Mit ihm sprechen wir über die Entstehungsgeschichte von Prometheus bei SoundCloud, wie sich das System von traditionellen Monitoring-Systemen unterscheidet, warum mit PromQL eine eigene Query-Language ins leben gerufen wurde aber auch welche Flaws er nach 12 Jahren Entwicklung gerne beheben würde.</p><p>Bonus: Wer kennt noch Nagios, Ganglia oder Graphite?</p><p><br></p><p><strong>Das schnelle Feedback zur Episode:</strong></p><p><a href="https://api.openpodcast.dev/feedback/132/upvote" rel="nofollow"><strong>👍 (top)</strong></a><strong> </strong><a href="https://api.openpodcast.dev/feedback/132/downvote" rel="nofollow"><strong>👎 (geht so)</strong></a></p><p><br></p><p>Feedback</p><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>Buy us a coffee: <a href="https://engineeringkiosk.dev/kaffee">https://engineeringkiosk.dev/kaffee</a></li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>LinkedIn: <a href="https://www.linkedin.com/company/engineering-kiosk/" rel="nofollow">https://www.linkedin.com/company/engineering-kiosk/</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach die Audiodatei per<a href="https://engineeringkiosk.dev/kontakt/"> Email</a> an <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a>.</p><p><br></p><h3 id="links">Links</h3><ul><li>Julius Volz: <a href="https://juliusv.com/" rel="nofollow">https://juliusv.com/</a></li><li>Prometheus: <a href="https://prometheus.io/" rel="nofollow">https://prometheus.io/</a></li><li>PromLabs YouTube Channel: <a href="https://www.youtube.com/@PromLabs" rel="nofollow">https://www.youtube.com/@PromLabs</a></li><li>PromLabs Trainngs: <a href="https://training.promlabs.com/" rel="nofollow">https://training.promlabs.com/</a></li><li>Large-scale cluster management at Google with Borg: <a href="https://static.googleusercontent.com/media/research.google.com/de//pubs/archive/43438.pdf" rel="nofollow">https://static.googleusercontent.com/media/research.google.com/de//pubs/archive/43438.pdf</a></li><li>Soundcloud: <a href="https://soundcloud.com/" rel="nofollow">https://soundcloud.com/</a></li><li>Prometheus: The Documentary <a href="https://www.youtube.com/watch?v=rT4fJNbfe14" rel="nofollow">https://www.youtube.com/watch?v=rT4fJNbfe14</a> </li><li>Demo-Metriken: <a href="https://demo.promlabs.com/metrics" rel="nofollow">https://demo.promlabs.com/metrics</a></li><li>Engineering Kiosk Episode #101 Observability und OpenTelemetry mit Severin Neumann: <a href="https://engineeringkiosk.dev/ep101">https://engineeringkiosk.dev/ep101</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) Prometheus mit Julius Volz</p><p>(00:07:58) Was ist Prometheus?</p><p>(00:16:24) Observability, Service Discovery</p><p>(00:21:04) Selbstentwicklung eines Monitoring-Tools innerhalb einer Audio-Firma</p><p>(00:27:33) MVP und Inspiration von Borgmon</p><p>(00:34:17) Pull- vs. Push-Modell</p><p>(00:53:28) PromQL und der Vergleich zu SQL</p><p>(01:01:59) Visualisierung von Metriken</p><p>(01:04:48) Flaws in Prometheus</p><p>(01:13:30) Wie steige ich ins Thema Prometheus ein?</p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://mastodon.social/@woolf" rel="nofollow">https://mastodon.social/@woolf</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><h3 id="feedback">Feedback</h3><ul><li>EngKiosk Community: <a href="https://engineeringkiosk.dev/join-discord">https://engineeringkiosk.dev/join-discord</a> </li><li>Buy us a coffee: <a href="https://engineeringkiosk.dev/kaffee">https://engineeringkiosk.dev/kaffee</a></li><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>LinkedIn: <a href="https://www.linkedin.com/company/engineering-kiosk/" rel="nofollow">https://www.linkedin.com/company/engineering-kiosk/</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li></ul>