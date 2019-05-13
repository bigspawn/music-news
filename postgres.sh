#!/bin/bash

docker run --name postgres -p 8532:5432 -e POSTGRES_USER=go-music -e POSTGRES_PASSWORD=mysecretpassword -d postgres
