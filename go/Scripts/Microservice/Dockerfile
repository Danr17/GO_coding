FROM golang:1.12.4-alpine3.9 AS build-env

# Allow Go to retrive the dependencies for the build step
RUN apk add --no-cache git

RUN mkdir /gopherconuk/

WORKDIR /gopherconuk/
ADD . /gopherconuk/

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o /gopherconuk/gcuk .

# final stage
FROM alpine:3.8

WORKDIR /
COPY --from=build-env /gopherconuk/certs/docker.localhost.* /
COPY --from=build-env /gopherconuk/gcuk /

EXPOSE 8080

CMD ["/gcuk"]