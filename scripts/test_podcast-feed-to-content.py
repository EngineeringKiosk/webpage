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