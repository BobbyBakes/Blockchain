package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "CopyrightedPhrases",
		Short: "Copyrighted Phrases Manager",
		Long:  `Copyrighted Phrases Manager`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
