FROM golang:latest AS builder

ENV PORT=8081
ENV GO111MODULE=on

WORKDIR $GOPATH/src/pokemonwiki

COPY . /go/src/pokemonwiki

RUN wget -qO - https://www.mongodb.org/static/pgp/server-4.4.asc | apt-key add -
RUN echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/4.4 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-4.4.list
RUN apt-get update
RUN apt-get install -y mongodb-org

RUN mongod


RUN go get .
RUN go build

EXPOSE $PORT

ENTRYPOINT [ "./pokemonwiki" ]