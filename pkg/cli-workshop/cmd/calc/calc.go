package calc

import "github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/core"

import "github.com/spf13/cobra"

type Options struct {
	*core.Options

	CircleRadius float64
	RectWidth    float64
	RectHeight   float64
}

func NewCalcOptions(o *core.Options) *Options {
	return &Options{Options: o}
}

func NewCalcCmd(o *Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "calc",
		Short: "Calculate some numbers",
		Long:  `Calculate some basic things based on internal geometry implementation`,
	}

	return cmd
}
