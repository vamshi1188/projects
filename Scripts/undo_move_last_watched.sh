#!/bin/bash

# Define the paths
current_dir="/media/vamshi/HDD/resources/tech with nana/TechworldwithNana - DevOps Bootcamp Updated 1-2024"
watched_dir="/media/vamshi/HDD/resources/tech with nana/TechworldwithNana - DevOps Bootcamp Updated 1-2024/watched_dir"

# Check if the watched directory exists
if [ -d "$watched_dir" ]; then
  # Get the last watched video file in the watched directory
  last_watched=$(ls -1 "$watched_dir"/*.mp4 | tail -n 1)

  # Check if a watched video file was found
  if [ -n "$last_watched" ]; then
    # Move the last watched video file back to the current directory
    mv "$last_watched" "$current_dir"
    echo "Moved: $last_watched to $current_dir"
  else
    echo "No watched videos found in $watched_dir"
  fi
else
  echo "Watched directory not found: $watched_dir"
fi
