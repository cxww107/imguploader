FROM golang:alpine AS build

COPY . /go/src/github.com/cxww107/imguploader

RUN \
    cd /go/src/github.com/cxww107/imguploader/cmd/srv/ && \ 
    apk add --no-cache git mercurial &&\
    go get -v && \
    apk del git mercurial && \
    go build -o /srv/imguploader/cmd/srv/app && \
    rm -rf /go/src/*


FROM alpine:latest

WORKDIR /srv/imguploader/cmd/srv/
COPY --from=build /srv/imguploader/cmd/srv/ ./

RUN apk add --no-cache ca-certificates && \
    apk add --update curl gnupg tzdata

COPY ./cmd/srv/certs/ /srv/imguploader/cmd/srv/certs

EXPOSE 8888

WORKDIR /srv/imguploader/

CMD ["/srv/imguploader/cmd/srv/app"]



