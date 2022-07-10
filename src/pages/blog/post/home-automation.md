---
layout: ../../../layouts/blog-post.astro
title: "Der einfache Einstieg in die Home Automation"
subtitle: Dies ist ein Untertitel, der mal groß werden möchte.
description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum."
tags: [
    "Automation",
    "Zuhause",
    "Internet of Things"
]
date: 2022-05-29
thumbnail: /images/blog/macbook3.jpg
headerimage: /images/blog-content/content-photo3.jpg
draft: true
---



## Content zu verwerten

Episode 22: Home Automation


Sprungmarken


Links:







Wie man Huawei ausspricht

Sind wir noch ein Software Engineering Podcast?
Static Site Generator

Frontend ist komplex geworden

Relevant
- Internet of Things (IoT) ist überall
- Smart Devices hat jeder im Raum
- Software Engineers haben einen Hang zur Spielerei
- AUtomatisierung Automatisierung Automatisierung

Einstiegsdroge für viele
- Robotter: Staugsauber-Robotter, Mährobotter
- Andys Einstieg: Vernetzter Rauchmelder

Andys Staubsauger-Robotter mit China connected ist und der Grundriss nun in Chrina ist
Warum Teppich-Werbung für Andy relevant sein könnte

Wer möchte denn heutzutage den Lichtschalter

Anwendbar für existierende Strukturen
- Wohnung (gekauft oder gemietet)
- Haus (gekauft oder gemietet)
- WG Zimmer oder Hostel

Primärer Fokus: I.d.R. Kabellose Komponenten

Anwendungstechnisch: Hobby-Bereich, nichts professionelles

Ziel: Schnell einsteigen, möglichst lean los starten

Home Automation ist ein Zeit-Grab

Wolfgang und Andy im Influencer Game
Warum bald alle Höhrerinnen nackig durch die Wohnung laufen werden

Wie startet man the lean way?
Zum Ikea / Zum Lidl -> Kauft man sich ein Device und man hat bereits eine Automation
    -> Achtung: Beim Auswahl der Komponenten darauf achten, dass man das auch ohne Cloud betreiben kann
    -> zB Smarte Glühbrne
Zentraler Server (zB Raspberry Pi) und Home Assistant installieren
Zigbee Gateway (zB Deconz Conbee Dresden Elektronik)
 
Wahrscheinlich habt Ihr schon Devices
- iPhone / Android
- Modernes Auto (zB Kia)
- Amazon Alexa / Apple Homepod
- Soundsystem
- TV
- Drucker
- Fritzbox

Was ist Home Assistant?
- Platform für Hardware Komponenten und Protokolle
- Data Privacy first
- Python
- Einer der populärsten Projekte
- Holländer
- Hoher Fokus auf Benutzerfreundlichkeit
- Sehr stabil

Installations-Methoden von Home Assistant: Docker, Home Assistant OS, etc.

FHEM: Homematic Rauchmelder - Virtueller Rauchmelder

Herausforderungen
- Viele Protokolle im eigenen Netzwerk (viele EINgangstüren, baut Mesh auf oder nicht, ist stabil oder nicht stabil, Reichweite, Frequenzbereiche bei viel Nachbarschaft)
    - Tip: Macht euch Gedanken, wie viel Protokolle ihr haben wollt, unsere Empfehlung: Limitiert dies stark (Andy: WLAN, Zigbee und Homematic)
    - Gedankengang: Brauch ich einen Zusätzlichen Dongle / Ne Bride in das Netz oder kann ich das mit meinem vorhandenem Setup abbilden?

- Herausforderungen bei der Home Automation: (Manuelle) Nuztung trotz Ausfall (nach Restart oder Stromausfall, zB korrupte SD Karte)
    - Das ganze muss auch noch ohne Automation funktionieren
    - Andere Leute im Haushalt: Die Frau / Partner / Partnerin die nicht technisch sind

- Herausforderungen bei der Home Automation: Box of Shame
    - Geht tierisch ins Geld
    - Kiste im Haushalt mit unverbaute Home Automation Komponenten

- Herausforderungen bei der Home Automation: Ist das ganze ohne Internet / Cloud lauffähig?
    - zB Video Kamera zur Erkennung von Katze oder Fahrzeug

- Herausforderungen bei der Home Automation: Monitoring von Batterie-Komponenten
    - Können einfach ausfallen, weil die Batterie leer ist
    - Dann keine Daten mehr

Typische Fehler:
- Smartlampe wird der Strom weggenommen (weil man zB die Lampe ausgemacht hat/Schalter deaktiviert hat) -> Kann man nicht mehr einschalten


Wolfgangs setup
- Raspberry Pi mit Home Assistant
- Setzt sehr stark auf Zigbee
- Primär Lampen, Temperator- und Tür Sensoren und Heizungssteuerung

Andys Setup
- Intel NUC
- Home Assistant in Docker Container
- FHEM
- MQTT

Single Chip Comuter
- Raspberry Pi
- Ordroid
- Home Assistant Yellow

Zigbee
- "Dummes Protokoll"
- Mesh
- Geringe Reichweite
- Jedes Zigbee-Gerät ist auch ein Router / Repeater (sofern unter Spannung)
- Unidirektional
- Low Power Consumption Mesh
- Meisten Komponenten via Knopfzelle
- Gegründet von der Zigbee Alliance
- Bekommt man überall
- Wichtig: Man braucht die Zigbee Gateway von den Herstellern nicht! Solang Zigbee da steht, braucht man nicht

HomeMatic
- 8XX Mhz
- Bidirektional
- Meistens mit Batterien
- Ihr braucht einen eigenen Dongle

Z-Wave

Bluetooth
- Low Energy

W-LAN
- Bidirektional, TCP-IP
- Zieht enorm an der Batterie
- Beispiel: Shelly, Relai

Thread
- IPv6
- Recht neu
- Thread devices aktuell noch recht teuer (ca. Faktor 5 teurer)
- Eher so die Zukunft
- Allianz von Google, Amazon, Apple, Meta, Zigbee Alliance

Pro-Tip:
Testbrett bauen

Philip Wimmers wird bei Min. 40 erwähnt
Florian Hümbs wird bei Minute 54 erwähnt

Automation Usecases
- Regle die Heizung runter, wenn du nicht zuhause bist (Wolfgang)
- Batterie-Check wenn die Batterie unter 25% geht - Home Assistant Blueprint

Welche Automations haben wir implementiert?
- Heiz-Körper regulieren, wenn man nicht zuhause ist
- PV-Anlage? Waschmaschine nur starten, wenn es Sonnig ist
- Alarmanlage: Wenn die los geht, springt auch das Sound-System an
- Telegram Bot
- Außenbeleuchtung an und morgens automatisch aus
- Automation auf der Tür, wenn keiner zuhause
- 