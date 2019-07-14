# Multi-stage docker build file
# Build go app
FROM golang:alpine
ADD app /app
WORKDIR /app
ENV GOFLAGS="-mod=vendor"
RUN go build -o music-news .

# Build container
FROM golang:alpine
WORKDIR /app
COPY --from=0 /app/music-news .
CMD ["./music-news"]