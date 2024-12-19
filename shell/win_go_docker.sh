#!/bin/bash

buildDir="build"
if [ ! -d $buildDir ]; then
    mkdir $buildDir
fi

if [ -z "$1" ]; then
  echo "参数为空"
  exit 1
fi


GOOS=linux go build -trimpath -o  build/"$1" "$2"
cmd="\"./$1\",\"-c\",\"./config/$1.toml\""
dockerfilepath=build/Dockerfile
cat <<EOF > $dockerfilepath
 FROM jybl/timezone AS tz

 FROM frolvlad/alpine-glibc

 #修改容器时区
 ENV TZ=Asia/Shanghai LANG=C.UTF-8
 COPY --from=tz /usr/share/zoneinfo/\$TZ /usr/share/zoneinfo/\$TZ
 RUN echo \$TZ > /etc/timezone && ln -sf /usr/share/zoneinfo/\$TZ /etc/localtime

 WORKDIR /app

 ADD ./$1 /app

 CMD [$cmd]
EOF

#image=jybl/$1:$(date "+%y%m%d%H%M")
image=jybl/$1
echo "docker build -t $image -f $dockerfilepath $buildDir; docker push $image"
wsl bash -c "cd /mnt/$PWD; pwd; docker build -t $image -f $dockerfilepath $buildDir; docker push $image"
