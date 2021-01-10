FROM golang:1-alpine as build
RUN apk update && apk add build-base
WORKDIR /app
ADD . ./
RUN go build -o cloudquery

FROM alpine:latest
WORKDIR /app

COPY --from=build /app/cloudquery ./cloudquery
ENTRYPOINT [ "/app/cloudquery" ]