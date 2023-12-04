#!/bin/bash

redis-cli ACL SETUSER profile-service-rfg \>profile-service-pass on allkeys allcommands

go run main.go -ldflags="-X 'main.ALLOWED_ORIGINS="*"' -X 'main.DATABASE_ADDRESS="localhost:6379"' -X 'main.DATABASE_USER="profile-service-rfg"' -X 'main.DATABASE_PASSWORD="profile-service-pass"'"
