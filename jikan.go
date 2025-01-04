package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiResponse struct {
	Data AnimeData `json:"data"`
}

func fetchAnime() (*ApiResponse, error) {
	url := fmt.Sprintf("https://api.jikan.moe/v4/random/anime")

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error Code")
	}

	apiResponse := &ApiResponse{}
	if err := json.NewDecoder(resp.Body).Decode(apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}
