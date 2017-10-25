package message

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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
