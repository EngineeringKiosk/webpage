---
layout: ../../../layouts/blog-post.astro
title: "Was ist Infrastructure as Code?"
subtitle: Dies ist ein Untertitel, der mal groß werden möchte.
description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum."
tags: [
    "IaC",
    "Configuration Management"
]
date: 2022-05-29
thumbnail: /images/blog/macbook3.jpg
headerimage: /images/blog-content/content-photo3.jpg
draft: true
---



## Content zu verwerten

Episode 24: Infrastructure as Code


Old man yells at cloud - Oder: Wie managed man seine Infrastruktur mit Stil (und Software)

Anders als gewohnt nimmt in dieser Episode Andy die Dozenten-Rolle ein und beantwortet Wolfgang all seine Fragen zum Thema Infrastructure as Code. Wir klären wozu man das ganze eigentlich braucht, was Terraform und Pulumi ist, klären über einen weit verbreiteten Mythos auf, wo der Unterschied zwischen Infrastructure Orchestration und COnfiguration Management ist, was das beste Configuration Management Tool ist und wo es Herausforderungen bei der Verwendung von Infrastructure as Code gibt.

Bonus: Was ein Data Engineer ist, ob Wolfgang Holz-Klocks trägt und wie Deutschland mit dem 9€ Ticket umgeht.




Was ist eigentlich Infrastructure as Code?
-----------------------
- Cloud Infrastruktur
    Muss nicht AWS, GCP, Azure, Digital Ocean sein
    Kann auch Inhouse sein
    Man managed Hardware im Endeffekt
    Zur Einfachheit fokussieren wir uns auf Public Cloud Provider
    Cloud kann auch eine Sammlung von servern sein

Bedeutet eigentlich nur, dass du deine komplette Infrastrktur (Domains, Router, GitHub Repos, Server, Object Storage, user Management, User rechte management)
mit Quellcode managed - Netzwerk gehört auch dazu

Vorteil von Code hier?
- Klassischer Source Code
- Managebar im Software Development Lifecycle
- Versioniert, Testbar, Pull Requests, linter, etc.

Endziel:
- Definition von Infrastruktur
- Drückt auf Play
- Infrastruktur ist hochgefahren, 2 Server, etc.

Terraform
----------
HCL ist Human und Machine readable
Von HashiCorp
Inspiriert von der nginx configuration syntax

HCLs wie YAML
- Kann auch Schleifen und co - Language Features

Pulumi
----------
Nicht so populär wie Terraform

Warum kein Cloud Formation von AWS?
- Unterstützt die neusten produkte von AWS - Gibt es da kein delay
- Für Google ist es Terraform
- Terraforms Plugin (Provider) werden von den Platform-Betreibern selbst maintained
    zB AWS, GCP, GitHub, etc.


Was bringt mir Infrastructure as Code eigentlich?
-----------------------
- Dokumentation (vs. in der Web UI zusammen zu klicken)
- Ermöglicht die Team-Arbeit: Warum wurde dies geändert? (Git-Commits durch Versionierung)
- Security und Compliance
- Wiederverwendbarkeit (für andere Regionen oder andere Environments)
- Erholung von Hacker-Angriffen und Incidents
- Developer Experience: Infrastructure Management via Pull Request - Enabler für das DevOps mindset - Devs an System Adnins heranführen


Großer Mythos
-----------
Code ist einmal definiert, und auf alle Cloud Provider anwendbar
- Nein: Cloud spezifische Resourcen
- Verschiedene Settings pro Provider
- Kleinster gemeinsamer Nenner
- Du wählst ja einen Cloud Provider um das volle Potential ausnutzen

Bin ich schneller, wenn ich die Cloud migrieren möchte und die Infrastruktur bereits als Code definiert habe?
- Jede Cloud hat ihre eigenheiten
- Besonders im Bereich Access und Identity Management, Netzwerk und Projekt/Account Management
- Schneller? Ja, weil selbe Sprache (Terraform)
- Struktur recht gut Dokumentiert


Wo ist der Unterschied zwischen Ansible und Terraform?
Ansible ist nur ein Kandidat aus dem Bereich des Configuration Management. Andere sind Saltstack, Puppe, Chef, CFEngine
Terraform und Pulumi ist eher Infrastructure Orchestration

Kannst du mit Ansible Server hochfahren? Ja
Kannst du mit Terraform auch Config Management machen? Ja

Die Frage: Willst du das auch?
Usecase - Terraform: Server hochfahren
Usecase - Config Management: Pakete installieren, Uhrzeit einstellen, etc.

Perfekte Kombination? Beides!
Erst Terraform, dann Ansible play

Konzepte:
IasCode: Immutable, hält state
Config Mgmt: Oft Mutable, hält keinen state

Wird Configuration Management wie Ansible heutzutage noch gebraucht?
- Ja, es gibt immer noch Dinge die in die Verantwortung des Config Managements fallen
- zB Passöwrter für deine Docker Container registry hinterlegen

Ggf. auch ganz gut:
- Starten mit Terraform
- Schauen wie sich das entwickelt, wie viel geändert wird
- Ansible später dazu schalten

Was ist ein Terraform Provider
- Source-Code in Go
- QUatscht mit den vendor APIs
- Terraform files werden in HCL
- "Übersetzung" von HCL auf API-Calls

Was ist denn das beste Configuration Management-Tool: Ansible, Salt, Chef, Puppet, CFEngine oder Bahs-Scripte?
- Ist das ein Pull oder Push-Prinzip?
- Pull:
    Definiere den Code
    Pushe es in ein Repo
    Server fragen das Repo ab
    Executen es irgendwann
    Vorteil: Skaliert wie hulle
    Nachteil: Du weißt nicht wann + Tooling ob das erfolgreich

- Push
    Du sagst, wann es ausgeführt wird
    Nachteil: Dauert bei großer Infrastruktur

- Salt ist sehr schön, gute Features
- Side projects: Ansible
    SSH Connections
    Agentless


Was sind Herausforderung von Infrastructure as Code?
- Änderungen können zur Löschung von Resourcen (zb Servern) führen
    Wenn es keine In-Place Updates sind
    Jede Änderung kann zum Neuanlegen der Resource führen
    Spezielle Resourcen unterstützne In-Place upgrades

    Aber Cloud heißt ja nicht, dass der Server immer da ist ... Hardware failure

- Großes Repo mit vielen Terraform Resourcen
    - Anlage von Projekten, von User, von Gruppen
    - Template pro Projekt
    - Ein Terraform-Run dauert 20-40 Min.
    - Provider gehen gegen die Google API
    - Ablauf läuft in die Google API Rate Limit und 1000ende Resourcen durchgegangen


Dominik Siebel wird bei Minute 50 erwähnt

----------------------------------------------------------------------

Old man yells at cloud

Warum Wolfgang die holländischen Holzschlappen trägt

Infrastructure as Code kann auch ein Bash Script sein