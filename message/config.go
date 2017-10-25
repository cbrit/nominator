package message

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Config contains the unmarshalled config.json
type Config struct {
	Username string `json:"username"`
	URL      string `json:"url"`
	Secrets  struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		VToken       string `json:"verification_token"`
		APIToken     string `json:"api_token"`
	} `json:"secrets"`
}

// LoadConfig unmarshals config.json into a Config
func LoadConfig() Config {
	filepath, _ := filepath.Abs("./config.json")
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	var c Config
	json.Unmarshal(raw, &c)
	return c
}
