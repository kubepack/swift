//go:generate go-extpoints pkg/extpoints
package main

import (
	"os"

	logs "github.com/appscode/log/golog"
	"github.com/appscode/swift/pkg/cmds"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewRootCmd(Version).Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
