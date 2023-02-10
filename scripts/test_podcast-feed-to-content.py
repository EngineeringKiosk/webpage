from podcast_feed_to_content import *


def test_modify_openpodcast_up_down_voting():
    html_content = """<h3>
    <a href="https://api.openpodcast.dev/feedback/57/upvote" rel="nofollow"><strong>👍</strong> (sehr cool)</a>
    <a href="https://api.openpodcast.dev/feedback/18/downvote" rel="nofollow"><strong>&nbsp;</strong></a>
    <a href="https://api.openpodcast.dev/feedback/57/downvote" rel="nofollow"><strong>👎</strong> (geht so)</a>
    </h3>"""

    # Modify the HTML content
    new_html_content = modify_openpodcast_up_down_voting(html_content)

    #assert that the new HTML content starts with <p class="openpodcast-voting">
    assert new_html_content.startswith('<p class="openpodcast-voting">')

    #assert that there is no h3 tag in the new HTML content
    assert "<h3>" not in new_html_content

    #asser that there is a <strong>👍</strong> in the new HTML content and wasn't replaced by mistake
    assert "<strong>👍</strong>" in new_html_content

def test_modify_openpodcast_up_down_voting_check_greedy():
    html_content = """<h3>
    <a href="https://api.openpodcast.dev/feedback/57/upvote" rel="nofollow"><strong>👍</strong> (sehr cool)</a>
    </h3>
    other stuff here
    <h3>
    <a href="https://api.openpodcast.dev/feedback/57/downvote" rel="nofollow"><strong>👍</strong> (sehr cool)</a>
    </h3>"""

    # Modify the HTML content
    new_html_content = modify_openpodcast_up_down_voting(html_content)

    #assert that there are still two starting p tags and the algo wasn't greedy
    assert new_html_content.count("<p") == 2

def test_modify_openpodcast_up_down_voting_dont_touch_other_stuff():
    html_content = """<h3>
    <a href="https://otherurl" rel="nofollow"><strong>👍</strong> (sehr cool)</a>
    </h3>
    other stuff here
    <h3>
    <a href="https://api.openpodcast.dev/feedback/57/downvote" rel="nofollow"><strong>👍</strong> (sehr cool)</a>
    </h3>"""

    # Modify the HTML content
    new_html_content = modify_openpodcast_up_down_voting(html_content)

    #assert that there are still two h3 tags
    assert new_html_content.count("h3>") == 2

