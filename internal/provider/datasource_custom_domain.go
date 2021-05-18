package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare-com/terraform-provider-statusflare/statusflare"
)

func dataSourceCustomDomain() *schema.Resource {
	fields := map[string]*schema.Schema{
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the custom domain.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Status (active/pending) of the domain.",
		},
	}

	return &schema.Resource{
		ReadContext: dataSourceCustomDomainRead,
		Schema:      fields,
	}
}

func dataSourceCustomDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *statusflare.Client = meta.(*statusflare.Client)

	customDomains, err := client.AllCustomDomains()
	if err != nil {
		return diag.FromErr(err)
	}

	// find the custom domain by given domain
	domain := d.Get("domain").(string)
	for _, customDomain := range customDomains {
		if customDomain.Domain == domain {
			d.Set("id", customDomain.Id)
			d.SetId(customDomain.Id)
			return diag.Diagnostics{}
		}
	}

	return diag.Errorf("no custom domain for domain %s", domain)
}
