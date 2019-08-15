package main

import (
	"github.com/devopsdunkin/terraform-provider-nagios/nagios"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return nagios.Provider()
		},
	})
}
