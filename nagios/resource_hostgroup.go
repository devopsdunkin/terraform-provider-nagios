package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Hostgroup contains all info needed to create a hostgroup in Nagios
type Hostgroup struct {
	name  string `json:"hostgroup_name"`
	alias string `json:"alias"`
}

func resourceHostGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the hostgroup",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of the hostgroup",
			},
		},
		Create: resourceCreateHostGroup,
		Read:   resourceReadHostGroup,
		Update: resourceUpdateHostGroup,
		Delete: resourceDeleteHostGroup,
		// Exists: resourceExistsHostGroup,  #TODO: Need to figure out how to define this
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	hostgroup := &Hostgroup{
		d.Get("name").(string),
		d.Get("alias").(string),
	}

	err := nagiosClient.NewHostgroup(hostgroup)

	if err != nil {
		return err
	}

	return resourceReadHostGroup(d, m)
}

func resourceReadHostGroup(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUpdateHostGroup(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDeleteHostGroup(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceExistsHostGroup(d *schema.ResourceData, m interface{}) error {
	return nil
}
