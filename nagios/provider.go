package nagios

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

// Provider function defines the schema and resources for this Nagios provider
func NagiosProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NAGIOS_URL", ""),
				Description: "The URL of the Nagios application",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", ""),
				Description: "API token to authenticate to Nagios",
				Sensitive:   true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"nagios_hostgroup":    resourceHostGroup(),
			"nagios_host":         resourceHost(),
			"nagios_service":      resourceService(),
			"nagios_servicegroup": resourceServiceGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	token := d.Get("token").(string)

	if url == "" {
		return nil, errors.New("Invalid or no value supplied for URL")
	}

	if token == "" {
		return nil, errors.New("Invalid or no value supplied for token")
	}

	return NewClient(url, token), nil
}
