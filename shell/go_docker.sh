#!/bin/bash

buildDir="build"
if [ ! -d $buildDir ]; then
    mkdir $buildDir
fi

if [ -z "$1" ]; then
  echo "目标参数为空"
  exit 1
fi

if [ -z "$2" ]; then
  echo "源码参数为空"
  exit 1
fi

if [ -n "$3" ]; then
register="$3/"
fi


# build
echo GOOS=linux go build -trimpath -o  build/"$1" "$2"
GOOS=linux go build -trimpath -o  build/"$1" "$2"
cmd="\"./$1\",\"-c\",\"./config/$1.toml\""
dockerfilepath=build/Dockerfile
rundir=$(realpath $(dirname "${BASH_SOURCE[0]}"))
echo $rundir
source ${rundir}/dockerfile.sh $dockerfilepath $1 $cmd $register

#docker run --rm -v $GOPATH:/go -v $PWD:/work -w /work -e GOPROXY=$GOPROXY $GOIMAGE go build  -trimpath -o /work/build/$output /work/$1
image=${register}jybl/$1
source ${rundir}/deployyaml.sh build/${1}.yaml $1 $image
source ${rundir}/serviceyaml.sh build/${1}_service.yaml $1 9000
docker build -t $image -f $dockerfilepath $buildDir; docker push $image
docker push $image
