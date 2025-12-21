---
advertiser: ""
amazon_music: https://music.amazon.com/podcasts/c35a09fe-4116-4e04-8f68-77d61b112e46/episodes/503f44c9-7a01-4b14-b7ee-acd70d03d040/engineering-kiosk-48-der-layer-unter-docker-containerd-kubernetes-container-runtime-interface-cri-o-und-open-container-initiative-oci
apple_podcasts: https://podcasts.apple.com/us/podcast/48-der-layer-unter-docker-containerd-kubernetes-container/id1603082924?i=1000589004404&uo=4
audio: https://audio1.redcircle.com/episodes/ad2d93d9-918f-4a76-9026-0ad264afc94d/stream.mp3
chapter:
- start: 00:00:00
  title: Intro
- start: 00:00:58
  title: Computerbild, c't und iX und einen Layer tiefer als "Was ist Docker"
- start: 00:04:31
  title: "Der nächsttiefere Layer nach docker run: Highlevel Runtime containerd"
- start: 00:07:05
  title: Was ist denn ein Container-Lebenszyklus und der Unterschied zwischen Highlevel und Lowlevel-Runtime?
- start: 00:09:37
  title: Ist containerd noch ein Teil von docker?
- start: 00:10:35
  title: Kann containerd auch von anderen Clients als docker angesprochen werden?
- start: 00:13:06
  title: Verwendet Kubernetes auch docker? docker shim, CoreOS, rkt
- start: 00:16:01
  title: Container Runtime Interface (CRI) von Kubernetes
- start: 00:19:44
  title: Kubernetes, API-Server, kubelet und die Highlevel-Runtime
- start: 00:20:58
  title: "Eine alternative Highlevel-Container-Runtme zu containerd: CRI-O"
- start: 00:24:15
  title: Die Cloud Native Computing Foundation (CNCF)
- start: 00:28:16
  title: Was ist besser? Containerd oder CRI-O?
- start: 00:29:57
  title: docker shim und unterstützt Kubernetes noch docker?
- start: 00:32:10
  title: Docker container sind eigentlich Open Container Initiative (OCI) Images und Container
- start: 00:33:21
  title: Woher kommt die Open Container Initiative (OCI) und die Standards?
- start: 00:34:41
  title: Zusammenfassung und die Möglichkeiten eure eigene Runtime zu schreiben
- start: 00:36:31
  title: Ausblick auf die nächste Container-Folge und Feedback
deezer: https://www.deezer.com/episode/459224287
description: "Wie sieht eigentlich der Layer unter Docker aus? Und wie interagiert Kubernetes mit Containern? In Episode 46 haben wir geklärt, welches Problem Docker eigentlich löst. Das Container-Ecosystem ist jedoch weit größer. Deswegen widmet sich diese Folge der darunter liegenden Ebene. Wir besprechen die Modularisierung von Docker, die herausgetrennte Highlevel Runtime containerd, wie Kubernetes mit Docker-Containern umgeht, ob Docker Container die einzige Art von Containern ist, die Kubernetes unterstützt, was ein Container Runtime Interface (CRI) ist, was die Open Container Initiative (OCI) und ob auch du dir deine eigene Highlevel Container Runtime programmieren kannst. Bonus: Was die Linux-Mafia ist und wieso es bald eine österreichische Container Runtime gibt."
headlines: links::Links||sprungmarken::Sprungmarken||hosts::Hosts||feedback-gerne-auch-als-voice-message::Feedback (gerne auch als Voice Message)
image: ./48-der-layer-unter-docker-containerd-kubernetes-container-runtime-interface-cri-o-und-open-container-initiative-oci.jpg
length_second: 2264
pubDate: 2022-12-06 09:21:47+00:00
speaker:
- name: Andy Grunwald
  transcriptLetter: A
- name: Wolfi Gassler
  transcriptLetter: B
