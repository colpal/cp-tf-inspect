package cmd

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func TestUniqueModuleSources(t *testing.T) {
	module := &tfconfig.Module{
		ModuleCalls: map[string]*tfconfig.ModuleCall{
			"foo": {Source: "./modules/foo"},
			"bar": {Source: "git::https://example.com/bar.git"},
			"baz": {Source: "./modules/foo"},
		},
	}

	got := UniqueModuleSources(module)
	want := []string{"./modules/foo", "git::https://example.com/bar.git"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UniqueModuleSources() = %v, want %v", got, want)
	}
}
