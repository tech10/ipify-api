#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -s -extldflags "-static"' .
