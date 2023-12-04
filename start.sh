#!/bin/bash

redis-cli ACL SETUSER profile-service-rfg \>profile-service-pass on allkeys allcommands

go run main.go
