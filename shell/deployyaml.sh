#!/bin/bash

if [ -n "$1" ]; then
   filepath=$1
fi

if [ -n "$2" ]; then
   app=$2
   group=$2
fi

if [ -n "$3" ]; then
  image="$3"
fi
confdir=/var/sdk/config
if [ -n "$4" ]; then
  confdir="$4"
fi
datadir=/var/data
if [ -n "$5" ]; then
  datadir="$5"
fi

cat <<EOF > $filepath
apiVersion: apps/v1
kind: Deployment
metadata:
 labels:
   app: ${app}
   group: ${group}
 name: ${app}
 namespace: default
spec:
 replicas: 1
 selector:
   matchLabels:
     app: ${app}
 minReadySeconds: 5
 strategy:
   type: RollingUpdate
   rollingUpdate:
     maxSurge: 1
     maxUnavailable: 1
 template:
   metadata:
     labels:
       app: ${app}
       group: ${group}
   spec:
     containers:
       - name: ${app}
         image: ${image}
         resources:
           requests:
             memory: "10Mi"
             cpu: "10m"
           limits:
             memory: "50Mi"
         imagePullPolicy: IfNotPresent
         volumeMounts:
           - mountPath: /app/config
             name: config
           - mountPath: /data
             name: data
     volumes:
       - name: config
         hostPath:
           path: ${confdir}
           type: DirectoryOrCreate
       - name: data
         hostPath:
           path: ${datadir}/${group}
           type: DirectoryOrCreate


EOF