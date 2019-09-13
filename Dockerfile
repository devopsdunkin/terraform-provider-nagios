FROM golang:alpine

WORKDIR $GOPATH
RUN apk add --update git bash
RUN export GO111MODULE=auto

COPY . $GOPATH
RUN go build -o /tf-provider/terraform-provider-nagios

ENTRYPOINT ["/bin/bash"]