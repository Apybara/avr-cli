FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/avr-cli

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN make build && go mod download

# Deployment container
FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/avr-cli/avr-cli /avr-cli
ENTRYPOINT ["/avr-cli"]
CMD []
