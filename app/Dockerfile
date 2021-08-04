FROM golang:1.16-alpine

ARG ENV
ENV ENV ${ENV}

RUN apk update \
    && apk upgrade \
    && apk add --no-cache bash openssl git curl

WORKDIR /app

COPY src/ .

RUN tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-v0.6.1.tar.gz \
    && rm dockerize-alpine-linux-amd64-v0.6.1.tar.gz
RUN go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct
RUN go mod download

CMD go get -u github.com/cosmtrek/air \
            && dockerize -wait tcp://postgres:5432 -timeout 60s -wait tcp://redis:6379 -timeout 60s air

