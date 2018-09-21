#!/bin/sh

IMAGE=$1

cp ../apiproject ./

docker build -t registry.api.weibo.com/111test/$IMAGE .
docker push registry.api.weibo.com/111test/$IMAGE