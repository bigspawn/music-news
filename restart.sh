#!/bin/bash

docker-compose stop news notifier && \
docker-compose rm -f news notifier && \
docker-compose pull news && \
docker-compose up -d news notifier

