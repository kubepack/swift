package cmds

import (
	_ "net/http/pprof"

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
