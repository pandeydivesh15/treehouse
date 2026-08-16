package main

import (
	"os"
	"runtime/debug"

	"github.com/kunchenguid/treehouse/cmd"
)

var version = ""

func init() {
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
			version = info.Main.Version
		} else {
			version = "dev"
		}
	}
}

func main() {
	// Upstream handles a --update-check flag here, the entry point for the
	// background update child process. Nothing spawns it in this build.

	cmd.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
