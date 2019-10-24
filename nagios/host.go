package nagios

import (
	"fmt"
	"net/url"
	"strings"
)

// TODO: Need to figure out how most of the funcs should be scoped. Thinking we don't need to expose most of these globally
func (c *Client) newHost(host *Host) ([]byte, error) {
	nagiosURL, err := c.buildURL("host", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	data := setURLValuesFromHost(host)

	body, err := c.post(data, nagiosURL)

	if err != nil {
		return nil, err
	}

	err = c.applyConfig()

	if err != nil {
		return nil, err
	}

	err = c.applyConfig()

	if err != nil {
		return nil, err
	}

	return body, nil
}

// TODO: Need to refactor get, update and delete to accomodtae contacts being an array
func (c *Client) getHost(name string) (*Host, error) {
	var hostArray = []Host{}
	var host Host

	nagiosURL, err := c.buildURL("host", "GET", "host_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("host_name", name)

	err = c.get(data, &hostArray, nagiosURL)

	if err != nil {
		return nil, err
	}

	for i, _ := range hostArray {
		host.Name = hostArray[i].Name
		host.Alias = hostArray[i].Alias
		host.Address = hostArray[i].Address
		host.MaxCheckAttempts = hostArray[i].MaxCheckAttempts
		host.CheckPeriod = hostArray[i].CheckPeriod
		host.NotificationInterval = hostArray[i].NotificationInterval
		host.NotificationPeriod = hostArray[i].NotificationPeriod
		host.Contacts = hostArray[i].Contacts
		host.Templates = hostArray[i].Templates
		host.CheckCommand = hostArray[i].CheckCommand
		host.ContactGroups = hostArray[i].ContactGroups
		host.Notes = hostArray[i].Notes
		host.NotesURL = hostArray[i].NotesURL
		host.ActionURL = hostArray[i].ActionURL
		host.InitialState = hostArray[i].InitialState
		host.RetryInterval = hostArray[i].RetryInterval
		host.PassiveChecksEnabled = hostArray[i].PassiveChecksEnabled
		host.ActiveChecksEnabled = hostArray[i].ActiveChecksEnabled
		host.ObsessOverHost = hostArray[i].ObsessOverHost
		host.EventHandler = hostArray[i].EventHandler
		host.EventHandlerEnabled = hostArray[i].EventHandlerEnabled
		host.FlapDetectionEnabled = hostArray[i].FlapDetectionEnabled
		host.FlapDetectionOptions = hostArray[i].FlapDetectionOptions
		host.LowFlapThreshold = hostArray[i].LowFlapThreshold
		host.HighFlapThreshold = hostArray[i].HighFlapThreshold
		host.ProcessPerfData = hostArray[i].ProcessPerfData
		host.RetainStatusInformation = hostArray[i].RetainStatusInformation
		host.RetainNonstatusInformation = hostArray[i].RetainNonstatusInformation
		host.CheckFreshness = hostArray[i].CheckFreshness
		host.FreshnessThreshold = hostArray[i].FreshnessThreshold
		host.FirstNotificationDelay = hostArray[i].FirstNotificationDelay
		host.NotificationOptions = hostArray[i].NotificationOptions
		host.NotificationsEnabled = hostArray[i].NotificationsEnabled
		host.StalkingOptions = hostArray[i].StalkingOptions
		host.IconImage = hostArray[i].IconImage
		host.IconImageAlt = hostArray[i].IconImageAlt
		host.VRMLImage = hostArray[i].VRMLImage
		host.StatusMapImage = hostArray[i].StatusMapImage
		host.TwoDCoords = hostArray[i].TwoDCoords
		host.ThreeDCoords = hostArray[i].ThreeDCoords

		if i > 1 { // Nagios should only return 1 object during a GET with the way we are manipulating it. So only grab the first object and break if we have more than 1
			break
		}
	}

	return &host, nil
}

func (c *Client) updateHost(host *Host, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("host", "PUT", "host_name", host.Name, oldVal.(string), "")

	if err != nil {
		return err
	}

	nagiosURL = setUpdateURLHostParams(nagiosURL, host)

	_, err = c.put(nagiosURL)

	if err != nil {
		// If the error is this specific message, we want to "catch" it
		// and create a new host, then we can proceed on. Otherwise, we
		// can return the error and exit
		if strings.Contains(err.Error(), "Does the host exist?") {
			c.newHost(host)
		} else {
			return err
		}
	}

	err = c.applyConfig()

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) deleteHost(name string) ([]byte, error) {
	nagiosURL, err := c.buildURL("host", "DELETE", "host_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("host_name", name)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		return nil, err
	}

	err = c.applyConfig()

	if err != nil {
		return nil, err
	}

	return body, nil
}

func setURLValuesFromHost(host *Host) *url.Values {
	// Required attributes
	data := &url.Values{}
	data.Set("host_name", host.Name)
	data.Set("address", host.Address)
	data.Set("max_check_attempts", host.MaxCheckAttempts)
	data.Set("check_period", host.CheckPeriod)
	data.Set("notification_interval", host.NotificationInterval)
	data.Set("notification_period", host.NotificationPeriod)
	data.Set("contacts", mapArrayToString(host.Contacts))

	// Optional attributes
	if host.Templates != nil {
		data.Set("use", mapArrayToString(host.Templates))
	}
	if host.Alias != "" {
		data.Set("alias", host.Alias)
	}
	if host.CheckCommand != "" {
		data.Set("check_command", host.CheckCommand)
	}
	if host.ContactGroups != nil {
		data.Set("contact_groups", mapArrayToString(host.ContactGroups))
	}

	if host.Notes != "" {
		data.Set("notes", host.Notes)
	}

	if host.NotesURL != "" {
		data.Set("notes_url", host.NotesURL)
	}

	if host.ActionURL != "" {
		data.Set("action_url", host.ActionURL)
	}

	if host.InitialState != "" {
		data.Set("initial_state", host.InitialState)
	}

	if host.RetryInterval != "" {
		data.Set("retry_interval", host.RetryInterval)
	}

	if host.PassiveChecksEnabled != "" {
		data.Set("passive_checks_enabled", host.PassiveChecksEnabled)
	}

	if host.ActiveChecksEnabled != "" {
		data.Set("active_checks_enabled", host.ActiveChecksEnabled)
	}

	if host.ObsessOverHost != "" {
		data.Set("obsess_over_host", host.ObsessOverHost)
	}

	if host.EventHandler != "" {
		data.Set("event_handler", host.EventHandler)
	}

	if host.EventHandlerEnabled != "" {
		data.Set("event_handler_enabled", host.EventHandlerEnabled)
	}

	if host.FlapDetectionEnabled != "" {
		data.Set("flap_detection_enabled", host.FlapDetectionEnabled)
	}

	if host.FlapDetectionOptions != nil {
		data.Set("flap_detection_options", mapArrayToString(host.FlapDetectionOptions))
	}

	if host.LowFlapThreshold != "" {
		data.Set("low_flap_threshold", host.LowFlapThreshold)
	}

	if host.HighFlapThreshold != "" {
		data.Set("high_flap_threshold", host.HighFlapThreshold)
	}

	if host.ProcessPerfData != "" {
		data.Set("process_perf_data", host.ProcessPerfData)
	}

	if host.RetainStatusInformation != "" {
		data.Set("retain_status_information", host.RetainStatusInformation)
	}

	if host.RetainNonstatusInformation != "" {
		data.Set("retain_nonstatus_information", host.RetainNonstatusInformation)
	}

	if host.CheckFreshness != "" {
		data.Set("check_freshness", host.CheckFreshness)
	}

	if host.FreshnessThreshold != "" {
		data.Set("freshness_threshold", host.FreshnessThreshold)
	}

	if host.FirstNotificationDelay != "" {
		data.Set("first_notification_delay", host.FirstNotificationDelay)
	}

	if host.NotificationOptions != "" {
		data.Set("notification_options", host.NotificationOptions)
	}

	if host.NotificationsEnabled != "" {
		data.Set("notifications_enabled", host.NotificationsEnabled)
	}

	if host.StalkingOptions != "" {
		data.Set("stalking_options", host.StalkingOptions)
	}

	if host.IconImage != "" {
		data.Set("icon_image", host.IconImage)
	}

	if host.IconImageAlt != "" {
		data.Set("icon_image_alt", host.IconImageAlt)
	}

	if host.VRMLImage != "" {
		data.Set("vrml_image", host.VRMLImage)
	}

	if host.StatusMapImage != "" {
		data.Set("statusmap_image", host.StatusMapImage)
	}

	if host.TwoDCoords != "" {
		data.Set("2d_coords", host.TwoDCoords)
	}

	if host.ThreeDCoords != "" {
		data.Set("3d_coords", host.ThreeDCoords)
	}

	return data
}

func setUpdateURLHostParams(originalURL string, host *Host) string {
	var nagiosURL strings.Builder

	nagiosURL.WriteString(originalURL)
	nagiosURL.WriteString("&host_name=" + host.Name + "&alias=" + host.Alias + "&address=" + host.Address + "&max_check_attempts=" + host.MaxCheckAttempts +
		"&check_period=" + host.CheckPeriod + "&notification_interval=" + host.NotificationInterval +
		"&notification_period=" + host.NotificationPeriod + "&contacts=" + mapArrayToString(host.Contacts))

	// Optional attributes
	if host.Templates != nil {
		nagiosURL.WriteString("&use=")
		nagiosURL.WriteString(mapArrayToString(host.Templates))
	}
	if host.CheckCommand != "" {
		nagiosURL.WriteString("&check_command=")
		nagiosURL.WriteString(fmt.Sprint(host.CheckCommand))
	}

	if host.ContactGroups != nil {
		nagiosURL.WriteString("&contact_groups=")
		nagiosURL.WriteString(mapArrayToString(host.ContactGroups))
	}

	if host.Notes != "" {
		nagiosURL.WriteString("&notes=")
		nagiosURL.WriteString(host.Notes)
	}

	if host.NotesURL != "" {
		nagiosURL.WriteString("&notes_url=")
		nagiosURL.WriteString(host.NotesURL)
	}
  
	if host.ActionURL != "" {
		nagiosURL.WriteString("&action_url=")
		nagiosURL.WriteString(host.ActionURL)
	}

	if host.InitialState != "" {
		nagiosURL.WriteString("&initial_state=")
		nagiosURL.WriteString(host.InitialState)
	}

	if host.RetryInterval != "" {
		nagiosURL.WriteString("&retry_interval=")
		nagiosURL.WriteString(host.RetryInterval)
	}

	if host.PassiveChecksEnabled != "" {
		nagiosURL.WriteString("&passive_checks_enabled=")
		nagiosURL.WriteString(host.PassiveChecksEnabled)
	}

	if host.ActiveChecksEnabled != "" {
		nagiosURL.WriteString("&active_checks_enabled=")
		nagiosURL.WriteString(host.ActiveChecksEnabled)
	}

	if host.ObsessOverHost != "" {
		nagiosURL.WriteString("&obsess_over_host=")
		nagiosURL.WriteString(host.ObsessOverHost)
	}

	if host.EventHandler != "" {
		nagiosURL.WriteString("&event_handler=")
		nagiosURL.WriteString(host.EventHandler)
	}

	if host.EventHandlerEnabled != "" {
		nagiosURL.WriteString("&event_handler_enabled=")
		nagiosURL.WriteString(host.EventHandlerEnabled)
	}

	if host.FlapDetectionEnabled != "" {
		nagiosURL.WriteString("&flap_detection_enabled=")
		nagiosURL.WriteString(host.FlapDetectionEnabled)
	}

	if host.FlapDetectionOptions != nil {
		nagiosURL.WriteString("&flap_detection_options=")
		nagiosURL.WriteString(mapArrayToString(host.FlapDetectionOptions))
	}

	if host.LowFlapThreshold != "" {
		nagiosURL.WriteString("&low_flap_threshold=")
		nagiosURL.WriteString(host.LowFlapThreshold)
	}

	if host.HighFlapThreshold != "" {
		nagiosURL.WriteString("&high_flap_threshold=")
		nagiosURL.WriteString(host.HighFlapThreshold)
	}

	if host.ProcessPerfData != "" {
		nagiosURL.WriteString("&process_perf_data=")
		nagiosURL.WriteString(host.ProcessPerfData)
	}

	if host.RetainStatusInformation != "" {
		nagiosURL.WriteString("&retain_status_information=")
		nagiosURL.WriteString(host.RetainStatusInformation)
	}

	if host.RetainNonstatusInformation != "" {
		nagiosURL.WriteString("&retain_nonstatus_information")
		nagiosURL.WriteString(host.RetainNonstatusInformation)
	}

	if host.CheckFreshness != "" {
		nagiosURL.WriteString("&check_freshness=")
		nagiosURL.WriteString(host.CheckFreshness)
	}

	if host.FreshnessThreshold != "" {
		nagiosURL.WriteString("&freshness_threshold=")
		nagiosURL.WriteString(host.FreshnessThreshold)
	}

	if host.FirstNotificationDelay != "" {
		nagiosURL.WriteString("&first_notification_delay=")
		nagiosURL.WriteString(host.FirstNotificationDelay)
	}

	if host.NotificationOptions != "" {
		nagiosURL.WriteString("&notification_options=")
		nagiosURL.WriteString(host.NotificationOptions)
	}

	if host.NotificationsEnabled != "" {
		nagiosURL.WriteString("&notifications_enabled=")
		nagiosURL.WriteString(host.NotificationsEnabled)
	}

	if host.StalkingOptions != "" {
		nagiosURL.WriteString("&stalking_options=")
		nagiosURL.WriteString(host.StalkingOptions)
	}

	if host.IconImage != "" {
		nagiosURL.WriteString("&icon_image=")
		nagiosURL.WriteString(host.IconImage)
	}

	if host.IconImage != "" {
		nagiosURL.WriteString("&icon_image_alt=")
		nagiosURL.WriteString(host.IconImageAlt)
	}

	if host.VRMLImage != "" {
		nagiosURL.WriteString("&vrml_image=")
		nagiosURL.WriteString(host.VRMLImage)
	}

	if host.StatusMapImage != "" {
		nagiosURL.WriteString("&statusmap_image=")
		nagiosURL.WriteString(host.StatusMapImage)
	}

	if host.TwoDCoords != "" {
		nagiosURL.WriteString("&2d_coords=")
		nagiosURL.WriteString(host.TwoDCoords)
	}

	if host.ThreeDCoords != "" {
		nagiosURL.WriteString("&3d_coords=")
		nagiosURL.WriteString(host.ThreeDCoords)
	}

	return nagiosURL.String()
}

func setURLValuesFromHost(host *Host) *url.Values {
	// Required attributes
	data := &url.Values{}
	data.Set("host_name", host.Name)
	data.Set("address", host.Address)
	data.Set("max_check_attempts", host.MaxCheckAttempts)
	data.Set("check_period", host.CheckPeriod)
	data.Set("notification_interval", host.NotificationInterval)
	data.Set("notification_period", host.NotificationPeriod)
	data.Set("contacts", mapArrayToString(host.Contacts))
	data.Set("templates", mapArrayToString(host.Templates))

	// Optional attributes
	if host.Alias != "" {
		data.Set("alias", host.Alias)
	}
	if host.ContactGroups != nil {
		data.Set("contact_groups", mapArrayToString(host.ContactGroups))
	}

	if host.Notes != "" {
		data.Set("notes", host.Notes)
	}

	if host.NotesURL != "" {
		data.Set("notes_url", host.NotesURL)
	}

	if host.ActionURL != "" {
		data.Set("action_url", host.ActionURL)
	}

	if host.InitialState != "" {
		data.Set("initial_state", host.InitialState)
	}

	if host.RetryInterval != "" {
		data.Set("retry_interval", host.RetryInterval)
	}

	if host.PassiveChecksEnabled != "" {
		data.Set("passive_checks_enabled", host.PassiveChecksEnabled)
	}

	if host.ActiveChecksEnabled != "" {
		data.Set("active_checks_enabled", host.ActiveChecksEnabled)
	}

	if host.ObsessOverHost != "" {
		data.Set("obsess_over_host", host.ObsessOverHost)
	}

	if host.EventHandler != "" {
		data.Set("event_handler", host.EventHandler)
	}

	if host.EventHandlerEnabled != "" {
		data.Set("event_handler_enabled", host.EventHandlerEnabled)
	}

	if host.FlapDetectionEnabled != "" {
		data.Set("flap_detection_enabled", host.FlapDetectionEnabled)
	}

	if host.FlapDetectionOptions != nil {
		data.Set("flap_detection_options", mapArrayToString(host.FlapDetectionOptions))
	}

	if host.LowFlapThreshold != "" {
		data.Set("low_flap_threshold", host.LowFlapThreshold)
	}

	if host.HighFlapThreshold != "" {
		data.Set("high_flap_threshold", host.HighFlapThreshold)
	}

	if host.ProcessPerfData != "" {
		data.Set("process_perf_data", host.ProcessPerfData)
	}

	if host.RetainStatusInformation != "" {
		data.Set("retain_status_information", host.RetainStatusInformation)
	}

	if host.RetainNonstatusInformation != "" {
		data.Set("retain_nonstatus_information", host.RetainNonstatusInformation)
	}

	if host.CheckFreshness != "" {
		data.Set("check_freshness", host.CheckFreshness)
	}

	if host.FreshnessThreshold != "" {
		data.Set("freshness_threshold", host.FreshnessThreshold)
	}

	if host.FirstNotificationDelay != "" {
		data.Set("first_notification_delay", host.FirstNotificationDelay)
	}

	if host.NotificationOptions != nil {
		data.Set("notification_options", mapArrayToString(host.NotificationOptions))
	}

	if host.NotificationsEnabled != "" {
		data.Set("notifications_enabled", host.NotificationsEnabled)
	}

	if host.StalkingOptions != nil {
		data.Set("stalking_options", mapArrayToString(host.StalkingOptions))
	}

	if host.IconImage != "" {
		data.Set("icon_image", host.IconImage)
	}

	if host.IconImageAlt != "" {
		data.Set("icon_image_alt", host.IconImageAlt)
	}

	if host.VRMLImage != "" {
		data.Set("vrml_image", host.VRMLImage)
	}

	if host.StatusMapImage != "" {
		data.Set("statusmap_image", host.StatusMapImage)
	}

	if host.TwoDCoords != "" {
		data.Set("2d_coords", host.TwoDCoords)
	}

	if host.ThreeDCoords != "" {
		data.Set("3d_coords", host.ThreeDCoords)
	}

	return data
}