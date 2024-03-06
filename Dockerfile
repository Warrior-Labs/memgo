FROM golang:1.22 as builder
WORKDIR /memgo

# Check if MEMGO_PORT is set, if not set it to 8000
ENV MEMGO_PORT=8000

COPY . .
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o memgo-server

FROM alpine:latest

# Create an App Directory to work from
RUN mkdir -p /app
WORKDIR /app

# Copy the Executable from the Build Stage
COPY --from=builder /memgo/memgo-server .

# Run the Executable
CMD  /app/memgo-server