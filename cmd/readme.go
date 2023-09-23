package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// readmeCmd represents the readme command
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "asks for a repo and returns the readme file",
	Long: `This command will ask for a repo and return the readme file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Project name")
	},
}

func init() {
	rootCmd.AddCommand(readmeCmd)
}
