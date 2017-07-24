package cmds

import (
	_ "net/http/pprof"

	"github.com/appscode/go/hold"
	apiCmd "github.com/appscode/wheel/pkg/apiserver/cmd"
	"github.com/appscode/wheel/pkg/apiserver/cmd/options"
	_ "github.com/appscode/wheel/pkg/app"
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
