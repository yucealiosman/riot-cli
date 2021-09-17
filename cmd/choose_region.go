package cmd

import (
	"github.com/spf13/cobra"
)

// chooseRegionCmd represents the chooseRegion command
var chooseRegionCmd = &cobra.Command{
	Use:   "choose-region",
	Short: "Choose your region",

	Run: func(cmd *cobra.Command, args []string) {
		setRegion()

	},
}

func init() {
	rootCmd.AddCommand(chooseRegionCmd)

}
