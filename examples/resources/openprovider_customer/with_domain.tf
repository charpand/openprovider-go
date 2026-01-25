# Create multiple customers for different contact roles
resource "openprovider_customer" "owner" {
  email = "owner@example.com"

  phone {
    country_code = "1"
    area_code    = "555"
    number       = "1111111"
  }

  address {
    street  = "Main St"
    number  = "100"
    city    = "New York"
    country = "US"
  }

  name {
    first_name = "John"
    last_name  = "Owner"
  }
}

resource "openprovider_customer" "admin" {
  email = "admin@example.com"

  phone {
    country_code = "1"
    area_code    = "555"
    number       = "2222222"
  }

  address {
    street  = "Main St"
    number  = "100"
    city    = "New York"
    country = "US"
  }

  name {
    first_name = "Jane"
    last_name  = "Admin"
  }
}

resource "openprovider_customer" "tech" {
  email = "tech@example.com"

  phone {
    country_code = "1"
    area_code    = "555"
    number       = "3333333"
  }

  address {
    street  = "Main St"
    number  = "100"
    city    = "New York"
    country = "US"
  }

  name {
    first_name = "Bob"
    last_name  = "Tech"
  }
}

# Use all three customer handles in a domain
resource "openprovider_domain" "example" {
  domain         = "example.com"
  owner_handle   = openprovider_customer.owner.handle
  admin_handle   = openprovider_customer.admin.handle
  tech_handle    = openprovider_customer.tech.handle
  billing_handle = openprovider_customer.admin.handle
  period         = 1
  autorenew      = true
}
