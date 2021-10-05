FROM golang:latest AS builder

ENV PORT=8081
ENV GO111MODULE=on

WORKDIR $GOPATH/src/pokemonwiki

COPY . /go/src/pokemonwiki

RUN go get .
RUN go build

EXPOSE $PORT

ENTRYPOINT [ "./pokemonwiki" ]