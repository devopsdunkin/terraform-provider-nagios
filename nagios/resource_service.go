package nagios

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

// Service contains all info needed to create a service in Nagios
// Using tag with both JSON and schema because a POST uses URL encoding to send data

// TODO: Need to add in all of the other fields. What we have right now will work for initial testing
type Service struct {
	ServiceName          string        `json:"config_name" schema:"config_name"`
	HostName             []interface{} `json:"host_name" schema:"host_name"`
	Description          string        `json:"service_description" schema:"service_description"`
	CheckCommand         string        `json:"check_command" schema:"check_command"`
	MaxCheckAttempts     string        `json:"max_check_attempts" schema:"max_check_attempts"`
	CheckInterval        string        `json:"check_interval" schema:"check_interval"`
	RetryInterval        string        `json:"retry_interval" schema:"retry_interval"`
	CheckPeriod          string        `json:"check_period" schema:"check_period"`
	NotificationInterval string        `json:"notification_interval" schema:"notification_interval"`
	NotificationPeriod   string        `json:"notification_period" schema:"notification_period"`
	Contacts             []interface{} `json:"contacts" schema:"contacts"`
	Templates            []interface{} `json:"use" schema:"use"`
}

/* TODO: Need to figure out the dependencies here
1. A service must have a host attached to it in order to validate the config when restarting Nagios core
2. A service can have a template attached to it that has a host attached to the template
3. A service can have a template attached to it that has a hostgroup attached to the template

A user must provide either a host_name or template. If no host name is provided, we will have to check
the template to make sure a host or hostgrup is attached to it
*/

func resourceService() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the service that provides a general idea of what the service is checking",
			},
			"host_name": {
				// TODO: Unsure of how I want to handle this. We ultimately want this field to be optional since the better way to manage services
				// would be to put them in a service group. Users should have the option of what they want to do. Probably need to use the &force=1
				// command in API call.
				Type:        schema.TypeSet, // TODO: This needs to be a list
				Optional:    true,
				Description: "The name of the host to apply service to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the service",
			},
			"check_command": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The command to run to perform a check of the service",
			},
			"max_check_attempts": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The maximum number of times it will check the service", // TODO: Need clarification of this description and what this attr does
			},
			"check_interval": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How often should a check be performed on the service",
			},
			"retry_interval": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How often should we retry the check while the service is down",
			},
			"check_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "What time period should the service be checked",
			},
			"notification_interval": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How frequent should we notify that the service is down",
			},
			"notification_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "What time period should the service be alerted on",
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
				Description: "List of templates to apply to the service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Create: resourceCreateService,
		Read:   resourceReadService,
		Update: resourceUpdateService,
		Delete: resourceDeleteService,
		// Exists: resourceExistsService,  // TODO: Need to figure out how to define this
		// Importer: &schema.ResourceImporter{ // TODO: Need to figure out what is needed here
		// 	State: schema.ImportStatePassthrough,
		// },
	}
}

func resourceCreateService(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Start of resourceCreateService")
	nagiosClient := m.(*Client)

	service := &Service{
		ServiceName:          d.Get("service_name").(string),
		HostName:             d.Get("host_name").(*schema.Set).List(),
		Description:          d.Get("description").(string),
		CheckCommand:         d.Get("check_command").(string),
		MaxCheckAttempts:     d.Get("max_check_attempts").(string),
		CheckInterval:        d.Get("check_interval").(string),
		RetryInterval:        d.Get("retry_interval").(string),
		CheckPeriod:          d.Get("check_period").(string),
		NotificationInterval: d.Get("notification_interval").(string),
		NotificationPeriod:   d.Get("notification_period").(string),
		Contacts:             d.Get("contacts").(*schema.Set).List(),
		Templates:            d.Get("templates").(*schema.Set).List(),
	}

	// if service.HostName == "" {
	// TODO: Need to add hostgroup membership to schema. Then we will check if hostgroup has been provided or is a member in Nagios
	// }

	_, err := nagiosClient.newService(service)

	if err != nil {
		return err
	}

	d.SetId(service.ServiceName)

	return resourceReadService(d, m)
}

// TODO: When no changes are done, it still says "apply complete". Believe it should say "Infrastructure up-to-date"
func resourceReadService(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Start of resourceReadService")
	nagiosClient := m.(*Client)

	service, err := nagiosClient.getService(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error reading service - %s", err.Error())

		return err
	}

	if service == nil {
		// service not found in Nagios. Update terraform state
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] Starting to set state...")

	d.Set("service_name", service.ServiceName)
	d.Set("host_name", service.HostName)
	d.Set("description", service.Description)
	d.Set("check_command", service.CheckCommand)
	d.Set("max_check_attempts", service.MaxCheckAttempts)
	d.Set("check_interval", service.CheckInterval)
	d.Set("retry_interval", service.RetryInterval)
	d.Set("check_period", service.CheckPeriod)
	d.Set("notification_interval", service.NotificationInterval)
	d.Set("notification_period", service.NotificationPeriod)
	d.Set("contacts", service.Contacts)
	d.Set("templates", service.Templates)

	log.Printf("[DEBUG] Finished setting state...")

	return nil
}

func resourceUpdateService(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Start of resourceUpdateService")
	nagiosClient := m.(*Client)

	log.Printf("[DEBUG] resourceUpdateService => name - %s", d.Get("name").(string))

	service := &Service{
		ServiceName:          d.Get("service_name").(string),
		HostName:             d.Get("host_name").(*schema.Set).List(),
		Description:          d.Get("description").(string),
		CheckCommand:         d.Get("check_command").(string),
		MaxCheckAttempts:     d.Get("max_check_attempts").(string),
		CheckInterval:        d.Get("check_interval").(string),
		RetryInterval:        d.Get("retry_interval").(string),
		CheckPeriod:          d.Get("check_period").(string),
		NotificationInterval: d.Get("notification_interval").(string),
		NotificationPeriod:   d.Get("notification_period").(string),
		Contacts:             d.Get("contacts").(*schema.Set).List(),
		Templates:            d.Get("templates").(*schema.Set).List(),
	}

	oldVal, _ := d.GetChange("service_name")

	if oldVal == "" { // No change, but perhaps the resource was manually deleted and need to update it so pass in the same name
		oldVal = d.Get("service_name").(string)
		log.Printf("[DEBUG] resourceUpdateService => oldVal was blank, so should be same name as current - %s", oldVal)
	}

	// HTTP PUT for a Nagios service is weirder than the rest. Requires /api/v1/config/service/<service_name>/<service_description>?<rest of url>

	err := nagiosClient.updateService(service, oldVal)

	if err != nil {
		log.Printf("[ERROR] Error updating service in Nagios - %s", err.Error())
		return err
	}

	log.Printf("[DEBUG] resourceUpdateservice => Updateservice successful - %s", *service)

	// TODO: name and alias are not getting set.
	d.SetId(service.ServiceName)
	d.Set("service_name", service.ServiceName)
	d.Set("host_name", service.HostName)
	d.Set("description", service.Description)
	d.Set("check_command", service.CheckCommand)
	d.Set("max_check_attempts", service.MaxCheckAttempts)
	d.Set("check_interval", service.CheckInterval)
	d.Set("retry_interval", service.RetryInterval)
	d.Set("check_period", service.CheckPeriod)
	d.Set("notification_interval", service.NotificationInterval)
	d.Set("notification_period", service.NotificationPeriod)
	d.Set("contacts", service.Contacts)
	d.Set("templates", service.Templates)

	return resourceReadService(d, m)
}

func resourceDeleteService(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteService(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error trying to delete resource - %s", err.Error())
		return err
	}

	return nil
}
