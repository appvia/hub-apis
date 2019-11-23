NAME=hub-apis
AUTHOR=appvia
AUTHOR_EMAIL=info@appvia.io
REGISTRY=quay.io
GOVERSION ?= 1.13.0
ROOT_DIR=${PWD}
HARDWARE=$(shell uname -m)
GIT_SHA=$(shell git --no-pager describe --always --dirty)
BUILD_TIME=$(shell date '+%s')
VERSION ?= $(shell awk '/release.*=/ { print $$3 }' doc.go | sed 's/"//g')
DEPS=$(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
PACKAGES=$(shell go list ./...)
LFLAGS ?= -X main.gitsha=${GIT_SHA} -X main.compiled=${BUILD_TIME}
VETARGS ?= -asmdecl -atomic -bool -buildtags -copylocks -methods -nilfunc -printf -rangeloops -shift -structtags -unsafeptr
APIS=config/v1 core/v1 org/v1 rbac/v1 clusters/v1 store/v1

.PHONY: test authors changelog build docker static release lint cover vet glide-install

default: all

golang:
	@echo "--> Go Version"
	@go version

all: golang
	@echo "--> Generating Clientsets & Deepcopies"
	@rm -rf pkg/client 2>/dev/null
	@hack/update-codegen.sh
	@${MAKE} openapi-gen
	@${MAKE} register-gen
	@${MAKE} crd-gen

openapi-gen:
	@echo "--> Generating OpenAPI files"
	@echo "--> packages $(APIS)"
	@$(foreach api,$(APIS), \
		openapi-gen -h hack/custom-boilerplate.go.txt \
			--output-file-base zz_generated_openapi \
			-i github.com/appvia/hub-apis/pkg/apis/$(api) \
			-p github.com/appvia/hub-apis/pkg/apis/$(api); )

register-gen:
	@echo "--> Generating Schema register.go"
	@echo "--> packages $(APIS)"
	@$(foreach api,$(APIS), \
		register-gen -h hack/custom-boilerplate.go.txt \
			--output-file-base zz_generated_register \
			-i github.com/appvia/hub-apis/pkg/apis/$(api) \
			-p github.com/appvia/hub-apis/pkg/apis/$(api); )

crd-gen:
	@echo "--> Generating CRD deployment files"
	@mkdir -p deploy
	@rm -f deploy/* 2>/dev/null || true
	@controller-gen crd:trivialVersions=true,preserveUnknownFields=false paths=./pkg/apis/...  output:dir=deploy
