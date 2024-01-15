package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/undo-k/datamosh/internal/fileio"
	"github.com/undo-k/datamosh/internal/headers"
	"github.com/undo-k/datamosh/internal/moshers"
	"log"
)

var blendFilename string
var blendData []byte
var blendFileType string
var blendStart uint
var blendEnd uint
var blendThreshold float32
var blendPercent int

var blendCmd = &cobra.Command{
	Use:   "blend",
	Short: "plays swappies with the bytes between two input files, pretty crazy stuff",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if blendFilename == "None" {
			log.Fatal(errors.New("no blend file specified"))
		}
		blendData = fileio.ReadData(&blendFilename)
		blendFileType = headers.GetFileType(blendData)
		blendStart, blendEnd = headers.GetBounds(blendData, blendFileType)
		moshers.Blender(&data, &blendData, startIndex, endIndex, blendStart, blendEnd, blendThreshold, blendPercent)
	},
}

func init() {
	blendCmd.Flags().StringVarP(&blendFilename, "blender", "b", "None", "name of the file you want to collide with main input")
	blendCmd.Flags().Float32VarP(&blendThreshold, "threshold", "t", 10000, "lower the threshold the more damage")
	blendCmd.Flags().IntVarP(&blendPercent, "percentage", "p", 25, "0 - 100 size of the image slice you want blended")
	rootCmd.AddCommand(blendCmd)
}
