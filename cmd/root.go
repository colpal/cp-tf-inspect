package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cp-tf-inspect",
	Short: "A CLI tool to inspect Terraform workspaces",
	Long:  `cp-tf-inspect provides commands to query Terraform workspaces (e.g., listing module sources).`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listModuleSourceCmd)
}
