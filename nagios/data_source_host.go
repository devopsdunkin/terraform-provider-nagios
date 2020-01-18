package nagios

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHostRead,

		Schema: map[string]*schema.Schema{
			"host_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address of the host",
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Another name for the host that will be displayed in the web interface. If left blank, the value from `name` will be displayed",
			},
			"max_check_attempts": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How many times to retry the host check before alerting when the state is anything other than OK",
			},
			"check_period": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time period during which active checks of the host can be made",
			},
			"notification_interval": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How long to wait before sending another notification to a contact that the host is down",
			},
			"notification_period": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time period during which notifications can be sent for a host alert",
			},
			"contacts": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The list of users that Nagios should alert when a host is down",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alias": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A longer name to describe the host",
			},
			"templates": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A list of Nagios templates to apply to the host",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"check_command": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the command that should be used to check if the host is up or down",
			},
			"contact_groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A list of the contact groups that should be notified if the host goes down",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Notes about the host that may assist with troubleshooting",
			},
			"notes_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to a third-party documentation respoitory containing more information about the host",
			},
			"action_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to a third-party documentation repository containing actions to take in the event the host goes down",
			},
			"initial_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the host when it is first added to Nagios. Valid options are: 'd' down, 's' up or 'u' unreachable",
			},
			"retry_interval": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How often should Nagios try to check the host after the initial down alert",
			},
			"passive_checks_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not passive checks are enabled for the host",
			},
			"active_checks_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not active checks are enabled for the host",
			},
			"obsess_over_host": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not Nagios 'obsesses' over the host using the ochp_command",
			},
			"event_handler": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The command that should be run whenver a change in the state of the host is detected",
			},
			"event_handler_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not event handlers should be enabled for the host",
			},
			"flap_detection_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not flap detection is enabled for the host",
			},
			"flap_detection_options": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Determines what flap detection logic will be used for the host. One or more of the following valid options can be provided: 'd' down, 'o' up, or 'u' unreachable.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"process_perf_data": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if Nagios should process performance data",
			},
			"retain_status_information": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not status related information should be kept for the host",
			},
			"retain_nonstatus_information": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not non-status related information should be kept for the host",
			},
			"check_freshness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether or not freshness checks are enabled for the host",
			},
			"freshness_threshold": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The freshness threshold used for the host",
			},
			"first_notification_delay": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The amount of time to wait to send out the first notification when a host enters a non-UP state",
			},
			"notification_options": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Determines when Nagios should alert if a host is one or more of the following option: 'o' up, 'd' down, 'u' unreachable, 'r' recovery, 'f' flapping or 's' scheduled downtime",
			},
			"notifications_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if Nagios should send notifications",
			},
			"stalking_options": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A list of options to determine which states, if any, should be stalked by Nagios. Refer to the [Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/stalking.html) for more information on stalking",
			},
			"icon_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The icon to display in Nagios",
			},
			"icon_image_alt": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The text to display when hovering over the ",
			},
			"vrml_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The image that will be used as a texture map for the specified host",
			},
			"statusmap_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the image that should be used in the statusmap CGI in Nagios",
			},
			"2d_coords": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The coordinates to use when drawing the host in the statusmap CGI",
			},
			"3d_coords": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The coordinates to use when drawing the host in the statuswrl CGI",
			},
			"register": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if the host will be marked as active or inactive",
			},
			"free_variables": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "A key/value pair of free variables to add to the host. The key must begin with an underscore.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceHostRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	hostName := d.Get("host_name").(string)

	log.Printf("[DEBUG] hostName = %s", hostName)

	host, err := client.getHost(hostName)

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] host: %s", host)

	setDataFromHost(d, host)

	return nil
}
