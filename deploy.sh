#!/bin/sh

binFile=hugo-server
serverPath=/home/dev/hugo-server/bin/
echo "start deploy"

echo "cross compile"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

ssh $qcloud "sudo supervisorctl stop hugo-server"
echo "stop server"
scp $binFile $qcloud:$serverPath
echo "file copied!"
ssh $qcloud "sudo supervisorctl restart hugo-server"
echo "server started!"

rm -f $binFile
echo "finish"
