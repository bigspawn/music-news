#!/bin/bash

rsync -r .archive /media/external/backups/music-news/.archive
rsync docker-compose-prod.yml /media/external/backups/music-news
