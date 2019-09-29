package nagios

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

// Hostgroup contains all info needed to create a hostgroup in Nagios
// TODO: Test to see if we need both JSON and schema tags
// EWe tag with both JSON and schema because a POST uses URL encoding to send data
// A GET returns data in JSON format
type Hostgroup struct {
	Name  string `json:"hostgroup_name" schema:"hostgroup_name"`
	Alias string `json:"alias" schema:"alias"`
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
		},
		Create: resourceCreateHostGroup,
		Read:   resourceReadHostGroup,
		Update: resourceUpdateHostGroup,
		Delete: resourceDeleteHostGroup,
		// Exists: resourceExistsHostGroup,  // TODO: Need to figure out how to define this
		// Importer: &schema.ResourceImporter{ // TODO: Need to figure out what is needed here
		// 	State: schema.ImportStatePassthrough,
		// },
	}
}

func resourceCreateHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	hostgroup := &Hostgroup{
		Name:  d.Get("name").(string),
		Alias: d.Get("alias").(string),
	}

	_, err := nagiosClient.NewHostgroup(hostgroup)

	if err != nil {
		return err
	}

	log.Printf("[TRACE] Passed NewHostgroup err check")

	d.SetId(hostgroup.Name)

	log.Printf("[TRACE] Passed setting state")

	return resourceReadHostGroup(d, m)
}

// TODO: When no changes are done, it still says "apply complete". Believe it should say "Infrastructure up-to-date"
func resourceReadHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	hostgroup, err := nagiosClient.GetHostgroup(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error reading hostgroup - %s", err.Error())

		return err
	}

	if hostgroup == nil {
		// Hostgroup not found in Nagios. Update terraform state
		d.SetId("")
		return nil
	}

	d.Set("name", hostgroup.Name)
	d.Set("alias", hostgroup.Alias)

	return nil
}

func resourceUpdateHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	log.Printf("[DEBUG] name - %s", d.Get("name").(string))

	hostgroup := &Hostgroup{
		Name:  d.Get("name").(string),
		Alias: d.Get("alias").(string),
	}

	oldVal, _ := d.GetChange("name")

	if oldVal == "" { // No change, but perhaps the resource was manually deleted and need to update it so pass in the same name
		oldVal = d.Get("name").(string)
	}

	err := nagiosClient.UpdateHostgroup(hostgroup, oldVal)

	if err != nil {
		log.Printf("[ERROR] Error updating hostgroup in Nagios - %s", err.Error())
		return err
	}

	d.SetId(hostgroup.Name)
	d.Set("name", hostgroup.Name)
	d.Set("alias", hostgroup.Alias)

	return resourceReadHostGroup(d, m)
}

func resourceDeleteHostGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.DeleteHostgroup(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error trying to delete resource - %s", err.Error())
		return err
	}

	// Update Terraform state that we have deleted the resource
	d.SetId("")

	return nil
}

// TODO: Need to determine if this needs implemented. Need more understanding of this
// func resourceExistsHostGroup(d *schema.ResourceData, m interface{}) error {
// 	return resourceReadHostGroup(d, m)
// }
