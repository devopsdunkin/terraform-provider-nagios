package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Hostgroup contains all info needed to create a hostgroup in Nagios
// TODO: Test to see if we need both JSON and schema tags
// EWe tag with both JSON and schema because a POST uses URL encoding to send data
// A GET returns data in JSON format
type Hostgroup struct {
	Name    string        `json:"hostgroup_name" schema:"hostgroup_name"`
	Alias   string        `json:"alias" schema:"alias"`
	Members []interface{} `json:"members" schema:"members"`
}

func resourceHostGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the hostgroup. It can be up to 255 characters long.",
				// ValidateFunc: validation.StringLenBetween(1, 255),
			},
			"alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of the hostgroup",
				// ValidateFunc: validation.StringLenBetween(1, 255),
			},
			"members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of hosts to be members of this hostgroup",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Create: resourceCreateHostGroup,
		Read:   resourceReadHostGroup,
		Update: resourceUpdateHostGroup,
		Delete: resourceDeleteHostGroup,
		// Importer: &schema.ResourceImporter{ // TODO: Need to figure out what is needed here
		// 	State: schema.ImportStatePassthrough,
		// },
	}
}

func resourceCreateHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	hostgroup := &Hostgroup{
		Name:    d.Get("name").(string),
		Alias:   d.Get("alias").(string),
		Members: d.Get("members").(*schema.Set).List(),
	}

	_, err := nagiosClient.newHostgroup(hostgroup)

	if err != nil {
		return err
	}

	d.SetId(hostgroup.Name)

	return resourceReadHostGroup(d, m)
}

// TODO: When no changes are done, it still says "apply complete". Believe it should say "Infrastructure up-to-date"
func resourceReadHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	hostgroup, err := nagiosClient.getHostgroup(d.Id())

	if err != nil {
		return err
	}

	if hostgroup == nil {
		// Hostgroup not found in Nagios. Update terraform state
		d.SetId("")
		return nil
	}

	d.Set("name", hostgroup.Name)
	d.Set("alias", hostgroup.Alias)
	d.Set("members", hostgroup.Members)

	return nil
}

func resourceUpdateHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	hostgroup := &Hostgroup{
		Name:    d.Get("name").(string),
		Alias:   d.Get("alias").(string),
		Members: d.Get("members").(*schema.Set).List(),
	}

	oldVal, _ := d.GetChange("name")

	if oldVal == "" { // No change, but perhaps the resource was manually deleted and need to update it so pass in the same name
		oldVal = d.Get("name").(string)
	}

	err := nagiosClient.updateHostgroup(hostgroup, oldVal)

	if err != nil {
		return err
	}

	d.SetId(hostgroup.Name)
	d.Set("name", hostgroup.Name)
	d.Set("alias", hostgroup.Alias)
	d.Set("members", hostgroup.Members)

	return resourceReadHostGroup(d, m)
}

func resourceDeleteHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteHostgroup(d.Id())

	if err != nil {
		return err
	}

	return nil
}
