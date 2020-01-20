package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceHostgroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHostgroupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the hostgroup. It can be up to 255 characters long.",
			},
			"alias": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the hostgroup",
			},
			"members": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Members of this hostgroup",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Notes about the servicegroup that may assist with troubleshooting",
			},
			"notes_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to a third-party documentation repository containing more information about the servicegroup",
			},
			"action_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to a third-party documentation repository containing actions to take in the event the servicegroup goes down",
			},
		},
	}
}

func dataSourceHostgroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	hostgroupName := d.Get("name").(string)

	hostgroup, err := client.getHostgroup(hostgroupName)

	if err != nil {
		return err
	}

	setDataFromHostgroup(d, hostgroup)

	return nil
}
