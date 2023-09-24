package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var _logoBanner string = `
__        _______    ___     _______ ____  
\ \      / / ____|  / \ \   / / ____|  _ \ 
 \ \ /\ / /|  _|   / _ \ \ / /|  _| | |_) |
  \ V  V / | |___ / ___ \ V / | |___|  _ < 
   \_/\_/  |_____/_/   \_\_/  |_____|_| \_\   %s
`

var _version string = "v0.0.1"

func doServe(cmd *cobra.Command, args []string) {
	fmt.Printf(_logoBanner, _version)
	fmt.Println("execute serve", args)
}
