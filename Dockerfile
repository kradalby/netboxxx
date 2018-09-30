FROM golang:1.11.0-stretch as builder

RUN mkdir -p /app/bin
ADD . /app
WORKDIR /app
RUN go get github.com/mitchellh/gox
RUN gox -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"

