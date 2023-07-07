#!/bin/bash
cluster=$1

mkdir cert/$cluster
echo $CACRT |base64 -d > cert/$cluster/ca.crt
echo $DEVCRT |base64 -d > cert/$cluster/dev.crt
echo $DEVKEY |base64 -d > cert/$cluster/dev.key


kubectl config set-cluster k8s --server=$server --certificate-authority=cert/$cluster/ca.crt --embed-certs=true --kubeconfig=/root/.kube/config
kubectl config set-credentials dev --client-certificate=cert/$cluster/dev.crt --client-key=cert/$cluster/dev.key --embed-certs=true --kubeconfig=/root/.kube/config
kubectl config set-context dev --cluster=k8s --user=dev --kubeconfig=/root/.kube/config
kubectl config use-context dev --kubeconfig=/root/.kube/config