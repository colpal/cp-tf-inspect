package cmd

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/spf13/cobra"
)

var dir string

func UniqueModuleSources(module *tfconfig.Module) []string {
	sources := make(map[string]struct{}, len(module.ModuleCalls))
	for _, mod := range module.ModuleCalls {
		sources[mod.Source] = struct{}{}
	}
	sourcesList := make([]string, 0, len(sources))
	for src := range sources {
		sourcesList = append(sourcesList, src)
	}
	sort.Strings(sourcesList)
	return sourcesList
}

var listModuleSourceCmd = &cobra.Command{
	Use:   "list-module-source",
	Short: "List unique module sources in a Terraform workspace (JSON output only)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if dir == "" {
			return fmt.Errorf("--dir is required")
		}
		module, diag := tfconfig.LoadModule(dir)
		if diag.Err() != nil {
			return fmt.Errorf("failed to load module: %w", diag.Err())
		}
		sourcesList := UniqueModuleSources(module)
		b, err := json.MarshalIndent(sourcesList, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		fmt.Println(string(b))
		return nil
	},
}

func init() {
	listModuleSourceCmd.Flags().StringVar(&dir, "dir", "", "Path to the Terraform workspace directory (required)")
	listModuleSourceCmd.MarkFlagRequired("dir")
}
