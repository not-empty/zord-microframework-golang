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