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


GOOS=linux go build -trimpath -o  build/"$1" "$2"
cmd="\"./$1\",\"-c\",\"./config/$1.toml\""
dockerfilepath=build/Dockerfile
rundir=$(dirname $0)
source ${rundir}/dockerfile.sh $dockerfilepath $1 $cmd $register


#image=jybl/$1:$(date "+%y%m%d%H%M")
image=${register}jybl/$1
source ${rundir}/deployyaml.sh build/${1}.yaml $1 $image /root/config /data
source ${rundir}/service.sh build/${1}_service.yaml $1 9000
echo "docker build -t $image -f $dockerfilepath $buildDir; docker push $image"
wsl bash -c "cd /mnt/$PWD; pwd; docker build -t $image -f $dockerfilepath $buildDir; docker push $image"
