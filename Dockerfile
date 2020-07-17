FROM golang:1.14 as build

ENV GOFLAGS="-mod=vendor"

ADD . /build

WORKDIR /build/app

RUN go test -v -race .

RUN go build -o music-news .

# Build container
FROM golang:1.14-alpine

WORKDIR /srv

COPY --from=build /build/app/music-news /srv/music-news

CMD ["/srv/music-news"]
