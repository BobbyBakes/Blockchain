package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new copyright phrase",
	Long:  `Create a new copyright phrase`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Copyright Created")
	},
}
