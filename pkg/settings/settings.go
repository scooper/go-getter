package settings

import (
	"encoding/json"
	"io/ioutil"

	"github.com/scooper/go-getter/pkg/utils"
)

type Settings struct {
	Debug bool   `json:"debug"`
	Port  string `json:"port"`
}

// TODO: better filepath code
func Get() (*Settings, error) {

	settingsFile, fileErr := utils.GetSetting("settings.json")
	if fileErr != nil {
		return nil, fileErr
	}
	defer settingsFile.Close()

	var settings *Settings
	settingsAsBytes, _ := ioutil.ReadAll(settingsFile)
	jsonErr := json.Unmarshal(settingsAsBytes, &settings)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return settings, nil
}