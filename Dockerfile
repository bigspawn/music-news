FROM golang:1.15 as test

ENV GOFLAGS="-mod=vendor"

ADD . /build

WORKDIR /build/app

RUN go test -v -race .



FROM golang:1.15-alpine as build

ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0

ADD . /build

WORKDIR /build/app

RUN go build -o music-news .



FROM golang:1.15-alpine

WORKDIR /srv

COPY --from=build /build/app/music-news /srv/music-news

CMD ["/srv/music-news"]
