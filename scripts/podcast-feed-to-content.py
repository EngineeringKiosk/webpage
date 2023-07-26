from xml.sax.handler import feature_validation
import argparse
import requests
import xml.etree.ElementTree as ET
import unicodedata
import re
import mimetypes
import datetime
import yaml
import html
import toml
import os
from os.path import exists, isfile, join
import logging
import sys
from urllib.parse import urlparse
import pathlib

# External libraries
from bs4 import BeautifulSoup
from slugify import slugify
import frontmatter
import spotipy
from spotipy.oauth2 import SpotifyClientCredentials
from PIL import Image
import deezer

# Global variables
PODCAST_RSS_FEED = "https://feeds.redcircle.com/0ecfdfd7-fda1-4c3d-9515-476727f9df5e"
PATH_MARKDOWN_FILES = 'src/pages/podcast/episode'
PATH_IMAGE_FILES = 'public/images/podcast/episode'
TOML_FILE = 'netlify.toml'
REDIRECT_PREFIX = '/episodes/'

# URLs from Podcast sites
PODCAST_APPLE_URL = "https://itunes.apple.com/lookup?id=1603082924&media=podcast&entity=podcastEpisode&limit=100"
SPOTIFY_SHOW_ID = "0tJRC0UsObPCWLmmzmOkIs"
PODCAST_GOOGLE_URL = "https://podcasts.google.com/feed/aHR0cHM6Ly9mZWVkcy5yZWRjaXJjbGUuY29tLzBlY2ZkZmQ3LWZkYTEtNGMzZC05NTE1LTQ3NjcyN2Y5ZGY1ZQ"
DEEZER_PODCAST_ID = 3330122
# TODO Add episode single view link retrieval for Amazon Music

# From the Django project
# See https://docs.djangoproject.com/en/2.1/_modules/django/utils/text/#slugify
def slugify(value, allow_unicode=False):
    """
    Convert to ASCII if 'allow_unicode' is False. Convert spaces to hyphens.
    Remove characters that aren't alphanumerics, underscores, or hyphens.
    Convert to lowercase. Also strip leading and trailing whitespace.
    """
    value = str(value)
    if allow_unicode:
        value = unicodedata.normalize('NFKC', value)
    else:
        value = unicodedata.normalize('NFKD', value).encode('ascii', 'ignore').decode('ascii')
    value = re.sub(r'[^\w\s-]', '', value).strip().lower()
    return re.sub(r'[-\s]+', '-', value)


def download_file(url, local_filename):
    """
    Downloads {url} and stores it into {local_filename}.
    Original file extension will be guessed based on the mime type.

    Returns the full local file name incl. file extension.

    Exceptions are bubbled up, once raised.
    """
    with requests.get(url, stream=True) as r:
        r.raise_for_status()

        # Determine file extension
        content_type = r.headers['content-type']
        extension = mimetypes.guess_extension(content_type)
        file_name = f'{local_filename}{extension}'

        with open(file_name, 'wb') as f:
            for chunk in r.iter_content(chunk_size=8192):
                f.write(chunk)

    return file_name


# Not an awesome logic, but enough for our usecase.
# See https://stackoverflow.com/questions/9662346/python-code-to-remove-html-tags-from-a-string for more.
def remove_html_tags(raw_html):
    """
    Removing all HTML tags.
    """
    # First we replace "</p>" with "</p> " (adding a single whitespace).
    # Primarily to ensure that after the removal of each HTML element, 
    # we have a whitespace at the end of each sentence/paragraph.
    # Otherwise, our text would look like
    #
    #   [...] im Engineering Kiosk.Indirekt knüpfen [...]
    #
    # Take attention to the missing whitespace between the sentences
    cleantext = raw_html.replace("</p>", "</p> ")

    html_tags_regex = re.compile('<.*?>')
    cleantext = re.sub(html_tags_regex, '', cleantext)
    return cleantext


def parse_headlines_from_html(raw_html):
    """
    We get HTML from the Podcast platforms Rich Text Editor.
    Often, this is not the best generated HTML, but yeah. Lets keep it for now.

    TODO Link sprungmarken
    """
    html = raw_html

    # Get all headlines
    headline_slugs = {}
    found_headlines = re.findall("<h3>(<span>)?(.*?)(</span>)?</h3>", html)
    for h in found_headlines:
        line = h[1]

        # Sometimes we match <br>
        # Here, we ensure to skip it
        if "<" in line and ">" in line:
            continue

        slug = slugify(line)
        #html = html.replace(f"<h3>{h}", f'<h3 class="mb-4 text-2xl md:text-3xl font-semibold text-coolGray-800" id="{slug}">{h}')
        html = html.replace(f"<h3>{h[0]}{h[1]}{h[2]}", f'<h3 id="{slug}">{line}')
        headline_slugs[slug] = line

    return html, headline_slugs


