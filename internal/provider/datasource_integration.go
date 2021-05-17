package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare-com/terraform-provider-statusflare/statusflare"
)

func dataSourceIntegration() *schema.Resource {
	fields := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the integration.",
		},
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the integration.",
		},
	}

	return &schema.Resource{
		ReadContext: dataSourceIntegrationRead,
		Schema:      fields,
	}
}

func dataSourceIntegrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *statusflare.Client = meta.(*statusflare.Client)

	integrations, err := client.AllIntegrations()
	if err != nil {
		return diag.FromErr(err)
	}

	// find the integration by given name
	name := d.Get("name").(string)
	for _, integration := range integrations {
		if integration.Name == name {
			d.Set("id", integration.Id)
			d.SetId(integration.Id)
			return diag.Diagnostics{}
		}
	}

	return diag.Errorf("no integration for name %s", name)
}
