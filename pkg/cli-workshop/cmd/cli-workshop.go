package cmd

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/cmd/calc"
	"github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/core"
)

const (
	sleep = 10 * time.Second
)

//NewWorkshopCmd creates a new workshop command
func NewWorkshopCmd(o *core.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workshop",
		Short: "Enables git repository browsing",
		Long:  `This is a CLI built in a workshop to demonstrate how to built CLIs: https://github.com/tehcyx/go-cli-workshop`,
	}

	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "v", false, "verbose output")

	versionCmd := NewVersionCmd(NewVersionOptions(o))
	cmd.AddCommand(versionCmd)

	calcCmd := calc.NewCalcCmd(calc.NewCalcOptions(o))
	cmd.AddCommand(calcCmd)

	// circleCmd := calc.NewCircleCmd(calc.NewCalcOptions(o))
	// calcCmd.AddCommand(circleCmd)

	arbCmd := calc.NewArbCmd(calc.NewCalcOptions(o))
	calcCmd.AddCommand(arbCmd)

	return cmd
}
