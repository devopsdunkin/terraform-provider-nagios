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
	Name    string        `json:"servicegroup_name" schema:"servicegroup_name"`
	Alias   string        `json:"alias" schema:"alias"`
	Members []interface{} `json:"members" schema:"members"`
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
			"members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The hosts that the grouping of services should run on",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
		Name:    d.Get("name").(string),
		Alias:   d.Get("alias").(string),
		Members: d.Get("members").(*schema.Set).List(),
	}

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

	d.SetId(servicegroup.Name)
	d.Set("name", servicegroup.Name)
	d.Set("alias", servicegroup.Alias)
	d.Set("members", servicegroup.Members)

	return nil
}

func resourceUpdateServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	servicegroup := &Servicegroup{
		Name:    d.Get("name").(string),
		Alias:   d.Get("alias").(string),
		Members: d.Get("members").(*schema.Set).List(),
	}

	oldVal, _ := d.GetChange("name")

	err := nagiosClient.updateServicegroup(servicegroup, oldVal)

	if err != nil {
		return err
	}

	// TODO: name and alias are not getting set.
	d.SetId(servicegroup.Name)
	d.Set("name", servicegroup.Name)
	d.Set("alias", servicegroup.Alias)
	d.Set("members", servicegroup.Members)

	return resourceReadServiceGroup(d, m)
}

func resourceDeleteServiceGroup(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteServicegroup(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error trying to delete resource - %s", err.Error())
		return err
	}

	// Update Terraform state that we have deleted the resource
	d.SetId("")

	return nil
}
