package calc

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tehcyx/go-cli-workshop/internal/geometry"
)

const (
	radius = 1.2
)

func NewArbCmd(o *Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "arb [OPTIONS]",
		Short: "Calculate some numbers",
		Long:  `Calculate some basic things based on internal geometry implementation`,
		RunE:  func(_ *cobra.Command, _ []string) error { return o.Run() },
	}

	cmd.Flags().Float64VarP(&o.CircleRadius, "radius", "r", 0.0, "--radius 1.37")
	cmd.Flags().Float64VarP(&o.RectWidth, "height", "x", 0.0, "--height 1.37")
	cmd.Flags().Float64VarP(&o.RectHeight, "width", "w", 0.0, "--width 1.37")

	return cmd
}

func (o *Options) Run() error {
	var g geometry.Geometry
	if o.CircleRadius > 0.0 {
		g = geometry.Circle{Radius: o.CircleRadius}
	} else if o.RectWidth > 0.0 && o.RectHeight > 0.0 {
		g = geometry.Rect{Width: o.RectWidth, Height: o.RectHeight}
	} else {
		return fmt.Errorf("can't initiate any shape")
	}
	geometry.Measure(g)
	return nil
}
