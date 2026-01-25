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

func TestCreateCustomer(t *testing.T) {
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

	req := &customers.CreateCustomerRequest{
		Email: "test@example.com",
		Phone: customers.Phone{
			CountryCode: "1",
			AreaCode:    "555",
			Number:      "1234567",
		},
		Address: customers.Address{
			Street:  "Main St",
			Number:  "123",
			City:    "New York",
			Country: "US",
			Zipcode: "10001",
		},
		Name: customers.Name{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	handle, err := customers.Create(apiClient, req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if handle == "" {
		t.Log("Note: No handle returned by mock server")
	}
}
