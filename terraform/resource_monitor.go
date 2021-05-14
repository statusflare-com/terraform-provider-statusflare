package terraform

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare/terraform-provider-statusflare/statusflare"
)

func resourceMonitor() *schema.Resource {

	fields := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the monitor. Must be unique",
		},
		"url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "URL Address but  without schema. It might be www.example.com",
		},
		"scheme": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "https",
			Description: "The scheme might be http or https. The default value is https.",
		},
		"method": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "GET",
			Description: "The HTTP method. The default is 'GET'",
		},
		"expect_status": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     200,
			Description: "The expected HTTP status code. The default is 200.",
		},
		"retries": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "Retries or also 'notify_after' field in API. The default is 1.",
		},
		"worker": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "managed",
			Description: "Don't know purpose of this field but default value is 'managed'",
		},
		"integrations": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "IDs of integrations attached to this monitor.",
		},
	}

	return &schema.Resource{
		CreateContext: resourceMonitorCreate,
		ReadContext:   resourceMonitorRead,
		UpdateContext: resourceMonitorUpdate,
		DeleteContext: resourceMonitorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: fields,
	}
}

func resourceMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var client *statusflare.Client = meta.(*statusflare.Client)

	monitor, err := client.GetMonitor(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	monitorToData(monitor, d)
	return diags
}

func resourceMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var client *statusflare.Client = meta.(*statusflare.Client)

	monitor := statusflare.Monitor{}
	dataToMonitor(d, &monitor)

	err := client.CreateMonitor(&monitor)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(monitor.ID)
	resourceMonitorRead(ctx, d, meta)
	return diags
}

func resourceMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *statusflare.Client = meta.(*statusflare.Client)

	monitor, err := client.GetMonitor(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	dataToMonitor(d, monitor)
	err = client.SaveMonitor(monitor)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceMonitorRead(ctx, d, meta)
}

func resourceMonitorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		err    error
		diags  diag.Diagnostics
		client *statusflare.Client = m.(*statusflare.Client)
	)

	err = client.DeleteMonitor(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func dataToMonitor(src *schema.ResourceData, dst *statusflare.Monitor) {
	dst.Name = src.Get("name").(string)
	dst.URL = src.Get("url").(string)
	dst.Scheme = src.Get("scheme").(string)
	dst.Method = src.Get("method").(string)
	dst.ExpectStatus = src.Get("expect_status").(int)
	dst.NotifyAfter = src.Get("retries").(int)
	dst.Worker = src.Get("worker").(string)
	dst.Integrations = toStrArray(src.Get("integrations").([]interface{}))
}

func monitorToData(src *statusflare.Monitor, dst *schema.ResourceData) {
	dst.Set("name", src.Name)
	dst.Set("url", src.URL)
	dst.Set("scheme", src.Scheme)
	dst.Set("method", src.Method)
	dst.Set("expect_status", src.ExpectStatus)
	dst.Set("worker", src.Worker)
	dst.Set("retries", src.NotifyAfter)
	dst.Set("integrations", src.Integrations)
}
