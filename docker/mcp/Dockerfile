FROM golang:1.22-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the MCP server
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/mcp cmd/mcp/main.go

# Set the entrypoint to run the MCP server
ENTRYPOINT ["/usr/local/bin/mcp"] 