def get_chapter_from_description(description):
    # Chapter entries look like
    #       <p><span>(00:00:00) Intro</span></p>
    found_chapters = re.findall("<p><span>\(([0-9:]*)\) ([^<]*)</span></p>", description)
    if len(found_chapters) == 0:
        found_chapters = re.findall("<p>\(([0-9:]*)\) ([^<]*)</p>", description)

    chapter = []
    for c in found_chapters:
        start_time = c[0]
        if start_time.count(":") == 1:
            start_time = f"00:{start_time}"

        # In line with chaptermarks from podcast player
        # See https://github.com/podigee/podigee-podcast-player/blob/master/docs/configuration.md
        entry = {
            "start": start_time,
            "title": html.unescape(c[1]),
        }
        chapter.append(entry)

    return chapter


def remove_rel_nofollow_from_internal_links(html_content):
    """
    Removes all `rel="nofollow"` entries from https://engineeringkiosk.dev/*
    links in {html_content}.
    """
    pattern = r'<a href="(https:\/\/engineeringkiosk.dev\/[^\s]*)" rel="nofollow">'
    replacement = r'<a href="\1">'
    new_html_content = re.sub(pattern, replacement, html_content)

    pattern_jump = r'<a href="(https:\/\/jump\.engineeringkiosk.dev\/[^\s]*)" rel="nofollow">'
    replacement = r'<a href="\1">'
    new_html_content = re.sub(pattern_jump, replacement, new_html_content)

    return new_html_content


def modify_openpodcast_up_down_voting(html_content):
    """
    Modify the HTML output added for the Thumbs Up/Thumbs Down voting by OpenPodcast.

    Yep, this is super dirty, and hacky and does not scale and you should not do this at home.
    But i am not at home. And it is Friday evening. And this works.
    We just are not allowed to change anything on the HTML :D
    But this is fine (dog sitting on a chair inside a house of fire).
    """
    new_html_content = html_content.replace("<h3><strong>Deine ", "<p><strong>Deine ")
    new_html_content = new_html_content.replace("</strong></h3><h3><a href=\"https://api.openpodcast.dev", "</strong></p><p><a href=\"https://api.openpodcast.dev")
    new_html_content = new_html_content.replace("</strong> (geht so)</a></h3>", "</strong> (geht so)</a></p>")

    return new_html_content


