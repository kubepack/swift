package cmds

import (
	_ "net/http/pprof"

	"github.com/appscode/go/hold"
	"github.com/appscode/swift/pkg/extpoints"
	"github.com/appscode/swift/pkg/factory"
	_ "github.com/appscode/swift/pkg/release"
	apiCmd "github.com/appscode/swift/pkg/server/cmd"
	"github.com/appscode/swift/pkg/server/cmd/options"
	"github.com/spf13/cobra"
)

func NewCmdRun(version string) *cobra.Command {
	opt := options.New()
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Run swift apis",
		DisableAutoGenTag: true,
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
