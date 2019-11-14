package calc

// import (
// 	"fmt"

// 	"github.com/spf13/cobra"
// 	"github.com/tehcyx/go-cli-workshop/internal/geometry"
// )

// func NewCircleCmd(o *Options) *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "circle [OPTIONS]",
// 		Short: "Calculate some numbers",
// 		Long:  `Calculate some basic things based on internal geometry implementation`,
// 		RunE:  func(_ *cobra.Command, _ []string) error { return o.Run() },
// 	}

// 	cmd.Flags().Float64VarP(&o.CircleRadius, "radius", "r", 0.0, "--radius 1.37")

// 	return cmd
// }

// func (o *Options) Run() error {
// 	if o.CircleRadius == 0.0 {
// 		return fmt.Errorf("Circle radius cannot be zero")
// 	}
// 	circle := geometry.NewCircle(o.CircleRadius)
// 	geometry.Measure(circle)
// 	return nil
// }
