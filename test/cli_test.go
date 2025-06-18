package test

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIListModuleSource_JSON(t *testing.T) {
	binary := "../cp-tf-inspect"
	if _, err := exec.LookPath(binary); err != nil {
		t.Skip("cp-tf-inspect binary not found; build it before running integration tests")
	}

	workspace := filepath.Join("fixtures", "basic")
	cmd := exec.Command(binary, "list-module-source", "--dir", workspace)
	outBytes, err := cmd.CombinedOutput()
	out := string(outBytes)
	if err != nil {
		t.Fatalf("error running CLI: %s\noutput: %s", err, out)
	}

	want := []string{
		"../modules/storage",
		"git::https://example.com/bar.git",
		"terraform-aws-modules/vpc/aws",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Errorf("Output missing expected module source: %s\nfull output:\n%s", w, out)
		}
	}
}
