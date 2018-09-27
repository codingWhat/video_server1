#!/bin/bash

# Build Web UI
cd ~/github/src/video_server/web
go install
mkdir -p ~/github/bin/video_server_web_ui
cp ~/github/bin/web ~/github/bin/video_server_web_ui/web
cp -R ~/github/src/video_server/templates ~/github/bin/video_server_web_ui/