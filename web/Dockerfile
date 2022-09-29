FROM golang:1.19 AS builder

WORKDIR /app
COPY . /app
RUN go build -o /web_bin /app/main.go

## Deploy
FROM gcr.io/distroless/base-debian10 AS runner

WORKDIR /

COPY --from=builder /web_bin /
COPY ./www /www

EXPOSE 8088

ENTRYPOINT ["/web_bin"]
