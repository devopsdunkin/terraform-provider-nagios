package nagios

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

// Service contains all info needed to create a service in Nagios
// Using tag with both JSON and schema because a POST uses URL encoding to send data

// TODO: Need to add in all of the other fields. What we have right now will work for initial testing
type Service struct {
	ServiceName                string        `json:"config_name" schema:"config_name"`
	HostName                   []interface{} `json:"host_name" schema:"host_name"`
	DisplayName                string        `json:"display_name" schema:"display_name"`
	Description                string        `json:"service_description" schema:"service_description"`
	CheckCommand               string        `json:"check_command" schema:"check_command"`
	MaxCheckAttempts           string        `json:"max_check_attempts" schema:"max_check_attempts"`
	CheckInterval              string        `json:"check_interval" schema:"check_interval"`
	RetryInterval              string        `json:"retry_interval" schema:"retry_interval"`
	CheckPeriod                string        `json:"check_period" schema:"check_period"`
	NotificationInterval       string        `json:"notification_interval" schema:"notification_interval"`
	NotificationPeriod         string        `json:"notification_period" schema:"notification_period"`
	Contacts                   []interface{} `json:"contacts" schema:"contacts"`
	Templates                  []interface{} `json:"use" schema:"use"`
	IsVolatile                 string        `json:"is_volatile" schema:"is_volatile"`
	InitialState               string        `json:"initial_state" schema:"initial_state"`
	ActiveChecksEnabled        string        `json:"active_checks_enabled" schema:"active_checks_enabled"`
	PassiveChecksEnabled       string        `json:"passive_checks_enabled" schema:"passive_checks_enabled"`
	ObsessOverService          string        `json:"obsess_over_service" schema:"obsess_over_service"`
	CheckFreshness             string        `json:"check_freshness" schema:"check_freshness"`
	FreshnessThreshold         string        `json:"freshness_threshold" schema:"freshness_threshold"`
	EventHandler               string        `json:"event_handler" schema:"event_handler"`
	EventHandlerEnabled        string        `json:"event_handler_enabled" schema:"event_handler_enabled"`
	LowFlapThreshold           string        `json:"low_flap_threshold" schema:"low_flap_threshold"`
	HighFlapThreshold          string        `json:"high_flap_threshold" schema:"high_flap_threshold"`
	FlapDetectionEnabled       string        `json:"flap_detection_enabled" schema:"flap_detection_enabled"`
	FlapDetectionOptions       []interface{} `json:"flap_detection_options" schema:"flap_detection_options"`
	ProcessPerfData            string        `json:"process_perf_data" schema:"process_perf_data"`
	RetainStatusInformation    string        `json:"retain_status_information" schema:"retain_status_information"`
	RetainNonStatusInformation string        `json:"retain_nonstatus_information" schema:"retain_nonstatus_information"`
	FirstNotificationDelay     string        `json:"first_notification_delay" schema:"first_notification_delay"`
	NotificationOptions        []interface{} `json:"notification_options" schema:"notification_options"`
	NotificationsEnabled       string        `json:"notifications_enabled" schema:"notifications_enabled"`
	ContactGroups              []interface{} `json:"contact_groups" schema:"contact_groups"`
	Notes                      string        `json:"notes" schema:"notes"`
	NotesURL                   string        `json:"notes_url" schema:"notes_url"`
	ActionURL                  string        `json:"action_url" schema:"action_url"`
	IconImage                  string        `json:"icon_image" schema:"icon_image"`
	IconImageAlt               string        `json:"icon_image_alt" schema:"icon_image_alt"`
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
		},
		Create: resourceCreateService,
		Read:   resourceReadService,
		Update: resourceUpdateService,
		Delete: resourceDeleteService,
		// Importer: &schema.ResourceImporter{ // TODO: Need to figure out what is needed here
		// 	State: schema.ImportStatePassthrough,
		// },
	}
}

func resourceCreateService(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	log.Printf("[DEBUG] Right before calling setServiceFromSchema")

	service := setServiceFromSchema(d)
	log.Printf("[DEBUG] Completed setServiceFromSchema")

	// if service.HostName == "" {
	// TODO: Need to add hostgroup membership to schema. Then we will check if hostgroup has been provided or is a member in Nagios
	// }

	body, err := nagiosClient.newService(service)
	log.Printf("[DEBUG] newService completed. Body: %s", body)

	if err != nil {
		return err
	}

	d.SetId(service.ServiceName)

	log.Printf("[DEBUG] Service struct - %s", service)

	return resourceReadService(d, m)
}

