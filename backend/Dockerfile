FROM golang:1.8.3

ENV PACKAGE_NAME bitbucket.org/andreychernih/tweemote
ENV APP_DIR ${GOPATH}/src/${PACKAGE_NAME}
ENV VENDOR_DIR ${APP_DIR}/vendor

RUN mkdir -p ${APP_DIR}
WORKDIR ${APP_DIR}

# TODO lock govendor version
RUN go get github.com/kardianos/govendor
CMD go-wrapper run

RUN mkdir ${VENDOR_DIR}
COPY vendor/vendor.json ${VENDOR_DIR}/vendor.json
RUN govendor sync
RUN set -eu; \
    for dir in $(govendor list|awk '{ print $2 }'); do \
      echo "Installing ${dir}"; \
      cd "${VENDOR_DIR}/${dir}"; \
      go install || true; \
    done

COPY . ${APP_DIR}
RUN go-wrapper install




FROM alpine:latest

ENV APP_DIR=/tweemote \
    PATH=$PATH:/tweemote

# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN set -ex && \
      apk add --no-cache ca-certificates

COPY config.yaml $APP_DIR/config.yaml
COPY --from=0 /go/bin/tweemote $APP_DIR/tweemote
COPY migrations $APP_DIR/migrations

WORKDIR $APP_DIR
