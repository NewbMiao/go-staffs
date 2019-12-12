#!/bin/bash

DOCKERPATH="https://gist.githubusercontent.com/NewbMiao/f4340b483e3dfc057911cba8e7a37562/raw/5e88cbf6edb16af3d78162a69917591e274df8af/Dockerfile"
IMAGETAG="divan/golang:gotrace1.8"

if [ $# -eq 0 ]; then
    echo -e "Usage: Require go file pathname(some code enable trace goroutines)\n\
    Shell : sh gotrace.sh goroutines.go"
    exit;
fi

exist=$(docker images|grep divan/golang |grep gotrace1.8|wc -l)
if [ $exist -eq 0 ]; then
    echo "build image $IMAGETAG"
    {
        workspace=$(cd $(dirname $0) && pwd -P)
        cd $workspace
        if [ ! -f "Dockerfile" ]; then 
            curl -o Dockerfile "$DOCKERPATH"
        fi
        docker build . -t  $IMAGETAG
    }
fi

docker run --rm -it  -p 2000:2000  -v $PWD:/go $IMAGETAG $1