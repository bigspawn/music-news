FROM golang:1.16 as test

ENV GO111MODULE=on

ADD . /build

WORKDIR /build/internal

RUN go test -v -race .



FROM golang:1.16-alpine as build

ENV GO111MODULE=on
ENV CGO_ENABLED=0

ADD . /build

WORKDIR /build/cmd

RUN go mod download
RUN go build -o music-news .


FROM golang:1.16-alpine

WORKDIR /srv

COPY --from=build /build/cmd/music-news /srv/music-news

CMD ["/srv/music-news"]