// TODO: When no changes are done, it still says "apply complete". Believe it should say "Infrastructure up-to-date"
func resourceReadService(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	service, err := nagiosClient.getService(d.Get("service_name").(string))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] getService, post call. service: %s", service)

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

func setDataFromService(d *schema.ResourceData, service *Service) {
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
}

func setServiceFromSchema(d *schema.ResourceData) *Service {
	log.Printf("[DEBUG] ServiceName: %s", d.Get("service_name").(string))
	log.Printf("[DEBUG] HostName: %s", d.Get("host_name").(*schema.Set).List())
	log.Printf("[DEBUG] Description: %s", d.Get("description").(string))
	log.Printf("[DEBUG] CheckCommand: %s", d.Get("check_command").(string))
	log.Printf("[DEBUG] MaxCheckAttempts: %s", d.Get("max_check_attempts").(string))
	log.Printf("[DEBUG] CheckInterval: %s", d.Get("check_interval").(string))
	log.Printf("[DEBUG] RetryInterval: %s", d.Get("retry_interval").(string))
	log.Printf("[DEBUG] CheckPeriod: %s", d.Get("check_period").(string))
	log.Printf("[DEBUG] NotificationInterval: %s", d.Get("notification_interval").(string))
	log.Printf("[DEBUG] NotificationPeriod: %s", d.Get("notification_period").(string))
	log.Printf("[DEBUG] Contacts: %s", d.Get("contacts").(*schema.Set).List())
	log.Printf("[DEBUG] Templates: %s", d.Get("templates").(*schema.Set).List())
	log.Printf("[DEBUG] IsVolatile: %s", convertBoolToIntToString(d.Get("is_volatile").(bool)))
	log.Printf("[DEBUG] InitialState: %s", d.Get("initial_state").(string))
	log.Printf("[DEBUG] ActiveChecksEnabled: %s", convertBoolToIntToString(d.Get("active_checks_enabled").(bool)))
	log.Printf("[DEBUG] PassiveChecksEnabled: %s", convertBoolToIntToString(d.Get("passive_checks_enabled").(bool)))
	log.Printf("[DEBUG] ObsessOverService: %s", convertBoolToIntToString(d.Get("obsess_over_service").(bool)))
	log.Printf("[DEBUG] CheckFreshness: %s", convertBoolToIntToString(d.Get("check_freshness").(bool)))
	log.Printf("[DEBUG] FreshnessThreshold: %s", d.Get("freshness_threshold").(string))
	log.Printf("[DEBUG] EventHandler: %s", d.Get("event_handler").(string))
	log.Printf("[DEBUG] EventHandlerEnabled: %s", convertBoolToIntToString(d.Get("event_handler_enabled").(bool)))
	log.Printf("[DEBUG] LowFlapThreshold: %s", d.Get("low_flap_threshold").(string))
	log.Printf("[DEBUG] HighFlapThreshold: %s", d.Get("high_flap_threshold").(string))
	log.Printf("[DEBUG] FlapDetectionEnabled: %s", convertBoolToIntToString(d.Get("flap_detection_enabled").(bool)))
	log.Printf("[DEBUG] FlapDetectionOptions: %s", d.Get("flap_detection_options").(*schema.Set).List())
	log.Printf("[DEBUG] ProcessPerfData: %s", convertBoolToIntToString(d.Get("process_perf_data").(bool)))
	log.Printf("[DEBUG] RetainStatusInformation: %s", convertBoolToIntToString(d.Get("retain_status_information").(bool)))
	log.Printf("[DEBUG] RetainNonstatusInformation: %s", convertBoolToIntToString(d.Get("retain_nonstatus_information").(bool)))
	log.Printf("[DEBUG] FirstNotificationDelay: %s", d.Get("first_notification_delay").(string))
	log.Printf("[DEBUG] NotificationOptions: %s", d.Get("notification_options").(*schema.Set).List())
	log.Printf("[DEBUG] NotificationsEnabled: %s", convertBoolToIntToString(d.Get("notifications_enabled").(bool)))
	log.Printf("[DEBUG] ContactGroups: %s", d.Get("contact_groups").(*schema.Set).List())
	log.Printf("[DEBUG] Notes: %s", d.Get("notes").(string))
	log.Printf("[DEBUG] NotesURL: %s", d.Get("notes_url").(string))
	log.Printf("[DEBUG] ActionURL: %s", d.Get("action_url").(string))
	log.Printf("[DEBUG] IconImage: %s", d.Get("icon_image").(string))
	log.Printf("[DEBUG] IconImageAlt: %s", d.Get("icon_image_alt").(string))

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
	}

	return service
}
