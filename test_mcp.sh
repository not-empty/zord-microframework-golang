#!/bin/bash

# Build the MCP server
docker compose build mcp

# Start the MCP server in detached mode
docker compose up -d mcp

# Wait for the server to start
sleep 2

# Test the server with a valid JSON-RPC request
echo '{"jsonrpc":"2.0","method":"inspect","params":{},"id":1}' | docker compose exec -T mcp mcp

# Test another request
echo '{"jsonrpc":"2.0","method":"migrate","params":{"tenant":"test"},"id":2}' | docker compose exec -T mcp mcp

# Keep the server running for manual testing
echo "MCP server is running. Press Ctrl+C to stop."
docker compose logs -f mcp 