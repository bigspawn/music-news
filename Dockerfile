FROM golang:1.16 as test_build

ENV GO111MODULE=on
ENV CGO_ENABLED=0

ADD . /build
WORKDIR /build/internal
RUN go mod download
WORKDIR /build/cmd
RUN go build -o music-news .


FROM golang:1.16-alpine

WORKDIR /srv
COPY --from=test_build /build/cmd/music-news /srv/music-news
CMD ["/srv/music-news"]
