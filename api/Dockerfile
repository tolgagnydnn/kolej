FROM golang:1.12.6

RUN mkdir -p /go/src/api
RUN go get github.com/gorilla/mux

WORKDIR /go/src/api

COPY . /go/src/api

CMD [ "go", "run", "main.go", "model.go"]
