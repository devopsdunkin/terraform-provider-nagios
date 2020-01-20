package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServiceRead,

		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the service",
			},
			"host_name": {
				Type:        schema.TypeSet,
				Computed:    true,
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
				Computed:    true,
				Description: "The name of the command that should be used to check the status of the service",
			},
			"max_check_attempts": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How many times to retry the service check before alerting when the state is anything other than OK",
			},
			"check_interval": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The number of minutes to wait until the next regular check of the service",
			},
			"retry_interval": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The number of minutes to wait until re-checking the service",
			},
			"check_period": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time period during which active checks of the service can be made",
			},
			"notification_interval": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How long to wait before sending another notification to a contact that the service is down",
			},
			"notification_period": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time period during which notifications can be sent for a service alert",
			},
			"contacts": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The list of users that Nagios should alert when a service is down",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"templates": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A list of service templates to apply to the service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_volatile": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if the service is 'volatile'. Services typically are not volatile and this should be disabled. This accepts either true or false. The deault value is false",
			},
			"initial_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "By default, Nagios will assume the service are in an OK state. Valid options are: 'd' down, 's' up or 'u' unreachable",
			},
			"active_checks_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not active checks are enabled for the service",
			},
			"passive_checks_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not passive checks are enabled for the service",
			},
			"obsess_over_service": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not Nagios 'obsesses' over the service using the ocsp_command",
			},
			"check_freshness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether ot not freshness checks are enabled for the service",
			},
			"freshness_threshold": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The freshness threshold used for the service",
			},
			"event_handler": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The command that should be run whenever a change in the state of the service is detected",
			},
			"event_handler_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not event handlers should be enabled for the service",
			},
			"low_flap_threshold": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The minimum threshold that should be used when detecting if flapping is occurring",
			},
			"high_flap_threshold": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The maximum threshold that should be used when detecting if flapping is occurring",
			},
			"flap_detection_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not flap detection is enabled for the service",
			},
			"flap_detection_options": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Determines what flap detection logic will be used for the service. One or more of the following valid options can be provided: 'd' down, 'o' up, or 'u' unreachable",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"process_perf_data": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if Nagios should process performance data",
			},
			"retain_status_information": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not status related information should be kept for the service",
			},
			"retain_nonstatus_information": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not non-status related information should be kept for the service",
			},
			"first_notification_delay": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The amount of time to wait to send out the first notification when a service enters a non-UP state",
			},
			"notification_options": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Determines when Nagios should alert if a host is one or more of the following options: 'o' up, 'd' down, 'u' unreachable, 'r' recovery, 'f' flapping or 's' scheduled downtime",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notifications_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if Nagios should send notifications",
			},
			"contact_groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A list of the contact groups that should be notified if the service goes down",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Notes about the service that may assist with troubleshooting",
			},
			"notes_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to a third-party documentation repository containing more information about the service",
			},
			"action_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to a third-party documentation repository containing actions to take in the event the service goes down",
			},
			"icon_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The icon to display in Nagios",
			},
			"icon_image_alt": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The text to display when hovering over the icon_image or the text to display if the icon_image is unavailable",
			},
			"register": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if the host will be marked as active or inactive",
			},
			"free_variables": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "A key/value pair of free variables to add to the service. The key must begin with an underscore.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceServiceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	service, err := client.getService(d.Get("service_name").(string))

	if err != nil {
		return err
	}

	setDataFromService(d, service)

	return nil
}
