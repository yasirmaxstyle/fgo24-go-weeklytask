package api

import (
	"encoding/json"
	"fmt"
	"go-cli/models"
	"io"
	"net/http"
)

const MenuURL = "https://raw.githubusercontent.com/yasirmaxstyle/fgo24-node-datasource/refs/heads/master/menulist.json"

type APIMenuService struct{}

func NewAPIMenuService() *APIMenuService {
	return &APIMenuService{}
}

func (a *APIMenuService) FetchMenu() (*models.MenuData, error) {
	resp, err := http.Get(MenuURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch menu: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var menuData models.MenuData
	err = json.Unmarshal(body, &menuData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	fmt.Println("App loaded successfully!")
	return &menuData, nil
}
