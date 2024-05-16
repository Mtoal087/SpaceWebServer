package main

import (
	"encoding/json"
	"net/http"
)

type NasaPhoto struct {
	Copyright   string `json:"copyright"`
	Date        string `json:"date"`
	Explanation string `json:"explanation"`
	Hdurl       string `json:"hdurl"`
	Title       string `json:"title"`
}

func FetchNasaPhoto(apiKey string) (NasaPhoto, error) {
	url := "https://api.nasa.gov/planetary/apod?api_key=" + apiKey
	response, err := http.Get(url)
	if err != nil {
		return NasaPhoto{}, err
	}
	defer response.Body.Close()

	var photo NasaPhoto
	if err := json.NewDecoder(response.Body).Decode(&photo); err != nil {
		return NasaPhoto{}, err
	}

	return photo, nil
}
