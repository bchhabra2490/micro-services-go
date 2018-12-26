# vessel-service/Dockerfile
FROM golang:1.9.0 as builder

WORKDIR /go/src/services/vessel-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

# Start a new build

FROM debian:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/services/vessel-service/vessel-service .

CMD ["./vessel-service"]