//go:generate go-extpoints pkg/extpoints
package main

import (
	"os"

	"github.com/appscode/swift/pkg/cmds"
	"kmodules.xyz/client-go/logs"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
