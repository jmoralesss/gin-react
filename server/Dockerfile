FROM golang:alpine

WORKDIR /server

RUN apk add build-base

RUN mkdir /server/.cert

COPY /.cert ./.cert
COPY / ./

RUN go mod download
RUN go mod verify
RUN go build -o array-server

CMD ./array-server

# docker build -t array-server -f Dockerfile.server .
# docker run -it -p 443:443 array-server 