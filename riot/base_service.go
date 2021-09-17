package riot

import (
	"encoding/json"

	"github.com/yucealiosman/riot-cli/pkg"

	"github.com/spf13/viper"
)

func GetService(client *pkg.Client, queryParams map[string]string, urlPath string, object interface{}) error {
	headers := map[string]string{"X-Riot-Token": viper.GetString("riotToken"), "Content-Type": "application/json"}
	request, err := pkg.NewGetRequest(headers, queryParams, urlPath)
	if err != nil {
		return err
	}
	response, err := client.Send(request)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(response.Body), object)

	if err != nil {
		return err
	}
	return err
}
