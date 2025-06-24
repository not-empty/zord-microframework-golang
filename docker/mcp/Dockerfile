FROM golang:1.22-alpine AS builder

WORKDIR /build

# Install required packages
RUN apk add --no-cache \
    git \
    curl \
    bash \
    sudo

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the MCP server
RUN go build -o /build/mcp cmd/mcp/main.go

# Final stage
FROM golang:1.22-alpine

WORKDIR /app

# Install required packages
RUN apk add --no-cache \
    git \
    curl \
    bash \
    sudo

# Set up Atlas
ENV ATLAS_VERSION=v0.29.0
RUN curl -sSf https://atlasgo.sh | sh

# Copy the binary from builder
COPY --from=builder /build/mcp /usr/local/bin/mcp

# Create data directory for mounted volumes
RUN mkdir -p /app/data

# Set the entrypoint
ENTRYPOINT ["/usr/local/bin/mcp"] 