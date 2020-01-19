package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Host contains all info needed to create a host in Nagios
// TODO: Test to see if we need both JSON and schema tags
// Using tag with both JSON and schema because a POST uses URL encoding to send data

// Contact contains all attributes that can be set in Nagios
type Contact struct {
	ContactName                 string        `json:"contact_name"`
	HostNotificationsEnabled    string        `json:"host_notifications_enabled,omitempty"`
	ServiceNotificationsEnabled string        `json:"service_notifications_enabled,omitempty"`
	HostNotificationPeriod      string        `json:"host_notification_period,omitempty"`
	ServiceNotificationPeriod   string        `json:"service_notification_period,omitempty"`
	HostNotificationOptions     string        `json:"host_notification_options,omitempty"`
	ServiceNotificationOptions  string        `json:"service_notification_options,omitempty"`
	HostNotificationCommands    []interface{} `json:"host_notification_commands,omitempty"`
	ServiceNotificationCommands []interface{} `json:"service_notification_commands,omitempty"`
	Alias                       string        `json:"alias,omitempty"`
	ContactGroups               []interface{} `json:"contact_groups,omitempty"`
	Templates                   []interface{} `json:"use,omitempty"`
	Email                       string        `json:"email,omitempty"`
	Pager                       string        `json:"pager,omitempty"`
	Address1                    string        `json:"address1,omitempty"`
	Address2                    string        `json:"address2,omitempty"`
	Address3                    string        `json:"address3,omitempty"`
	CanSubmitCommands           string        `json:"can_submit_commands,omitempty"`
	RetainStatusInformation     string        `json:"retain_status_information,omitempty"`
	RetainNonstatusInformation  string        `json:"retain_nonstatus_information,omitempty"`
}

/*
	For any bool value, we allow the user to provide a true/false value, but you will notice
	that we immediately convert it to its integer form and then to a string. We want to provide
	the user with an easy to use schema, but Nagios wants the data as a one or zero in string format.
	This seemed to be the easiest way to accomplish that and I wanted to note why it was done that way.
*/

func resourceContact() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"contact_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the contact",
			},
			"host_notifications_enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Determines whether or not the contact will receive notifications about host problems and recoveries",
			},
			"service_notifications_enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Determines whether or not the contact will receive notifications about service problems and recoveries",
			},
			"host_notification_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The short name of the time period during which the contact can be notified about host problems or recoveries",
			},
			"service_notification_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The short name of the time period during which the contact can be notified about service problems or recoveries",
			},
			"host_notification_options": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The host states for which notifications can be sent out to this contact. Valid options are a combination of one or more of the following: d = notify on DOWN host states, u = notify on UNREACHABLE host states, r = notify on host recoveries (UP states), f = notify when the host starts and stops flapping, and s = send notifications",
			},
			"service_notification_options": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The service states for which notifications can be sent out to this contact. Valid options are a combination of one or more of the following: w = notify on WARNING service states, u = notify on UNKNOWN service states, c = notify on CRITICAL service states, r = notify on service recoveries (OK states), and f = notify when the service starts and stops flapping.",
			},
			"host_notification_commands": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "A list of the short names of the commands used to notify the contact of a host problem or recovery. Multiple notification commands should be separated by commas. All notification commands are executed when the contact needs to be notified",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_notification_commands": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "A list of the short names of the commands used to notify the contact of a service problem or recovery. Multiple notification commands should be separated by commas. All notification commands are executed when the contact needs to be notified",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alias": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A longer name or description for the contact",
			},
			"contact_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The short name(s) of the contactgroup(s) that the contact belongs to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"templates": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The contact templates to apply to the contact",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defines an email address for the contact",
			},
			"pager": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defines a pager number for the contact",
			},
			"address1": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defines additional 'addresses' for the contact",
			},
			"address2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defines additional 'addresses' for the contact",
			},
			"address3": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defines additional 'addresses' for the contact",
			},
			"can_submit_commands": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines whether or not the contact can submit external commands to Nagios from the CGIs",
			},
			"retain_status_information": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines whether or not status-related information about the contact is retained across program restarts",
			},
			"retain_nonstatus_information": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines whether or not non-status information about the contact is retained across program restarts.",
			},
		},
		Create: resourceCreateContact,
		Read:   resourceReadContact,
		Update: resourceUpdateContact,
		Delete: resourceDeleteContact,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateContact(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	contact := setContactFromSchema(d)

	_, err := nagiosClient.newContact(contact)

	if err != nil {
		return err
	}

	d.SetId(contact.ContactName)

	return resourceReadContact(d, m)
}

