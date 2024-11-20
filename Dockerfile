#Docker Pipeline
FROM golang:1.23.2-alpine3.20 as builder
WORKDIR /app
COPY . /app
RUN go mod download && \
    go build -o ./bin/music_library

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/bin/music_library ./bin/music_library
EXPOSE 9000
ENTRYPOINT ["./bin/music_library"]
