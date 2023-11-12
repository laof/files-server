FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
ADD . .

ENV GOPROXY=https://goproxy.cn

RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]