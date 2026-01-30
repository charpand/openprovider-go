# Domain Transfer Resource

This resource allows you to manage domain transfers in Openprovider.

## Example

```hcl
resource "openprovider_domain_transfer" "example" {
  domain = "example.com"
  auth_code = "example_auth_code"
}
```

## Arguments

- **domain** - The domain name you want to transfer.
- **auth_code** - The authorization code required for the transfer.