func resourceReadContact(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	contact, err := nagiosClient.getContact(d.Id())

	if err != nil {
		return err
	}

	if contact == nil {
		// contact not found. Let Terraform know to delete the state
		d.SetId("")
		return nil
	}

	setDataFromContact(d, contact)

	return nil
}

func resourceUpdateContact(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	contact := setContactFromSchema(d)

	oldVal, _ := d.GetChange("contact_name")

	if oldVal == "" {
		oldVal = d.Get("contact_name").(string)
	}

	err := nagiosClient.updateContact(contact, oldVal)

	if err != nil {
		return err
	}

	setDataFromContact(d, contact)

	return resourceReadContact(d, m)
}

func resourceDeleteContact(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteContact(d.Id())

	if err != nil {
		return err
	}

	return nil
}

func setDataFromContact(d *schema.ResourceData, contact *Contact) error {
	d.SetId(contact.ContactName)
	d.Set("contact_name", contact.ContactName)
	d.Set("host_notifications_enabled", contact.HostNotificationsEnabled)
	d.Set("service_notifications_enabled", contact.ServiceNotificationsEnabled)
	d.Set("host_notification_period", contact.HostNotificationPeriod)
	d.Set("service_notification_period", contact.ServiceNotificationPeriod)
	d.Set("host_notification_options", contact.HostNotificationOptions)
	d.Set("service_notification_options", contact.ServiceNotificationOptions)
	d.Set("host_notification_commands", contact.HostNotificationCommands)
	d.Set("service_notification_commands", contact.ServiceNotificationCommands)

	// Optional attributes
	if contact.Alias != "" {
		d.Set("alias", contact.Alias)
	}

	if contact.ContactGroups != nil {
		d.Set("contact_groups", contact.ContactGroups)
	}

	if contact.Templates != nil {
		d.Set("templates", contact.Templates)
	}

	if contact.Email != "" {
		d.Set("email", contact.Email)
	}

	if contact.Pager != "" {
		d.Set("pager", contact.Pager)
	}

	if contact.Address1 != "" {
		d.Set("address1", contact.Address1)
	}

	if contact.Address2 != "" {
		d.Set("address2", contact.Address2)
	}

	if contact.Address3 != "" {
		d.Set("address3", contact.Address3)
	}

	if contact.CanSubmitCommands != "" {
		d.Set("can_submit_commands", contact.CanSubmitCommands)
	}

	if contact.RetainStatusInformation != "" {
		d.Set("retain_status_information", contact.RetainStatusInformation)
	}

	if contact.RetainNonstatusInformation != "" {
		d.Set("retain_nonstatus_information", contact.RetainNonstatusInformation)
	}

	return nil
}

func setContactFromSchema(d *schema.ResourceData) *Contact {
	contact := &Contact{
		ContactName:                 d.Get("contact_name").(string),
		HostNotificationsEnabled:    convertBoolToIntToString(d.Get("host_notifications_enabled").(bool)),
		ServiceNotificationsEnabled: convertBoolToIntToString(d.Get("service_notifications_enabled").(bool)),
		HostNotificationPeriod:      d.Get("host_notification_period").(string),
		ServiceNotificationPeriod:   d.Get("service_notification_period").(string),
		HostNotificationOptions:     d.Get("host_notification_options").(string),
		ServiceNotificationOptions:  d.Get("service_notification_options").(string),
		HostNotificationCommands:    d.Get("host_notification_commands").(*schema.Set).List(),
		ServiceNotificationCommands: d.Get("service_notification_commands").(*schema.Set).List(),
		Alias:                       d.Get("alias").(string),
		ContactGroups:               d.Get("contact_groups").(*schema.Set).List(),
		Templates:                   d.Get("templates").(*schema.Set).List(),
		Email:                       d.Get("email").(string),
		Pager:                       d.Get("pager").(string),
		Address1:                    d.Get("address1").(string),
		Address2:                    d.Get("address2").(string),
		Address3:                    d.Get("address3").(string),
		CanSubmitCommands:           convertBoolToIntToString(d.Get("can_submit_commands").(bool)),
		RetainStatusInformation:     convertBoolToIntToString(d.Get("retain_status_information").(bool)),
		RetainNonstatusInformation:  convertBoolToIntToString(d.Get("retain_nonstatus_information").(bool)),
	}

	return contact
}
