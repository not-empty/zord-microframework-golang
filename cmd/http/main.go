package main

import (
	"go-skeleton/cmd"
	"go-skeleton/cmd/http/server"

	_ "github.com/go-sql-driver/mysql"
	"go-skeleton/pkg/registry"
)

var (
	reg *registry.Registry
)

func main() {
	cmd.Setup()
	serverInstance := server.NewServer(cmd.Reg, cmd.ApiPrefix)
	serverInstance.Start()
}
