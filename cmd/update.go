package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update treehouse to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Upstream downloads a release binary over the running one. This
		// build is compiled from this checkout, so applying a release would
		// silently discard whatever is patched here.
		return fmt.Errorf(
			"self-update is disabled in this build\n" +
				"treehouse is built from source; pull the checkout and " +
				"rebuild instead")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
