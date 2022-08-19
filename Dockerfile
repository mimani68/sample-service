#
# 1. Build Container
#
FROM golang:1.16.3-buster AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN apt update && apt install git

# First add modules list to better utilize caching
COPY go.sum go.mod /app/

# Download dependencies
RUN go mod download

# Adding sourc eof project
COPY . /app
RUN git rev-parse --short HEAD > /app/VERSION

# Build components.
# Put built binaries and runtime resources in /app dir ready to be copied over or used.
RUN go build -o ./build/app main.go

#
# 2. Runtime Container
#
FROM alpine:3.15

LABEL maintainer="Mahdi Imani <imani.mahdi@gmail.com>"

ENV TZ=Asia/Tehran

RUN apk update && \
    apk add --update --no-cache tzdata openssl && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app

COPY --from=build /app/build/app /app/
COPY --from=build /app/logs /app/logs
COPY --from=build /app/VERSION /app/VERSION

EXPOSE 3000

CMD ["./app"]