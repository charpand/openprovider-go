resource "openprovider_customer" "company_owner" {
  company_name = "Example Corp"
  email        = "contact@example.com"
  locale       = "en_US"
  comments     = "Primary domain owner for company domains"

  phone {
    country_code = "1"
    area_code    = "555"
    number       = "9876543"
  }

  address {
    street  = "Tech Boulevard"
    number  = "456"
    suffix  = "Suite 100"
    city    = "San Francisco"
    state   = "CA"
    zipcode = "94105"
    country = "US"
  }

  name {
    first_name = "Jane"
    last_name  = "Smith"
    prefix     = "Ms."
  }
}

# Use the customer handle in a domain resource
resource "openprovider_domain" "company_domain" {
  domain       = "example.com"
  owner_handle = openprovider_customer.company_owner.handle
  period       = 1
}
