FROM golang:1.22.5-alpine3.19
ARG ENV
WORKDIR /usr/src/
COPY . .
RUN apk --no-cache add npm nodejs bash curl
RUN chmod u+x ./apps/web/start.sh
RUN chmod u+x install.sh && ./install.sh

