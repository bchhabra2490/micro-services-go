# consignment-service/Dockerfile
FROM golang:1.9.0 

WORKDIR /go/src/services/consignment-service/

# Copy the current code into our workdir
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN rm -rf vendor
# Create a dep project, and run `ensure`, which will pull in all of the dependencies
# within this directory.
RUN dep init && dep ensure
RUN go get -u google.golang.org/grpc && go get -u github.com/golang/protobuf/protoc-gen-go
# Build the binary, with a few flags which will allow
# us to run this binary in Alpine. 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo main.go repository.go handler.go datastore.go

# Start a new build process

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/services/consignment-service/consignment-service .

CMD ["./consignment-service"]