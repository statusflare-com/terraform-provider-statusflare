# Statusflare Terraform Provider

- Website: [statusflare.com](https://statusflare.com)
- Registry: [https://registry.terraform.io](https://registry.terraform.io/providers/statusflare-com/statusflare/latest)
- Join our [Discord server](https://discord.gg/psfJKMCN4v) for help (or just to chat)!

## Requirements

 - [Terraform](https://https://www.terraform.io/downloads.html) 0.14.x
 - [Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)


## Using the provider

```
data "statusflare_integration" "slack" {
  name = "slack-integration-name"
}

resource "statusflare_monitor" "example" {
  name = "Example monitor"
  url  = "www.example.com"

  integrations = [
    data.statusflare_integration.slack.id
  ]
}
```

## Building The Provider

Ensure you have Go 1.16 and Terraform 0.14.x present in your environment. Clone 
this repository and run the command:

```
make install
```

This command will build and copy provider binary into your `~/terraform.d/plugins` 
folder. After successful installation you can start using provider like this:


```
terraform {
  required_providers {
    statusflare = {
      version = "~> 0.4"
      source  = "github.com/statusflare-com/terraform-provider-statusflare"
    }
  }
}

provider "statusflare" {
}
```

