import os
from os.path import isfile, join
import json
import logging
import re

from functions import (
    build_correct_file_path,
    get_podcast_episode_number_from_filename_number,
    has_podcast_episode_a_transcript,
    get_podcast_episode_transcript_by_number
)

# External libraries
import frontmatter

# Global variables
PATH_EPISODE_MARKDOWN_FILES = 'src/content/podcast'
PATH_BLOG_MARKDOWN_FILES = 'src/content/blog'
PATH_PODCAST_INFO_JSON_FILE = 'src/data/podcast-info.json'
PATH_TAGS_JSON_FILE = 'src/data/tags.json'

intervals = (
    ('weeks', 604800),  # 60 * 60 * 24 * 7
    ('days', 86400),    # 60 * 60 * 24
    ('hours', 3600),    # 60 * 60
    ('minutes', 60),
    ('seconds', 1),
)

def build_episode_statistics(path_md_files) -> dict:
    """
    Going through all Podcast Episode Markdown files and collect
    and calculate some basic statistics.
    """
    stats = dict()
    stats['number_of_episodes'] = 0
    stats['total_length_seconds'] = 0

    # Get existing podcast episodes
    episodes = [f for f in os.listdir(path_md_files) if isfile(join(path_md_files, f)) and f.endswith('.md')]
    for episode in episodes:
        file_path = f"{path_md_files}/{episode}"
        with open(file_path) as f:
            episode_frontmatter = frontmatter.load(f)
            
            stats['number_of_episodes'] += 1
            stats['total_length_seconds'] += episode_frontmatter['length_second']
            
    return stats


def build_episode_speaking_time_statistics(path_md_files) -> dict:
    """
    Going through all Podcast Episode transcript files
    and calculate speaking time per speaker.
    """

    # Get existing podcast episodes
    episode_speaking_time = []
    overall_speaking_time = {
        'total_length_s': 0,
        'speaker': {},
    }
    overall_speaking_time['speaker']['Unbekannt'] = 0

    episodes = [f for f in os.listdir(path_md_files) if isfile(join(path_md_files, f)) and f.endswith('.md')]
    for episode in episodes:
        episode_number = get_podcast_episode_number_from_filename_number(episode)

        if has_podcast_episode_a_transcript(episode_number) is False:
            continue

        # Read podcast episode data
        file_path = f"{path_md_files}/{episode}"
        with open(file_path) as f:
            episode_frontmatter = frontmatter.load(f)

        # Read transcript data
        transcript_data = get_podcast_episode_transcript_by_number(episode_number)

        speaker_map = {}
        for s in episode_frontmatter['speaker']:
            speaker_map[s['transcriptLetter']] = s['name']

        overall_speaking_time['total_length_s'] += episode_frontmatter['length_second']

        ms_sum = 0
        speaking_time = {}
        for u in transcript_data['utterances']:
            length = u['end'] - u['start']
            ms_sum += length
            speaker = speaker_map[u['speaker']]

            if speaker not in speaking_time:
                speaking_time[speaker] = 0

            if speaker not in overall_speaking_time['speaker']:
                overall_speaking_time['speaker'][speaker] = 0

            speaking_time[speaker] += length
            overall_speaking_time['speaker'][speaker] += length

        episode_length_ms = episode_frontmatter['length_second'] * 1000
        if ms_sum < episode_length_ms:
            leftover_ms = episode_length_ms - ms_sum
            speaking_time['Unbekannt'] = leftover_ms
            overall_speaking_time['speaker']['Unbekannt'] += leftover_ms

        # Sort overall speaking time
        overall_speaking_time['speaker'] = dict(sorted(overall_speaking_time['speaker'].items(), key=lambda item: item[1], reverse=True))

        # Sort speaking time
        sorted_speaking_time = dict(sorted(speaking_time.items(), key=lambda item: item[1], reverse=True))
        episode_speaking_time.append({
            'speaking': sorted_speaking_time,
            'title': episode_frontmatter['title'],
            'pubDate': episode_frontmatter['pubDate'],
            'length_second': episode_frontmatter['length_second'],
        })

    # Sort episodes by publishing date
    sorted_episode_speaking_time = sorted(episode_speaking_time, key=lambda d: d['pubDate'], reverse=True) 

    return sorted_episode_speaking_time, overall_speaking_time

def build_tags_statistics(file_path) -> dict:
    """
    Collect and calculate some basic statistics about the
    tags used in the podcast.
    """

    stats = dict()
    stats['top_5_tags'] = 0
    stats['total_num_tags'] = 0

    with open(file_path) as f:
        data = json.load(f)

        # Sort tags
        # >= Python 3.7 dicts are sorted
        tags = dict(sorted(data.items(), key=lambda x: x[1]['usage_count'], reverse=True))
        stats['total_num_tags'] = len(tags)

        top_five = []
        for k, v in tags.items():
            s = f"{k} ({v['usage_count']} times)"
            top_five.append(s)

            if len(top_five) == 5:
                break

        stats['top_5_tags'] = top_five

    return stats


