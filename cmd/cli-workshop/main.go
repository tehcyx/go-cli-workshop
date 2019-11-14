package main

import (
	"fmt"
	"os"

	"github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/cmd"
	"github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/core"
)

func main() {
	command := cmd.NewWorkshopCmd(core.NewOptions())

	err := command.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
