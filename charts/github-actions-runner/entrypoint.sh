#!/usr/bin/env bash

python3 -m http.server 8080 &

./run-helper.sh "$@"
