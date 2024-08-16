package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go-skeleton/cmd"
	"go-skeleton/cmd/http/server"
	"go-skeleton/pkg/registry"
)

var (
	reg *registry.Registry
)

// @title Swagger Zord API
// @version 1.0
// @description This is the Zord backend server.
func main() {
	cmd.Setup()
	serverInstance := server.NewServer(cmd.Reg, cmd.ApiPrefix)
	serverInstance.Start()
}
