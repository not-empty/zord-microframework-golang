package main

import (
	"fmt"
	"go-skeleton/kernel"
)

func main() {
	k := kernel.NewKernel()
	k.Boot()
	err := k.Start()
	if err != nil {
		fmt.Println("Unable to Boot")
		return
	}
}