def sync_podcast_episodes(rss_feed, path_md_files, path_img_files, spotify_client):
    """
    Syncs the Podcast Episodes from the RSS feed down to disk
    and prepares the content to match the structure of the used
    static site generator.

    This means in detail:
    - Read the RSS feed
    - Extracts the information into a Frontmatter structure
    - Downloads the image
    - Writes a Markdown and image file to disk

    """

    logging.info(f"Reading Podcast RSS Feed {rss_feed} ...")
    # Get the XML Feed content
    try:
        feed_response = requests.get(rss_feed)
        feed_response.raise_for_status()

        # Catching this exception is enough, because this is the
        # parent exceptions for Connection, Timeouts, Redirect and HTTP errors.
        # See https://docs.python-requests.org/en/latest/user/quickstart/#errors-and-exceptions
    except requests.exceptions.RequestException as e:
        raise SystemExit(e)

    logging.info(f"Reading Podcast RSS Feed {rss_feed} ... Successful")

    # requests makes an educated guess on the response encoding.
    # Here we overwrite the encoding to UTF-8
    feed_response.encoding = 'utf-8'

    logging.info("Requesting content from Podcast sites ...")
    apple_podcast_content = get_json_content_from_url(PODCAST_APPLE_URL)
    spotify_episodes = spotify_client.show_episodes(SPOTIFY_SHOW_ID, limit=50, offset=0, market="DE")
    google_podcast_content = get_raw_content_from_url(PODCAST_GOOGLE_URL)

    # Pagination and auth not respected.
    # Right now it works, because a) we don't make that much requests and
    # b) don't have that much episodes.
    # If we have more and more episodes, this might look different and need adjustment.
    #
    # Query quota (2022-07-17)
    # The number of requests per second is limited to 50 requests / 5 seconds.
    deezer_client = deezer.Client()
    deezer_podcast = deezer_client.get_podcast(DEEZER_PODCAST_ID)
    deezer_episodes = deezer_podcast.get_episodes()
    logging.info("Requesting content from Podcast sites ... Successful")

    logging.info("Processing Podcast Episode items ...")

    # Parse the XML and process all items
    parsed_xml = ET.fromstring(feed_response.text)
    channel = parsed_xml.find("channel")
    for item in channel.findall('item'):
        # Other fields that are available
        # - itunes:episodeType
        # - itunes:title
        # - itunes:episode
        # - itunes:author
        # - content:encoded
        # - itunes:duration
        # - link
        title = item.find('title').text
        description = item.find('description').text

        # Get all headlines from description and add
        # jump markers to <h[1-6]> tags
        html_content, headlines = parse_headlines_from_html(description)
        description_html = html_content

        # Previously we had a logic here to pretty print
        # the HTML via BeautifulSoup and oup.prettify().
        # We removed it, because Astro (the used static side generator)
        # had some issues with HTML parsing within markdown files like
        #   - https://github.com/withastro/astro/issues/3529
        #   - https://github.com/withastro/astro/issues/3642
        # Hence we removed the logic.
        #
        # In the end, it was only for us humans to make it a bit more readable
        # in the markdown files. The content is managed in our Podcast platform
        # (RedCircle) and we don't modify the Markdown files manually at all.
        #
        # Thats why we got rid of the prettify logic.
        html_content = description_html.strip()
        html_content = remove_rel_nofollow_from_internal_links(html_content)
        html_content = modify_openpodcast_up_down_voting(html_content)

        chapter = get_chapter_from_description(description)

        # Parse headlines
        headline_info = '||'.join([f'{slug}::{headline}' for slug, headline in headlines.items()])

        description_text_only = remove_html_tags(description)

        # Cut text after the Intro text
        # The string is hardcoded. Bad? Yep, maybe. Works? Yep.
        str_to_split = ""
        if "Feedback an stehtisch" in description_text_only:
            str_to_split = "Feedback an stehtisch@engineeringkiosk.dev"
        elif "Feedback (gerne" in description_text_only:
            str_to_split = "Feedback (gerne"

        description_short = description_text_only
        if str_to_split:
            description_short = description_text_only.split(str_to_split)
            description_short = description_short[0]

        description_short = description_short.strip()
        description_short = html.unescape(description_short)

        # Date format: Tue, 05 Apr 2022 04:25:00 +0000
        pub_date = item.find('pubDate').text
        # See full list of format codes here
        # https://docs.python.org/3/library/datetime.html#strftime-strptime-behavior
        date_parsed = datetime.datetime.strptime(pub_date, '%a, %d %b %Y %H:%M:%S %z').strftime('%Y-%m-%d')

        mp3 = item.find('enclosure')
        mp3_link = mp3.attrib.get('url')

        # If an image is available,
        # download it and store it in the public folder
        image = item.find('{http://www.itunes.com/dtds/podcast-1.0.dtd}image')
        if image is not None:
            image = image.attrib.get('href')
            file_ext = pathlib.Path(image).suffix
            image_filename = slugify(title, True)
            # File extension will be determined in download_file
            image_filename = f"{path_img_files}/{image_filename}"

            # If the cover doesnt exist, download and resize.
            if not os.path.isfile(f"{image_filename}{file_ext}"):
                image_filename = download_file(image, image_filename)
                # Resize image
                cover_image = Image.open(image_filename)
                resized_cover_image = cover_image.resize((700, 700))
                resized_cover_image.save(image_filename)
            else:
                image_filename = f"{image_filename}{file_ext}"

            # Remove prefix "public"
            image_filename = image_filename.split("public")
            image_filename = image_filename[1]

        else:
            image = ""

        # Build the filename and the content
        filename = slugify(title, True)
        filename = f'{filename}.md'

        spotify_episode = get_episode_from_spotify(spotify_episodes, title)
        length_second = 0
        spotify_link = ""
        if spotify_episode is not None:
            length_second = spotify_episode["duration_ms"] / 1000
            length_second = int(length_second)

            spotify_link = spotify_episode["external_urls"]["spotify"]

        data = {
            'layout': '../../../layouts/podcast-episode.astro',
            'title': title,
            'audio': mp3_link,
            'pubDate': date_parsed,
            'image': image_filename,
            'description': description_short,
            'headlines': headline_info,
            'chapter': chapter,
            'spotify': spotify_link,
            'google_podcasts': get_episode_link_from_google(google_podcast_content, title),
            'apple_podcasts': get_episode_link_from_apple(apple_podcast_content, title),
            'amazon_music': '',
            'deezer': get_episode_link_from_deezer(deezer_episodes, title),
            'rtlplus': '',
            'youtube': '',
            'tags': [],
            'length_second': length_second,
        }

        full_file_path = f'{path_md_files}/{filename}'

        # If the file exists, we want to read the frontmatter, extract
        # the spotify, google podcasts, etc. links and write this to the
        # new data.
        # Why? Because the Podcast player links are added manual.
        # The rest of the data (title, description, etc.) are parsed
        # out of the RSS feed and maintained in a different application.
        # This way, we "update" our local data based on the management application,
        # but keep the manual parts.
        if exists(full_file_path):
            with open(full_file_path) as f:
                episode = frontmatter.load(f)

                keys_to_keep = [
                    'spotify',
                    'google_podcasts',
                    'apple_podcasts',
                    'amazon_music',
                    'deezer',
                    'rtlplus',
                    'youtube',
                    'tags',
                    'length_second',
                ]
                for key in keys_to_keep:
                    val = episode.get(key)
                    if val:
                        data[key] = episode.get(key)

        content_yaml = yaml.dump(data)
        content = (
            '---\n'
            f'{content_yaml}\n'
            '---\n'
            f'{html_content}'
        )

        # Write file to disk as a new podcast episode
        f = open(full_file_path, 'w', encoding='utf8')
        f.write(content)
        f.close()

    logging.info("Processing Podcast Episode items ... Successful")


