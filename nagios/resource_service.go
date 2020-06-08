package nagios

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// Service contains all info needed to create a service in Nagios
// Using tag with both JSON and schema because a POST uses URL encoding to send data

// TODO: Need to add in all of the other fields. What we have right now will work for initial testing
type Service struct {
	ServiceName                string                 `json:"config_name"`
	HostName                   []interface{}          `json:"host_name"`
	DisplayName                string                 `json:"display_name,omitempty"`
	Description                string                 `json:"service_description"`
	CheckCommand               string                 `json:"check_command"`
	MaxCheckAttempts           string                 `json:"max_check_attempts"`
	CheckInterval              string                 `json:"check_interval"`
	RetryInterval              string                 `json:"retry_interval"`
	CheckPeriod                string                 `json:"check_period"`
	NotificationInterval       string                 `json:"notification_interval"`
	NotificationPeriod         string                 `json:"notification_period"`
	Contacts                   []interface{}          `json:"contacts"`
	Templates                  []interface{}          `json:"use,omitempty"`
	IsVolatile                 string                 `json:"is_volatile,omitempty"`
	InitialState               string                 `json:"initial_state,omitempty"`
	ActiveChecksEnabled        string                 `json:"active_checks_enabled,omitempty"`
	PassiveChecksEnabled       string                 `json:"passive_checks_enabled,omitempty"`
	ObsessOverService          string                 `json:"obsess_over_service,omitempty"`
	CheckFreshness             string                 `json:"check_freshness,omitempty"`
	FreshnessThreshold         string                 `json:"freshness_threshold,omitempty"`
	EventHandler               string                 `json:"event_handler,omitempty"`
	EventHandlerEnabled        string                 `json:"event_handler_enabled,omitempty"`
	LowFlapThreshold           string                 `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold          string                 `json:"high_flap_threshold,omitempty"`
	FlapDetectionEnabled       string                 `json:"flap_detection_enabled,omitempty"`
	FlapDetectionOptions       []interface{}          `json:"flap_detection_options,omitempty"`
	ProcessPerfData            string                 `json:"process_perf_data,omitempty"`
	RetainStatusInformation    string                 `json:"retain_status_information,omitempty"`
	RetainNonStatusInformation string                 `json:"retain_nonstatus_information,omitempty"`
	FirstNotificationDelay     string                 `json:"first_notification_delay,omitempty"`
	NotificationOptions        []interface{}          `json:"notification_options,omitempty"`
	NotificationsEnabled       string                 `json:"notifications_enabled,omitempty"`
	ContactGroups              []interface{}          `json:"contact_groups,omitemptys"`
	Notes                      string                 `json:"notes,omitempty"`
	NotesURL                   string                 `json:"notes_url,omitempty"`
	ActionURL                  string                 `json:"action_url,omitempty"`
	IconImage                  string                 `json:"icon_image,omitempty"`
	IconImageAlt               string                 `json:"icon_image_alt,omitempty"`
	Register                   string                 `json:"register,omitempty"`
	FreeVariables              map[string]interface{} `json:"free_variables,omitempty"`
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
				Description: "The name of the service",
			},
			"host_name": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The hosts that the service should run on",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Defines the description of the service. It may contain spaces, dashes and colons (avoid using semicolons, apostrophes and quotation marks)",
			},
			"check_command": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the command that should be used to check the status of the service",
			},
			"max_check_attempts": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How many times to retry the service check before alerting when the state is anything other than OK",
			},
			"check_interval": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The number of minutes to wait until the next regular check of the service",
			},
			"retry_interval": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The number of minutes to wait until re-checking the service",
			},
			"check_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The time period during which active checks of the service can be made",
			},
			"notification_interval": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How long to wait before sending another notification to a contact that the service is down",
			},
			"notification_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The time period during which notifications can be sent for a service alert",
			},
			"contacts": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The list of users that Nagios should alert when a service is down",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"templates": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of service templates to apply to the service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_volatile": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Determines if the service is 'volatile'. Services typically are not volatile and this should be disabled. This accepts either true or false. The deault value is false",
			},
			"initial_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "By default, Nagios will assume the service are in an OK state. Valid options are: 'd' down, 's' up or 'u' unreachable",
			},
			"active_checks_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Sets whether or not active checks are enabled for the service",
			},
			"passive_checks_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Sets whether or not passive checks are enabled for the service",
			},
			"obsess_over_service": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not Nagios 'obsesses' over the service using the ocsp_command",
			},
			"check_freshness": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether ot not freshness checks are enabled for the service",
			},
			"freshness_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The freshness threshold used for the service",
			},
			"event_handler": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The command that should be run whenever a change in the state of the service is detected",
			},
			"event_handler_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not event handlers should be enabled for the service",
			},
			"low_flap_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The minimum threshold that should be used when detecting if flapping is occurring",
			},
			"high_flap_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The maximum threshold that should be used when detecting if flapping is occurring",
			},
			"flap_detection_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not flap detection is enabled for the service",
			},
			"flap_detection_options": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Determines what flap detection logic will be used for the service. One or more of the following valid options can be provided: 'd' down, 'o' up, or 'u' unreachable",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"process_perf_data": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if Nagios should process performance data",
			},
			"retain_status_information": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not status related information should be kept for the service",
			},
			"retain_nonstatus_information": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not non-status related information should be kept for the service",
			},
			"first_notification_delay": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The amount of time to wait to send out the first notification when a service enters a non-UP state",
			},
			"notification_options": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Determines when Nagios should alert if a host is one or more of the following options: 'o' up, 'd' down, 'u' unreachable, 'r' recovery, 'f' flapping or 's' scheduled downtime",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notifications_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if Nagios should send notifications",
			},
			"contact_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of the contact groups that should be notified if the service goes down",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Notes about the service that may assist with troubleshooting",
			},
			"notes_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL to a third-party documentation repository containing more information about the service",
			},
			"action_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL to a third-party documentation repository containing actions to take in the event the service goes down",
			},
			"icon_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The icon to display in Nagios",
			},
			"icon_image_alt": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The text to display when hovering over the icon_image or the text to display if the icon_image is unavailable",
			},
			"register": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Determines if the host will be marked as active or inactive",
			},
			"free_variables": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "A key/value pair of free variables to add to the service. The key must begin with an underscore.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Create: resourceCreateService,
		Read:   resourceReadService,
		Update: resourceUpdateService,
		Delete: resourceDeleteService,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateService(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	service := setServiceFromSchema(d)

	_, err := nagiosClient.newService(service)

	if err != nil {
		return err
	}

	d.SetId(service.ServiceName)

	return resourceReadService(d, m)
}

func resourceReadService(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	service, err := nagiosClient.getService(d.Get("service_name").(string))

	if err != nil {
		return err
	}

	if service == nil {
		// service not found in Nagios. Update terraform state
		d.SetId("")
		return nil
	}

	setDataFromService(d, service)

	return nil
}

func resourceUpdateService(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	service := setServiceFromSchema(d)

	oldVal, _ := d.GetChange("service_name")
	oldDesc, _ := d.GetChange("description")

	if oldVal == "" { // No change, but perhaps the resource was manually deleted and need to update it so pass in the same name
		oldVal = d.Get("service_name").(string)
	}

	if oldDesc == "" {
		oldDesc = d.Get("description").(string)
	}

	// HTTP PUT for a Nagios service is weirder than the rest. Requires /api/v1/config/service/<service_name>/<service_description>?<rest of url>
	err := nagiosClient.updateService(service, oldVal, oldDesc)

	if err != nil {
		return err
	}

	setDataFromService(d, service)

	return resourceReadService(d, m)
}

func resourceDeleteService(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteService(mapArrayToString(d.Get("host_name").(*schema.Set).List()), d.Get("description").(string))

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func setDataFromService(d *schema.ResourceData, service *Service) error {
	// required attributes
	d.SetId(service.ServiceName)
	d.Set("service_name", service.ServiceName)
	d.Set("host_name", service.HostName)
	d.Set("description", service.Description)
	d.Set("max_check_attempts", service.MaxCheckAttempts)
	d.Set("check_interval", service.CheckInterval)
	d.Set("retry_interval", service.RetryInterval)
	d.Set("check_period", service.CheckPeriod)
	d.Set("notification_interval", service.NotificationInterval)
	d.Set("notification_period", service.NotificationPeriod)
	d.Set("contacts", service.Contacts)

	// optionsl attributes
	if service.CheckCommand != "" {
		d.Set("check_command", service.CheckCommand)
	}
	if service.Templates != nil {
		d.Set("templates", service.Templates)
	}

	if service.IsVolatile != "" {
		d.Set("is_volatile", service.IsVolatile)
	}

	if service.InitialState != "" {
		d.Set("initial_state", service.InitialState)
	}

	if service.ActiveChecksEnabled != "" {
		d.Set("active_checks_enabled", service.ActiveChecksEnabled)
	}

	if service.PassiveChecksEnabled != "" {
		d.Set("passive_checks_enabled", service.PassiveChecksEnabled)
	}

	if service.ObsessOverService != "" {
		d.Set("obsess_over_service", service.ObsessOverService)
	}

	if service.CheckFreshness != "" {
		d.Set("check_freshness", service.CheckFreshness)
	}

	if service.FreshnessThreshold != "" {
		d.Set("freshness_threshold", service.FreshnessThreshold)
	}

	if service.EventHandler != "" {
		d.Set("event_handler", service.EventHandler)
	}

	if service.EventHandlerEnabled != "" {
		d.Set("event_handler_enabled", service.EventHandlerEnabled)
	}

	if service.LowFlapThreshold != "" {
		d.Set("low_flap_threshold", service.LowFlapThreshold)
	}

	if service.HighFlapThreshold != "" {
		d.Set("high_flap_threshold", service.HighFlapThreshold)
	}

	if service.FlapDetectionEnabled != "" {
		d.Set("flap_detection_enabled", service.FlapDetectionEnabled)
	}

	if service.FlapDetectionOptions != nil {
		d.Set("flap_detection_options", service.FlapDetectionOptions)
	}

	if service.ProcessPerfData != "" {
		d.Set("process_perf_data", service.ProcessPerfData)
	}

	if service.RetainStatusInformation != "" {
		d.Set("retain_status_information", service.RetainStatusInformation)
	}

	if service.RetainNonStatusInformation != "" {
		d.Set("retain_nonstatus_information", service.RetainNonStatusInformation)
	}

	if service.FirstNotificationDelay != "" {
		d.Set("first_notification_delay", service.FirstNotificationDelay)
	}

	if service.NotificationOptions != nil {
		d.Set("notification_options", service.NotificationOptions)
	}

	if service.NotificationsEnabled != "" {
		d.Set("notifications_enabled", service.NotificationsEnabled)
	}

	if service.ContactGroups != nil {
		d.Set("contact_groups", service.ContactGroups)
	}

	if service.Notes != "" {
		d.Set("notes", service.Notes)
	}

	if service.NotesURL != "" {
		d.Set("notes_url", service.NotesURL)
	}

	if service.ActionURL != "" {
		d.Set("action_url", service.ActionURL)
	}

	if service.IconImage != "" {
		d.Set("icon_image", service.IconImage)
	}

	if service.IconImageAlt != "" {
		d.Set("icon_image_alt", service.IconImageAlt)
	}

	if service.Register != "" {
		d.Set("register", service.Register)
	}

	if service.FreeVariables != nil {
		if err := d.Set("free_variables", service.FreeVariables); err != nil {
			return fmt.Errorf("Error setting free variables for resource %s: %s", d.Id(), err)
		}
	}

	return nil
}

func setServiceFromSchema(d *schema.ResourceData) *Service {
	service := &Service{
		ServiceName:                d.Get("service_name").(string),
		HostName:                   d.Get("host_name").(*schema.Set).List(),
		Description:                d.Get("description").(string),
		CheckCommand:               d.Get("check_command").(string),
		MaxCheckAttempts:           d.Get("max_check_attempts").(string),
		CheckInterval:              d.Get("check_interval").(string),
		RetryInterval:              d.Get("retry_interval").(string),
		CheckPeriod:                d.Get("check_period").(string),
		NotificationInterval:       d.Get("notification_interval").(string),
		NotificationPeriod:         d.Get("notification_period").(string),
		Contacts:                   d.Get("contacts").(*schema.Set).List(),
		Templates:                  d.Get("templates").(*schema.Set).List(),
		IsVolatile:                 convertBoolToIntToString(d.Get("is_volatile").(bool)),
		InitialState:               d.Get("initial_state").(string),
		ActiveChecksEnabled:        convertBoolToIntToString(d.Get("active_checks_enabled").(bool)),
		PassiveChecksEnabled:       convertBoolToIntToString(d.Get("passive_checks_enabled").(bool)),
		ObsessOverService:          convertBoolToIntToString(d.Get("obsess_over_service").(bool)),
		CheckFreshness:             convertBoolToIntToString(d.Get("check_freshness").(bool)),
		FreshnessThreshold:         d.Get("freshness_threshold").(string),
		EventHandler:               d.Get("event_handler").(string),
		EventHandlerEnabled:        convertBoolToIntToString(d.Get("event_handler_enabled").(bool)),
		LowFlapThreshold:           d.Get("low_flap_threshold").(string),
		HighFlapThreshold:          d.Get("high_flap_threshold").(string),
		FlapDetectionEnabled:       convertBoolToIntToString(d.Get("flap_detection_enabled").(bool)),
		FlapDetectionOptions:       d.Get("flap_detection_options").(*schema.Set).List(),
		ProcessPerfData:            convertBoolToIntToString(d.Get("process_perf_data").(bool)),
		RetainStatusInformation:    convertBoolToIntToString(d.Get("retain_status_information").(bool)),
		RetainNonStatusInformation: convertBoolToIntToString(d.Get("retain_nonstatus_information").(bool)),
		FirstNotificationDelay:     d.Get("first_notification_delay").(string),
		NotificationOptions:        d.Get("notification_options").(*schema.Set).List(),
		NotificationsEnabled:       convertBoolToIntToString(d.Get("notifications_enabled").(bool)),
		ContactGroups:              d.Get("contact_groups").(*schema.Set).List(),
		Notes:                      d.Get("notes").(string),
		NotesURL:                   d.Get("notes_url").(string),
		ActionURL:                  d.Get("action_url").(string),
		IconImage:                  d.Get("icon_image").(string),
		IconImageAlt:               d.Get("icon_image_alt").(string),
		Register:                   convertBoolToIntToString(d.Get("register").(bool)),
		FreeVariables:              d.Get("free_variables").(map[string]interface{}),
	}

	return service
}
