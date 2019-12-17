package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Host contains all info needed to create a host in Nagios
// TODO: Test to see if we need both JSON and schema tags
// Using tag with both JSON and schema because a POST uses URL encoding to send data

// TODO: Need to add in all of the other fields. What we have right now will work for initial testing
type Contactgroup struct {
	ContactgroupName    string        `json:"contactgroup_name"`
	Alias               string        `json:"alias"`
	Members             []interface{} `json:"members,omitempty"`
	ContactgroupMembers []interface{} `json:"contactgroup_members,omitempty"`
}

/*
	For any bool value, we allow the user to provide a true/false value, but you will notice
	that we immediately convert it to its integer form and then to a string. We want to provide
	the user with an easy to use schema, but Nagios wants the data as a one or zero in string format.
	This seemed to be the easiest way to accomplish that and I wanted to note why it was done that way.
*/

func resourceContactgroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"contactgroup_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the contact group",
			},
			"alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Determines whether or not the contact will receive notifications about host problems and recoveries",
			},
			"members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of the short names of the commands used to notify the contact of a host problem or recovery. Multiple notification commands should be separated by commas. All notification commands are executed when the contact needs to be notified",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"contactgroup_members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of the short names of the commands used to notify the contact of a service problem or recovery. Multiple notification commands should be separated by commas. All notification commands are executed when the contact needs to be notified",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Create: resourceCreateContactgroup,
		Read:   resourceReadContactgroup,
		Update: resourceUpdateContactgroup,
		Delete: resourceDeleteContactgroup,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateContactgroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	contactgroup := setContactgroupFromSchema(d)

	_, err := nagiosClient.newContactgroup(contactgroup)

	if err != nil {
		return err
	}

	d.SetId(contactgroup.ContactgroupName)

	return resourceReadContactgroup(d, m)
}

func resourceReadContactgroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	contactgroup, err := nagiosClient.getContactgroup(d.Id())

	if err != nil {
		return err
	}

	if contactgroup == nil {
		// contact not found. Let Terraform know to delete the state
		d.SetId("")
		return nil
	}

	setDataFromContactgroup(d, contactgroup)

	return nil
}

func resourceUpdateContactgroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	contactgroup := setContactgroupFromSchema(d)

	oldVal, _ := d.GetChange("contactgroup_name")

	if oldVal == "" {
		oldVal = d.Get("contactgroup_name").(string)
	}

	err := nagiosClient.updateContactgroup(contactgroup, oldVal)

	if err != nil {
		return err
	}

	setDataFromContactgroup(d, contactgroup)

	return resourceReadContactgroup(d, m)
}

func resourceDeleteContactgroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteContactgroup(d.Id())

	if err != nil {
		return err
	}

	return nil
}

func setDataFromContactgroup(d *schema.ResourceData, contactgroup *Contactgroup) {
	d.SetId(contactgroup.ContactgroupName)
	d.Set("contactgroup_name", contactgroup.ContactgroupName)
	d.Set("alias", contactgroup.Alias)

	if contactgroup.Members != nil {
		d.Set("members", contactgroup.Members)
	}

	if contactgroup.ContactgroupMembers != nil {
		d.Set("contactgroup_members", contactgroup.ContactgroupMembers)
	}
}

func setContactgroupFromSchema(d *schema.ResourceData) *Contactgroup {
	contactgroup := &Contactgroup{
		ContactgroupName:    d.Get("contactgroup_name").(string),
		Alias:               d.Get("alias").(string),
		Members:             d.Get("members").(*schema.Set).List(),
		ContactgroupMembers: d.Get("contactgroup_members").(*schema.Set).List(),
	}

	return contactgroup
}
