##
## Build
##
FROM golang:1.18-alpine3.15 AS build

WORKDIR /app

RUN apk update && apk add git
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /app/go-src

##
## Deploy
##
FROM alpine:3.15

WORKDIR /app

COPY --from=build /app/go-src /app

EXPOSE 3000

CMD ["/app/go-src"]