package nagios

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

// servicegroup contains all info needed to create a servicegroup in Nagios
// TODO: Test to see if we need both JSON and schema tags
// EWe tag with both JSON and schema because a POST uses URL encoding to send data
// A GET returns data in JSON format
type Servicegroup struct {
	Name  string `json:"servicegroup_name" schema:"servicegroup_name"`
	Alias string `json:"alias" schema:"alias"`
}

func resourceServiceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the servicegroup",
			},
			"alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of the servicegroup",
			},
		},
		Create: resourceCreateServiceGroup,
		Read:   resourceReadServiceGroup,
		Update: resourceUpdateServiceGroup,
		Delete: resourceDeleteServiceGroup,
		// Exists: resourceExistsServiceGroup,  // TODO: Need to figure out how to define this
		// Importer: &schema.ResourceImporter{ // TODO: Need to figure out what is needed here
		// 	State: schema.ImportStatePassthrough,
		// },
	}
}

func resourceCreateServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	servicegroup := &Servicegroup{
		Name:  d.Get("name").(string),
		Alias: d.Get("alias").(string),
	}

	_, err := nagiosClient.NewServicegroup(servicegroup)

	if err != nil {
		return err
	}

	d.SetId(servicegroup.Name)
	d.Set("name", servicegroup.Name)
	d.Set("alias", servicegroup.Alias)

	return resourceReadServiceGroup(d, m)
}

// TODO: When no changes are done, it still says "apply complete". Believe it should say "Infrastructure up-to-date"
func resourceReadServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)
	log.Printf("[DEBUG] name - %s", d.Id())

	servicegroup, err := nagiosClient.GetServicegroup(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error reading servicegroup - %s", err.Error())

		return err
	}

	if servicegroup == nil {
		// servicegroup not found in Nagios. Update terraform state
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] d.Set on servicegroup.Name - %s", servicegroup.Name)
	log.Printf("[DEBUG] d.Set on servicegroup.Alias - %s", servicegroup.Alias)
	log.Printf("[DEBUG] d.Id - %s", d.Id())

	d.SetId(servicegroup.Name)
	d.Set("name", servicegroup.Name)
	d.Set("alias", servicegroup.Alias)

	return nil
}

func resourceUpdateServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	log.Printf("[DEBUG] name - %s", d.Get("name").(string))

	servicegroup := &Servicegroup{
		Name:  d.Get("name").(string),
		Alias: d.Get("alias").(string),
	}

	oldVal, _ := d.GetChange("name")

	log.Printf("[DEBUG] Old value - %s", oldVal.(string))

	err := nagiosClient.UpdateServicegroup(servicegroup, oldVal) // TODO: Alias is not getting updated. It is blank

	if err != nil {
		log.Printf("[ERROR] Error updating servicegroup in Nagios - %s", err.Error())
		return err
	}

	// TODO: name and alias are not getting set.
	d.SetId(servicegroup.Name)
	d.Set("name", servicegroup.Name)
	d.Set("alias", servicegroup.Alias)

	return resourceReadServiceGroup(d, m)
}

func resourceDeleteServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.DeleteServicegroup(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error trying to delete resource - %s", err.Error())
		return err
	}

	// Update Terraform state that we have deleted the resource
	d.SetId("")

	return nil
}

// TODO: Need to determine if this needs implemented. Need more understanding of this
// func resourceExistsservicegroup(d *schema.ResourceData, m interface{}) error {
// 	return resourceReadservicegroup(d, m)
// }
