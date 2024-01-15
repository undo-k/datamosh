package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undo-k/datamosh/internal/moshers"
)

var decimateThreshold float32

var decimateCmd = &cobra.Command{
	Use:   "decimate",
	Short: "just wrecks a percentage of the bytes with nonsense (very dangerous)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		moshers.Decimate(&data, startIndex, endIndex, decimateThreshold)
	},
}

func init() {
	decimateCmd.Flags().Float32VarP(&decimateThreshold, "threshold", "t", 10000, "lower the threshold the more damage")

	rootCmd.AddCommand(decimateCmd)
}
