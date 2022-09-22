FROM golang:1.19-alpine as builder

WORKDIR /go/src/bre

COPY . .

RUN go get -d -v ./...

RUN go build -o /app ./cmd/.

FROM alpine:latest

COPY --from=builder /app /app

WORKDIR /

ENTRYPOINT ["/app"]