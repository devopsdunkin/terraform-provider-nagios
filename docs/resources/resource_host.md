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

`name`: The name of the host
`address`: The IP address of the host
`display_name`: The pretty formatted name if you want to use something besides `name`
`max_check_attempts`:
`check_period`: Determines when Nagios should be monitoring the host. 
`notification_interval`: How often should Nagios send notifications about a host
`notification_period`: When should Nagios alert administrators of a down host. An example value of this would be `24x7`
`contacts`: The list of users that Nagios should alert when a host is down
`alias`: Another name for the host if it is known as something else
`templates`: A list of Nagios templates to apply to the host
`contact_groups`: A list of groups that should be notified when the host goes down
`notes`: Notes about the host that may assist with troubleshooting
`notes_url`: URL to a third-party documentation respoitory containing more information about the host
`action_url`: URL to a third-party documentation repository containing actions to take in the event the host goes down
`initial_state`: The state of the host when it is first added to Nagios
`retry_interval`: How often should Nagios try to check the host after the initial down alert
`passive_checks_enabled`:
`active_checks_enabled`:
`obsess_over_host`:
`event_handler`:
`event_handler_enabled`:
`flap_detection_enabled`:
`flap_detection_options`:
`low_flap_threshold`:
`high_flap_threshold`:
`process_perf_data`: Determines if Nagios should process performance dat
`retain_status_information`:
`retain_nonstatus_information`:
`check_freshness`:
`freshness_threshold`:
`first_notification_delay`:
`notification_options`: A list of notification options. Determines if Nagios should alert if a host is down, up, or unreachable
`notifications_enabled`: Determines if Nagios should send notifications
`stalking_options`:
`icon_image`: The icon to display in Nagios
`icon_image_alt`:  The text to display when hovering over the `icon_image` or the text to display if the `icon_image` is unavailable
`vrml_image`:
`statusmap_image`:
`2d_coords`:
`3d_coords`: