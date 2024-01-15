package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undo-k/datamosh/internal/moshers"
)

var quadraticCmd = &cobra.Command{
	Use:   "quadratic",
	Short: "byte swapping using quadratic function",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		moshers.Quadratic(&data, startIndex, endIndex)
	},
}

func init() {
	rootCmd.AddCommand(quadraticCmd)
}
