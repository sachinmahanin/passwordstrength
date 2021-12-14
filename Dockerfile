FROM golang:alpine as builder

ARG BUILD_VERSION
ARG COMMIT_SHA
ARG BUILD_SHORT_VERSION

ENV GOARCH=amd64
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV APP_VERSION=1
ENV HOSTNAME=localhost
ENV APP_PORT=18605

COPY . /go/src/github.com/sachinmahanin/passwordStrength
WORKDIR /go/src/github.com/sachinmahanin/passwordStrength

RUN go build -mod=vendor -o bin/passwordStrength
RUN ls /go/src/github.com/sachinmahanin/passwordStrength
RUN chmod +x /go/src/github.com/sachinmahanin/passwordStrength
EXPOSE 18605/tcp
CMD ./bin/passwordStrength