def test_modify_openpodcast_up_down_voting_full_test():
    """
    Raw content of episode 57
    """
    html_content = """<p>Servant Leadership, die dienende Führung: Heiße Luft oder ein neuer Trend?</p><p>Mit Leadership-Stilen ist es wie mit JavaScript-Frameworks: Jede Woche kommt ein neuer. Servant Leadership, der dienende Führungsstil, kann auch als solch einer bezeichnet werden. Speziell mit dem Einzug der Generation Z in die Arbeitswelt, könnte dieser Leadership-Style in Zukunft eine besondere Bedeutung bekommen. Doch was ist Servant Leadership überhaupt? Wie unterscheidet sich dieser Stil von anderen Stilen wie Laissez-faire? Warum ist Servant Leadership gerade ein Trend und in aller Munde? Für wen beziehungsweise in welchen Situationen ist Servant Leadership unangebracht? Und warum spielt die Frage &#34;Warum?&#34; dabei eine bedeutende Rolle? All das und viel mehr in dieser Episode.</p><p>Bonus: Was JavaScript-Frameworks, Krokodile und Greta Thunberg mit Servant Leadership zu tun haben.</p><p><br></p><h3><strong>Deine &#34;schnelle&#34; Rückmeldung zur Episode?</strong></h3><h3><a href="https://api.openpodcast.dev/feedback/57/upvote" rel="nofollow"><strong>👍</strong> (sehr cool)</a> <a href="https://api.openpodcast.dev/feedback/18/downvote" rel="nofollow"><strong> </strong></a><a href="https://api.openpodcast.dev/feedback/57/downvote" rel="nofollow"><strong>👎</strong> (geht so)</a></h3><p><br></p><p>Feedback (gerne auch als Voice Message)</p><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li><li>WhatsApp +49 15678 136776</li></ul><p><br></p><p>Gerne behandeln wir auch euer Audio Feedback in einer der nächsten Episoden, einfach Audiodatei per <a href="https://engineeringkiosk.dev/kontakt/">Email</a> oder WhatsApp Voice Message an +49 15678 136776</p><h3 id="links">Links</h3><ul><li>Überblick Leadership Styles: <a href="https://asana.com/de/resources/leadership-styles" rel="nofollow">https://asana.com/de/resources/leadership-styles</a></li><li>Buch &#34;Servant Leadership&#34; von Robert K. Greenleaf: <a href="https://www.amazon.de/Servant-Leadership-Journey-Legitimate-Greatness/dp/0809105543" rel="nofollow">https://www.amazon.de/Servant-Leadership-Journey-Legitimate-Greatness/dp/0809105543</a></li><li>Engineering Kiosk #47 Wer Visionen hat, soll zum Arzt!?: <a href="https://engineeringkiosk.dev/podcast/episode/47-wer-visionen-hat-soll-zum-arzt/">https://engineeringkiosk.dev/podcast/episode/47-wer-visionen-hat-soll-zum-arzt/</a></li><li>Simon Sinek: <a href="https://simonsinek.com/" rel="nofollow">https://simonsinek.com/</a></li><li>Simon Sinek: Wie große Führungspersönlichkeiten zum Handeln inspirieren: <a href="https://www.ted.com/talks/simon_sinek_how_great_leaders_inspire_action?language=de" rel="nofollow">https://www.ted.com/talks/simon_sinek_how_great_leaders_inspire_action?language=de</a></li><li>Buch &#34;Turn The Ship Around!: A True Story of Turning Followers Into Leaders&#34;: <a href="https://www.amazon.de/Turn-Ship-Around-Building-Breaking/dp/0241250943/" rel="nofollow">https://www.amazon.de/Turn-Ship-Around-Building-Breaking/dp/0241250943/</a></li><li>Engineering Kiosk #17 Was können wir beim Incident Management von der Feuerwehr lernen?: <a href="https://engineeringkiosk.dev/podcast/episode/17-was-k%C3%B6nnen-wir-beim-incident-management-von-der-feuerwehr-lernen/">https://engineeringkiosk.dev/podcast/episode/17-was-k%C3%B6nnen-wir-beim-incident-management-von-der-feuerwehr-lernen/</a></li><li>Engineering Kiosk #44 Der Weg zum hochperformanten Team: <a href="https://engineeringkiosk.dev/podcast/episode/44-der-weg-zum-hochperformanten-team/">https://engineeringkiosk.dev/podcast/episode/44-der-weg-zum-hochperformanten-team/</a></li></ul><h3 id="sprungmarken">Sprungmarken</h3><p>(00:00:00) Intro</p><p>(00:00:46) Wann hast du das erste Mal in deinem Leben Leadership gezeigt?</p><p>(00:04:19) Servant Leadership und welche Leadership-Styles gibt es denn noch?</p><p>(00:06:47) Was ist der Unterschied zwischen Servant Leadership und Laissez-faire?</p><p>(00:10:31) Definition von Servant Leadership</p><p>(00:11:42) Fokus auf das Team, auf die Gemeinschaft und Weitblick sowie visionäres Denken</p><p>(00:15:00) Überzeugungskraft und was Warum: Warum machen wir das alles hier?</p><p>(00:22:51) Verantwortungsbewusstsein für das was man tut</p><p>(00:27:09) Wachstum der einzelnen Teammitglieder und vom Team</p><p>(00:29:00) Sich selbst nicht vernachlässigen: Bewusstsein, Stärken und Schwächen</p><p>(00:30:51) Wieso ist Servant Leadership gerade aktuell und der neue Trend? (Millennials und Gen-Z)</p><p>(00:38:05) Gesicht von Servant Leadership: Simon Sinek</p><p>(00:39:15) Was sind die Nachteile von Servant Leadership?</p><p>(00:46:39) Für wen ist Servant Leadership nicht geeignet?</p><p>(00:55:54) Outro</p><p><br></p><p><br></p><h3 id="hosts">Hosts</h3><ul><li>Wolfgang Gassler (<a href="https://mastodon.social/@woolf" rel="nofollow">https://mastodon.social/@woolf</a>)</li><li>Andy Grunwald (<a href="https://twitter.com/andygrunwald" rel="nofollow">https://twitter.com/andygrunwald</a>)</li></ul><h3 id="feedback-gerne-auch-als-voice-message">Feedback (gerne auch als Voice Message)</h3><ul><li>Email: <a href="mailto:stehtisch@engineeringkiosk.dev" rel="nofollow">stehtisch@engineeringkiosk.dev</a></li><li>Mastodon: <a href="https://podcasts.social/@engkiosk" rel="nofollow">https://podcasts.social/@engkiosk</a></li><li>Twitter: <a href="https://twitter.com/EngKiosk" rel="nofollow">https://twitter.com/EngKiosk</a></li></ul><p>WhatsApp +49 15678 136776</p>"""
    expected     = """What is the expected result?"""

    # Modify the HTML content
    actual = modify_openpodcast_up_down_voting(html_content)

    assert actual == expected