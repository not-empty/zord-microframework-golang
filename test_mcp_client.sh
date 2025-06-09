#!/bin/bash

# Testa o MCP server com uma requisição ListOfferings via stdio

# Criar um pipe nomeado
PIPE="/tmp/mcp_pipe"
rm -f "$PIPE"
mkfifo "$PIPE"

# Iniciar o servidor em background
./mcp < "$PIPE" > "$PIPE" &
SERVER_PID=$!

# Aguardar um momento para o servidor iniciar
sleep 1

# Enviar a requisição
REQ='{"jsonrpc":"2.0","id":1,"method":"ListOfferings","params":{}}'
echo "Enviando requisição ListOfferings para o MCP server..."
echo "$REQ" > "$PIPE"

# Aguardar a resposta
cat "$PIPE"

# Limpar
kill $SERVER_PID
rm -f "$PIPE"

# Se o servidor não estiver rodando, inicie-o em outro terminal com:
# ./mcp 