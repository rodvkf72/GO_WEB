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

RUN cd $GOPATH/src && git clone https://github.com/rodvkf72/GO_WEB.git
RUN go get github.com/labstack/echo
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/dgrijalva/jwt-go

WORKDIR $GOPATH