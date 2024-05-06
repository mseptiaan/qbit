FROM golang:1.22-alpine AS builder

LABEL maintainer="Muhammad Septian <muhammad.septian@anteraja.id>"

WORKDIR /build

COPY . ./
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main ./cmd

FROM alpine:3.15.5

RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

COPY --from=builder ["/build/main", "/"]
EXPOSE 8080
EXPOSE 50000
CMD ["./main"]