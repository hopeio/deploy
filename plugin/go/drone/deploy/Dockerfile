#FROM golang:1.22-alpine3.16 AS builder
#
#ENV GOPROXY https://goproxy.io,https://goproxy.cn,direct
#WORKDIR /build
# ADD . /build
#RUN go build -trimpath -o /build/deploy
#

FROM bitnami/kubectl AS kubectl

FROM docker:20.10.19-cli-alpine3.16

#修改容器时区
ENV TZ=Asia/Shanghai LANG=C.UTF-8

RUN echo -n 'VFppZjIAAAAAAAAAAAAAAAAAAAAAAAADAAAAAwAAAAAAAAAdAAAAAwAAAAyAAAAAoJeigKF5BPDIWV6AyQn5cMnTvQDLBYrwy3xAANI7PvDTi3uA1EKt8NVFIgDWTL/w1zy/ANgGZnDZHfKA2UF88B66UiAfaZuQIH6EoCFJfZAiZ6EgIylfkCRHgyAlEnwQJidlICbyXhAoB0cgKNJAEAIBAgECAQIBAgECAQIBAgECAQIBAgECAQIBAgECAABx1wAAAAB+kAEEAABwgAAITE1UAENEVABDU1QAAAAAAAAAVFppZjIAAAAAAAAAAAAAAAAAAAAAAAADAAAAAwAAAAAAAAAdAAAAAwAAAAz/////fjZDKf////+gl6KA/////6F5BPD/////yFlegP/////JCflw/////8nTvQD/////ywWK8P/////LfEAA/////9I7PvD/////04t7gP/////UQq3w/////9VFIgD/////1ky/8P/////XPL8A/////9gGZnD/////2R3ygP/////ZQXzwAAAAAB66UiAAAAAAH2mbkAAAAAAgfoSgAAAAACFJfZAAAAAAImehIAAAAAAjKV+QAAAAACRHgyAAAAAAJRJ8EAAAAAAmJ2UgAAAAACbyXhAAAAAAKAdHIAAAAAAo0kAQAgECAQIBAgECAQIBAgECAQIBAgECAQIBAgECAQIAAHHXAAAAAH6QAQQAAHCAAAhMTVQAQ0RUAENTVAAAAAAAAAAKQ1NULTgK'|base64 -d > /etc/localtime && mkdir -p /usr/share/zoneinfo/Asia && ln -s /etc/localtime /usr/share/zoneinfo/Asia/Shanghai && echo -n 'Asia/Shanghai' > /etc/timezone


COPY --from=kubectl /opt/bitnami/kubectl/bin/kubectl /bin/

ADD ./tpl /tpl

COPY ./build/deploy /bin

CMD ["deploy"]