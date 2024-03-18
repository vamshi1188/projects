#!/bin/bash

# Define the paths
current_dir="/media/vamshi/HDD/resources/tech with nana/TechworldwithNana - DevOps Bootcamp Updated 1-2024"
watched_dir="/media/vamshi/HDD/resources/tech with nana/TechworldwithNana - DevOps Bootcamp Updated 1-2024/watched_dir"

# Check if the watched directory exists, create it if not
if [ ! -d "$watched_dir" ]; then
  mkdir -p "$watched_dir"
fi

# Get the first video file in the current directory
first_video=$(ls -1 "$current_dir"/*.mp4 | head -n 1)

# Check if a video file was found
if [ -n "$first_video" ]; then
  # Move the first video file to the watched directory
  mv "$first_video" "$watched_dir"
  echo "Moved: $first_video to $watched_dir"
else
  echo "No videos found in $current_dir"
fi
