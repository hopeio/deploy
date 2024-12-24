#!/bin/bash

if [ -n "$1" ]; then
   filepath=$1
fi

if [ -n "$2" ]; then
   app=$2
fi

if [ -n "$3" ]; then
  port="$3"
fi

cat <<EOF > $filepath
  apiVersion: v1
  kind: Service
  metadata:
    name: ${app}
    namespace: default
    labels:
      app: ${app}
  spec:
    type: ClusterIP
    ports:
      - name: http
        port: 80
        protocol: TCP
        targetPort: ${port}
    selector:
      app: ${app}
EOF