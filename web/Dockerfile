FROM frolvlad/alpine-glibc:alpine-3.7_glibc-2.26

ENV GO_VER 1.15.2
ENV GO_OS linux
ENV GO_ARCH amd64

ENV PATH $GOPATH/bin:/opt/go/bin:$PATH
ENV GOPATH /go

RUN mkdir -p $GOPATH/src $GOPATH/bin
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache git
RUN mkdir -p /opt
RUN wget -q -O - https://dl.google.com/go/go$GO_VER.$GO_OS-$GO_ARCH.tar.gz \
        | tar -C /opt/ -zxf -
RUN set -eux; \
	export GOROOT = "$ (go env GOROOT)" ; \
	./run-tool-which-requires-GOROOT-set.sh

WORKDIR $GOPATH