package cmd

import (
	"github.com/spf13/cobra"
)

// chooseRegionCmd represents the chooseRegion command
var chooseRegionCmd = &cobra.Command{
	Use:   "region-set",
	Short: "Choose your region",

	Run: func(cmd *cobra.Command, args []string) {
		setRegion(regionPrompt)

	},
}

func init() {
	rootCmd.AddCommand(chooseRegionCmd)

}
