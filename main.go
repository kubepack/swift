//go:generate go-extpoints pkg/extpoints
package main

import (
	"os"

	"kmodules.xyz/client-go/logs"
	"kubepack.dev/swift/pkg/cmds"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
