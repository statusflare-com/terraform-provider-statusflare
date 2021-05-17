# Statusflare Terraform Provider

## Requirements
 - Terraform 0.14.x
 - Go 1.16 (to build the provider plugin)

# Building The Provider

Ensure you have Go 1.16 and Terraform 0.14.x present in your environment. Clone 
this repository and run the command:

```
make install
```

This command will build and copy provider binary into your `~/terraform.d/plugins` 
folder. After successful installation you can start using provider by:

```
terraform {
  required_providers {
    statusflare = {
      version = "~> 1.0"
      source  = "github.com/statusflare-com/terraform-provider-statusflare"
    }
  }
}

provider "statusflare" {
}
```

