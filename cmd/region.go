package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RegionCmd represents the region command
var RegionCmd = &cobra.Command{
	Use:   "region",
	Short: "Riot region",

	Run: func(cmd *cobra.Command, args []string) {
		showResult(cmd, viper.GetString("region"))
	},
}

func init() {
	rootCmd.AddCommand(RegionCmd)

}
