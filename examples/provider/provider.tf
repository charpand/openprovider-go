terraform {
  required_providers {
    openprovider = {
      source  = "charpand/openprovider"
      version = ">= 0.1.0"
    }
  }
}

provider "openprovider" {
  username = var.openprovider_username
  password = var.openprovider_password
}
