package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare-com/terraform-provider-statusflare/statusflare"
)

// This is provider's 'main' entry point.
func New(version string) *schema.Provider {

	configFields := map[string]*schema.Schema{
		"api_url": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SF_API_URL", "https://api.statusflare.com"),
			Description: "Statusflare API URL.",
		},
		"account_id": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SF_ACCOUNT_ID", nil),
			Description: "Your Statusflare Account ID. This can also be specified with the `SF_ACCOUNT_ID` env. variable.",
		},
		"key_id": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SF_KEY_ID", nil),
			Description: "Your token's key ID. This can also be specified with the `SF_KEY_ID` env. variable.",
		},
		"token": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SF_TOKEN", nil),
			Description: "Token's secret part. This can also be specified with the `SF_TOKEN` env. variable.",
		},
	}

	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"statusflare_monitor":     resourceMonitor(),
			"statusflare_integration": resourceIntegration(),
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
	apiUrl := d.Get("api_url").(string)
	accountId := d.Get("account_id").(string)
	keyId := d.Get("key_id").(string)
	token := d.Get("token").(string)

	client := statusflare.NewClient(apiUrl, accountId, keyId, token)
	return client, diag.Diagnostics{}
}
