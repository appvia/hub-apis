FROM alpine:3.10
LABEL Name=hub-apiserver \
      Release=https://github.com/appvia/hub-apiserver \
      Maintainer=gambol99@gmail.com \
      Url=https://github.com/appvia/hub-apiserver \
      Help=https://github.com/appvia/hub-apiserver/issues

RUN apk add --no-cache ca-certificates bash curl

RUN curl -sL https://storage.googleapis.com/kubernetes-release/release/v1.16.3/bin/linux/amd64/kubectl -o /bin/kubectl && \
    chmod +x /bin/kubectl

COPY deploy/ /deploy
COPY hack/entrypoint.sh /entrypoint.sh

USER 65534

ENTRYPOINT [ "/entrypoint.sh" ]
