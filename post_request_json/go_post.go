package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	//var data = strings.NewReader(`{"content": "From Go. Wooo!"}`)
	var data = strings.NewReader(`{"content": "", "embeds": [{"title": "Hello!"}]}`)

    var webhook_url = "some_url_discord_webhook"
	req, err := http.NewRequest("POST", webhook_url, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
