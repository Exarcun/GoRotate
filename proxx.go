package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func getPublicIP(client *http.Client) (string, error) {
	resp, err := client.Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return result["query"].(string), nil
}

func main() {
	proxyURL, err := url.Parse("http://199.233.238.6:2452")
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}

	for {
		ip, err := getPublicIP(client)
		if err != nil {
			fmt.Println("Error getting public IP:", err)
		} else {
			fmt.Println("Public IP:", ip)
		}

		time.Sleep(1 * time.Second) // Check every 30 seconds
	}
}
