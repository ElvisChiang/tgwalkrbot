#!/bin/sh

ssh $BOTHOST "sh -c 'tail -f $BOTPATH/log.txt'"
