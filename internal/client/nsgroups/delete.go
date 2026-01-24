// Package nsgroups provides functionality for working with nameserver groups.
package nsgroups

import (
	"fmt"
	"net/http"

	"github.com/charpand/terraform-provider-openprovider/internal/client"
)

// Delete deletes a nameserver group by name via the Openprovider API.
//
// Endpoint: DELETE https://api.openprovider.eu/v1beta/dns/nameservers/groups/{ns_group}
func Delete(c *client.Client, name string) error {
	path := fmt.Sprintf("/v1beta/dns/nameservers/groups/%s", name)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", c.BaseURL, path), nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to delete nameserver group: status code %d", resp.StatusCode)
	}

	return nil
}
