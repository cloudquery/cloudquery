FROM golang:1-alpine as build
RUN apk update && apk add build-base
WORKDIR /app
ADD . ./
RUN go build -o cloudquery

FROM alpine:latest
WORKDIR /app
# TODO: better way to pass in config
COPY config.yml ./app/config.yml
COPY --from=build /app/cloudquery ./cloudquery
ENTRYPOINT [ "/app/cloudquery" ]
