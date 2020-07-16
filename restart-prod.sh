#!/bin/bash

docker-compose -f docker-compose-prod.yml up --force-recreate -d news notifier
