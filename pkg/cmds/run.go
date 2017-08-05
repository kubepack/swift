package cmds

import (
	_ "net/http/pprof"

	"github.com/appscode/go/hold"
	"github.com/appscode/wheel/pkg/analytics"
	"github.com/appscode/wheel/pkg/extpoints"
	"github.com/appscode/wheel/pkg/factory"
	_ "github.com/appscode/wheel/pkg/release"
	apiCmd "github.com/appscode/wheel/pkg/server/cmd"
	"github.com/appscode/wheel/pkg/server/cmd/options"
	"github.com/spf13/cobra"
)

func NewCmdRun(version string) *cobra.Command {
	opt := options.New()
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
			extpoints.Connectors.Register(&factory.InClusterConnector{}, factory.UIDInClusterConnector)
			extpoints.Connectors.Register(&factory.DirectConnector{TillerEndpoint: opt.TillerEndpoint}, factory.UIDDirectConnector)
			extpoints.Connectors.Register(&factory.KubeconfigConnector{Context: opt.KubeContext}, factory.UIDKubeconfigConnector)

			apiCmd.Run(opt)
			hold.Hold()
		},
	}

	opt.AddFlags(cmd.Flags())
	return cmd
}
