package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/spf13/cobra"
)

var dir string
var recursive bool

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

func collectModuleSources(rootDir, currentDir string, module *tfconfig.Module, acc map[string]struct{}, visited map[string]struct{}) {
	for _, call := range module.ModuleCalls {
		src := call.Source
		if strings.HasPrefix(src, "./") || strings.HasPrefix(src, "../") {
			absPath := filepath.Join(currentDir, src)
			if rel, err := filepath.Rel(rootDir, absPath); err == nil {
				src = filepath.ToSlash(rel)
			}
			childDir := filepath.Join(currentDir, call.Source)
			childDir = filepath.Clean(childDir)
			if _, seen := visited[childDir]; !seen {
				visited[childDir] = struct{}{}
				childModule, diag := tfconfig.LoadModule(childDir)
				if diag.Err() == nil {
					collectModuleSources(rootDir, childDir, childModule, acc, visited)
				}
			}
		}
		acc[src] = struct{}{}
	}
}

func RecursiveModuleSources(rootDir string, module *tfconfig.Module) []string {
	acc := make(map[string]struct{})
	visited := make(map[string]struct{})
	rootDir = filepath.Clean(rootDir)
	visited[rootDir] = struct{}{}
	collectModuleSources(rootDir, rootDir, module, acc, visited)
	sourcesList := make([]string, 0, len(acc))
	for src := range acc {
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
		var sourcesList []string
		if recursive {
			sourcesList = RecursiveModuleSources(dir, module)
		} else {
			sourcesList = UniqueModuleSources(module)
		}
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
	listModuleSourceCmd.Flags().BoolVar(&recursive, "recursive", false, "Recursively follow local module calls and list all underlying module sources")
	listModuleSourceCmd.MarkFlagRequired("dir")
}
