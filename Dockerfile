FROM golang:1.19 as build

ENV GO111MODULE=on
ENV CGO_ENABLED=1

ADD . /build
WORKDIR /build

RUN \
  cd cmd && go build -o /build/music-news


FROM golang:1.19

WORKDIR /srv

COPY --from=build /build/music-news /srv/music-news

CMD ["/srv/music-news"]
