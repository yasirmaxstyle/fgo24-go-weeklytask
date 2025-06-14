package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const MenuURL = "https://raw.githubusercontent.com/yasirmaxstyle/fgo24-node-datasource/refs/heads/master/menulist.json"

func FetchMenuFromAPI() (*MenuData, error) {
	resp, err := http.Get(MenuURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch menu: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var menuData MenuData
	err = json.Unmarshal(body, &menuData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	fmt.Println("App loaded successfully!")
	return &menuData, nil
}
