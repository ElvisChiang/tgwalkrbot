#!/bin/sh

ssh -n -f $BOTHOST "sh -c 'cd $BOTPATH; nohup ./tgwalkrbot >> log.txt 2>&1 &'"
