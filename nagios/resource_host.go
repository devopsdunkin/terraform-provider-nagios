package nagios

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

// Host contains all info needed to create a host in Nagios
// TODO: Test to see if we need both JSON and schema tags
// Using tag with both JSON and schema because a POST uses URL encoding to send data

// TODO: Need to add in all of the other fields. What we have right now will work for initial testing
type Host struct {
	Name                       string                 `json:"host_name"`
	Address                    string                 `json:"address"`
	DisplayName                string                 `json:"display_name,omitempty"`
	MaxCheckAttempts           string                 `json:"max_check_attempts"`
	CheckPeriod                string                 `json:"check_period"`
	NotificationInterval       string                 `json:"notification_interval"`
	NotificationPeriod         string                 `json:"notification_period"`
	Contacts                   []interface{}          `json:"contacts"`
	Alias                      string                 `json:"alias,omitempty"`
	Templates                  []interface{}          `json:"use,omitempty"`
	CheckCommand               string                 `json:"check_command,omitempty"`
	ContactGroups              []interface{}          `json:"contact_groups,omitempty"`
	Notes                      string                 `json:"notes,omitempty"`
	NotesURL                   string                 `json:"notes_url,omitempty"`
	ActionURL                  string                 `json:"action_url,omitempty"`
	InitialState               string                 `json:"initial_state,omitempty"`
	RetryInterval              string                 `json:"retry_interval,omitempty"`
	PassiveChecksEnabled       string                 `json:"passive_checks_enabled,omitempty"`
	ActiveChecksEnabled        string                 `json:"active_checks_enabled,omitempty"`
	ObsessOverHost             string                 `json:"obsess_over_host,omitempty"`
	EventHandler               string                 `json:"event_handler,omitempty"`
	EventHandlerEnabled        string                 `json:"event_handler_enabled,omitempty"`
	FlapDetectionEnabled       string                 `json:"flap_detection_enabled,omitempty"`
	FlapDetectionOptions       []interface{}          `json:"flap_detection_options,omitempty"`
	LowFlapThreshold           string                 `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold          string                 `json:"high_flap_threshold,omitempty"`
	ProcessPerfData            string                 `json:"process_perf_data,omitempty"`
	RetainStatusInformation    string                 `json:"retain_status_information,omitempty"`
	RetainNonstatusInformation string                 `json:"retain_nonstatus_information,omitempty"`
	CheckFreshness             string                 `json:"check_freshness,omitempty"`
	FreshnessThreshold         string                 `json:"freshness_threshold,omitempty"`
	FirstNotificationDelay     string                 `json:"first_notification_delay,omitempty"`
	NotificationOptions        string                 `json:"notification_options,omitempty"`
	NotificationsEnabled       string                 `json:"notifications_enabled,omitempty"`
	StalkingOptions            string                 `json:"stalking_options,omitempty"`
	IconImage                  string                 `json:"icon_image,omitempty"`
	IconImageAlt               string                 `json:"icon_image_alt,omitempty"`
	VRMLImage                  string                 `json:"vrml_image,omitempty"`
	StatusMapImage             string                 `json:"statusmap_image,omitempty"`
	TwoDCoords                 string                 `json:"2d_coords,omitempty"`
	ThreeDCoords               string                 `json:"3d_coords,omitempty"`
	Register                   string                 `json:"register,omitempty"`
	FreeVariables              map[string]interface{} `json:"free_variables,omitempty"`
}

/*
	For any bool value, we allow the user to provide a true/false value, but you will notice
	that we immediately convert it to its integer form and then to a string. We want to provide
	the user with an easy to use schema, but Nagios wants the data as a one or zero in string format.
	This seemed to be the easiest way to accomplish that and I wanted to note why it was done that way.
*/

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
				Description: "Another name for the host that will be displayed in the web interface. If left blank, the value from `name` will be displayed",
			},
			"max_check_attempts": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How many times to retry the host check before alerting when the state is anything other than OK",
			},
			"check_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The time period during which active checks of the host can be made",
			},
			"notification_interval": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How long to wait before sending another notification to a contact that the host is down",
			},
			"notification_period": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The time period during which notifications can be sent for a host alert",
			},
			"contacts": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The list of users that Nagios should alert when a host is down",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alias": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "A longer name to describe the host",
				ValidateFunc: validation.StringLenBetween(1, 255),
			},
			"templates": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of Nagios templates to apply to the host",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"check_command": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the command that should be used to check if the host is up or down",
			},
			"contact_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of the contact groups that should be notified if the host goes down",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Notes about the host that may assist with troubleshooting",
			},
			"notes_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL to a third-party documentation respoitory containing more information about the host",
			},
			"action_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL to a third-party documentation repository containing actions to take in the event the host goes down",
			},
			"initial_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The state of the host when it is first added to Nagios. Valid options are: 'd' down, 's' up or 'u' unreachable",
			},
			"retry_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "How often should Nagios try to check the host after the initial down alert",
			},
			"passive_checks_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not passive checks are enabled for the host",
			},
			"active_checks_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not active checks are enabled for the host",
			},
			"obsess_over_host": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not Nagios 'obsesses' over the host using the ochp_command",
			},
			"event_handler": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The command that should be run whenver a change in the state of the host is detected",
			},
			"event_handler_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not event handlers should be enabled for the host",
			},
			"flap_detection_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not flap detection is enabled for the host",
			},
			"flap_detection_options": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Determines what flap detection logic will be used for the host. One or more of the following valid options can be provided: 'd' down, 'o' up, or 'u' unreachable.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"process_perf_data": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if Nagios should process performance data",
			},
			"retain_status_information": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not status related information should be kept for the host",
			},
			"retain_nonstatus_information": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not non-status related information should be kept for the host",
			},
			"check_freshness": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether or not freshness checks are enabled for the host",
			},
			"freshness_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The freshness threshold used for the host",
			},
			"first_notification_delay": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The amount of time to wait to send out the first notification when a host enters a non-UP state",
			},
			"notification_options": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Determines when Nagios should alert if a host is one or more of the following option: 'o' up, 'd' down, 'u' unreachable, 'r' recovery, 'f' flapping or 's' scheduled downtime",
			},
			"notifications_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if Nagios should send notifications",
			},
			"stalking_options": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A list of options to determine which states, if any, should be stalked by Nagios. Refer to the [Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/stalking.html) for more information on stalking",
			},
			"icon_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The icon to display in Nagios",
			},
			"icon_image_alt": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The text to display when hovering over the ",
			},
			"vrml_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The image that will be used as a texture map for the specified host",
			},
			"statusmap_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the image that should be used in the statusmap CGI in Nagios",
			},
			"2d_coords": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The coordinates to use when drawing the host in the statusmap CGI",
			},
			"3d_coords": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The coordinates to use when drawing the host in the statuswrl CGI",
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
				Description: "A key/value pair of free variables to add to the host. The key must begin with an underscore.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Create: resourceCreateHost,
		Read:   resourceReadHost,
		Update: resourceUpdateHost,
		Delete: resourceDeleteHost,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

// func validateCognitoSupportedLoginProviders(v interface{}, k string) (ws []string, errors []error) {
// 	value := v.(string)
// 	if len(value) < 1 {
// 		errors = append(errors, fmt.Errorf("%q cannot be less than 1 character", k))
// 	}

// 	if len(value) > 128 {
// 		errors = append(errors, fmt.Errorf("%q cannot be longer than 128 characters", k))
// 	}

// 	if !regexp.MustCompile(`^[\w.;_/-]+$`).MatchString(value) {
// 		errors = append(errors, fmt.Errorf("%q must contain only alphanumeric characters, dots, semicolons, underscores, slashes and hyphens", k))
// 	}

// 	return
// }

func resourceCreateHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	host := getHostSchema(d)

	_, err := nagiosClient.newHost(host)

	if err != nil {
		return err
	}

	d.SetId(host.Name)

	return resourceReadHost(d, m)
}

// TODO: When no changes are done, it still says "apply complete". Believe it should say "Infrastructure up-to-date"
func resourceReadHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	host, err := nagiosClient.getHost(d.Id())

	if err != nil {
		return err
	}

	if host == nil {
		// host not found in Nagios. Update terraform state
		d.SetId("")
		return nil
	}

	err = setDataFromHost(d, host)

	if err != nil {
		return err
	}

	return nil
}

func resourceUpdateHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	host := getHostSchema(d)

	oldVal, _ := d.GetChange("name")

	if oldVal == "" { // No change, but perhaps the resource was manually deleted and need to update it so pass in the same name
		oldVal = d.Get("name").(string)
	}

	err := nagiosClient.updateHost(host, oldVal)

	if err != nil {
		return err
	}

	err = setDataFromHost(d, host)

	if err != nil {
		return err
	}

	return resourceReadHost(d, m)
}

func resourceDeleteHost(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteHost(d.Id())

	if err != nil {
		return err
	}

	// Update Terraform state that we have deleted the resource
	d.SetId("")

	return nil
}

func setDataFromHost(d *schema.ResourceData, host *Host) error {
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

	// Optional attributes
	if host.CheckCommand != "" {
		d.Set("check_command", host.CheckCommand)
	}

	if host.ContactGroups != nil {
		d.Set("contact_groups", host.ContactGroups)
	}

	if host.Notes != "" {
		d.Set("notes", host.Notes)
	}

	if host.NotesURL != "" {
		d.Set("notes_url", host.NotesURL)
	}

	if host.ActionURL != "" {
		d.Set("action_url", host.ActionURL)
	}

	if host.InitialState != "" {
		d.Set("initial_state", host.InitialState)
	}

	if host.RetryInterval != "" {
		d.Set("retry_interval", host.RetryInterval)
	}

	if host.PassiveChecksEnabled != "" {
		d.Set("passive_checks_enabled", host.PassiveChecksEnabled)
	}

	if host.ActiveChecksEnabled != "" {
		d.Set("active_checks_enabled", host.ActiveChecksEnabled)
	}

	if host.ObsessOverHost != "" {
		d.Set("obsess_over_host", host.ObsessOverHost)
	}

	if host.EventHandler != "" {
		d.Set("event_handler", host.EventHandler)
	}

	if host.EventHandlerEnabled != "" {
		d.Set("event_handler_enabled", host.EventHandlerEnabled)
	}

	if host.FlapDetectionEnabled != "" {
		d.Set("flap_detection_enabled", host.FlapDetectionEnabled)
	}

	if host.FlapDetectionOptions != nil {
		d.Set("flap_detection_options", host.FlapDetectionOptions)
	}

	if host.LowFlapThreshold != "" {
		d.Set("low_flap_threshold", host.LowFlapThreshold)
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

	if host.NotificationOptions != "" {
		d.Set("notification_options", host.NotificationOptions)
	}

	if host.NotificationsEnabled != "" {
		d.Set("notifications_enabled", host.NotificationsEnabled)
	}

	if host.StalkingOptions != "" {
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

	if host.Register != "" {
		d.Set("register", host.Register)
	}

	if host.FreeVariables != nil {
		if err := d.Set("free_variables", host.FreeVariables); err != nil {
			return fmt.Errorf("Error setting free variables for resource %s: %s", d.Id(), err)
		}
	}

	return nil
}

// getHostSchema retrieves the values provided from the user in their TF files and sets the Host struct fields to its values
func getHostSchema(d *schema.ResourceData) *Host {
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
		CheckCommand:               d.Get("check_command").(string),
		ContactGroups:              d.Get("contact_groups").(*schema.Set).List(),
		Notes:                      d.Get("notes").(string),
		NotesURL:                   d.Get("notes_url").(string),
		ActionURL:                  d.Get("action_url").(string),
		InitialState:               d.Get("initial_state").(string),
		RetryInterval:              d.Get("retry_interval").(string),
		PassiveChecksEnabled:       convertBoolToIntToString(d.Get("passive_checks_enabled").(bool)),
		ActiveChecksEnabled:        convertBoolToIntToString(d.Get("active_checks_enabled").(bool)),
		ObsessOverHost:             convertBoolToIntToString(d.Get("obsess_over_host").(bool)),
		EventHandler:               d.Get("event_handler").(string),
		EventHandlerEnabled:        convertBoolToIntToString(d.Get("event_handler_enabled").(bool)),
		FlapDetectionEnabled:       convertBoolToIntToString(d.Get("flap_detection_enabled").(bool)),
		FlapDetectionOptions:       d.Get("flap_detection_options").(*schema.Set).List(),
		LowFlapThreshold:           d.Get("low_flap_threshold").(string),
		HighFlapThreshold:          d.Get("high_flap_threshold").(string),
		ProcessPerfData:            convertBoolToIntToString(d.Get("process_perf_data").(bool)),
		RetainStatusInformation:    convertBoolToIntToString(d.Get("retain_status_information").(bool)),
		RetainNonstatusInformation: convertBoolToIntToString(d.Get("retain_nonstatus_information").(bool)),
		CheckFreshness:             convertBoolToIntToString(d.Get("check_freshness").(bool)),
		FreshnessThreshold:         d.Get("freshness_threshold").(string),
		FirstNotificationDelay:     d.Get("first_notification_delay").(string),
		NotificationOptions:        d.Get("notification_options").(string),
		NotificationsEnabled:       convertBoolToIntToString(d.Get("notifications_enabled").(bool)),
		StalkingOptions:            d.Get("stalking_options").(string),
		IconImage:                  d.Get("icon_image").(string),
		IconImageAlt:               d.Get("icon_image_alt").(string),
		VRMLImage:                  d.Get("vrml_image").(string),
		StatusMapImage:             d.Get("statusmap_image").(string),
		TwoDCoords:                 d.Get("2d_coords").(string),
		ThreeDCoords:               d.Get("3d_coords").(string),
		Register:                   convertBoolToIntToString(d.Get("register").(bool)),
		FreeVariables:              d.Get("free_variables").(map[string]interface{}),
	}

	return host
}
