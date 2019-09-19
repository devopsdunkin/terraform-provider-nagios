package nagios

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

// Host contains all info needed to create a host in Nagios
// TODO: Test to see if we need both JSON and schema tags
// Using tag with both JSON and schema because a POST uses URL encoding to send data

// TODO: Need to add in all of the other fields. What we have right now will work for initial testing
type Host struct {
	Name                 string        `json:"host_name" schema:"host_name"`
	Alias                string        `json:"alias" schema:"alias"`
	Address              string        `json:"address" schema:"address"`
	MaxCheckAttempts     string        `json:"max_check_attempts" schema:"max_check_attempts"`
	CheckPeriod          string        `json:"check_period" schema:"check_period"`
	NotificationInterval string        `json:"notification_interval" schema:"notification_interval"`
	NotificationPeriod   string        `json:"notification_period" schema:"notification_period"`
	Contacts             []interface{} `json:"contacts" schema:"contacts"`
	Templates            []interface{} `json:"use" schema:"use"`
}

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the host",
			},
			"alias": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The alias of the host",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IP address of the host",
			},
			"max_check_attempts": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The maximum number of times it will check the host",
			},
			"check_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The check period that the host should belong to",
			},
			"notification_interval": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How often the host should be checked",
			},
			"notification_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "What time periods should the host be alerted on",
			},
			"contacts": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of users or groups to notify when an alert is triggered",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"templates": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of templates to apply to the host",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Create: resourceCreateHost,
		Read:   resourceReadHost,
		Update: resourceUpdateHost,
		Delete: resourceDeleteHost,
		// Exists: resourceExistsHost,  // TODO: Need to figure out how to define this
		// Importer: &schema.ResourceImporter{ // TODO: Need to figure out what is needed here
		// 	State: schema.ImportStatePassthrough,
		// },
	}
}

func resourceCreateHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	host := &Host{
		Name:                 d.Get("name").(string),
		Alias:                d.Get("alias").(string),
		Address:              d.Get("address").(string),
		MaxCheckAttempts:     d.Get("max_check_attempts").(string),
		CheckPeriod:          d.Get("check_period").(string),
		NotificationInterval: d.Get("notification_interval").(string),
		NotificationPeriod:   d.Get("notification_period").(string),
		Contacts:             d.Get("contacts").(*schema.Set).List(),
		Templates:            d.Get("templates").(*schema.Set).List(),
	}

	log.Printf("[DEBUG] host.name - %s", host.Name)
	log.Printf("[DEBUG] host.address - %s", host.Address)
	log.Printf("[DEBUG] host.max_check_attempts - %s", host.MaxCheckAttempts)
	log.Printf("[DEBUG] host.check_period - %s", host.CheckPeriod)
	log.Printf("[DEBUG] host.notification_interval - %s", host.NotificationInterval)
	log.Printf("[DEBUG] host.notification_period - %s", host.NotificationPeriod)
	log.Printf("[DEBUG] host.contacts - %s", host.Contacts)

	_, err := nagiosClient.NewHost(host)

	if err != nil {
		return err
	}

	d.SetId(host.Name)
	d.Set("name", host.Name)
	d.Set("alias", host.Alias)
	d.Set("address", host.Address)
	d.Set("max_check_attempts", host.MaxCheckAttempts)
	d.Set("check_period", host.CheckPeriod)
	d.Set("notification_interval", host.NotificationInterval)
	d.Set("notification_period", host.NotificationPeriod)
	d.Set("contacts", host.Contacts) // TODO: If contact does not exist in Nagios, it should not add it. Applies fine through API but causes validation errors when trying to apply config manually
	d.Set("templates", host.Templates)

	return resourceReadHost(d, m)
}

// TODO: When no changes are done, it still says "apply complete". Believe it should say "Infrastructure up-to-date"
func resourceReadHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	host, err := nagiosClient.GetHost(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error reading host - %s", err.Error())

		return err
	}

	if host == nil {
		// host not found in Nagios. Update terraform state
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] d.Set within Read func on host.Contacts - %s", host.Contacts[0])
	log.Printf("[DEBUG] d.Id - %s", d.Id())

	d.SetId(host.Name)
	d.Set("name", host.Name)
	d.Set("alias", host.Alias)
	d.Set("address", host.Address)
	d.Set("max_check_attempts", host.MaxCheckAttempts)
	d.Set("check_period", host.CheckPeriod)
	d.Set("notification_interval", host.NotificationInterval)
	d.Set("notification_period", host.NotificationPeriod)
	d.Set("contacts", host.Contacts)
	d.Set("templates", host.Templates)

	return nil
}

func resourceUpdateHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	log.Printf("[DEBUG] name - %s", d.Get("name").(string))

	host := &Host{
		Name:                 d.Get("name").(string),
		Alias:                d.Get("alias").(string),
		Address:              d.Get("address").(string),
		MaxCheckAttempts:     d.Get("max_check_attempts").(string),
		CheckPeriod:          d.Get("check_period").(string),
		NotificationInterval: d.Get("notification_interval").(string),
		NotificationPeriod:   d.Get("notification_period").(string),
		Contacts:             d.Get("contacts").(*schema.Set).List(),
		Templates:            d.Get("templates").(*schema.Set).List(),
	}

	oldVal, _ := d.GetChange("name")

	log.Printf("[DEBUG] Old value - %s", oldVal.(string))

	err := nagiosClient.UpdateHost(host, oldVal)

	if err != nil {
		log.Printf("[ERROR] Error updating host in Nagios - %s", err.Error())
		return err
	}

	// TODO: name and alias are not getting set.
	d.SetId(host.Name)
	d.Set("name", host.Name)
	d.Set("alias", host.Alias)
	d.Set("address", host.Address)
	d.Set("max_check_attempts", host.MaxCheckAttempts)
	d.Set("check_period", host.CheckPeriod)
	d.Set("notification_interval", host.NotificationInterval)
	d.Set("notification_period", host.NotificationPeriod)
	d.Set("contacts", host.Contacts)
	d.Set("templates", host.Templates)

	return resourceReadHost(d, m)
}

func resourceDeleteHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.DeleteHost(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error trying to delete resource - %s", err.Error())
		return err
	}

	// Update Terraform state that we have deleted the resource
	d.SetId("")

	return nil
}

// TODO: Need to determine if this needs implemented. Need more understanding of this
// func resourceExistshost(d *schema.ResourceData, m interface{}) error {
// 	return resourceReadhost(d, m)
// }
