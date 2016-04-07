#!/bin/sh

CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build && scp tgwalkrbot $BOTHOST:$BOTPATH
