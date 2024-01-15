package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undo-k/datamosh/internal/moshers"
)

var displace int

var shiftCmd = &cobra.Command{
	Use:  "shift",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		moshers.ShiftSubslice(data, startIndex, endIndex, displace)
	},
}

func init() {
	shiftCmd.Flags().IntVarP(&displace, "displace", "d", 210000, "number of places to shift bytes")
	rootCmd.AddCommand(shiftCmd)
}