def trim_prefix(line, prefix):
    """
    Removes prefix from line.
    """
    if line.startswith(prefix):
        line = line[len(prefix):]

    return line


def create_redirects(file_to_parse, path_md_files, redirect_prefix):
    """
    Creates redirects for every Podcast episode to /episodes/123 by

    1. Reading the toml file
    2. Processing all podcast episodes
    3. Writing the toml file
    """
    # Read existing toml file
    parsed_toml = ""
    with open(file_to_parse) as f:
        content = f.read()
        parsed_toml = toml.loads(s=content)

    # Restructure existing redirects into a hashmap
    # for easier lookup
    redirect_episode_number_regex = re.compile(f"^{redirect_prefix}([-\d]*)$")
    redirect_map = {}
    for redirect in parsed_toml['redirects']:
        # Find the number of the episode
        redirect_episode_number = re.findall(redirect_episode_number_regex, redirect["from"])
        # When there is no match, the list is empty
        if not redirect_episode_number:
            continue

        redirect_episode_number = redirect_episode_number[0]

        redirect_map[redirect_episode_number] = redirect


    # Get existing podcast episodes
    episodes = [f for f in os.listdir(path_md_files) if isfile(join(path_md_files, f)) and f.endswith('.md')]

    episode_number_regex = re.compile('([-\d]*)-')
    for episode in episodes:
        # Find the number of the episode
        episode_number = re.findall(episode_number_regex, episode)[0]

        # If we have a number loke 00 or 05, remove the first 0
        episode_number = trim_prefix(episode_number, "0")

        # Check if we have a redirect for this episode already
        # If yes, skip ip
        if episode_number in redirect_map:
            logging.info(f"Skipping redirect processing for episode {episode_number}: Redirect exists already.")
            continue

        logging.info(f"Adding redirect for episode {episode_number}")

        episode_file = episode.removesuffix(".md")
        new_redirect_shortlink = {
            "from": f"/episodes/{episode_number}",
            "to": f"/podcast/episode/{episode_file}?pkn=shortlink",
            "status": 301,
            "force": True,
        }
        parsed_toml['redirects'].append(new_redirect_shortlink)

        # We don't add a campaign part here
        # Netlify forwards the query params as well.
        # This way, we can dynamically decide what param to use
        # e.g:
        #   - https://engineeringkiosk.dev/ep5?pkn=shownotes
        #   - https://engineeringkiosk.dev/ep5?pkn=twit_init
        new_redirect_episode_shortlink = {
            "from": f"/ep{episode_number}",
            "to": f"/podcast/episode/{episode_file}",
            "status": 301,
            "force": True,
        }
        parsed_toml['redirects'].append(new_redirect_episode_shortlink)

    # Write new file
    with open(file_to_parse, 'w') as f:
        toml.dump(o=parsed_toml, f=f)


def get_json_content_from_url(u):
    """
    Retrieves the JSON content from address u.
    """
    content = ""
    with requests.get(u, stream=True) as r:
        r.raise_for_status()
        content = r.json()

    return content


def get_raw_content_from_url(u):
    """
    Retrieves the raw content from address u.
    """
    content = ""
    with requests.get(u, stream=True) as r:
        r.raise_for_status()
        content = r.content

    return content


