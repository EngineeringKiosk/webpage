import argparse
import json
import os
from os.path import isfile, join
import sys
import logging

from functions import (
    build_correct_file_path,
    configure_global_logger,
    read_json_file
)

from constants import (
    EPISODES_STORAGE_DIR,
    BLOGPOST_CONTENT_FILES,
    GERMAN_TECH_PODCAST_CONTENT_FILES,
    TAG_FILE_CONTENT,
    TAG_FILE_GERMAN_TECH_PODCASTS,
    TAG_FILE_AWESOME_SOFTWARE_ENGINEERING_GAMES,
    AWESOME_SOFTWARE_ENGINEERING_GAMES_CONTENT_FILES
)

import frontmatter

def read_all_tags_from_content_files(pathes, json_key="tags"):
    """
    Reads all tags from all content files of this website.
    A content file are Markdown (md) or MDX files like
    blog posts or podcast episodes.

    Returns a dictonary with the tag as key and the
    usage count as value.

    @param: Array of folder pathes to search for files
    @return: Dict with tags as keys and usage count as values
    """
    tags = {}
    for p in pathes:
        logging.info(f"Reading tags from files in {p} ...")

        # Get existing content files in Markdown or MDX
        content_files = [f for f in os.listdir(p) if isfile(join(p, f)) and (f.endswith('.md') or f.endswith('.mdx'))]
        for content_file in content_files:
            full_file_path = f'{p}/{content_file}'
            logging.info(f"Processing file {full_file_path} ...")

            with open(full_file_path) as f:
                fm = frontmatter.load(f)
                frontmatter_tags = fm.get(json_key)

                for tag in frontmatter_tags:
                    if tag in tags:
                        tags[tag] += 1
                    else:
                        tags[tag] = 1

        # Get existing content files in JSON
        content_files = [f for f in os.listdir(p) if isfile(join(p, f)) and f.endswith('.json')]
        for content_file in content_files:
            full_file_path = f'{p}/{content_file}'
            logging.info(f"Processing file {full_file_path} ...")

            with open(full_file_path) as f:
                fm = json.load(f)

                json_key_parts = json_key.split(".")
                for part in json_key_parts:
                    frontmatter_tags = fm.get(part, {})
                    fm = fm.get(part, {})

                if frontmatter_tags is not None:
                    for tag in frontmatter_tags:
                        if tag in tags:
                            tags[tag] += 1
                        else:
                            tags[tag] = 1

    return tags

def get_all_tags_without_description(tag_descriptions, tags):
    tags_with_description = set(())
    tags_without_descriptions = set(())

    # Collecting tags that are already part of the tag file
    for tag_name, tag_descriptions in tag_descriptions.items():
        # If one of the keys don't exist
        if "short_desc" not in tag_descriptions or "long_desc" not in tag_descriptions:
            tags_without_descriptions.add(tag_name)
            continue

        # If both keys exists and are filled
        if ("short_desc" in tag_descriptions and tag_descriptions['short_desc']) \
            and ("long_desc" in tag_descriptions and tag_descriptions['long_desc']):
            tags_with_description.add(tag_name)
            continue

        tags_without_descriptions.add(tag_name)

    # Checking for new tags
    for tag_name in tags.keys():
        if tag_name not in tags_with_description and tag_name not in tags_without_descriptions:
            tags_without_descriptions.add(tag_name)

    return tags_without_descriptions


if __name__ == "__main__":
    # Argument and parameter parsing
    cli_parser = argparse.ArgumentParser(description='Find tags without descriptions. Without `-write-file`, the tags get printed to stdout.')
    cli_parser.add_argument(
        '-write-file',
        action='store_true',
        help='Modifies the local tag storage file and adds missing tags (default: %(default)s)')

    # We have two modes:
    #   - website-content: Things like blog posts or podcast markdown files
    #   - german-tech-podcasts: Imported content like src/data/german-tech-podcasts.json
    cli_parser.add_argument('mode',
        metavar='mode',
        type=str,
        default='website-content',
        const='website-content',
        nargs='?',
        choices=['website-content', 'german-tech-podcasts', 'awesome-software-engineering-games'],
        help='Mode to execute. Supported: website-content, german-tech-podcasts, awesome-software-engineering-games (default: %(default)s)')

    args = cli_parser.parse_args()

    # Setup logger
    configure_global_logger()

    tags = {}
    content_pathes = []
    tag_file_path = ""
    json_key = "tags"
    if args.mode == "website-content":
        tag_file_path = build_correct_file_path(TAG_FILE_CONTENT)
        content_pathes = [
            build_correct_file_path(EPISODES_STORAGE_DIR),
            build_correct_file_path(BLOGPOST_CONTENT_FILES),
        ]

    if args.mode == "german-tech-podcasts":
        tag_file_path = build_correct_file_path(TAG_FILE_GERMAN_TECH_PODCASTS)
        content_pathes = [
            build_correct_file_path(GERMAN_TECH_PODCAST_CONTENT_FILES)
        ]

    if args.mode == "awesome-software-engineering-games":
        # Those are genres, not tags
        tag_file_path = build_correct_file_path(TAG_FILE_AWESOME_SOFTWARE_ENGINEERING_GAMES)
        content_pathes = [
            build_correct_file_path(AWESOME_SOFTWARE_ENGINEERING_GAMES_CONTENT_FILES)
        ]
        json_key = "german_content.genres"

    tags = read_all_tags_from_content_files(content_pathes, json_key=json_key)
    len_content_tags = len(tags)

    logging.info(f"Reading existing tag descriptions from {tag_file_path} ...")
    tag_descriptions = read_json_file(tag_file_path)
    logging.info(f"Determining tags with missing descriptions out of {len_content_tags} unique content tags ...")
    tags_without_desc = get_all_tags_without_description(tag_descriptions, tags)
    logging.info(f"Found {len(tags_without_desc)} tags with missing descriptions out of {len_content_tags} unique content tags ...")

    # Should we print out the files?
    if not args.write_file:
        for t in tags_without_desc:
            print(t)

        # We let it fail, because in a perfect world, there should be
        # no missing descriptions. However, if we let this run as a CI
        # job, the CI job failure pings us to take action.
        if len(tags_without_desc):
            sys.exit(1)

        sys.exit(0)

    logging.info(f"Writing missing tag structures into file {tag_file_path} ...")

    # Updating the local JSON tag file:
    #   - write the missing tags
    for t in tags_without_desc:
        # If the tag already exists, mostly a subkey is missing
        if t not in tag_descriptions:
            tag_descriptions[t] = {}

        subKeys = ["short_desc", "long_desc", "usage_count"]
        for k in subKeys:
            if k not in tag_descriptions[t]:
                tag_descriptions[t][k] = ""

    # Updating the local JSON tag file:
    #   - updating the usage count
    for t in list(tag_descriptions):
        usage_count = 0
        if t in tags:
            usage_count = tags[t]

        # Delete tags that don't have any content
        if usage_count == 0:
            del tag_descriptions[t]
            continue

        tag_descriptions[t]["usage_count"] = usage_count

    # Finally write the file
    with open(tag_file_path, 'w') as fp:
        json.dump(tag_descriptions, fp, indent=4)

    logging.info(f"Writing missing tag structures into file {tag_file_path} ... done")
    sys.exit(0)
