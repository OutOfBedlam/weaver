package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var (
	dataFlag string
)

func init() {
	rootCmd = &cobra.Command{
		Use:   "weaver",
		Short: "Weaver is a data processing pipeline server",
		Long: `A flexible and elegant way to build data middleware.
which provides a simple yet powerful script based pipe definition.`,
	}

	versionCmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver"},
		Short:   "Show version",
		Run:     doVersion,
	}
	rootCmd.AddCommand(versionCmd)

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start weaver server",
		Long:  `Start weaver server process`,
		Run:   doServe,
	}
	serveCmd.PersistentFlags().StringVarP(&dataFlag, "data", "d", "", "path to data directory")
	rootCmd.AddCommand(serveCmd)

	shellCmd := &cobra.Command{
		Use:   "shell",
		Short: "Weaver shell interface",
		Long:  `Start weaver shell command line interface`,
		Run:   doShell,
	}
	rootCmd.AddCommand(shellCmd)

	shellShowCmd := &cobra.Command{
		Use:   "show",
		Short: "show informations",
	}
	shellCmd.AddCommand(shellShowCmd)

	shellShowInfoCmd := &cobra.Command{
		Use:   "info",
		Short: "show informations",
		Run:   doShellInfo,
	}
	shellShowCmd.AddCommand(shellShowInfoCmd)

	shellShowTablesCmd := &cobra.Command{
		Use:   "tables",
		Short: "show tables",
		Run:   doShellTables,
	}
	shellShowCmd.AddCommand(shellShowTablesCmd)
}
