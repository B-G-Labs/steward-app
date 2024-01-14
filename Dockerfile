FROM golang:latest

WORKDIR /usr/src/api

RUN go install github.com/mitranim/gow@latest

COPY . .

RUN cd ./apps/api && go mod tidy