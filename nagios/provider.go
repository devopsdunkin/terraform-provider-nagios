package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider function defines the schema and resources for this Nagios provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("URL", ""),
				Description: "The URL of the Nagios application",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TOKEN", ""),
				Description: "The API token used to authenticate to Nagios",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"nagios_hostgroup": resourceHostGroup(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	token := d.Get("token").(string)
	return NewClient(url, token), nil
}