def get_episode_link_from_apple(content, title: str) -> str:
    """
    Parses the Apple Episode Single View link (matching with title) from content.
    content is a JSON representation of the Apple Podcast Engineering Kiosk site.
    title is the full title of a single episode.

    If no title matches, it will return an empty string.
    """
    u = ""
    tracks = content["results"]
    for track in tracks:
        if track["trackName"] == title:
            u = track["trackViewUrl"]

    return u


def create_spotify_client(app_id: str, app_secret: str):
    """
    Creates a spotify api client based on the application
    secrets app_id and app_secret.

    Docs: https://github.com/plamere/spotipy
    Docs: https://developer.spotify.com/documentation/web-api/reference/#/
    """
    spotify_client = spotipy.Spotify(
        auth_manager=SpotifyClientCredentials(
            client_id=app_id,
            client_secret=app_secret
        )
    )
    return spotify_client


def get_episode_from_spotify(episodes, title: str) -> dict:
    """
    Parses the Spotify Episode Single View link (matching with title) from episodes list.
    episodes is a JSON representation from the Spotify API / Engineering Kiosk Show.
    title is the full title of a single episode.

    If no title matches, it will return an empty string.
    """
    e = None
    for episode in episodes["items"]:
        if episode["name"] == title:
            e = episode

    return e


def get_episode_link_from_deezer(episodes, title: str) -> str:
    """
    Parses the Deezer Episode Single View link (matching with title) from episodes list.
    episodes is a JSON representation from the Deezer API / Engineering Kiosk Show.
    title is the full title of a single episode.

    If no title matches, it will return an empty string.
    """
    u = ""
    for episode in episodes:
        if episode.title == title:
            u = episode.link

    return u


def get_episode_link_from_google(content, title: str) -> str:
    """
    Google Podcast does not offer an API, xml feed or anything like this.

    Hence we do typical HTML link scraping.
    There is no error checking at all, because we want this function to
    fail if there is anything changing.

    If no title matches, it will return an empty string.
    """
    scheme = "https"
    hostname = "podcasts.google.com"

    soup = BeautifulSoup(content, features="html.parser")
    items = soup.findAll('div', string = title)

    u = ""
    for item in items:
        link = item.findParent("a").get('href')
        o = urlparse(link)

        # The links we get are relative like
        #   ./feed/...
        # We need absolute URLs.
        o = o._replace(path=trim_prefix(o.path, "./"))
        u = o._replace(scheme=scheme, netloc=hostname).geturl()

        # First link is enough.
        break

    return u


if __name__ == "__main__":
    # Argument and parameter parsing
    cli_parser = argparse.ArgumentParser(description='Automate new Podcast Episide parsing')
    # See https://stackoverflow.com/questions/40324356/python-argparse-choices-with-a-default-choice
    cli_parser.add_argument('Mode',
        metavar='mode',
        type=str,
        default='sync',
        const='sync',
        nargs='?',
        choices=['sync', 'redirect'],
        help='Mode to execute. Supported: sync, redirect (default: %(default)s)')

    args = cli_parser.parse_args()

    # Setup logger
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s [%(levelname)s] %(message)s",
        handlers=[
            logging.StreamHandler()
        ]
    )

    # Determine if the script is called from root
    # or from the scripts directory.
    directory_path = os.getcwd()
    folder_name = os.path.basename(directory_path)
    folder_prefix = ""
    if folder_name == "scripts":
        folder_prefix = "../"

    match args.Mode:
        case "sync":
            # Bootstrapping Spotify API client
            SPOTIFY_APP_CLIENT_ID = os.getenv('SPOTIFY_APP_CLIENT_ID')
            SPOTIFY_APP_CLIENT_SECRET = os.getenv('SPOTIFY_APP_CLIENT_SECRET')
            if not SPOTIFY_APP_CLIENT_ID or not SPOTIFY_APP_CLIENT_SECRET:
                logging.error("Env vars SPOTIFY_APP_CLIENT_ID or SPOTIFY_APP_CLIENT_SECRET are not set properly.")
                logging.error("Please double check and restart.")
                sys.exit(1)

            spotify_client = create_spotify_client(SPOTIFY_APP_CLIENT_ID, SPOTIFY_APP_CLIENT_SECRET)

            sync_podcast_episodes(PODCAST_RSS_FEED, f"{folder_prefix}{PATH_MARKDOWN_FILES}", f"{folder_prefix}{PATH_IMAGE_FILES}", spotify_client)
        case "redirect":
            # TODO Once Python 3.11 is out, replace toml library with stdlib
            # See https://peps.python.org/pep-0680/
            create_redirects(TOML_FILE, PATH_MARKDOWN_FILES, REDIRECT_PREFIX)