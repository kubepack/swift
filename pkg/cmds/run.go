package cmds

import (
	_ "net/http/pprof"

	"github.com/appscode/go/hold"
	_ "github.com/appscode/wheel/pkg/release"
	apiCmd "github.com/appscode/wheel/pkg/server/cmd"
	"github.com/appscode/wheel/pkg/server/cmd/options"
	"github.com/spf13/cobra"
	"github.com/appscode/wheel/pkg/analytics"
)

func NewCmdRun(version string) *cobra.Command {
	opt := options.NewConfig()
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run wheel apis",
		PreRun: func(cmd *cobra.Command, args []string) {
			if opt.EnableAnalytics {
				analytics.Enable()
			}
			analytics.SendEvent("wheel", "started", version)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			analytics.SendEvent("wheel", "stopped", version)
		},
		Run: func(cmd *cobra.Command, args []string) {
			apiCmd.Run(opt)
			hold.Hold()
		},
	}

	opt.AddFlags(cmd.Flags())
	return cmd
}
