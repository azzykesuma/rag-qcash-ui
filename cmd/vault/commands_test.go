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
