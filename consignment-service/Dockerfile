# consignment-service/Dockerfile


FROM golang:1.9.0 as builder

WORKDIR /go/src/services/consignment-service

# Copy the current code into our workdir
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

# Create a dep project, and run `ensure`, which will pull in all of the dependencies
# within this directory.
RUN dep init && dep ensure

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine. 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# Start a new build process

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/services/consignment-service .

CMD ["./consignment-service"]