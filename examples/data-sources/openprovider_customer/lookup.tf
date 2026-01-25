data "openprovider_customer" "existing" {
  handle = "XX123456-XX"
}

# Use an existing customer handle in a domain
resource "openprovider_domain" "example" {
  domain       = "example.com"
  owner_handle = data.openprovider_customer.existing.handle
  period       = 1
}
