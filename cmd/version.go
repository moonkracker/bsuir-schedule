package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bsuir-schedule",
	Long:  `All software has versions. This is bsuir-schedule's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("bsuir-schedule %s\n", rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
