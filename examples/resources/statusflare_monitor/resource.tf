resource "statusflare_monitor" "example" {
  name = "Example monitor"
  url  = "www.example.com"
  
  integrations = [
    data.statusflare_integration.slack.id
  ]
}
