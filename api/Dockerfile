FROM golang:1.12.6

WORKDIR /go/src/api

RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get github.com/gorilla/mux

RUN mkdir -p /go/src/api
COPY . /go/src/api

CMD [ "go", "run", "main.go", "model.go"]
