FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
ADD . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]