data "statusflare_custom_domain" "example" {
  domain = "mydomain.example.com"
}

resource "statusflare_status_page" "example" {
  name           = "Example status page"
  custom_domain  = statusflare_custom_domain.example.domain
}
