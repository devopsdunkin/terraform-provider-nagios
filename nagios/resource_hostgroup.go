package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Hostgroup contains all info needed to create a hostgroup in Nagios
type Hostgroup struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"hostgroup_name"`
	Alias string `json:"alias"`
}

func resourceHostGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the hostgroup",
			},
			"alias": {
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
		d.Id(),
		d.Get("name").(string),
		d.Get("alias").(string),
	}

	err := nagiosClient.NewHostgroup(hostgroup)

	if err != nil {
		return err
	}

	d.SetId(hostgroup.Id)

	return resourceReadHostGroup(d, m)
}

func resourceReadHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	hostgroup, err := nagiosClient.GetHostgroup(d.Get("name").(string))

	if err != nil {
		return err
	}

	d.SetId(hostgroup.Id)
	d.Set("name", hostgroup.Name)
	d.Set("alias", hostgroup.Alias)

	return nil
}

func resourceUpdateHostGroup(d *schema.ResourceData, m interface{}) error {
	return resourceReadHostGroup(d, m)
}

func resourceDeleteHostGroup(d *schema.ResourceData, m interface{}) error {
	return resourceReadHostGroup(d, m)
}

func resourceExistsHostGroup(d *schema.ResourceData, m interface{}) error {
	return resourceReadHostGroup(d, m)
}
