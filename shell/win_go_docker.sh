#!/bin/bash

buildDir="build"
if [ ! -d $buildDir ]; then
    mkdir $buildDir
fi

while getopts ":t:s:r:" opt; do
  case ${opt} in
    t )
      target=$OPTARG
      ;;
    s )
      source_path=$OPTARG
      ;;
    r )
      register="$OPTARG/"
      ;;
    \? )
      echo "无效参数: -$OPTARG" 1>&2
      exit 1
      ;;
    : )
      echo "参数 -$OPTARG 需要一个值" 1>&2
      exit 1
      ;;
  esac
done

GOOS=linux go build -trimpath -o  build/"$target" "$source_path"
cmd="\"./$target\",\"-c\",\"./config/$target.toml\""
dockerfilepath=build/Dockerfile
rundir=$(realpath $(dirname "${BASH_SOURCE[0]}"))
echo $rundir
source ${rundir}/dockerfile.sh -f $dockerfilepath -a $target -c $cmd -r $register


#image=jybl/$target:$(date "+%y%m%d%H%M")
image=${register}jybl/$target
source ${rundir}/deployyaml.sh -f build/${target}.yaml -a $target -i $image -c /root/config -d /data
source ${rundir}/serviceyaml.sh -f build/${target}_service.yaml -a $target -p 9000
echo "docker build -t $image -f $dockerfilepath $buildDir; docker push $image"
wsl bash -c "cd /mnt/$PWD; pwd; docker build -t $image -f $dockerfilepath $buildDir; docker push $image"
