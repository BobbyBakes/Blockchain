package cmd

import (
	"fmt"

	"github.com/BobbyBakes/CopyrightedPhrases/app/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new copyright phrase",
	Long:  `Create a new copyright phrase`,
	Run:   createPhrase,
}

func createPhrase(cmd *cobra.Command, args []string) {
	ok := service.MsgService.CreatePhrase()
	if ok {
		fmt.Println("Copyright Created")
	} else {
		fmt.Println("Copyright Creation Failed")
	}

}
