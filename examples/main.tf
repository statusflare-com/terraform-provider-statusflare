terraform {
  required_providers {
    statusflare = {
      version = "~> 1.0"
      source  = "statusflare.com/statusflare/statusflare"
    }
  }
}

provider "statusflare" {
}

data "statusflare_integration" "slack" {
  name = "some-slack-integration"
}

resource "statusflare_monitor" "first" {
  name = "hello-world"
  url  = "www.helloworld.com"
  integrations = [
    data.statusflare_integration.slack.id
  ]
}
