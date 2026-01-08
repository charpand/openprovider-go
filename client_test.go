package openprovider

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("Default configuration", func(t *testing.T) {
		client := NewClient(Config{})

		if client.baseURL != DefaultBaseURL {
			t.Errorf("Expected baseURL %s, got %s", DefaultBaseURL, client.baseURL)
		}
		if client.httpClient == nil {
			t.Error("Expected default httpClient to be initialized, got nil")
		}
	})

	t.Run("Custom BaseURL", func(t *testing.T) {
		customURL := "http://localhost:4010"
		client := NewClient(Config{
			BaseURL: customURL,
		})

		if client.baseURL != customURL {
			t.Errorf("Expected baseURL %s, got %s", customURL, client.baseURL)
		}
	})
}
