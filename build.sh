#!/bin/bash

go build .

docker build -t bigspawn/music-news:1.2.1 .

docker push bigspawn/music-news:1.2.1