#!/bin/bash

git pull
docker-compose -f docker-compose-prod.yml rm -s -f news notifier
docker-compose -f docker-compose-prod.yml up --build -d news notifier
