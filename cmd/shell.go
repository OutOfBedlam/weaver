package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func doShell(cmd *cobra.Command, args []string) {
	fmt.Println("doShell()")
}

func doShellInfo(cmd *cobra.Command, args []string) {
	fmt.Println("doShellInfo()")
}

func doShellTables(cmd *cobra.Command, args []string) {
	fmt.Println("doShellTable()")
}
