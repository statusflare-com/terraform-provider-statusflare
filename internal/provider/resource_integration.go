package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare-com/terraform-provider-statusflare/statusflare"
)

func resourceIntegration() *schema.Resource {

	fields := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the integration",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "webhook",
			Description: "Type of the integration, e.g. webhook",
		},
		"secret": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Sensitive:   true,
			Description: "The secret of the integration, e.g. webhook URL",
		},
	}

	return &schema.Resource{
		CreateContext: resourceIntegrationCreate,
		ReadContext:   resourceIntegrationRead,
		UpdateContext: resourceIntegrationUpdate,
		DeleteContext: resourceIntegrationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: fields,
	}
}

func resourceIntegrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var client *statusflare.Client = meta.(*statusflare.Client)

	integration, err := client.GetIntegration(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	integrationToData(integration, d)
	return diags
}

func resourceIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var client *statusflare.Client = meta.(*statusflare.Client)

	integration := statusflare.Integration{}
	dataToIntegration(d, &integration)

	err := client.CreateIntegration(&integration)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(integration.Id)
	d.Set("secret", integration.Secret)

	resourceIntegrationRead(ctx, d, meta)
	return diags
}

func resourceIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *statusflare.Client = meta.(*statusflare.Client)

	integration, err := client.GetIntegration(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	dataToIntegration(d, integration)
	err = client.SaveIntegration(integration)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIntegrationRead(ctx, d, meta)
}

func resourceIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		err    error
		diags  diag.Diagnostics
		client *statusflare.Client = m.(*statusflare.Client)
	)

	err = client.DeleteIntegration(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func dataToIntegration(src *schema.ResourceData, dst *statusflare.Integration) {
	dst.Name = src.Get("name").(string)
	dst.Type = src.Get("type").(string)
	dst.Secret = src.Get("secret").(string)
}

func integrationToData(src *statusflare.Integration, dst *schema.ResourceData) {
	dst.Set("name", src.Name)
	dst.Set("type", src.Type)
}