spotify: https://open.spotify.com/episode/3K3nMEPuTCd6sxi4JQXIVO
tags:
- Backend
- DevOps
- Cloud
title: "#48 Der Layer unter Docker: containerd, Kubernetes, Container Runtime Interface, CRI-O und Open Container Initiative (OCI)"
transcript_slim: src/data/transcripts/48-transcript-slim.json
youtube: https://www.youtube.com/watch?v=U0_i4N_xyCE
---
<p><span>Wie sieht eigentlich der Layer unter Docker aus? Und wie interagiert Kubernetes mit Containern?</span></p><p><span>In Episode 46 haben wir geklärt, welches Problem Docker eigentlich löst. Das Container-Ecosystem ist jedoch weit größer. Deswegen widmet sich diese Folge der darunter liegenden Ebene. Wir besprechen die Modularisierung von Docker, die herausgetrennte Highlevel Runtime containerd, wie Kubernetes mit Docker-Containern umgeht, ob Docker Container die einzige Art von Containern ist, die Kubernetes unterstützt, was ein Container Runtime Interface (CRI) ist, was die Open Container Initiative (OCI) und ob auch du dir deine eigene Highlevel Container Runtime programmieren kannst.</span></p><p><span>Bonus: Was die Linux-Mafia ist und wieso es bald eine österreichische Container Runtime gibt.</span></p><p><br></p><p><span>Feedback (gerne auch als Voice Message)</span></p><ul><li><span>Email: </span><a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li><span>Twitter: </span><a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li><span>WhatsApp </span>+49 15678 136776</li></ul><p><br></p><p><span>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per </span><a href="https://engineeringkiosk.dev/kontakt/">Email</a><span> oder WhatsApp Voice Message an </span>+49 15678 136776</p><p><br></p><p>Unsere aktuellen Werbepartner findest du auf <a href="https://engineeringkiosk.dev/partners">https://engineeringkiosk.dev/partners</a></p><p> </p><h3 id="links">Links</h3><ul><li><span>Engineering Kiosk #46 Welches Problem löst Docker?: </span><a href="https://engineeringkiosk.dev/podcast/episode/46-welches-problem-l%C3%B6st-docker/">https://engineeringkiosk.dev/podcast/episode/46-welches-problem-l%C3%B6st-docker/</a></li><li><span>c&#39;t - Magazin für Computer Technik: </span><a href="https://www.heise.de/ct/" rel="nofollow">https://www.heise.de/ct/</a></li><li><span>iX - Magazin für professionelle Informationstechnik: </span><a href="https://www.heise.de/ix/" rel="nofollow">https://www.heise.de/ix/</a></li><li><span>Highlevel container runtime containerd: </span><a href="https://containerd.io/" rel="nofollow">https://containerd.io/</a></li><li><span>Getting started with containerd: </span><a href="https://github.com/containerd/containerd/blob/main/docs/getting-started.md" rel="nofollow">https://github.com/containerd/containerd/blob/main/docs/getting-started.md</a></li><li><span>contaiNERD CTL - Docker-compatible CLI for containerd, with support for Compose, Rootless, eStargz, OCIcrypt, IPFS: </span><a href="https://github.com/containerd/nerdctl" rel="nofollow">https://github.com/containerd/nerdctl</a></li><li><span>shim: </span><a href="https://de.wikipedia.org/wiki/Shim_(Informatik)" rel="nofollow">https://de.wikipedia.org/wiki/Shim_(Informatik)</a></li><li><span>Core OS: </span><a href="https://de.wikipedia.org/wiki/Core_OS" rel="nofollow">https://de.wikipedia.org/wiki/Core_OS</a></li><li><span>rkt: </span><a href="https://github.com/rkt/rkt" rel="nofollow">https://github.com/rkt/rkt</a></li><li><span>Container Runtime Interface (CRI): </span><a href="https://kubernetes.io/docs/concepts/architecture/cri/" rel="nofollow">https://kubernetes.io/docs/concepts/architecture/cri/</a></li><li><span>CRI Plugin von containerd: </span><a href="https://github.com/containerd/containerd/tree/main/pkg/cri" rel="nofollow">https://github.com/containerd/containerd/tree/main/pkg/cri</a></li><li><span>Highlevel Container Runtime CRI-O: </span><a href="https://cri-o.io/" rel="nofollow">https://cri-o.io/</a></li><li><span>Cloud Native Computing Foundation: </span><a href="https://www.cncf.io/" rel="nofollow">https://www.cncf.io/</a></li><li><span>Linux Foundation: </span><a href="https://www.linuxfoundation.org/" rel="nofollow">https://www.linuxfoundation.org/</a></li><li><span>Apache Foundation: </span><a href="https://www.apache.org/" rel="nofollow">https://www.apache.org/</a></li><li><span>gVisor - Container Security Platform: </span><a href="https://gvisor.dev/" rel="nofollow">https://gvisor.dev/</a></li><li><span>minikube - local Kubernetes cluster: </span><a href="https://minikube.sigs.k8s.io/docs/" rel="nofollow">https://minikube.sigs.k8s.io/docs/</a></li><li><span>Open Container Initiative: </span><a href="https://opencontainers.org/" rel="nofollow">https://opencontainers.org/</a></li></ul><h3><br></h3><h3 id="sprungmarken">Sprungmarken</h3><p><span>(00:00:00) Intro</span></p><p><span>(00:00:58) Computerbild, c&#39;t und iX und einen Layer tiefer als &#34;Was ist Docker&#34;</span></p><p><span>(00:04:31) Der nächsttiefere Layer nach docker run: Highlevel Runtime containerd</span></p><p><span>(00:07:05) Was ist denn ein Container-Lebenszyklus und der Unterschied zwischen Highlevel und Lowlevel-Runtime?</span></p><p><span>(00:09:37) Ist containerd noch ein Teil von docker?</span></p><p><span>(00:10:35) Kann containerd auch von anderen Clients als docker angesprochen werden?</span></p><p><span>(00:13:06) Verwendet Kubernetes auch docker? docker shim, CoreOS, rkt</span></p><p><span>(00:16:01) Container Runtime Interface (CRI) von Kubernetes</span></p><p><span>(00:19:44) Kubernetes, API-Server, kubelet und die Highlevel-Runtime</span></p><p><span>(00:20:58) Eine alternative Highlevel-Container-Runtme zu containerd: CRI-O</span></p><p><span>(00:24:15) Die Cloud Native Computing Foundation (CNCF)</span></p><p><span>(00:28:16) Was ist besser? Containerd oder CRI-O?</span></p><p><span>(00:29:57) docker shim und unterstützt Kubernetes noch docker?</span></p><p><span>(00:32:10) Docker container sind eigentlich Open Container Initiative (OCI) Images und Container</span></p><p><span>(00:33:21) Woher kommt die Open Container Initiative (OCI) und die Standards?</span></p><p><span>(00:34:41) Zusammenfassung und die Möglichkeiten eure eigene Runtime zu schreiben</span></p><p><span>(00:36:31) Ausblick auf die nächste Container-Folge und Feedback</span></p><p><br></p><h3 id="hosts">Hosts</h3><ul><li><span>Wolfgang Gassler (</span><a href="https://twitter.com/schafele" rel="nofollow">https://twitter.com/schafele</a><span>)</span></li><li><span>Andy Grunwald (</span><a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a><span>)</span></li></ul><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li><span>Email: </span><a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li><span>Twitter: </span><a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li><span>WhatsApp </span>+49 15678 136776</li></ul>