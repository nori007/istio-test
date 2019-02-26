FROM golang:1.11-alpine

ENV DISTRIBUTION_DIR /go/src/github.com/docker/distribution
ENV DOCKER_BUILDTAGS include_oss include_gcs include_obs

ARG GOOS=linux
ARG GOARCH=amd64

RUN apk add --no-cache bash
RUN apk add build-base
RUN apk add git

# 소스 다운로드
WORKDIR /go/src/
RUN git clone https://github.com/nori007/istio-test.git

WORKDIR /go/src/go-server-2
ENV GO111MODULE=on
RUN go mod download
RUN go build

EXPOSE 9993
CMD ["./go-server-2"]