terraform {
  required_providers {
    statusflare = {
      version = "~> 1.0"
      source  = "statusflare.com/statusflare/statusflare"
    }
  }
}

provider "statusflare" {
  account_id = "..."
  key_id     = "..."
  token      = "***"
}


