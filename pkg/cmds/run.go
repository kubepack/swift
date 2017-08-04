package cmds

import (
	_ "net/http/pprof"

	"github.com/appscode/go/hold"
	_ "github.com/appscode/wheel/pkg/app"
	apiCmd "github.com/appscode/wheel/pkg/server/cmd"
	"github.com/appscode/wheel/pkg/server/cmd/options"
	"github.com/spf13/cobra"
)

func NewCmdRun() *cobra.Command {
	config := options.NewConfig()
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run wheel apis",
		Run: func(cmd *cobra.Command, args []string) {
			apiCmd.Run(config)
			hold.Hold()
		},
	}

	config.AddFlags(cmd.Flags())
	return cmd
}
