FROM golang:1.19 AS builder

WORKDIR /app
COPY . /app
RUN export GOPROXY=https://goproxy.io,direct; go build -o /web_bin /app/main.go

## Deploy
FROM debian:buster AS runner

WORKDIR /

COPY --from=builder /web_bin /
COPY ./www /www

EXPOSE 8088

ENTRYPOINT ["/web_bin"]
