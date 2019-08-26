#!/bin/bash

sn=portfolio
cd $GOPATH/ambiente/go/ReactGolangRestfullApiMongoJWT

tmux new-session  -s "$sn" -n devsession -d \; \
	send-keys "make hot-reload" C-m \; \
split-window -p 40 -t "$sn" \; \
	send-keys "make hot-server" C-m

tmux attach-session -t "$sn"
