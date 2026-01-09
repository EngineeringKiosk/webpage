---
advertiser: ""
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/eee8260f-3f06-4b10-9514-e114099f3024/engineering-kiosk-38-monitoring-metriken-tracing-alerting-observability
apple_podcasts: https://podcasts.apple.com/us/podcast/38-monitoring-metriken-tracing-alerting-observability/id1603082924?i=1000580748597&uo=4
audio: https://audio1.redcircle.com/episodes/d77dbefc-8553-4a49-81c2-27f9232b88ec/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:00:50
  title: Wolfgangs MySQL-Buch
- start: 00:02:11
  title: "Heutiges Thema: Wie würde Andy die Themen Monitoring, Alerting, Metriken und Logging bei einem Side Projekt angehen?"
- start: 00:04:49
  title: Warum brauchst du Logging, Monitoring, Metriken und Tracing?
- start: 00:07:29
  title: Logging von Exceptions, Warnings und anderen Fehler, Logging und der ELK-Stack
- start: 00:16:06
  title: Was sollte man eigentlich alles loggen?
- start: 00:19:22
  title: Log-Rotation und Log-Retention auf Object-Storage
- start: 00:27:30
  title: Metriken mit Prometheus
- start: 00:31:46
  title: Visualisierung von Metriken mit Grafana
- start: 00:34:25
  title: Intelligente Alerting Systeme und die richtigen Schwellenwerte finden
- start: 00:38:47
  title: Alerts senden und anrufen lassen
- start: 00:43:22
  title: "Tracing: Was ist das und brauchen wir das?"
- start: 00:48:49
  title: Was ist Observability?
- start: 00:51:42
  title: Iterativer Aufbau seiner Plattform und Alternativen
- start: 00:54:49
  title: Keine bezahlte Werbung
- start: 00:55:14
  title: Outro und Feedback
deezer: https://www.deezer.com/episode/437316457
description: "Wie würde heutzutage ein moderner Logging, Metriken, Monitoring, Alerting und Tracing-Stack aussehen? Im Infrastruktur-Bereich gibt es zu jedem Bereich etliche Tools. Cloud-Native ist das Buzzword der Stunde. In dieser Episode erzählt Andy, wie er einen modernen Stack für ein Side-Projekt für die Bereiche Logging, Metriken, Monitoring, Alerting und Tracing aufsetzen würde. Unter anderem geht es dabei um Fragen wie: Was sollte man eigentlich alles loggen? Wie kann man von einem Alert angerufen werden? Wie visualisiert man Daten in schönen Graphen? Brauchen wir Tracing? Und was ist Observability? Bonus: Engineering Porn und Buzzword-Bingo."
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./38-monitoring-metriken-tracing-alerting-observability.jpg
pubDate: 2022-09-27 04:00:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: B
- name: Wolfi Gassler
  transcriptLetter: A
