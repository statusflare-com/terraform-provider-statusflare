resource "statusflare_custom_domain" "example" {
  domain = "example.statusflare.app"

  lifecycle {
    create_before_destroy = true
  }
}

resource "statusflare_status_page" "example" {
  name     = "My example status page"
  custom_domain = [statusflare_custom_domain.example.domain]
}
