#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if ! test -x ${GOPATH}/src/k8s.io/code-generator/generate-groups.sh ; then
  echo "Installing k8s.io/code-generator/cmd/..."
  export GO111MODULE=off
  # We need to build without modules but with a specified version...
  go get -d k8s.io/code-generator/cmd/...
  (
    cd ${GOPATH}/src/k8s.io/code-generator
    git checkout kubernetes-1.14.1
    go build -o ${GOPATH}/bin/ ./cmd/register-gen/...
  )
fi

${GOPATH}/src/k8s.io/code-generator/generate-groups.sh all  \
  github.com/appvia/hub-apis/pkg/client \
  github.com/appvia/hub-apis/pkg/apis \
  "core:v1 org:v1 config:v1 rbac:v1 clusters:v1 store:v1" \
  --go-header-file hack/custom-boilerplate.go.txt
