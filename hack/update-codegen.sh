#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

${GOPATH}/src/k8s.io/code-generator/generate-groups.sh all  \
  github.com/appvia/hub-apis/pkg/client \
  github.com/appvia/hub-apis/pkg/apis \
  "core:v1 org:v1 config:v1 rbac:v1 clusters:v1 store:v1" \
  --go-header-file hack/custom-boilerplate.go.txt
