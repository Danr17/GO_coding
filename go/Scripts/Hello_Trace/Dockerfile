FROM golang as builder 
WORKDIR /home

COPY go.* ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o hello ./cmd

FROM alpine:latest
RUN apk update && apk --no-cache --update add ca-certificates

COPY --from=builder /home/hello /hello
RUN chmod a+x /hello

ENV PORT 8080

CMD ["/hello"]
