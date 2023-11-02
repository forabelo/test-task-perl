FROM golang:1.20

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

COPY . .

# Build the go app
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main ./cmd/app

EXPOSE 8441

#COPY .docker/common/docker-healthcheck.sh /usr/local/bin/docker-healthcheck
#RUN chmod +x /usr/local/bin/docker-healthcheck
#
#HEALTHCHECK --interval=30s --timeout=3s --retries=3 CMD ["docker-healthcheck"]

# Run the binary programs produced by `go build`
CMD ["/app/bin/main"]
