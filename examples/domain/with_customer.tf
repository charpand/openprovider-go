# Create a customer for the domain owner
resource "openprovider_customer" "domain_owner" {
  email = "owner@example.com"

  phone {
    country_code = "1"
    area_code    = "555"
    number       = "1234567"
  }

  address {
    street  = "Main Street"
    number  = "100"
    city    = "New York"
    state   = "NY"
    zipcode = "10001"
    country = "US"
  }

  name {
    first_name = "John"
    last_name  = "Doe"
  }
}

# Register a domain using the customer handle
resource "openprovider_domain" "example" {
  domain       = "example.com"
  owner_handle = openprovider_customer.domain_owner.handle
  period       = 1
}
