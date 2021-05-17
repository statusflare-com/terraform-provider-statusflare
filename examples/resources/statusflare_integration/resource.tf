resource "statusflare_integration" "example" {
  name   = "example-integration-name"
  type   = "webhook"
  secret = "https://webhook.example.com"
}
