# 0. Fetch ca-certificates
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# 1. build executable
FROM golang:1.13-alpine as builder
WORKDIR /go/src/app

# 2. install build dependencies
RUN apk add git
RUN apk add protoc
RUN go get github.com/golang/protobuf/protoc-gen-go

# 3. generate proto files
COPY *.proto .
RUN mkdir gen
RUN protoc --go_out=plugins=grpc:./gen *.proto

# 4. install dependencies
COPY go.mod .
COPY go.sum .
RUN go get -d -v ./...
RUN go install -v ./...

# 5. Build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o /tmp/main

# 5. Start process from scratch as appuser
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /tmp/main /bin/main

ENTRYPOINT [ "/bin/main" ]