spotify: https://open.spotify.com/episode/6o12Q11f95KuKYytqmis6u
tags:
- Backend
- DevOps
- Software Engineering
- Cloud
title: "#38 Monitoring, Metriken, Tracing, Alerting, Observability"
transcript_slim: src/data/transcripts/38-transcript-slim.json
youtube: https://www.youtube.com/watch?v=mnsIq8msouo
---
<p>Wie würde heutzutage ein moderner Logging, Metriken, Monitoring, Alerting und Tracing-Stack aussehen?</p><p>Im Infrastruktur-Bereich gibt es zu jedem Bereich etliche Tools. Cloud-Native ist das Buzzword der Stunde. In dieser Episode erzählt Andy, wie er einen modernen Stack für ein Side-Projekt für die Bereiche Logging, Metriken, Monitoring, Alerting und Tracing aufsetzen würde. Unter anderem geht es dabei um Fragen wie: Was sollte man eigentlich alles loggen? Wie kann man von einem Alert angerufen werden? Wie visualisiert man Daten in schönen Graphen? Brauchen wir Tracing? Und was ist Observability?</p><p>Bonus: Engineering Porn und Buzzword-Bingo.</p><p><br></p><p>Feedback (gerne auch als Voice Message)</p><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per <a href="https://engineeringkiosk.dev/kontakt/">Email</a> oder WhatsApp Voice Message an +49 15678 136776</p><p><br></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p> </p><h3 id="links">Links</h3><ul><li>Episode #37 Mit IT-Büchern Geld verdienen? Wer liest überhaupt noch Bücher?: <a href="https://engineeringkiosk.dev/podcast/episode/37-mit-it-b%C3%BCchern-geld-verdienen-wer-liest-%C3%BCberhaupt-noch-b%C3%BCcher/?pkn=shownotes">https://engineeringkiosk.dev/podcast/episode/37-mit-it-b%C3%BCchern-geld-verdienen-wer-liest-%C3%BCberhaupt-noch-b%C3%BCcher/?pkn=shownotes</a> </li><li>Episode #17 Was können wir beim Incident Management von der Feuerwehr lernen?: <a href="https://engineeringkiosk.dev/podcast/episode/17-was-k%C3%B6nnen-wir-beim-incident-management-von-der-feuerwehr-lernen/?pkn=shownotes">https://engineeringkiosk.dev/podcast/episode/17-was-k%C3%B6nnen-wir-beim-incident-management-von-der-feuerwehr-lernen/?pkn=shownotes</a> </li><li>Sentry: <a href="https://sentry.io/" rel="nofollow">https://sentry.io/</a></li><li>Datadog: <a href="https://www.datadoghq.com/" rel="nofollow">https://www.datadoghq.com/</a></li><li>Splunk: <a href="https://www.splunk.com/" rel="nofollow">https://www.splunk.com/</a></li><li>Elasticsearch: <a href="https://www.elastic.co/de/enterprise-search/" rel="nofollow">https://www.elastic.co/de/enterprise-search/</a></li><li>Logstash: <a href="https://github.com/elastic/logstash" rel="nofollow">https://github.com/elastic/logstash</a></li><li>Kibana: <a href="https://github.com/elastic/kibana" rel="nofollow">https://github.com/elastic/kibana</a></li><li>OpenSearch: <a href="https://opensearch.org/" rel="nofollow">https://opensearch.org/</a></li><li>Elastic Cloud: <a href="https://www.elastic.co/de/cloud/" rel="nofollow">https://www.elastic.co/de/cloud/</a></li><li>Aiven: <a href="https://aiven.io/" rel="nofollow">https://aiven.io/</a></li><li>Fluentd: <a href="https://www.fluentd.org/" rel="nofollow">https://www.fluentd.org/</a></li><li>Amazon S3 und S3 Glacier: <a href="https://aws.amazon.com/de/s3/" rel="nofollow">https://aws.amazon.com/de/s3/</a></li><li>Amazon Athena: <a href="https://aws.amazon.com/de/athena/" rel="nofollow">https://aws.amazon.com/de/athena/</a></li><li>Prometheus: <a href="https://prometheus.io/" rel="nofollow">https://prometheus.io/</a></li><li>VictoriaMetrics: <a href="https://github.com/VictoriaMetrics/VictoriaMetrics" rel="nofollow">https://github.com/VictoriaMetrics/VictoriaMetrics</a></li><li>InfluxDB: <a href="https://www.influxdata.com/" rel="nofollow">https://www.influxdata.com/</a></li><li>M3 Metrics Engine: <a href="https://m3db.io/" rel="nofollow">https://m3db.io/</a></li><li>Prometheus Node Exporter: <a href="https://github.com/prometheus/node_exporter" rel="nofollow">https://github.com/prometheus/node_exporter</a></li><li>Grafana: <a href="https://github.com/grafana/grafana" rel="nofollow">https://github.com/grafana/grafana</a></li><li>PromQL: <a href="https://prometheus.io/docs/prometheus/latest/querying/basics/" rel="nofollow">https://prometheus.io/docs/prometheus/latest/querying/basics/</a></li><li>OpsGenie: <a href="https://www.atlassian.com/de/software/opsgenie" rel="nofollow">https://www.atlassian.com/de/software/opsgenie</a></li><li>Jaeger: <a href="https://www.jaegertracing.io/" rel="nofollow">https://www.jaegertracing.io/</a></li><li>Zipkin: <a href="https://zipkin.io/" rel="nofollow">https://zipkin.io/</a></li><li>OpenTracing: <a href="https://opentracing.io/" rel="nofollow">https://opentracing.io/</a></li><li>OpenTelemetry: <a href="https://opentelemetry.io/" rel="nofollow">https://opentelemetry.io/</a></li><li>yak shaving: <a href="https://seths.blog/2005/03/dont_shave_that/" rel="nofollow">https://seths.blog/2005/03/dont_shave_that/</a></li><li>Cloud Native Computing Foundation: <a href="https://www.cncf.io/" rel="nofollow">https://www.cncf.io/</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) Intro</p><p>(00:00:50) Wolfgangs MySQL-Buch</p><p>(00:02:11) Heutiges Thema: Wie würde Andy die Themen Monitoring, Alerting, Metriken und Logging bei einem Side Projekt angehen?</p><p>(00:04:49) Warum brauchst du Logging, Monitoring, Metriken und Tracing?</p><p>(00:07:29) Logging von Exceptions, Warnings und anderen Fehler, Logging und der ELK-Stack</p><p>(00:16:06) Was sollte man eigentlich alles loggen?</p><p>(00:19:22) Log-Rotation und Log-Retention auf Object-Storage</p><p>(00:27:30) Metriken mit Prometheus</p><p>(00:31:46) Visualisierung von Metriken mit Grafana</p><p>(00:34:25) Intelligente Alerting Systeme und die richtigen Schwellenwerte finden</p><p>(00:38:47) Alerts senden und anrufen lassen</p><p>(00:43:22) Tracing: Was ist das und brauchen wir das?</p><p>(00:48:49) Was ist Observability?</p><p>(00:51:42) Iterativer Aufbau seiner Plattform und Alternativen</p><p>(00:54:49) Keine bezahlte Werbung</p><p>(00:55:14) Outro und Feedback</p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://twitter.com/schafele" rel="nofollow">https://twitter.com/schafele</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul>