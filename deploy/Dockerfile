############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
# Build the binary.
RUN export CGO_ENABLED=0 && go build -o /go/bin/cloudquery
############################
# STEP 2 build a small image
############################
FROM lambci/lambda:go1.x as build
# Copy our static executable.
COPY --from=builder /go/bin/cloudquery ./cloudquery
# Run the cloudquery binary.
ENTRYPOINT ["/var/runtime/aws-lambda-go", "./cloudquery"]