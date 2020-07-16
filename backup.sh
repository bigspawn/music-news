#!/bin/bash

DEST=/media/external/backups/music-news/

rsync -r -v .archive $DEST
rsync -v docker-compose-prod.yml $DEST
