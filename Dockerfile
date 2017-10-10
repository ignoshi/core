FROM golang:1.8
WORKDIR /go/src/github.com/ignoshi/core
COPY . /go/src/github.com/ignoshi/core
RUN go get
EXPOSE 8000
CMD ["go", "run", "main.go"]
