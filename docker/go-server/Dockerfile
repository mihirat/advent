FROM golang:1.7.3-onbuild

RUN mkdir -p /go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

CMD ["go-wrapper", "run"]
