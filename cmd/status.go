package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A high-level overview of portfolio performance",
	Long: `Calculate some basic statistics and charts to provide a high-level
overview of portfolio performance. Takes as input a portfolio spec or a
portfolio history.

The portfolio spec or portfolio history can be input as a file or as command
line arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

  // Inputs needed:
  // - portfolio
  //   - parse a properly-formatted file
  //   - parse a series of flags that contain the info needed for an asset
  //     - Provide a utility to parse a file that may not be in the correct format

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
