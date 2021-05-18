package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/statusflare-com/terraform-provider-statusflare/statusflare"
)

func resourceStatusPage() *schema.Resource {

	fields := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the status page.",
		},
		"monitors": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "IDs of monitors attached to this status page.",
		},
		"custom_domain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The custom domain attached to your status page.",
		},
		"custom_domain_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The path for your custom domain. The default is '/'.",
		},
		"hide_monitor_details": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Hide monitor details (URL, scheme, ..) on the status page. The default is false.",
		},
		"hide_statusflare": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Hide statusflare branding/links on the status page. The default is false.",
		},
		"histogram_days": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Number of days to render on status page for each monitor. The default is 90.",
		},
		"config": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "Additional configuration of the status page. See example for list of options.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	return &schema.Resource{
		CreateContext: resourceStatusPageCreate,
		ReadContext:   resourceStatusPageRead,
		UpdateContext: resourceStatusPageUpdate,
		DeleteContext: resourceStatusPageDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: fields,
	}
}

func resourceStatusPageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var client *statusflare.Client = meta.(*statusflare.Client)

	statusPage, err := client.GetStatusPage(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	statusPageToData(statusPage, d)
	return diags
}

func resourceStatusPageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var client *statusflare.Client = meta.(*statusflare.Client)

	statusPage := statusflare.StatusPage{}
	dataToStatusPage(d, &statusPage)

	err := client.CreateStatusPage(&statusPage)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(statusPage.Id)
	resourceStatusPageRead(ctx, d, meta)
	return diags
}

func resourceStatusPageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *statusflare.Client = meta.(*statusflare.Client)

	statusPage, err := client.GetStatusPage(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	dataToStatusPage(d, statusPage)
	err = client.SaveStatusPage(statusPage)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceStatusPageRead(ctx, d, meta)
}

func resourceStatusPageDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		err    error
		diags  diag.Diagnostics
		client *statusflare.Client = m.(*statusflare.Client)
	)

	err = client.DeleteStatusPage(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func dataToStatusPage(src *schema.ResourceData, dst *statusflare.StatusPage) {
	dst.Name = src.Get("name").(string)
	dst.Monitors = toStrArray(src.Get("monitors").([]interface{}))
	dst.CustomDomain = src.Get("custom_domain").(string)
	dst.CustomDomainPath = src.Get("custom_domain_path").(string)
	dst.HideMonitorDetails = src.Get("hide_monitor_details").(bool)
	dst.HideStatusflare = src.Get("hide_statusflare").(bool)

	dst.Config.Title = src.Get("config.title").(string)
	dst.Config.HistogramDays = src.Get("histogram_days").(int)
	dst.Config.LogoUrl = src.Get("config.logo_url").(string)
	dst.Config.FaviconUrl = src.Get("config.favicon_url").(string)
	dst.Config.AllMonitorsOperational = src.Get("config.all_monitors_operational").(string)
	dst.Config.NotAllMonitorsOperational = src.Get("config.not_all_monitors_operational").(string)
	dst.Config.MonitorOperationalLabel = src.Get("config.monitor_operational_label").(string)
	dst.Config.MonitorNotOperationalLabel = src.Get("config.monitor_not_operational_label").(string)
	dst.Config.MonitorNoDataLabel = src.Get("config.monitor_no_data_label").(string)
	dst.Config.HistogramNoData = src.Get("config.histogram_no_data").(string)
	dst.Config.HistogramNoIncidents = src.Get("config.histogram_no_incidents").(string)
	dst.Config.HistogramSomeIncidents = src.Get("config.histogram_some_incidents").(string)
}

func statusPageToData(src *statusflare.StatusPage, dst *schema.ResourceData) {
	dst.Set("name", src.Name)
	dst.Set("monitors", src.Monitors)
	dst.Set("custom_domain", src.CustomDomain)
	dst.Set("custom_domain_path", src.CustomDomainPath)
	dst.Set("hide_monitor_details", src.HideMonitorDetails)
	dst.Set("hide_statusflare", src.HideStatusflare)
	dst.Set("histogram_days", src.Config.HistogramDays)

	// TODO fix refreshing state
	// [ERROR] setting state: config: must be a map
	dst.Set("config", src.Config)
}
