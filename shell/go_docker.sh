#!/bin/bash

buildDir="build"
if [ ! -d $buildDir ]; then
    mkdir $buildDir
fi

if [ -z "$1" ]; then
  echo "参数为空"
  exit 1
fi

if [ -n "$3" ]; then
register="$3/"
fi


# build
GOOS=linux go build -trimpath -o  build/"$1" "$2"
cmd="\"./$1\",\"-c\",\"./config/$1.toml\""
dockerfilepath=build/Dockerfile
source $(dirname $0)/dockerfile.sh $dockerfilepath $1 $cmd $register

#docker run --rm -v $GOPATH:/go -v $PWD:/work -w /work -e GOPROXY=$GOPROXY $GOIMAGE go build  -trimpath -o /work/build/$output /work/$1
image=${register}jybl/$1
docker build -t $image -f $dockerfilepath $buildDir; docker push $image
docker push $image
