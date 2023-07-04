FROM bitnami/kubectl AS kubectl

FROM docker:20.10.19-cli-alpine3.16

COPY --from=kubectl /opt/bitnami/kubectl/bin/kubectl /bin/

ADD ./tpl /tpl

COPY ./deploy/deploy /bin

CMD ["deploy"]