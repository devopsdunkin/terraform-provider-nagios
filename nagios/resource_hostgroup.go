package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Hostgroup contains all info needed to create a hostgroup in Nagios
// TODO: Test to see if we need both JSON and schema tags
// EWe tag with both JSON and schema because a POST uses URL encoding to send data
// A GET returns data in JSON format
type Hostgroup struct {
	Name      string        `json:"hostgroup_name" schema:"hostgroup_name"`
	Alias     string        `json:"alias" schema:"alias"`
	Members   []interface{} `json:"members" schema:"members"`
	Notes     string        `json:"notes" schema:"notes"`
	NotesURL  string        `json:"notes_url" schema:"notes_url"`
	ActionURL string        `json:"action_url" schema:"action_url"`
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

	hostgroup := setHostgroupFromSchema(d)

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

	setDataFromHostgroup(d, hostgroup)

	return nil
}

func resourceUpdateHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	hostgroup := setHostgroupFromSchema(d)

	oldVal, _ := d.GetChange("name")

	if oldVal == "" { // No change, but perhaps the resource was manually deleted and need to update it so pass in the same name
		oldVal = d.Get("name").(string)
	}

	err := nagiosClient.updateHostgroup(hostgroup, oldVal)

	if err != nil {
		return err
	}

	setDataFromHostgroup(d, hostgroup)

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

func setDataFromHostgroup(d *schema.ResourceData, hostgroup *Hostgroup) {
	// required attributes
	d.SetId(hostgroup.Name)
	d.Set("name", hostgroup.Name)
	d.Set("alias", hostgroup.Alias)

	// optional attributes
	if hostgroup.Members != nil {
		d.Set("members", hostgroup.Members)
	}

	if hostgroup.Notes != "" {
		d.Set("notes", hostgroup.Notes)
	}

	if hostgroup.NotesURL != "" {
		d.Set("notes_url", hostgroup.NotesURL)
	}

	if hostgroup.ActionURL != "" {
		d.Set("action_url", hostgroup.ActionURL)
	}
}

func setHostgroupFromSchema(d *schema.ResourceData) *Hostgroup {
	hostgroup := &Hostgroup{
		Name:      d.Get("name").(string),
		Alias:     d.Get("alias").(string),
		Members:   d.Get("members").(*schema.Set).List(),
		Notes:     d.Get("notes").(string),
		NotesURL:  d.Get("notes_url").(string),
		ActionURL: d.Get("action_url").(string),
	}

	return hostgroup
}
