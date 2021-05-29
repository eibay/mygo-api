FROM golang:alpine AS build

RUN apk add --no-cache curl git alpine-sdk

WORKDIR $GOPATH/src/github.com/eibay/mygo-api

COPY . $GOPATH/src/github.com/eibay/mygo-api/
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM alpine:latest
WORKDIR /main
COPY --from=build /main /main
EXPOSE 8080
ENTRYPOINT [ "./main" ]      