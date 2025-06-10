FROM golang:1.22-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the MCP server
RUN go build -o /app/mcp cmd/mcp/main.go

# Set the entrypoint to run the MCP server
ENTRYPOINT ["/app/mcp"] 