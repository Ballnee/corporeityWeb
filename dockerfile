FROM golang:1.19.5-bullseye AS builder

ENV GOARCH=amd64    \
    GOOS=linux  \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY . .

RUN go build -o app ./backend/.

FROM ubuntu:20.04

RUN apt-get -y update && DEBIAN_FRONTEND=noninteractive apt-get -y install tzdata

ENV TZ="Asia/Shanghai"

WORKDIR /app

COPY --from=builder /build/app .
COPY --from=builder /build/backend/config.yaml .

EXPOSE 8080
ENTRYPOINT ["/app/app"]