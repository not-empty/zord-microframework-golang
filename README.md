# Zord Microframework
## Build your mecha

GOlang base repository with code gen to create a fast golang project based on hexagonal architeture

---

# Development
> Remember to create your .env file based on .env.example

### 1. Using Docker Compose
Up mysql and zord project:

``` SHELL
docker compose up
```

<br />

#### 2. Using raw go build

You will need to build the http/main.go file:

``` SHELL
go build -o server cmd/http/main.go
```

Then run the server

``` SHELL
./server
```

<br />

#### 3. Running from go file

``` SHELL
go run cmd/http/main.go
```

<br />

**Note:** To run the local build as described in the second or third option, a MySQL server must be running. This is necessary for the application to interact with its database. The easiest way to set up a MySQL server locally is by using Docker. Below is a command to start a MySQL server container using Docker:

``` SHELL
docker compose up mysql -d
```
This command will ensure that a MySQL server is running in the background, allowing you to execute the local build successfully.

---

### Cli

#### build cli

to build cli into binary file run
``` SHELL
go build -o cli cmd/cli/main.go
```

then you can run all cli commands with the binary file
``` SHELL
./cli -h
```

if you`re developing something in the cli the best way is run it directly to all changes 
``` SHELL
go run cmd/cli/main.go
```

---

#### Cli Commands

create new domain (crud):
``` SHELL
./cli create-domain {{domain}}
```

destroy domain:
``` SHELL
./cli destroy-domain {{domain}}
```

migrate all domains:
``` SHELL
./cli migrate
```

**Obs:** If you`re generating code inside docker container you need to change generated folder and file permissions to code out of docker container.

run the follow command to edit generated files:
``` SHELL
sudo chown $USER:$USER -R .
```

if you have a group name different from username change the command accordingly

---

#### Run tests
Run all tests:
``` SHELL
go test ./...
```

Verify code coverage:
``` SHELL
// Generate coverage output
go test ./... -coverprofile=coverage.out

// Generate HTML file
go tool cover -html=coverage.out
```

### Docs (WIP):
https://github.com/not-empty/zord-microframework-golang/wiki

### Development

Want to contribute? Great!

The project using a simple code.
Make a change in your file and be careful with your updates!
**Any new code will only be accepted with all validations.**


**Not Empty Foundation - Free codes, full minds**

---

### MCP Server

The MCP (Model-Controller-Provider) server provides a set of tools for code generation and database management through Claude Desktop integration.

#### Building the MCP Server

To build the MCP server into a binary file:

``` SHELL
go build -o mcp cmd/mcp/main.go
```

#### Configuring Claude Desktop

To use the MCP server with Claude Desktop, you need to configure it in your Claude Desktop settings. Add the following configuration to your Claude Desktop tools:

```json
{
  "name": "zord-mcp",
  "description": "Zord Microframework MCP Server",
  "command": "path/to/project/mcp",
  "transport": "stdio"
}
```

Replace `path/to/project/mcp` with the actual path to your MCP server binary.

#### Example: Running MCP Server via Docker

To build the MCP Docker image, run:

```sh
docker build -t zord-mcp -f docker/mcp/Dockerfile .
```

To use the MCP server with Claude Desktop via Docker, add the following configuration to your Claude Desktop tools:

```json
{
  "mcpServers": {
    "zord-mcp": {
      "name": "zord-mcp-docker",
      "description": "Zord MCP Server via Docker",
      "command": "docker run -i --rm --env-file /your/project/.env -v /your/project:/app zord-mcp",
      "transport": "stdio",
      "tools": []
    }
  }
}
```

Note: if you have another mcp server add it inside the mcpServers key with another name, like this: 

```json
{
  "mcpServers": {
    "zord-mcp": {
      "name": "zord-mcp-docker",
      "description": "Zord MCP Server via Docker",
      "command": "docker run -i --rm --env-file /your/project/.env -v /your/project:/app zord-mcp",
      "transport": "stdio",
      "tools": []
    },
    "zord-mcp2": {
      "name": "zord-mcp-docker",
      "description": "Zord MCP Server via Docker",
      "command": "docker run -i --rm --env-file /your/project/.env -v /your/project:/app zord-mcp",
      "transport": "stdio",
      "tools": []
    },
  }
}
```

This will run the MCP server inside a Docker container and connect it to Claude Desktop using stdio transport.

#### Available Tools

The MCP server provides the following tools:

1. **Create Domain**
   - Tool: `create-domain`
   - Description: Creates a new domain service
   - Arguments:
     - `name`: Name of the domain to create (required)
     - `validator`: Whether to create domain with validation (optional)
     - `domainType`: Type of domain to create (default: crud)

2. **Destroy Domain**
   - Tool: `destroy-domain`
   - Description: Destroys a domain service
   - Arguments:
     - `name`: Name of the domain to destroy (required)
     - `domainType`: Type of domain to destroy (default: crud)

3. **Create Domain from Schema**
   - Tool: `create-domain-from-schema`
   - Description: Creates a domain from database schema
   - Arguments:
     - `schemaName`: Name of the schema to use (required)
     - `tableName`: Optional specific table to generate

4. **Migrate**
   - Tool: `migrate`
   - Description: Migrates database schema
   - Arguments:
     - `tenant`: Optional tenant name for migration

5. **Inspect**
   - Tool: `inspect`
   - Description: Inspects database schema
   - No arguments required

6. **Generate Schema from DB**
   - Tool: `generate-schema-from-db`
   - Description: Generates schema from database
   - Arguments:
     - `schemaName`: Name for the generated schema (required)
     - `databaseName`: Optional specific database to generate from

#### Environment Variables

The MCP server requires the following environment variables to be set:

```env
ENVIRONMENT=development
APP=zord
VERSION=1.0.0
DB_USER=your_db_user
DB_PASS=your_db_password
DB_URL=localhost
DB_PORT=3306
DB_DATABASE=your_database
DB_TEST_DATABASE=your_test_database
```

Make sure to set these variables in your `.env` file before running the MCP server.

To test the MCP server directly, you can run:
```sh
docker run -i zord-mcp
```