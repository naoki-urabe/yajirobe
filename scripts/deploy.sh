#!/bin/bash
USER_HOST=`echo "$USER@$HOST"`
SYNC_PATH=`echo "$USER_HOST:/home/$USER/app/yajirobe"`
SSH_PORT="ssh -p $HOST_PORT"
rsync -e "$SSH_PORT" -acvz  --delete ./ $SYNC_PATH
ssh $USER_HOST -p $HOST_PORT "cd app/yajirobe;/usr/bin/docker-compose down -v;docker pull ghcr.io/naoki-urabe/yajirobe/api:prod;docker pull ghcr.io/naoki-urabe/yajirobe/front:prod;docker pull ghcr.io/naoki-urabe/yajirobe/task:prod; docker image prune -f;/usr/bin/docker-compose -f docker-compose.prod.yml up -d"