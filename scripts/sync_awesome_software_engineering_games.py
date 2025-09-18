import os
import json
import logging
import tempfile
import shutil

from functions import (
    build_correct_file_path,
    configure_global_logger
)

from constants import (
    AWESOME_SOFTWARE_ENGINEERING_GAMES_GIT_REPO,
    AWESOME_SOFTWARE_ENGINEERING_GAMES_GIT_REPO_NAME,
    AWESOME_SOFTWARE_ENGINEERING_GAMES_JSON_PATH_IN_GIT_REPO,
    AWESOME_SOFTWARE_ENGINEERING_GAMES_IMAGES_PATH_IN_GIT_REPO,
    AWESOME_SOFTWARE_ENGINEERING_GAMES_JSON_STORAGE,
    AWESOME_SOFTWARE_ENGINEERING_GAMES_IMAGE_STORAGE
)

from git import Repo
# from PIL import Image

def sync_awesome_software_engineering_games(json_storage_path, image_storage_path):
    tmp_dir = tempfile.gettempdir()
    tmp_clone_dir = os.path.join(tmp_dir, AWESOME_SOFTWARE_ENGINEERING_GAMES_GIT_REPO_NAME)
    
    # Cloning git repository
    logging.info(f"Cloning {AWESOME_SOFTWARE_ENGINEERING_GAMES_GIT_REPO} into {tmp_clone_dir}...")
    Repo.clone_from(AWESOME_SOFTWARE_ENGINEERING_GAMES_GIT_REPO, tmp_clone_dir)
    logging.info(f"Cloning {AWESOME_SOFTWARE_ENGINEERING_GAMES_GIT_REPO} into {tmp_clone_dir}... successful")
    
    # Reading JSON files
    json_file_dir = os.path.join(tmp_clone_dir, AWESOME_SOFTWARE_ENGINEERING_GAMES_JSON_PATH_IN_GIT_REPO)
    json_files = [json_file for json_file in os.listdir(json_file_dir) if json_file.endswith('.json')]
    logging.info(f"Found {len(json_files)} JSON files in {json_file_dir}")
    
    # Copy JSON files over
    for json_file in json_files:
        # Copy files over
        src = os.path.join(tmp_clone_dir, AWESOME_SOFTWARE_ENGINEERING_GAMES_JSON_PATH_IN_GIT_REPO, json_file)
        dst = os.path.join(json_storage_path, json_file)
        logging.info(f"Copying {json_file} from {src} to {dst}...")
        shutil.copy2(src, dst)

        # Modify file content
        with open(dst) as f:
            data = json.load(f)
            # Modify image key to only contain the filename
            data["image"] = f"./{os.path.basename(data['image'])}"

            # Replace some Genres to avoid duplicates
            if "MMO" in data["german_content"]["genres"]:
                data["german_content"]["genres"].remove("MMO")
                if "Massively Multiplayer" not in data["german_content"]["genres"]:
                    data["german_content"]["genres"].append("Massively Multiplayer")

            if "Simulationen" in data["german_content"]["genres"]:
                data["german_content"]["genres"].remove("Simulationen")
                if "Simulation" not in data["german_content"]["genres"]:
                    data["german_content"]["genres"].append("Simulation")

        # Write new file content
        with open(dst, 'w') as fp:
            json.dump(data, fp, indent=4)

    # Copy images over
    # Right now we only do it oneway.
    # If a game updates its image, we don't delete the old one.
    # Dirty? Maybe. However, fast for now and the assumption is, that this
    # will not happen very often. If this assumption is wrong, we will update the
    # piece of code below.
    images_file_dir = os.path.join(tmp_clone_dir, AWESOME_SOFTWARE_ENGINEERING_GAMES_IMAGES_PATH_IN_GIT_REPO)
    image_files = [image_file for image_file in os.listdir(images_file_dir) if not image_file.startswith(".")]
    logging.info(f"Found {len(image_files)} image files in {images_file_dir}")

    for image_file in image_files:
        # Copy files over
        src = os.path.join(tmp_clone_dir, AWESOME_SOFTWARE_ENGINEERING_GAMES_IMAGES_PATH_IN_GIT_REPO, image_file)
        dst = os.path.join(image_storage_path, image_file)
        logging.info(f"Copying {image_file} from {src} to {dst}...")
        shutil.copy2(src, dst)

        # Resize images
        #image_to_resize = Image.open(dst)
        #if image_to_resize.width > 700 and image_to_resize.height > 700:
        #    logging.info(f"Resizing {image_file} from {image_to_resize.width}x{image_to_resize.height} to 700x700...")
        #    resized_image = image_to_resize.resize((700, 700))
        #    resized_image.save(dst)

    # Removing git clone
    logging.info(f"Removing cloned repository from merged JSON file {tmp_clone_dir}...")
    shutil.rmtree(tmp_clone_dir)

    logging.info("Aaaaand we are done")


if __name__ == "__main__":
    # Setup logger
    configure_global_logger()

    image_storage_path = build_correct_file_path(AWESOME_SOFTWARE_ENGINEERING_GAMES_IMAGE_STORAGE)
    json_storage_path = build_correct_file_path(AWESOME_SOFTWARE_ENGINEERING_GAMES_JSON_STORAGE)

    sync_awesome_software_engineering_games(json_storage_path, image_storage_path)