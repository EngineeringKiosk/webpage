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
import pathlib

from episode_finder import EpisodeFinder

from functions import (
    build_correct_file_path,
    configure_global_logger,
    get_podcast_episode_transcript_slim_path_by_episode_number
)

from constants import (
    EPISODES_STORAGE_DIR,
    EPISODES_IMAGES_STORAGE_DIR,
    PODCAST_RSS_FEED,
    TOML_FILE,
    REDIRECT_PREFIX,
    DEFAULT_SPEAKER,
    PODCAST_APPLE_URL,
    SPOTIFY_SHOW_ID,
    DEEZER_PODCAST_ID
)

# External libraries
import frontmatter
import spotipy
from spotipy.oauth2 import SpotifyClientCredentials
from PIL import Image
import deezer


# Amazon Music is missing here.
# We are adding this link manually, because there is no API available
# and the HTML page of Amazon Music is full of async JavaScript.
# Manual operations might be easier here.

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
    found_chapters = re.findall(r"<p><span>\(([0-9:]*)\) ([^<]*)</span></p>", description)
    if len(found_chapters) == 0:
        found_chapters = re.findall(r"<p>\(([0-9:]*)\) ([^<]*)</p>", description)

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


def sync_podcast_episodes(rss_feed, path_md_files, path_img_files, no_api_calls=False, spotify_client=None):
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

    # Check if we should run the API calls to external podcast platforms
    apple_podcast_content = {}
    spotify_episodes = []
    deezer_episodes = []
    if no_api_calls:
        logging.info("Requesting content from Podcast sites is disabled via `--no-api-calls` flag!")

    else:
        logging.info("Requesting content from Podcast sites ...")
        apple_podcast_content = get_json_content_from_url(PODCAST_APPLE_URL)
        spotify_episodes = spotify_client.show_episodes(SPOTIFY_SHOW_ID, limit=50, offset=0, market="DE")

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

    episode_finder = EpisodeFinder()

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
        date_parsed = datetime.datetime.strptime(pub_date, '%a, %d %b %Y %H:%M:%S %z')

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

            # Get only the filename
            image_filename = f"./{os.path.basename(image_filename)}"

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

        episode_number = episode_finder.get_episode_number_from_filename(filename, leading_zero=True)
        data = {
            'advertiser': '',
            'amazon_music': '',
            'apple_podcasts': get_episode_link_from_apple(apple_podcast_content, title),
            'audio': mp3_link,
            'chapter': chapter,
            'deezer': get_episode_link_from_deezer(deezer_episodes, title),
            'description': description_short,
            'headlines': headline_info,
            'image': image_filename,
            'length_second': length_second,
            'pubDate': date_parsed,
            'speaker': DEFAULT_SPEAKER,
            'spotify': spotify_link,
            'tags': [],
            'title': title,
            'transcript_slim': get_podcast_episode_transcript_slim_path_by_episode_number(episode_number),
        }

        full_file_path = f'{path_md_files}/{filename}'

        # If the file exists, we want to read the frontmatter, extract
        # the spotify, apple podcasts, etc. links and write this to the
        # new data.
        # Why? Because some of the Podcast player links are added manual.
        # The rest of the data (title, description, etc.) are parsed
        # out of the RSS feed and maintained in a different application.
        # This way, we "update" our local data based on the management application,
        # but keep the manual parts.
        if exists(full_file_path):
            with open(full_file_path) as f:
                episode = frontmatter.load(f)

                keys_to_keep = [
                    'advertiser',
                    'amazon_music',
                    'apple_podcasts',
                    'deezer',
                    'length_second',
                    'speaker',
                    'spotify',
                    'tags',
                    'youtube',
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
    redirect_episode_number_regex = re.compile(rf"^{redirect_prefix}([-\d]*)$")
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

    episode_number_regex = re.compile(r'([-\d]*)-')
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
            "to": f"/podcast/episode/{episode_file}?pkn=shortlink",
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

    if "results" not in content:
        return u

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

    if episodes is None or "items" not in episodes:
        return e

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


if __name__ == "__main__":
    # Argument and parameter parsing
    cli_parser = argparse.ArgumentParser(description='Automate new Podcast Episide parsing')
    cli_parser.add_argument('Mode',
        metavar='mode',
        type=str,
        default='sync',
        const='sync',
        nargs='?',
        choices=['sync', 'redirect'],
        help='Mode to execute. Supported: sync, redirect (default: %(default)s)')
    cli_parser.add_argument("-n", "--no-api-calls", action="store_true", help='Avoids network calls to Platforms like Spotify, Apple, ... (default: %(default)s)')

    args = cli_parser.parse_args()

    # Setup logger
    configure_global_logger()

    match args.Mode:
        case "sync":
            # Bootstrapping API clients
            spotify_client = None
            if not args.no_api_calls:
                # Spotify
                SPOTIFY_APP_CLIENT_ID = os.getenv('SPOTIFY_APP_CLIENT_ID')
                SPOTIFY_APP_CLIENT_SECRET = os.getenv('SPOTIFY_APP_CLIENT_SECRET')
                if not SPOTIFY_APP_CLIENT_ID or not SPOTIFY_APP_CLIENT_SECRET:
                    logging.error("Env vars SPOTIFY_APP_CLIENT_ID or SPOTIFY_APP_CLIENT_SECRET are not set properly.")
                    logging.error("Please double check and restart.")
                    sys.exit(1)

                spotify_client = create_spotify_client(SPOTIFY_APP_CLIENT_ID, SPOTIFY_APP_CLIENT_SECRET)

            sync_podcast_episodes(
                PODCAST_RSS_FEED,
                build_correct_file_path(EPISODES_STORAGE_DIR),
                build_correct_file_path(EPISODES_IMAGES_STORAGE_DIR),
                no_api_calls=args.no_api_calls,
                spotify_client=spotify_client
            )
        case "redirect":
            create_redirects(TOML_FILE, EPISODES_STORAGE_DIR, REDIRECT_PREFIX)
