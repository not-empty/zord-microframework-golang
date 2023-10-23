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
go build .
```

to run local build you need a mysql server running, the easiest way is using docker

``` SHELL
docker compose up mysql -d
```

Then run the server

``` SHELL
./go-skeleton http
```

### Usage

#### Create new crud

helpers (All sub commands has their own helper using -h flag):
``` SHELL
./go-skeleton -h
```

create new domain (crud):
``` SHELL
./go-skeleton create-domain {{domain}}
```

destroy domain:
``` SHELL
./go-skeleton destroy-domain {{domain}}
```

migrate all domains:
``` SHELL
./go-skeleton migrate
```

### Development

Want to contribute? Great!

The project using a simple code.
Make a change in your file and be careful with your updates!
**Any new code will only be accepted with all validations.**


**Not Empty Foundation - Free codes, full minds**