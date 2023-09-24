package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func doVersion(cmd *cobra.Command, args []string) {
	fmt.Printf(_logoBanner, _version)
}
