package cmds

import (
	"flag"

	"github.com/appscode/go/flags"
	"github.com/appscode/go/signals"
	v "github.com/appscode/go/version"
	"github.com/spf13/cobra"
	"kmodules.xyz/client-go/logs"
	"kmodules.xyz/client-go/tools/cli"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "swift [command]",
		Short:             `Swift by Appscode - Ajax friendly Helm Tiller Proxy`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			flags.DumpAll(c.Flags())
			cli.SendAnalytics(c, v.Version.Version)
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	logs.ParseFlags()
	rootCmd.PersistentFlags().BoolVar(&cli.EnableAnalytics, "enable-analytics", cli.EnableAnalytics, "Send analytical events to Google Analytics")

	stopCh := signals.SetupSignalHandler()
	rootCmd.AddCommand(NewCmdRun(stopCh))
	rootCmd.AddCommand(v.NewCmdVersion())

	return rootCmd
}
