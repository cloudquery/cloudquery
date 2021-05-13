############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
# Build the binary.
RUN go build -o /go/bin/cloudquery
############################
# STEP 2 build a small image
############################
FROM public.ecr.aws/lambda/provided:al2 as build
# Copy our static executable.
COPY --from=builder /go/bin/cloudquery ./cloudquery
# Run the hello binary.
ENTRYPOINT ["./cloudquery"]