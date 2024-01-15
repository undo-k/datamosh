package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/undo-k/datamosh/internal/fileio"
	"github.com/undo-k/datamosh/internal/headers"
	"log"
)

var inputName string
var outputName string
var data []byte
var fileType string
var startIndex uint
var endIndex uint

var rootCmd = &cobra.Command{
	Use:  "datamosh",
	Args: cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		if inputName == "None" {
			log.Fatal(errors.New("no input file specified"))
		}
		data = fileio.ReadData(&inputName)
		fileType = headers.GetFileType(data)
		startIndex, endIndex = headers.GetBounds(data, fileType)
		setOutputName()
	},
	PersistentPostRun: func(cmd *cobra.Command, _ []string) {
		fileio.WriteData(&data, &outputName)
		fmt.Println("Successfully wrote to", outputName)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputName, "input", "i", "None", "input file")
	rootCmd.PersistentFlags().StringVarP(&outputName, "output", "o", "default", "output file (optional)")
	rootCmd.MarkFlagRequired("input")
}

func setOutputName() {
	if outputName == "default" {
		outputName = "output" + fileType
	}
}
