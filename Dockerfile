FROM golang:1.9

RUN go get github.com/mitchellh/gox

RUN go get github.com/Masterminds/glide
