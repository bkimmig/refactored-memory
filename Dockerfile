FROM golang:1.16.2-alpine3.13 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

FROM golang:1.16.2-alpine3.13 AS app

# Export necessary port
EXPOSE 8080

COPY --from=builder /dist/main /

# Command to run when starting the container -> CMD ["/dist/main"] (add this to
# your entrypoint of the deployment
