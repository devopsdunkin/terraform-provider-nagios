package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// servicegroup contains all info needed to create a servicegroup in Nagios
// TODO: Test to see if we need both JSON and schema tags
// EWe tag with both JSON and schema because a POST uses URL encoding to send data
// A GET returns data in JSON format
type Servicegroup struct {
	Name      string        `json:"servicegroup_name" schema:"servicegroup_name"`
	Alias     string        `json:"alias" schema:"alias"`
	Members   []interface{} `json:"members" schema:"members"`
	Notes     string        `json:"notes" schema:"notes"`
	NotesURL  string        `json:"notes_url" schema:"notes_url"`
	ActionURL string        `json:"action_url" schema:"action_url"`
}

func resourceServiceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Nagios servicegroup",
			},
			"alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description or other name that the servicegroup may be called. This field can be longer and more descriptive",
			},
			"members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of hosts and/or services that should be members of the servicegroup. The members must be valid hosts and services within Nagios and must be active",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Notes about the servicegroup that may assist with troubleshooting",
			},
			"notes_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL to a third-party documentation repository containing more information about the servicegroup",
			},
			"action_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL to a third-party documentation repository containing actions to take in the event the servicegroup goes down",
			},
		},
		Create: resourceCreateServiceGroup,
		Read:   resourceReadServiceGroup,
		Update: resourceUpdateServiceGroup,
		Delete: resourceDeleteServiceGroup,
		// Importer: &schema.ResourceImporter{ // TODO: Need to figure out what is needed here
		// 	State: schema.ImportStatePassthrough,
		// },
	}
}

func resourceCreateServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	servicegroup := setServicegroupFromSchema(d)

	_, err := nagiosClient.newServicegroup(servicegroup)

	if err != nil {
		return err
	}

	d.SetId(servicegroup.Name)

	return resourceReadServiceGroup(d, m)
}

func resourceReadServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	servicegroup, err := nagiosClient.getServicegroup(d.Id())

	if err != nil {
		return err
	}

	if servicegroup == nil {
		// servicegroup not found in Nagios. Update terraform state
		d.SetId("")
		return nil
	}

	setDataFromServicegroup(d, servicegroup)

	return nil
}

func resourceUpdateServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	servicegroup := setServicegroupFromSchema(d)

	oldVal, _ := d.GetChange("name")

	err := nagiosClient.updateServicegroup(servicegroup, oldVal)

	if err != nil {
		return err
	}

	// TODO: name and alias are not getting set.
	setDataFromServicegroup(d, servicegroup)

	return resourceReadServiceGroup(d, m)
}

func resourceDeleteServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteServicegroup(d.Id())

	if err != nil {
		return err
	}

	return nil
}

func setDataFromServicegroup(d *schema.ResourceData, servicegroup *Servicegroup) {
	// required attributes
	d.SetId(servicegroup.Name)
	d.Set("name", servicegroup.Name)
	d.Set("alias", servicegroup.Alias)

	// optional attributes
	if servicegroup.Members != nil {
		d.Set("members", servicegroup.Members)
	}

	if servicegroup.Notes != "" {
		d.Set("notes", servicegroup.Notes)
	}

	if servicegroup.NotesURL != "" {
		d.Set("notes_url", servicegroup.NotesURL)
	}

	if servicegroup.ActionURL != "" {
		d.Set("action_url", servicegroup.ActionURL)
	}
}

func setServicegroupFromSchema(d *schema.ResourceData) *Servicegroup {
	servicegroup := &Servicegroup{
		Name:      d.Get("name").(string),
		Alias:     d.Get("alias").(string),
		Members:   d.Get("members").(*schema.Set).List(),
		Notes:     d.Get("notes").(string),
		NotesURL:  d.Get("notes_url").(string),
		ActionURL: d.Get("action_url").(string),
	}

	return servicegroup
}
