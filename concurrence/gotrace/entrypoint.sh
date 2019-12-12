#!/bin/bash
set -euo pipefail

if [ $# -eq 0 ]; then
    echo -e "Usage: Require go file pathname(some code enable trace goroutines)\n\
    Shell : sh gotrace.sh goroutines.go\n\
    Docker: $GOTRACECMD goroutines.go\n\
            (you need build image locally first: docker build . -t divan/golang:gotrace1.8)"
    exit;
fi

if [ -f /.dockerenv ]; then # build on docker container
    gotrace $1
else
    echo "Only can run in container of divan/golang:gotrace1.8"
fi