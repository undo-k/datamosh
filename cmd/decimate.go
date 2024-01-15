package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undo-k/datamosh/internal/moshers"
)

var decimateCmd = &cobra.Command{
	Use:   "decimate",
	Short: "just wrecks 0.1% of the bytes with nonsense (very dangerous)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		moshers.Decimate(data, startIndex, endIndex)
	},
}

func init() {
	rootCmd.AddCommand(decimateCmd)
}
