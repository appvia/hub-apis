#!/bin/bash
#

KUBE_SERVER=${KUBE_SERVER:-"http://kube-apiserver:8080"}

echo "Applying the Hub API manifests"
if ! kubectl --server=${KUBE_SERVER} apply --validate=false -f /deploy; then
  echo "[error] failed to apply the hub api manifests"
  exit 1
fi

sleep 365d
