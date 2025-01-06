#!/bin/bash

if [ -n "$1" ]; then
   filepath=$1
fi

if [ -n "$2" ]; then
   app=$2
fi

if [ -n "$3" ]; then
   cmd=$3
fi

if [ -n "$4" ]; then
  register="$4/"
fi

cat <<EOF > $filepath
FROM ${register}jybl/timezone AS tz

FROM ${register}frolvlad/alpine-glibc

#修改容器时区
ENV TZ=Asia/Shanghai LANG=C.UTF-8
COPY --from=tz /usr/share/zoneinfo/\$TZ /usr/share/zoneinfo/\$TZ
RUN echo \$TZ > /etc/timezone && ln -sf /usr/share/zoneinfo/\$TZ /etc/localtime

WORKDIR /app

ADD ./${app} /app

CMD [$cmd]
EOF