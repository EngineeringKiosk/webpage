import logging
import os
import json

from constants import TRANSCRIPT_STORAGE_DIR

def build_correct_file_path(file_path) -> str:
    """
    Build the correct path to files.

    If a python script gets called from `scripts` folder, 
    we need to get a level up.
    This function takes care about the correct path.
    """
    p = file_path

    # Determine if the script is called from root
    # or from the scripts directory.
    directory_path = os.getcwd()
    folder_name = os.path.basename(directory_path)
    if folder_name == "scripts":
        p = f"../{file_path}"
    
    return p

def get_podcast_episode_transcript_slim_path_by_episode_number(episode_number) -> str:
    transcript_file = f"{episode_number}-transcript-slim.json"
    return get_podcast_episode_transcript_path_by_episode_number(transcript_file)

def get_podcast_episode_transcript_path_by_episode_number(transcript_file) -> str:
    file_path = build_correct_file_path(TRANSCRIPT_STORAGE_DIR) + '/' + transcript_file

    if os.path.exists(file_path):
        # Depending on which subfolder the script is called, we need to adjust the path.
        if file_path.startswith('../'):
            file_path = file_path[3:]
        return file_path

    return ''

def configure_global_logger():
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s [%(levelname)s] %(message)s",
        handlers=[
            logging.StreamHandler()
        ]
    )

def read_json_file(file_path):
    """
    Reads the JSON file {file_path} and returns
    the content as parsed JSON dict.
    If the file does not exist, returns an empty dict.
    """
    if not os.path.exists(file_path):
        return {}

    try:
        with open(file_path, "r", encoding="utf-8") as f:
            data = json.load(f)
        return data
    except (json.JSONDecodeError, OSError):
        # Falls Datei korrupt oder nicht lesbar -> leeres Dict zurückgeben
        return {}