# Multi-stage docker build file
# Build go app

FROM golang:1.14-alpine
ADD . /build
WORKDIR /build/app

ENV GOFLAGS="-mod=vendor"

RUN go build -o music-news .

# Build container
FROM golang:alpine
WORKDIR /app
COPY --from=0 /build/app/music-news .
CMD ["./music-news"]
