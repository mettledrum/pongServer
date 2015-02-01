#!/bin/bash

echo "making new gif"

rm /root/angular_server/pong_snapshots/pong_gif/pong.gif

gifsicle --delay=50 --loop /root/angular_server/pong_snapshots/*.gif > /root/angular_server/pong_snapshots/pong_gif/pong.gif