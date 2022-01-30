FROM golang:1.16 AS builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM builder as test
WORKDIR /go/src/app
CMD go test -v ./...

FROM alpine:latest as prod
WORKDIR /root/
COPY --from=builder /go/src/app/app ./
CMD ["./app"]