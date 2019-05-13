FROM golang:1.11.5

RUN mkdir /app
COPY music-news /app/
RUN chmod +x /app/music-news
WORKDIR /app

CMD ["./music-news"]