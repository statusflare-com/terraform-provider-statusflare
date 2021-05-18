package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare-com/terraform-provider-statusflare/statusflare"
)

func resourceCustomDomain() *schema.Resource {

	fields := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"domain": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The full domain of the custom domain. Can be either example.statusflare.app or your custom domain.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Status (active/pending) of the domain.",
		},
	}

	return &schema.Resource{
		CreateContext: resourceCustomDomainCreate,
		ReadContext:   resourceCustomDomainRead,
		DeleteContext: resourceCustomDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: fields,
	}
}

func resourceCustomDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var client *statusflare.Client = meta.(*statusflare.Client)

	customDomain, err := client.GetCustomDomain(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	customDomainToData(customDomain, d)
	return diags
}

func resourceCustomDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var client *statusflare.Client = meta.(*statusflare.Client)

	customDomain := statusflare.CustomDomain{}
	dataToCustomDomain(d, &customDomain)

	err := client.CreateCustomDomain(&customDomain)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(customDomain.Id)

	resourceCustomDomainRead(ctx, d, meta)
	return diags
}

func resourceCustomDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		err    error
		diags  diag.Diagnostics
		client *statusflare.Client = m.(*statusflare.Client)
	)

	err = client.DeleteCustomDomain(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func dataToCustomDomain(src *schema.ResourceData, dst *statusflare.CustomDomain) {
	dst.Domain = src.Get("domain").(string)
	dst.Type = "full_domain"
}

func customDomainToData(src *statusflare.CustomDomain, dst *schema.ResourceData) {
	dst.Set("status", src.Status)
	dst.Set("domain", src.Domain)
}
