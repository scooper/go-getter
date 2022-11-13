package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Settings struct {
	Debug bool   `json:"debug"`
	Port  string `json:"port"`
}

// TODO: better filepath code
func Get() (*Settings, error) {
	path, perr := filepath.Abs("./settings.json")
	if perr != nil {
		return nil, perr
	}
	settingsJson, ferr := os.Open(path)
	if ferr != nil {
		return nil, ferr
	}

	defer settingsJson.Close()
	var settings *Settings
	settingsAsBytes, _ := ioutil.ReadAll(settingsJson)
	jsonErr := json.Unmarshal(settingsAsBytes, &settings)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return settings, nil
}