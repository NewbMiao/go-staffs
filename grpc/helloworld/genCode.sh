#!/bin/bash
set -euo pipefail


workspace=$(cd $(dirname $0) && pwd -P)
cd $workspace

# check protoc
command -v protoc >/dev/null 2>&1 || { 
    echo >&2 "require protoc but it's not installed.  installing..."; 
    if [ "$(uname)"=="Darwin" ] ;then
        brew install protobuf
    else
        echo "not on Darwin env, pls install protoc first!"
        exit 1
    fi;
}

# check protoc-gen-go
command -v protoc-gen-go >/dev/null 2>&1 || { 
    echo >&2 "require protoc-gen-go but it's not installed.  installing..."; 
    go get -u github.com/golang/protobuf/protoc-gen-go; 
}

# gen .pb.go
protoc  --go_out=plugins=grpc:. *.proto

echo "done! see $(ls *.pb.go)"
