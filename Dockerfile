FROM golang:1.10.2
# FROM golang:1.10.2-alpine

WORKDIR /go/src/github.com/aliceblock/sample

COPY . /go/src/github.com/aliceblock/sample

# RUN go build -o app ./cmd/server

RUN go get -u github.com/codegangsta/gin
# RUN apk update && apk add --no-cache git \
#   && go get -u github.com/codegangsta/gin \
#   && apk del git \
#   && rm -rf /var/cache/apk/*

ENV ENVIRONMENT debug
ENV HOST 0.0.0.0
ENV PORT 8080

EXPOSE 8080

# CMD ["./app"]
CMD ["gin", "--path", "./cmd/server", "--port", "8080", "main.go"]