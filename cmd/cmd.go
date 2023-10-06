package cmd

import (
	"github.com/spf13/cobra"
	"go-skeleton/cmd/handlers/http"
)

type Command interface {
	Boot(*cobra.Command, []string)
	Start(*cobra.Command, []string)
}

var commandList = map[string]Command{
	"http": http.NewServer(),
}
