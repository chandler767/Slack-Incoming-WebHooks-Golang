package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	Webhook_URL = "" // Get a Webhook Url and learn more: https://api.slack.com/incoming-webhooks
	client := &http.Client{Timeout: (time.Second * 30)}
	payload := strings.NewReader("{\"channel\":\"general\",\"username\":\"testing-bot\",\"text\":\"This was a test\",\"icon_emoji\":\"ghost\"}")
	req, err := http.NewRequest("POST", Webhook_URL, payload)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(resp.Status) // 200 OK = successful post
}
