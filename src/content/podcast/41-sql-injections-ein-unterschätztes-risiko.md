---
advertiser: ""
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/1aa39b24-8c09-496d-9505-20c0ddb6981d/engineering-kiosk-41-sql-injections---ein-untersch%C3%A4tztes-risiko
apple_podcasts: https://podcasts.apple.com/us/podcast/41-sql-injections-ein-untersch%C3%A4tztes-risiko/id1603082924?i=1000582974713&uo=4
audio: https://audio1.redcircle.com/episodes/51d3390e-ffce-4ec6-a6d9-350dc0cd163b/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:00:42
  title: SQL-Injections aus den 90ern und die Vielfalt in Berlin
- start: 00:02:49
  title: "Das heutige Thema: Web-Security SQL-Injections in der Tiefe"
- start: 00:05:07
  title: Was sind SQL-Injections?
- start: 00:08:48
  title: Sind SQL-Injections auch im Jahr 2022 noch ein Problem?
- start: 00:10:56
  title: Wann gab es die erste SQL-Injection? Woher stammt diese Sicherheitslücke?
- start: 00:13:22
  title: Was sind die Gründe, dass SQL-Injections noch so ein großes Problem sind?
- start: 00:19:37
  title: "Verschiedene Arten von SQL-Injections: Output-Based, Error-Based, Blind-SQL-Injections, Time-Based-SQL-Injections, Out-of-Band-SQL-Injections"
- start: 00:27:42
  title: "Bug Bounty: 2-Channel SQL Injection-Attacke in Kombination mit Cross-Site-Scripting (XSS) bei trivago"
- start: 00:29:42
  title: Mehrstufige Attacken und Ausnutzung mehrerer Lücken nacheinander
- start: 00:33:16
  title: "Möglicher Schaden durch eine SQL-Injection: Daten verändern, Befehle auf dem Server ausführen, lokale Dateien lesen und schreiben, SQL-Funktionen ausführen, Denial of Service (DoS)"
- start: 00:39:09
  title: "Gegenmaßnahmen um SQL-Injections zu verhindern: Prepared Statements, Datenbank-Komponenten updaten, limitierte Rechte für Datenbank-User, Web Application Firewalls (WAF)"
- start: 00:56:42
  title: "Möglichkeiten um deine Anwendung automatisch zu testen: Unit-Tests, statische Analyse, dynamische Analyse mit sqlmap und Fuzzing"
- start: 01:02:51
  title: Maßnahmen um Sicherheit zu gewährleisten von Datenbank as a Service-Providern
- start: 01:06:51
  title: Outro
deezer: https://www.deezer.com/episode/444953397
description: "SQL Injections: Eine der weitverbreitetsten Sicherheitslücken im Web, auch im Jahr 2022 Der Großteil aller Applikationen interagiert in irgendeiner Art und Weise mit einer Datenbank. Deswegen werden die meisten Entwicklerinnen und Entwickler bereits von der Sicherheitslücke \"SQL Injection\" gehört haben. Seit 24 Jahren ist dies eine der weitverbreitetsten Sicherheitslücken im Internet und es ist kein Ende in Sicht. Was ist eigentlich eine SQL-Injection im Detail? Welche verschiedenen Arten gibt es? Was ist der Grund, dass uns dieses Einfallstor so lange beschäftigt? Woher kommt diese und wer hat sie entdeckt? Wie kann man sich schützen und seine Anwendung ausreichend testen? All das und noch viel mehr in dieser Episode. Bonus: Der Kontrast zwischen Duisburg und Berlin und wie die SQL-Injektion als Nebenprodukt entdeckt wurde."
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./41-sql-injections-ein-unterschätztes-risiko.jpg
length_second: 4117
pubDate: 2022-10-18 04:00:00+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: A
- name: Wolfi Gassler
  transcriptLetter: B
