package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config represents the format used for config.json.
// If adding user-configurable options, this structure is the place.
type Config struct {
	BaseDomain    string `json:"base_domain"`
	ListenAddress string `json:"listen_address"`
}

// globalConfig is the global config to be used.
var globalConfig Config

// LoadGlobalConfig loads the global config.
func LoadGlobalConfig() {
	configContents, err := os.ReadFile("./config.json")
	check(err)
	err = json.Unmarshal(configContents, &globalConfig)
	check(err)
}

// baseDomain emits "https://<base domain><your url>".
func (c Config) baseDomain(subpath string) string {
	return fmt.Sprintf("https://%s%s", c.BaseDomain, subpath)
}
