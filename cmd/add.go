package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/undo-k/datamosh/internal/fileio"
	"github.com/undo-k/datamosh/internal/headers"
	"github.com/undo-k/datamosh/internal/moshers"
	"log"
)

var addFilename string
var addData []byte
var addFileType string
var addStart uint
var addEnd uint

// var addThreshold float32
var addPercent int

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "plays swappies with the bytes between two input files, pretty crazy stuff",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if addFilename == "None" {
			log.Fatal(errors.New("no add file specified"))
		}
		addData = fileio.ReadData(&addFilename)
		addFileType = headers.GetFileType(addData)
		addStart, addEnd = headers.GetBounds(addData, addFileType)
		moshers.Appender(&data, &addData, startIndex, endIndex, addStart, addEnd, addPercent)
	},
}

func init() {
	addCmd.Flags().StringVarP(&addFilename, "adder", "a", "None", "name of the file you want to collide with main input")
	//addCmd.Flags().Float32VarP(&addThreshold, "threshold", "t", 10000, "lower the threshold the more damage")
	addCmd.Flags().IntVarP(&addPercent, "percentage", "p", 25, "0 - 100 size of the image slice you want added")
	rootCmd.AddCommand(addCmd)
}
