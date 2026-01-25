resource "openprovider_customer" "owner" {
  email = "owner@example.com"

  phone {
    country_code = "1"
    area_code    = "555"
    number       = "1234567"
  }

  address {
    street  = "Main Street"
    number  = "123"
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
