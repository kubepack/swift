package cmds

import (
	_ "net/http/pprof"

	"github.com/appscode/go/hold"
	"github.com/appscode/swift/pkg/connectors"
	"github.com/appscode/swift/pkg/extpoints"
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
			extpoints.Connectors.Register(&connectors.InClusterConnector{
				TillerCACertFile:     opt.TillerCACertFile,
				TillerClientCertFile: opt.TillerClientCertFile,
				TillerClientKeyFile:  opt.TillerClientKeyFile,
				InsecureSkipVerify:   opt.InsecureSkipVerify,
			}, connectors.UIDInClusterConnector)

			extpoints.Connectors.Register(&connectors.DirectConnector{
				TillerEndpoint:       opt.TillerEndpoint,
				TillerCACertFile:     opt.TillerCACertFile,
				TillerClientCertFile: opt.TillerClientCertFile,
				TillerClientKeyFile:  opt.TillerClientKeyFile,
				InsecureSkipVerify:   opt.InsecureSkipVerify,
			}, connectors.UIDDirectConnector)

			extpoints.Connectors.Register(&connectors.KubeconfigConnector{
				Context:            opt.KubeContext,
				InsecureSkipVerify: opt.InsecureSkipVerify,
			}, connectors.UIDKubeconfigConnector)

			apiCmd.Run(opt)
			hold.Hold()
		},
	}

	opt.AddFlags(cmd.Flags())
	return cmd
}
