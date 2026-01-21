package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Services struct {
}

func New() *Services {
	return &Services{}
}

func (s *Services) DogRandom(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://random.dog/woof.json", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type randomDogResponse struct {
		URL string `json:"url"`
	}

	var dogData randomDogResponse
	err = json.Unmarshal(body, &dogData)
	if err != nil {
		return "", err
	}

	return dogData.URL, nil
}
