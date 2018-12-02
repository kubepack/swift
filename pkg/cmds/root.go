package cmds

import (
	"flag"
	"log"

	"github.com/appscode/go/signals"
	v "github.com/appscode/go/version"
	"github.com/appscode/kutil/tools/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "swift [command]",
		Short:             `Swift by Appscode - Ajax friendly Helm Tiller Proxy`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.Flags().VisitAll(func(flag *pflag.Flag) {
				log.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
			})
			cli.SendAnalytics(c, v.Version.Version)
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})
	rootCmd.PersistentFlags().BoolVar(&cli.EnableAnalytics, "enable-analytics", cli.EnableAnalytics, "Send analytical events to Google Analytics")

	stopCh := signals.SetupSignalHandler()
	rootCmd.AddCommand(NewCmdRun(stopCh))
	rootCmd.AddCommand(v.NewCmdVersion())

	return rootCmd
}
