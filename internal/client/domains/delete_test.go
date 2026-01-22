// Package domains_test contains tests for the domains package.
package domains_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/charpand/terraform-provider-openprovider/internal/client"
	"github.com/charpand/terraform-provider-openprovider/internal/client/domains"
	"github.com/charpand/terraform-provider-openprovider/internal/testutils"
)

func TestDeleteDomain(t *testing.T) {
	baseURL := os.Getenv("TEST_API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:4010"
	}

	httpClient := &http.Client{
		Transport: &testutils.MockTransport{RT: http.DefaultTransport},
	}

	config := client.Config{
		BaseURL:    baseURL,
		Username:   "test",
		Password:   "test",
		HTTPClient: httpClient,
	}
	apiClient := client.NewClient(config)

	// Delete a test domain
	err := domains.Delete(apiClient, 123)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
