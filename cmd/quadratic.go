package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undo-k/datamosh/internal/moshers"
)

var quadThreshold float32

var quadraticCmd = &cobra.Command{
	Use:   "quadratic",
	Short: "byte swapping using quadratic function",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		moshers.Quadratic(&data, startIndex, endIndex, quadThreshold)
	},
}

func init() {
	quadraticCmd.Flags().Float32VarP(&quadThreshold, "threshold", "t", 10000, "lower the threshold the more damage")

	rootCmd.AddCommand(quadraticCmd)
}