def build_blog_statistics(path_md_files) -> dict:
    """
    Going through all Blog Markdown files and collect
    and calculate some basic statistics.
    """
    stats = dict()
    stats['number_of_blog_posts'] = 0

    # Get existing podcast episodes
    posts = [f for f in os.listdir(path_md_files) if isfile(join(path_md_files, f)) and f.endswith('.mdx')]
    for post in posts:
        file_path = f"{path_md_files}/{post}"
        with open(file_path) as f:
            post_frontmatter = frontmatter.load(f)
            
            stats['number_of_blog_posts'] += 1
            
    return stats


def build_podcast_statistics(file_path) -> dict:
    """
    Collect and calculate some basic statistics about the
    basic podcast data.
    """
    stats = dict()

    with open(file_path) as f:
        data = json.load(f)

        stats["number_of_plattforms"] = len(data["platformLinks"])
        stats["plattforms"] = sentence_case(", ".join(data["platformLinks"].keys()))
    
    return stats


def sentence_case(string):
    if string != '':
        result = re.sub('([A-Z])', r' \1', string)
        return result[:1].upper() + result[1:].lower()
    return


def get_time_human_readable(seconds, granularity=4) -> str:
    result = []

    for name, count in intervals:
        value = seconds // count
        if value:
            seconds -= value * count
            if value == 1:
                name = name.rstrip('s')
            result.append("{} {}".format(value, name))

    return ', '.join(result[:granularity])


def print_podcast_episode_stats(stats):
    seconds_per_episode_avg = stats['total_length_seconds'] / stats['number_of_episodes']

    print("")
    print("Podcast Episodes")
    print_headline_spacer()
    print(f"Number of episodes: {stats['number_of_episodes']}")
    print(f"Total length of content: {get_time_human_readable(stats['total_length_seconds'])}")
    print(f"Avg of episode length: {get_time_human_readable(seconds_per_episode_avg)}")
    print(f"Number of tags: {stats['total_num_tags']}")
    print(f"Top 5 tags:")
    for v in stats['top_5_tags']:
        print(f"\t{v}")


def print_podcast_stats(stats):
    print("Podcast")
    print_headline_spacer()
    print(f"Available on platforms: {stats['number_of_plattforms']}")
    print(f"Platforms: {stats['plattforms']}")


def print_blog_post_stats(stats):
    print("")
    print("Blog posts")
    print_headline_spacer()
    print(f"Number of blog posts: {stats['number_of_blog_posts']}")


def print_episode_speaking_time(episode_speaking_stats):
    print("")
    print("Podcast Episode Speaking Time")
    print_headline_spacer()

    for title, episode in enumerate(episode_speaking_stats):
        print(f"{episode['title']} ({episode['pubDate']})")

        for name, length_in_ms in episode['speaking'].items():
            length_in_s = round(length_in_ms / 1000, 0)
            length_in_s = int(length_in_s)

            percent = (length_in_s / episode['length_second']) * 100
            percent = round(percent, 2)

            print(f"\t{name}: {percent}% / {get_time_human_readable(length_in_s)}")

        print("")


def print_overall_speaking_time(overall_speaking_stats):
    print("")
    print("Podcast Overall Speaking Time")
    print_headline_spacer()

    for name, length_in_ms in overall_speaking_stats['speaker'].items():
        length_in_s = round(length_in_ms / 1000, 0)
        length_in_s = int(length_in_s)

        percent = (length_in_s / overall_speaking_stats['total_length_s']) * 100
        percent = round(percent, 2)

        print(f"\t{name}: {percent}% / {get_time_human_readable(length_in_s)}")


def print_headline_spacer():
    print('-----------------')

if __name__ == "__main__":
    # Setup logger
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s [%(levelname)s] %(message)s",
        handlers=[
            logging.StreamHandler()
        ]
    )

    # Collecting data
    episode_stats = build_episode_statistics(build_correct_file_path(PATH_EPISODE_MARKDOWN_FILES))
    episode_speaking_stats, overall_speaking_stats = build_episode_speaking_time_statistics(build_correct_file_path(PATH_EPISODE_MARKDOWN_FILES))
    blog_stats = build_blog_statistics(build_correct_file_path(PATH_BLOG_MARKDOWN_FILES))
    tags_stats = build_tags_statistics(build_correct_file_path(PATH_TAGS_JSON_FILE))
    podcast_stats = build_podcast_statistics(build_correct_file_path(PATH_PODCAST_INFO_JSON_FILE))
    
    # Merging dicts
    podcast_episode_stats = dict()
    podcast_episode_stats = podcast_episode_stats | episode_stats
    podcast_episode_stats = podcast_episode_stats | tags_stats

    print("Engineering Kiosk Summary and statistics")
    print("=========================")
    print_podcast_stats(podcast_stats)
    print_blog_post_stats(blog_stats)
    print_podcast_episode_stats(podcast_episode_stats)
    print_overall_speaking_time(overall_speaking_stats)
    print_episode_speaking_time(episode_speaking_stats)