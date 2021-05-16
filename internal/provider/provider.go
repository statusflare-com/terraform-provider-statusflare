package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare/terraform-provider-statusflare/statusflare"
)

// This is provider's 'main' entry point.
func New(version string) *schema.Provider {

	configFields := map[string]*schema.Schema{
		"account_id": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("STATUSFLARE_ACCOUNT_ID", nil),
			Description: "Your Statusflare Account ID. This can also be specified with the `STATUSFLARE_ACCOUNT_ID` env. variable.",
		},
		"key_id": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("STATUSFLARE_KEY_ID", nil),
			Description: "Your token's key ID. This can also be specified with the `STATUSFLARE_KEY_ID` env. variable.",
		},
		"token": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("STATUSFLARE_TOKEN", nil),
			Description: "Token's secret part. This can also be specified with the `STATUSFLARE_TOKEN` env. variable.",
		},
	}

	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"statusflare_monitor": resourceMonitor(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"statusflare_integration": dataSourceIntegration(),
		},
		ConfigureContextFunc: configure,
		Schema:               configFields,
	}
}

// this function initialize and configure the Statusflare client
func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accountId := d.Get("account_id").(string)
	keyId := d.Get("key_id").(string)
	token := d.Get("token").(string)

	client := statusflare.NewClient(accountId, keyId, token)
	return client, diag.Diagnostics{}
}
