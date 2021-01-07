FROM golang:1-alpine as build
RUN apk update && apk add build-base
WORKDIR /app
ADD . ./
RUN go build -o main

FROM alpine:latest
WORKDIR /app
# TODO: better way to pass in config
COPY config.yml ./
COPY --from=build /app/main ./main
ENTRYPOINT [ "/app/main" ]
