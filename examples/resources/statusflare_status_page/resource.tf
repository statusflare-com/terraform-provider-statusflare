resource "statusflare_monitor" "example" {
  name = "Example monitor"
  url  = "www.example.com"
}

resource "statusflare_status_page" "example" {
  name     = "My example status page"
  monitors = [statusflare_monitor.example.id]

  // following config options are defaults you can override
  config = {
    title                         = "Status Page",
    logo_url                      = "statusflare.png",
    favicon_url                   = "favicon.ico",
    all_monitors_operational      = "All Monitors Operational",
    not_all_monitors_operational  = "Not All Monitors Operational",
    monitor_operational_label     = "Operational",
    monitor_not_operational_label = "Not Operational",
    monitor_no_data_label         = "No data",
    histogram_no_data             = "No data",
    histogram_no_incidents        = "All good",
    histogram_some_incidents      = "incident(s)",
  }
}