# Resource: Hostgroup

## Overview

This resource manages Nagios hosts. Nagios monitors hosts based off health checks defined by an administrator.

Refer to the object definition for [hosts](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#host) from Nagios to learn more

## Best practices

This resource allows you to configure any attribute on a host, however, we recommend to manually create a host template that contains shared settings for the hosts, and then use the host resource to create your hosts and set host specific attributes. 

## Example

```hcl
resource "nagios_host" "host1" {
    name                            = "host1"
    address                         = "127.0.0.1"
    display_name                    = "host1 - Test server"
    max_check_attempts              = "2"
    check_period                    = "24x7"
    notification_interval           = "5"
    notification_period             = "24x7"
    contacts                        = [
        "nagiosadmin"
    ]
    alias                           = "host1"
    templates                       = "generic-host"
    contact_groups                  = [
        "noc_staff"
    ]
    notes                           = "If this host is down for more than 20 mninutes, page out to the operations on-call"
    notes_url                       = "http://docs.company.com/host_alert"
    action_url                      = "http://docs.company.com/host_alert/actions"
    initial_state                   = "s"
    retry_interval                  = "5"
    passive_checks_enabled          = 1
    active_checks_enabled           = 1
    obsess_over_host                = 0
    process_perf_data               = 1
    notification_options            = [
        "d"
    ]
    notifications_enabled           = 1
}
```

## Arguments

Below is a brief description of what each field is used for in Nagios. Refer to the [official Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html) for more detail

`name` (required): The name of the host
`address` (required): The IP address of the host
`display_name`: Another name for the host that will be displayed in the web interface. If left blank, the value from `name` will be displayed
`max_check_attempts` (required): How many times to retry the host check before alerting when the state is anything other than OK
`check_period` (required): The time period during which active checks of the host can be made
`notification_interval` (required): How long to wait before sending another notification to a contact that the host is down
`notification_period` (requireD): The time period during which notifications can be sent for a host alert
`contacts` (required): The list of users that Nagios should alert when a host is down
`alias`: A longer name to describe the host
`templates`: A list of Nagios templates to apply to the host
`contact_groups`: A list of the contact groups that should be notified if the host goes down
`notes`: Notes about the host that may assist with troubleshooting
`notes_url`: URL to a third-party documentation respoitory containing more information about the host
`action_url`: URL to a third-party documentation repository containing actions to take in the event the host goes down
`initial_state`: The state of the host when it is first added to Nagios
`retry_interval`: How often should Nagios try to check the host after the initial down alert
`passive_checks_enabled`: Sets whether or not passive checks are enabled for the host
`active_checks_enabled`: Sets whether or not active checks are enabled for the host
`obsess_over_host`: Sets whether or not Nagios "obsesses" over the host using the ochp_command
`event_handler`: The command that should be run whenver a change in the state of the host is detected
`event_handler_enabled`: Sets whether or not event handlers should be enabled for the host
`flap_detection_enabled`: Sets whether or not flap detection is enabled for the host
`flap_detection_options`: Determines what flap detection logic will be used for the host. Valid options for this attribute are o = UP, d = DOWN and u = UNREACHABLE
`low_flap_threshold`: The minimum threshold that should be used when detecting if flapping is occurring
`high_flap_threshold`: The maximum threshold that should be used when detecting if flapping is occurring
`process_perf_data`: Determines if Nagios should process performance dat
`retain_status_information`: Sets whether or not status related information should be kept for the host
`retain_nonstatus_information`: Sets whether or not non-status related information should be kept for the host
`check_freshness`: Sets whether or not freshness checks are enabled for the host
`freshness_threshold`: The freshness threshold used for the host
`first_notification_delay`: The amount of time to wait to send out the first notification when a host enters a non-UP state
`notification_options`: A list of notification options. Determines if Nagios should alert if a host is down, up, or unreachable
`notifications_enabled`: Determines if Nagios should send notifications
`stalking_options`: A list of options to determine which states, if any, should be stalked by Nagios. Refer to the [Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/stalking.html) for more information on stalking
`icon_image`: The icon to display in Nagios
`icon_image_alt`:  The text to display when hovering over the `icon_image` or the text to display if the `icon_image` is unavailable
`vrml_image`: The image that will be used as a texture map for the specified host
`statusmap_image`: The name of the image that should be used in the statusmap CGI in Nagios
`2d_coords`: The coordinates to use when drawing the host in the statusmap CGI
`3d_coords`: The coordinates to use when drawing the host in the statuswrl CGI