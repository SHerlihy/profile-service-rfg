#!/bin/bash

redis-cli ACL SETUSER profile-service-rfg \>profile-service-pass on allkeys allcommands

GOOS=linux GOARCH=amd64 CGO_ENABLED=0\
 go build\
 -o /home/thehuge/projects/profile-service-rfg/deployment/\
 -ldflags="-X 'main.ALLOWED_ORIGINS="*"'\
 -X 'main.DATABASE_ADDRESS="localhost:6380"'\
 -X 'main.DATABASE_USER="profile-service-rfg"'\
 -X 'main.DATABASE_PASSWORD="profile-service-pass"'"\
 ./main.go
