data "statusflare_integration" "example" {
  name = "example-integration-name"
}

resource "statusflare_monitor" "example" {
  name = "Example monitor"
  url  = "www.example.com"
  
  integrations = [
    data.statusflare_integration.example.id
  ]
}
