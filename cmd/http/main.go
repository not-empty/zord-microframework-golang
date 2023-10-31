package main

import (
	"go-skeleton/cmd/http/server"
	"go-skeleton/pkg"
)

func main() {
	serverInstance := server.NewServer()
	serverInstance.Start()
}

func init() {
	pkg.Mysql.Connect()
	pkg.Logger.Boot()
	pkg.Validator.Boot()
}
