package cmd

import (
	"fmt"
	"os"

	"github.com/yucealiosman/riot-cli/util"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func setRiotToken() string {
	prompt := promptui.Prompt{
		Label: "Enter Riot API Token Please. You can view more details at https://developer.riotgames.com/docs/portal#_getting-started",
		Templates: &promptui.PromptTemplates{
			Valid:   "{{ . | cyan }} ",
			Success: "{{ . | blue }} ",
		},
	}

	token, err := prompt.Run()

	if err != nil {
		showErrorAndExit(err)

	}

	fmt.Printf("You entered %q\n", token)
	util.WriteLocalConfigFile("riottoken", token)

	return token

}

func showErrorAndExit(err error) {
	color.Red("An error has occured %v\n", err)
	os.Exit(1)
}

func showResult(result string) {
	color.Cyan("%v\n", result)
}
