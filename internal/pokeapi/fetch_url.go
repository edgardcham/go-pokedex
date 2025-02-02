package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/edgardcham/go-pokedex/internal/pokecache"
	"io"
	"net/http"
)

func FetchURL(url string, target interface{}, cache *pokecache.Cache) error {
	if data, found := cache.Get(url); found {
		return json.Unmarshal(data, target)
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error getting response: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Error: %v", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response: %v", err)
	}

	cache.Add(url, body)
	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}
	return nil
}
