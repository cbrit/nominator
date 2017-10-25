package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
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

// SendMessage sends POST requests to a specified channel
func SendMessage(channel string, text string) {
	config := LoadConfig()
	data := url.Values{}
	data.Set("token", config.Secrets.APIToken)
	data.Add("channel", channel)
	data.Add("text", text)
	data.Add("username", config.Username)

	u, err := url.ParseRequestURI(config.URL)
	if err != nil {
		fmt.Println("Error while parsing request URI")
	}
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		// handle
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(req)
	if err != nil {
		// handle
	}
	fmt.Println(res.Status)
}
