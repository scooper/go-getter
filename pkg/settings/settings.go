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

// TODO: handle errors / log
// TODO: better filepath code
func Get() Settings {
	path, _ := filepath.Abs("./settings.json")
	settingsJson, _ := os.Open(path)
	defer settingsJson.Close()
	var settings Settings
	settingsAsBytes, _ := ioutil.ReadAll(settingsJson)
	json.Unmarshal(settingsAsBytes, &settings)
	return settings
}