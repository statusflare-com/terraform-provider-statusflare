terraform {
  required_providers {
    statusflare = {
      version = "~> 0.1"
      source  = "statusflare-com/statusflare"
    }
  }
}

provider "statusflare" {
  account_id = "..."
  key_id     = "..."
  token      = "***"
}
