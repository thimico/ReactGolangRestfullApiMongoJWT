#!/bin/bash

sn=portfolio
cd $GOPATH/src/ReactGolangRestfullApiMongoJWT
source /etc/environment

tmux new-session  -s "$sn" -n devsession -d \; \
	send-keys "make hot-reload" C-m \; \
split-window -p 40 -t "$sn" \; \
	send-keys "go run main.go" C-m

tmux attach-session -t "$sn"
