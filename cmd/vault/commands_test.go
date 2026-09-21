package main

import (
	"flag"
	"reflect"
	"testing"
)

func TestInterspersedOptions(t *testing.T) {
	for _, args := range [][]string{{"session expiration", "--project", "example", "--limit", "5"}, {"--project=example", "session expiration", "--limit=5"}} {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		project := fs.String("project", "", "")
		limit := fs.Int("limit", 10, "")
		if err := parseInterspersed(fs, args); err != nil {
			t.Fatal(err)
		}
		if *project != "example" || *limit != 5 || !reflect.DeepEqual(fs.Args(), []string{"session expiration"}) {
			t.Fatalf("bad options %s %d %v", *project, *limit, fs.Args())
		}
	}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	full := fs.Bool("full", false, "")
	if err := parseInterspersed(fs, []string{"source path", "--full"}); err != nil || !*full || fs.Arg(0) != "source path" {
		t.Fatal("boolean scan flag failed")
	}
	if err := parseInterspersed(fs, []string{"--unknown"}); err == nil {
		t.Fatal("unknown flag silently ignored")
	}
}

func TestResolveScanLimits(t *testing.T) {
	tests := []struct {
		name                   string
		full                   bool
		limit, perTool         int
		limitSet, perToolSet   bool
		wantLimit, wantPerTool int
		wantError              bool
	}{
		{name: "default global batch", limit: 10, wantLimit: 10},
		{name: "per tool batch", limit: 10, perTool: 10, perToolSet: true, wantPerTool: 10},
		{name: "explicit zero per tool keeps default", limit: 10, perToolSet: true, wantLimit: 10},
		{name: "unlimited", limit: 0, limitSet: true},
		{name: "full defaults unlimited", full: true, limit: 10},
		{name: "full bounded rejected", full: true, limit: 10, limitSet: true, wantError: true},
		{name: "two limit modes rejected", limit: 10, perTool: 10, limitSet: true, perToolSet: true, wantError: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limit, perTool, err := resolveScanLimits(tt.full, tt.limit, tt.perTool, tt.limitSet, tt.perToolSet)
			if (err != nil) != tt.wantError || limit != tt.wantLimit || perTool != tt.wantPerTool {
				t.Fatalf("got (%d, %d, %v), want (%d, %d, error=%v)", limit, perTool, err, tt.wantLimit, tt.wantPerTool, tt.wantError)
			}
		})
	}
}
