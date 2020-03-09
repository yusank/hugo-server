#!/bin/sh

binFile=hugo-server
serverPath=/home/dev/hugo-server/bin/
echo "start deploy"

echo "cross compile"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

scp $binFile $qcloud:$serverPath

rm -f $binFile
echo "finish"
