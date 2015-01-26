#!/bin/bash

epochTime=$(date +%s000)

# clean so not prompted to overwrite
rm /root/angular_server/pong_snapshots/output.raw
rm /root/angular_server/pong_snapshots/latest_picture/latest.gif

echo "capturing snapshot"

# get 1 frame of 1920x1080 h264
/root/angular_server/camera_control/capture -F -o -c 1 > /root/angular_server/pong_snapshots/output.raw

# convert to gif
ffmpeg -f h264 -i /root/angular_server/pong_snapshots/output.raw -s 480x270 -qscale:v 10 /root/angular_server/pong_snapshots/${epochTime}_%03d.gif

# copy to latest_picture folder
cp /root/angular_server/pong_snapshots/${epochTime}_001.gif /root/angular_server/pong_snapshots/latest_picture/latest.gif