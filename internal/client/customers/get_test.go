// Package customers_test contains tests for the customers package.
package customers_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/charpand/terraform-provider-openprovider/internal/client"
	"github.com/charpand/terraform-provider-openprovider/internal/client/customers"
	"github.com/charpand/terraform-provider-openprovider/internal/testutils"
)

func TestGetCustomer(t *testing.T) {
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

	customer, err := customers.Get(apiClient, "XX123456-XX")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if customer == nil {
		t.Log("Note: No customer returned by mock server")
		return
	}

	if customer.Handle == "" {
		t.Log("Note: Customer handle not populated by mock server")
	}
}