spotify: https://open.spotify.com/episode/33PP79EzBWrsQ5aZqqzlTw
tags:
- Backend
- Datenbanken
- Sicherheit
- Software Engineering
title: "#41 SQL Injections - Ein unterschätztes Risiko"
transcript_slim: src/data/transcripts/41-transcript-slim.json
youtube: https://www.youtube.com/watch?v=EPfWgGvFtM4
---
<p>SQL Injections: Eine der weitverbreitetsten Sicherheitslücken im Web, auch im Jahr 2022</p><p>Der Großteil aller Applikationen interagiert in irgendeiner Art und Weise mit einer Datenbank. Deswegen werden die meisten Entwicklerinnen und Entwickler bereits von der Sicherheitslücke &#34;SQL Injection&#34; gehört haben. Seit 24 Jahren ist dies eine der weitverbreitetsten Sicherheitslücken im Internet und es ist kein Ende in Sicht. Was ist eigentlich eine SQL-Injection im Detail? Welche verschiedenen Arten gibt es? Was ist der Grund, dass uns dieses Einfallstor so lange beschäftigt? Woher kommt diese und wer hat sie entdeckt? Wie kann man sich schützen und seine Anwendung ausreichend testen? All das und noch viel mehr in dieser Episode.</p><p>Bonus: Der Kontrast zwischen Duisburg und Berlin und wie die SQL-Injektion als Nebenprodukt entdeckt wurde.</p><p><br></p><p>Feedback (gerne auch als Voice Message)</p><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per <a href="https://engineeringkiosk.dev/kontakt/">Email</a> oder WhatsApp Voice Message an +49 15678 136776</p><p><br></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p> </p><h3 id="links">Links</h3><ul><li>Phrack Magazine Volume 8, Issue 54 Dec 25th, 1998, article 08 of 12: <a href="http://www.phrack.org/archives/issues/54/8.txt" rel="nofollow">http://www.phrack.org/archives/issues/54/8.txt</a></li><li>OWASP Top Ten 2021: <a href="https://owasp.org/www-project-top-ten/" rel="nofollow">https://owasp.org/www-project-top-ten/</a></li><li>CVE Details - Security Vulnerabilities Published In 2022(SQL Injection): <a href="https://www.cvedetails.com/vulnerability-list/year-2022/opsqli-1/sql-injection.html" rel="nofollow">https://www.cvedetails.com/vulnerability-list/year-2022/opsqli-1/sql-injection.html</a></li><li>Analyzing Prepared Statement Performance: <a href="https://orangematter.solarwinds.com/2014/11/19/analyzing-prepared-statement-performance/" rel="nofollow">https://orangematter.solarwinds.com/2014/11/19/analyzing-prepared-statement-performance/</a></li><li>SQL Injection Prevention Cheat Sheet: <a href="https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html" rel="nofollow">https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html</a></li><li>OWASP Top 10 (2021) - A03:2021 – Injection: <a href="https://owasp.org/Top10/A03_2021-Injection/" rel="nofollow">https://owasp.org/Top10/A03_2021-Injection/</a></li><li>CVE Details - Heartbleed (CVE-2014-0160): <a href="https://www.cvedetails.com/cve/CVE-2014-0160/" rel="nofollow">https://www.cvedetails.com/cve/CVE-2014-0160/</a></li><li>CVE Details - Log4Shell (CVE-2021-44228): <a href="https://www.cvedetails.com/cve/CVE-2021-44228/" rel="nofollow">https://www.cvedetails.com/cve/CVE-2021-44228/</a></li><li>xkcd &#34;Exploits of a Mom&#34;: <a href="https://xkcd.com/327/" rel="nofollow">https://xkcd.com/327/</a></li><li>HackerOne-Programm von trivago: <a href="https://hackerone.com/trivago" rel="nofollow">https://hackerone.com/trivago</a></li><li>Owncloud: <a href="https://owncloud.com/" rel="nofollow">https://owncloud.com/</a></li><li>TYPO3: <a href="https://typo3.org/" rel="nofollow">https://typo3.org/</a></li><li>Wordpress: <a href="https://wordpress.com/de/" rel="nofollow">https://wordpress.com/de/</a></li><li>SQL-Proxy: <a href="https://github.com/sysown/proxysql" rel="nofollow">https://github.com/sysown/proxysql</a></li><li>GitHub CodeQL: <a href="https://codeql.github.com/" rel="nofollow">https://codeql.github.com/</a></li><li>sqlmap: <a href="https://sqlmap.org/" rel="nofollow">https://sqlmap.org/</a></li><li>SQLi-Fuzzer: A SQL Injection Vulnerability Discovery Framework Based on Machine Learning: <a href="https://ieeexplore.ieee.org/document/9657925" rel="nofollow">https://ieeexplore.ieee.org/document/9657925</a></li><li>OWASP Zed Attack Proxy (ZAP): <a href="https://www.zaproxy.org/" rel="nofollow">https://www.zaproxy.org/</a></li><li>PlanetScale: <a href="https://planetscale.com/" rel="nofollow">https://planetscale.com/</a></li><li>Awesome static analysis: <a href="https://github.com/analysis-tools-dev/static-analysis" rel="nofollow">https://github.com/analysis-tools-dev/static-analysis</a></li></ul><p><br></p><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) Intro</p><p>(00:00:42) SQL-Injections aus den 90ern und die Vielfalt in Berlin</p><p>(00:02:49) Das heutige Thema: Web-Security SQL-Injections in der Tiefe</p><p>(00:05:07) Was sind SQL-Injections?</p><p>(00:08:48) Sind SQL-Injections auch im Jahr 2022 noch ein Problem?</p><p>(00:10:56) Wann gab es die erste SQL-Injection? Woher stammt diese Sicherheitslücke?</p><p>(00:13:22) Was sind die Gründe, dass SQL-Injections noch so ein großes Problem sind?</p><p>(00:19:37) Verschiedene Arten von SQL-Injections: Output-Based, Error-Based, Blind-SQL-Injections, Time-Based-SQL-Injections, Out-of-Band-SQL-Injections</p><p>(00:27:42) Bug Bounty: 2-Channel SQL Injection-Attacke in Kombination mit Cross-Site-Scripting (XSS) bei trivago</p><p>(00:29:42) Mehrstufige Attacken und Ausnutzung mehrerer Lücken nacheinander</p><p>(00:33:16) Möglicher Schaden durch eine SQL-Injection: Daten verändern, Befehle auf dem Server ausführen, lokale Dateien lesen und schreiben, SQL-Funktionen ausführen, Denial of Service (DoS)</p><p>(00:39:09) Gegenmaßnahmen um SQL-Injections zu verhindern: Prepared Statements, Datenbank-Komponenten updaten, limitierte Rechte für Datenbank-User, Web Application Firewalls (WAF)</p><p>(00:56:42) Möglichkeiten um deine Anwendung automatisch zu testen: Unit-Tests, statische Analyse, dynamische Analyse mit sqlmap und Fuzzing</p><p>(01:02:51) Maßnahmen um Sicherheit zu gewährleisten von Datenbank as a Service-Providern</p><p>(01:06:51) Outro</p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://twitter.com/schafele" rel="nofollow">https://twitter.com/schafele</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><p><br></p><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul>