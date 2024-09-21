#!/bin/sh
docker stop docker-go-local
docker rm docker-go-local
docker build --label docker-go-multi -t docker-go-multi:0.0.1 -t docker-go-multi:latest  -f dockerfile-go-multi .
docker run -d -l docker-go-local --name docker-go-local -p 8181:8181 docker-go-multi:0.0.1
docker image prune