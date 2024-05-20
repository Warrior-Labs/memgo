FROM golang:1.22 as builder
WORKDIR /memgo

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
EXPOSE 8000
CMD  /app/memgo-server