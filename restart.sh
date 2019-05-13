#!/bin/bash

VERSION=1.2.0
NAME=music-news-prod
#NAME=music-news

docker stop ${NAME}
docker rm ${NAME}
docker pull bigspawn/music-news:${VERSION}
docker run -d --name=${NAME} bigspawn/music-news:${VERSION}