#!/bin/bash

rsync -r -v .archive /media/external/backups/music-news/.archive
rsync -v docker-compose-prod.yml /media/external/backups/music-news
