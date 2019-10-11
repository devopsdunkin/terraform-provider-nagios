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
	Name                       string        `json:"host_name" schema:"host_name"`
	Address                    string        `json:"address" schema:"address"`
	DisplayName                string        `json:"display_name" schema:"display_name"`
	MaxCheckAttempts           string        `json:"max_check_attempts" schema:"max_check_attempts"`
	CheckPeriod                string        `json:"check_period" schema:"check_period"`
	NotificationInterval       string        `json:"notification_interval" schema:"notification_interval"`
	NotificationPeriod         string        `json:"notification_period" schema:"notification_period"`
	Contacts                   []interface{} `json:"contacts" schema:"contacts"`
	Alias                      string        `json:"alias" schema:"alias"`
	Templates                  []interface{} `json:"use" schema:"use"`
	ContactGroups              []interface{} `json:"contact_groups" schema:"contact_groups"`
	Notes                      string        `json:"notes" schema:"notes"`
	NotesURL                   string        `json:"notes_url" schema:"notes_url"`
	ActionURL                  string        `json:"action_url" schema:"action_url"`
	InitialState               string        `json:"initial_state" schema:"initial_state"`
	RetryInterval              string        `json:"retry_interval" schema:"retry_interval"`
	PassiveChecksEnabled       string        `json:"passive_checks_enabled" schema:"passive_checks_enabled"`
	ActiveChecksEnabled        string        `json:"active_checks_enabled" schema:"active_checks_enabled"`
	ObsessOverHost             string        `json:"obsess_over_host" schema:"obsess_over_host"`
	EventHandler               string        `json:"event_handler" schema:"event_handler"`
	EventHandlerEnabled        string        `json:"event_handler_enabled" schema:"event_handler_enabled"`
	FlapDetectionEnabled       string        `json:"flap_detection_enabled" schema:"flap_detection_enabled"`
	FlapDetectionOptions       []interface{} `json:"flap_detection_options" schema:"flap_detection_options"`
	LowFlapThreshold           string        `json:"low_flap_threshold" schema:"low_flap_threshold"`
	HighFlapThreshold          string        `json:"high_flap_threshold" schema:"high_flap_threshold"`
	ProcessPerfData            string        `json:"process_perf_data" schema:"process_perf_data"`
	RetainStatusInformation    string        `json:"retain_status_information" schema:"retain_status_information"`
	RetainNonstatusInformation string        `json:"retain_nonstatus_information" schema:"retain_nonstatus_information"`
	CheckFreshness             string        `json:"check_freshness" schema:"check_freshness"`
	FreshnessThreshold         string        `json:"freshness_threshold" schema:"freshness_threshold"`
	FirstNotificationDelay     string        `json:"first_notification_delay" schema:"first_notification_delay"`
	NotificationOptions        []interface{} `json:"notification_options" schema:"notification_options"`
	NotificationsEnabled       string        `json:"notifications_enabled" schema:"notifications_enabled"`
	StalkingOptions            []interface{} `json:"stalking_options" schema:"stalking_options"`
	IconImage                  string        `json:"icon_image" schema:"icon_image"`
	IconImageAlt               string        `json:"icon_image_alt" schema:"icon_image_alt"`
	VRMLImage                  string        `json:"vrml_image" schema:"vrml_image"`
	StatusMapImage             string        `json:"statusmap_image" schema:"statusmap_image"`
	TwoDCoords                 string        `json:"2d_coords" schema:"2d_coords"`
	ThreeDCoords               string        `json:"3d_coords" schema:"3d_coords"`
}

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the host",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IP address of the host",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The display name of the host",
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
			"alias": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The alias of the host",
			},
			"templates": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of templates to apply to the host",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"contact_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of contact groups to apply to the host",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     nil,
				Description: "Notes that mmay provide additional operational information about the host",
			},
			"notes_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "URL that may link to additional operational information about the host",
			},
			"action_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "URL that can link to steps to perform in the event of a host alert or event",
			},
			"initial_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The initial state of the host when it is first added to Nagios. It should be one of three values: 'd' (down), 's' (up) or 'u' (unreachable)",
			},
			"retry_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "How frequent Nagios should retry checking a host",
			},
			"passive_checks_enabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},
			"active_checks_enabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},
			"obsess_over_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"event_handler": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"event_handler_enabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Determines if an event handler should be enabled for the host",
			},
			"flap_detection_enabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Determines if flap detection should be enabled for the host",
			},
			"flap_detection_options": { // TODO: Unsure if this should be a list or comma separated value
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "u - unreachable, d - down, o - up (Can be multiple options)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"low_flap_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"high_flap_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"process_perf_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Determines if Nagios should process performance data for the host",
			},
			"retain_status_information": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"retain_nonstatus_information": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"check_freshness": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"freshness_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"first_notification_delay": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"notification_options": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "d - down, u - unreachable, r - recovery, f - flapping, s - scheduled downtime (Can be multiple options)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notifications_enabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"stalking_options": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "d - down, o - up, u - unreachable, N - notification, n - none (Can be multiple options)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"icon_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The icon image to display in Nagios for the host",
			},
			"icon_image_alt": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The alernative text to display when hovering over the icon image, or when the icon image is not available",
			},
			"vrml_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"statusmap_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"2d_coords": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"3d_coords": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
		Create: resourceCreateHost,
		Read:   resourceReadHost,
		Update: resourceUpdateHost,
		Delete: resourceDeleteHost,
		// Importer: &schema.ResourceImporter{ // TODO: Need to figure out what is needed here
		// 	State: schema.ImportStatePassthrough,
		// },
	}
}

func resourceCreateHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	host := setHostFromSchema(d)

	log.Printf("[DEBUG] CreateHost - After calling setHostFromSchema() - %s", host)

	_, err := nagiosClient.NewHost(host)

	if err != nil {
		return err
	}

	d.SetId(host.Name)

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

	setDataFromHost(d, host)

	return nil
}

func resourceUpdateHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	log.Printf("[DEBUG] resourceUpdateHost => name - %s", d.Get("name").(string))

	host := setHostFromSchema(d)

	log.Printf("[DEBUG] Host after setting it with setHostFromSchema - %s", host)

	oldVal, _ := d.GetChange("name")

	if oldVal == "" { // No change, but perhaps the resource was manually deleted and need to update it so pass in the same name
		oldVal = d.Get("name").(string)
		log.Printf("[DEBUG] resourceUpdateHost => oldVal was blank, so should be same name as current - %s", oldVal)
	}

	err := nagiosClient.UpdateHost(host, oldVal)

	if err != nil {
		log.Printf("[ERROR] Error updating host in Nagios - %s", err.Error())
		return err
	}

	log.Printf("[DEBUG] Right before calling setDataFromHost()")

	setDataFromHost(d, host)

	log.Printf("[DEBUG] Right after returning from setDataFromHost()")

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

func setDataFromHost(d *schema.ResourceData, host *Host) {
	// Required attributes
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

	log.Printf("[DEBUG] setDataFromHost() - Right before optional attributes")

	// Optional attributes
	if host.ContactGroups != nil {
		d.Set("contact_groups", host.ContactGroups)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - ContactGroups")
	}

	if host.Notes != "" {
		d.Set("notes", host.Notes)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - Notes - %s", host.Notes)
	}

	if host.NotesURL != "" {
		d.Set("notes_url", host.NotesURL)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - NotesURL - %s", host.NotesURL)
	}

	if host.ActionURL != "" {
		d.Set("action_url", host.ActionURL)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - ActionURL - %s", host.ActionURL)
	}

	if host.InitialState != "" {
		d.Set("initial_state", host.InitialState)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - InitialState - %s", host.InitialState)
	}

	if host.RetryInterval != "" {
		d.Set("retry_interval", host.RetryInterval)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - RetryInterval - %s", host.RetryInterval)
	}

	if host.PassiveChecksEnabled != "" {
		d.Set("passive_checks_enabled", host.PassiveChecksEnabled)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - PassiveChecksEnabled - %s", host.PassiveChecksEnabled)
	}

	if host.ActiveChecksEnabled != "" {
		d.Set("active_checks_enabled", host.ActiveChecksEnabled)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - ActiveChecksEnabled - %s", host.ActiveChecksEnabled)
	}

	if host.ObsessOverHost != "" {
		d.Set("obsess_over_host", host.ObsessOverHost)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - ObsessOverHost - %s", host.ObsessOverHost)
	}

	if host.EventHandler != "" {
		d.Set("event_handler", host.EventHandler)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - EventHandler - %s", host.EventHandler)
	}

	if host.EventHandlerEnabled != "" {
		d.Set("event_handler_enabled", host.EventHandlerEnabled)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - EventHandlerEnabled - %s", host.EventHandlerEnabled)
	}

	if host.FlapDetectionEnabled != "" {
		d.Set("flap_detection_enabled", host.FlapDetectionEnabled)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - FlapDetectionEnabled - %s", host.FlapDetectionEnabled)
	}

	if host.FlapDetectionOptions != nil {
		d.Set("flap_detection_options", host.FlapDetectionOptions)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - FlapDetectionOptions - %s", host.FlapDetectionOptions)
	}

	if host.LowFlapThreshold != "" {
		d.Set("low_flap_threshold", host.LowFlapThreshold)
		log.Printf("[DEBUG] Just seeing if we hit inside if statement - ContactGroups")
	}

	if host.HighFlapThreshold != "" {
		d.Set("high_flap_threshold", host.HighFlapThreshold)
	}

	if host.ProcessPerfData != "" {
		d.Set("process_perf_data", host.ProcessPerfData)
	}

	if host.RetainStatusInformation != "" {
		d.Set("retain_status_information", host.RetainStatusInformation)
	}

	if host.RetainNonstatusInformation != "" {
		d.Set("retain_nonstatus_information", host.RetainNonstatusInformation)
	}

	if host.CheckFreshness != "" {
		d.Set("check_freshness", host.CheckFreshness)
	}

	if host.FreshnessThreshold != "" {
		d.Set("freshness_threshold", host.FreshnessThreshold)
	}

	if host.FirstNotificationDelay != "" {
		d.Set("first_notification_delay", host.FirstNotificationDelay)
	}

	if host.NotificationOptions != nil {
		d.Set("notification_options", host.NotificationOptions)
	}

	if host.NotificationsEnabled != "" {
		d.Set("notifications_enabled", host.NotificationsEnabled)
	}

	if host.StalkingOptions != nil {
		d.Set("stalking_options", host.StalkingOptions)
	}

	if host.IconImage != "" {
		d.Set("icon_image", host.IconImage)
	}

	if host.IconImageAlt != "" {
		d.Set("icon_image_alt", host.IconImageAlt)
	}

	if host.VRMLImage != "" {
		d.Set("vrml_image", host.VRMLImage)
	}

	if host.StatusMapImage != "" {
		d.Set("statusmap_image", host.StatusMapImage)
	}

	if host.TwoDCoords != "" {
		d.Set("2d_coords", host.TwoDCoords)
	}

	if host.ThreeDCoords != "" {
		d.Set("3d_coords", host.ThreeDCoords)
	}
}

