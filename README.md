# Zord Microframework
## Build your mecha

GOlang base repository with code gen to create a fast golang project based on hexagonal architeture

### Installation

Remember to create your .env file based on .env.example

#### Using Docker
Up mysql and zord project:

``` SHELL
docker compose up
```

#### Using raw go build

``` SHELL
go build cmd/http/server.go
```

to run local build you need a mysql server running, the easiest way is using docker

``` SHELL
docker compose up mysql -d
```

Then run the server

``` SHELL
./server
```

#### Running from go file

to run local build you need a mysql server running, the easiest way is using docker
``` SHELL
docker compose up mysql -d
```

``` SHELL
go run cmd/http/server.go
```

### Cli

#### build cli

to build cli into binary file run
``` SHELL
go build cmd/cli/cli.go
```

then you can run all cli commands with the binary file
``` SHELL
./cli -h
```

if you`re developing something in the cli the best way is run it directly to all changes 
``` SHELL
go run cmd/cli/cli.go
```

#### Cli usage

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

#### Cli usage inside docker image

Enter in zord image:
``` SHELL
docker exec -it zord-http sh
```

Build cli binary:
``` SHELL
go build cmd/cli/cli.go
```

Use it:
``` SHELL
./cli -h
```

If you`re generating code inside docker container you need to change generated folder and file permissions to code out of docker container.

run the follow command to edit generated files:
``` SHELL
sudo chown $USER:$USER -R .
```

if you have a group name different from username change the command accordingly

### Development

Want to contribute? Great!

The project using a simple code.
Make a change in your file and be careful with your updates!
**Any new code will only be accepted with all validations.**


**Not Empty Foundation - Free codes, full minds**