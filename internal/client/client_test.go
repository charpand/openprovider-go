package client

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockAuthTransport struct {
	loginCalled       bool
	return400OnNoAuth bool
}

func (m *mockAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.HasSuffix(req.URL.Path, "/v1beta/auth/login") {
		m.loginCalled = true
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"code": 0, "data": {"token": "new-token"}}`)),
			Header:     make(http.Header),
		}, nil
	}

	auth := req.Header.Get("Authorization")
	if auth == "" {
		status := http.StatusUnauthorized
		body := "Unauthorized"
		if m.return400OnNoAuth {
			status = http.StatusBadRequest
			body = "Bad Request"
		}
		return &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	}

	if auth == "Bearer expired-token" {
		return &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       io.NopCloser(strings.NewReader(`Unauthorized`)),
			Header:     make(http.Header),
		}, nil
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"status": "ok"}`)),
		Header:     make(http.Header),
	}, nil
}

func TestNewClient(t *testing.T) {
	t.Run("Default configuration", func(t *testing.T) {
		client := NewClient(Config{})

		if client.BaseURL != DefaultBaseURL {
			t.Errorf("Expected BaseURL %s, got %s", DefaultBaseURL, client.BaseURL)
		}
		if client.HTTPClient == nil {
			t.Error("Expected default HTTPClient to be initialized, got nil")
		}
	})

	t.Run("Custom BaseURL", func(t *testing.T) {
		customURL := "http://localhost:4010"
		client := NewClient(Config{
			BaseURL: customURL,
		})

		if client.BaseURL != customURL {
			t.Errorf("Expected BaseURL %s, got %s", customURL, client.BaseURL)
		}
	})

	t.Run("Automatic login when token is missing", func(t *testing.T) {
		transport := &mockAuthTransport{return400OnNoAuth: true}
		hc := &http.Client{Transport: transport}
		client := NewClient(Config{
			Username:   "testuser",
			Password:   "testpass",
			HTTPClient: hc,
		})

		req, _ := http.NewRequest("GET", "http://example.com/test", nil)
		resp, err := client.Do(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %d", resp.StatusCode)
		}
		if !transport.loginCalled {
			t.Error("Expected Login to be called, but it wasn't")
		}
		if client.Token != "new-token" {
			t.Errorf("Expected token to be updated to 'new-token', got '%s'", client.Token)
		}
	})

	t.Run("Token expiration and retry", func(t *testing.T) {
		transport := &mockAuthTransport{}
		hc := &http.Client{Transport: transport}
		client := NewClient(Config{
			Username:   "testuser",
			Password:   "testpass",
			Token:      "expired-token",
			HTTPClient: hc,
		})

		req, _ := http.NewRequest("GET", "http://example.com/test", nil)
		resp, err := client.Do(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %d", resp.StatusCode)
		}
		if !transport.loginCalled {
			t.Error("Expected Login to be called on 401, but it wasn't")
		}
		if client.Token != "new-token" {
			t.Errorf("Expected token to be updated to 'new-token', got '%s'", client.Token)
		}
	})
}