func setHostFromSchema(d *schema.ResourceData) *Host {
	host := &Host{
		Name:                       d.Get("name").(string),
		Alias:                      d.Get("alias").(string),
		Address:                    d.Get("address").(string),
		MaxCheckAttempts:           d.Get("max_check_attempts").(string),
		CheckPeriod:                d.Get("check_period").(string),
		NotificationInterval:       d.Get("notification_interval").(string),
		NotificationPeriod:         d.Get("notification_period").(string),
		Contacts:                   d.Get("contacts").(*schema.Set).List(),
		Templates:                  d.Get("templates").(*schema.Set).List(),
		ContactGroups:              d.Get("contact_groups").(*schema.Set).List(),
		Notes:                      d.Get("notes").(string),
		NotesURL:                   d.Get("notes_url").(string),
		ActionURL:                  d.Get("action_url").(string),
		InitialState:               d.Get("initial_state").(string),
		RetryInterval:              d.Get("retry_interval").(string),
		PassiveChecksEnabled:       d.Get("passive_checks_enabled").(string),
		ActiveChecksEnabled:        d.Get("active_checks_enabled").(string),
		ObsessOverHost:             d.Get("obsess_over_host").(string),
		EventHandler:               d.Get("event_handler").(string),
		EventHandlerEnabled:        d.Get("event_handler_enabled").(string),
		FlapDetectionEnabled:       d.Get("flap_detection_enabled").(string),
		FlapDetectionOptions:       d.Get("flap_detection_options").(*schema.Set).List(),
		LowFlapThreshold:           d.Get("low_flap_threshold").(string),
		HighFlapThreshold:          d.Get("high_flap_threshold").(string),
		ProcessPerfData:            d.Get("process_perf_data").(string),
		RetainStatusInformation:    d.Get("retain_status_information").(string),
		RetainNonstatusInformation: d.Get("retain_nonstatus_information").(string),
		CheckFreshness:             d.Get("check_freshness").(string),
		FreshnessThreshold:         d.Get("freshness_threshold").(string),
		FirstNotificationDelay:     d.Get("first_notification_delay").(string),
		NotificationOptions:        d.Get("notification_options").(*schema.Set).List(),
		NotificationsEnabled:       d.Get("notifications_enabled").(string),
		StalkingOptions:            d.Get("stalking_options").(*schema.Set).List(),
		IconImage:                  d.Get("icon_image").(string),
		IconImageAlt:               d.Get("icon_image_alt").(string),
		VRMLImage:                  d.Get("vrml_image").(string),
		StatusMapImage:             d.Get("statusmap_image").(string),
		TwoDCoords:                 d.Get("2d_coords").(string),
		ThreeDCoords:               d.Get("3d_coords").(string),
	}

	return host
}
