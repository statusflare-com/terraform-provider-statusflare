package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare-com/terraform-provider-statusflare/statusflare"
)

func resourceMonitor() *schema.Resource {

	fields := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the monitor.",
		},
		"url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "URL Address but without scheme, e.g. www.example.com.",
		},
		"scheme": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "https",
			Description: "The scheme might be http, https, tcp, icmp. The default value is https.",
		},
		"method": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "GET",
			Description: "The HTTP method. The default is 'GET'.",
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
			Description: "ID of the worker to perform checks from. The default is 'managed'.",
		},
		"integrations": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "IDs of integrations attached to this monitor.",
		},
		"follow_redirects": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"insecure_skip_verify": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"timeout": {
			Type:        schema.TypeInt,
			Default:     30,
			Optional:    true,
			Description: "Timeout in seconds. The default is 30.",
		},
		"interval": {
			Type:        schema.TypeInt,
			Default:     300,
			Optional:    true,
			Description: "Check interval in seconds. The default is 300.",
		},
		"headers": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "HTTP headers for http(s) monitors.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"body": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "HTTP body for http(s) monitors.",
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

	d.SetId(monitor.Id)
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
	dst.FollowRedirects = src.Get("follow_redirects").(bool)
	dst.InsecureSkipVerify = src.Get("insecure_skip_verify").(bool)
	dst.Timeout = src.Get("timeout").(int)
	dst.Interval = src.Get("interval").(int)
	dst.Headers = src.Get("headers").(map[string]interface{})
	dst.Body = src.Get("body").(string)
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
	dst.Set("follow_redirects", src.FollowRedirects)
	dst.Set("insecure_skip_verify", src.InsecureSkipVerify)
	dst.Set("timeout", src.Timeout)
	dst.Set("interval", src.Interval)
	dst.Set("headers", src.Headers)
	dst.Set("body", src.Body)
}
