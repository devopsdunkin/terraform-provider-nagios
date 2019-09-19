package nagios

import (
	"errors"
	"log"

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
				Description: "API token to authenticate to Nagios",
				Sensitive:   true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"nagios_hostgroup": resourceHostGroup(),
			"nagios_host":      resourceHost(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	token := d.Get("token").(string)

	if url == "" || token == "" {
		log.Printf("[ERROR] Invalid or no value supplied for URL or token")
		return nil, errors.New("Invalid or no value supplied for URL or token")
	}

	return NewClient(url, token), nil
}
