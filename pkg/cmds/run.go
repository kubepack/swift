package cmds

import (
	_ "net/http/pprof"

	v "github.com/appscode/go/version"
	"github.com/appscode/kutil/tools/cli"
	"github.com/appscode/swift/pkg/cmds/server"
	_ "github.com/appscode/swift/pkg/release"
	"github.com/spf13/cobra"
)

func NewCmdRun(stopCh <-chan struct{}) *cobra.Command {
	o := server.NewSwiftOptions()

	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Run swift apis",
		DisableAutoGenTag: true,
		PreRun: func(c *cobra.Command, args []string) {
			cli.SendPeriodicAnalytics(c, v.Version.Version)
		},
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.AddFlags(flags)

	return cmd
}
