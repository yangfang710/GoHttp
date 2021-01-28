#!/bin/bash
set -e

function server_start() {
    /bin/server
}

process=$1
if [ "$process" = 'server' ]; then
    server_start
else
    exit 0
fi
