package nagios

import (
	"log"
	"net/url"
	"strings"
)

type UpdateResponse struct {
	StatusError   string `json:"error"`
	StatusSuccess string `json:"success"`
}

// TODO: Need to figure out how most of the funcs should be scoped. Thinking we don't need to expose most of these globally
func (c *Client) NewHost(host *Host) ([]byte, error) {
	nagiosURL, err := c.buildURL("host", "POST", "", "", "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	// contactList := mapArrayToString(host.Contacts)
	// templatesList := mapArrayToString(host.Templates)
	// contactGroupsList := mapArrayToString(host.ContactGroups)
	// flapDetectionOptionsList := mapArrayToString(host.FlapDetectionOptions)
	// notificationOptionsList := mapArrayToString(host.NotificationOptions)
	// stalkingOptionsList := mapArrayToString(host.StalkingOptions)

	data := setURLValuesFromHost(host)

	// data := &url.Values{}
	// data.Set("host_name", host.Name)
	// data.Set("alias", host.Alias)
	// data.Set("address", host.Address)
	// data.Set("max_check_attempts", host.MaxCheckAttempts)
	// data.Set("check_period", host.CheckPeriod)
	// data.Set("notification_interval", host.NotificationInterval)
	// data.Set("notification_period", host.NotificationPeriod)
	// data.Set("contacts", contactList)
	// data.Set("use", templatesList)
	// data.Set("contact_groups", contactGroupsList)
	// data.Set("notes", host.Notes)
	// data.Set("notes_url", host.NotesURL)
	// data.Set("action_url", host.ActionURL)
	// data.Set("initial_state", host.InitialState)
	// data.Set("retry_interval", host.RetryInterval)
	// data.Set("passive_checks_enabled", host.PassiveChecksEnabled)
	// data.Set("active_checks_enabled", host.ActiveChecksEnabled)
	// data.Set("obsess_over_host", host.ObsessOverHost)
	// data.Set("event_handler", host.EventHandler)
	// data.Set("event_handler_enabled", host.EventHandlerEnabled)
	// data.Set("flap_detection_enabled", host.FlapDetectionEnabled)
	// data.Set("flap_detection_options", flapDetectionOptionsList)
	// data.Set("low_flap_threshold", host.LowFlapThreshold)
	// data.Set("high_flap_threshold", host.HighFlapThreshold)
	// data.Set("process_perf_data", host.ProcessPerfData)
	// data.Set("retain_status_information", host.RetainStatusInformation)
	// data.Set("retain_nonstatus_information", host.RetainNonstatusInformation)
	// data.Set("check_freshness", host.CheckFreshness)
	// data.Set("freshness_threshold", host.FreshnessThreshold)
	// data.Set("first_notification_delay", host.FirstNotificationDelay)
	// data.Set("notification_options", notificationOptionsList)
	// data.Set("notifications_enabled", host.NotificationsEnabled)
	// data.Set("stalking_options", stalkingOptionsList)
	// data.Set("icon_image", host.IconImage)
	// data.Set("icon_image_alt", host.IconImageAlt)
	// data.Set("vrml_image", host.VRMLImage)
	// data.Set("statusmap_image", host.StatusMapImage)
	// data.Set("2d_coords", host.TwoDCoords)
	// data.Set("3d_coords", host.ThreeDCoords)

	body, err := c.post(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error occurred during HTTP POST - %s", err.Error())
		return nil, err
	}

	err = c.applyConfig()

	if err != nil {
		return nil, err
	}

	return body, nil
}

// TODO: Need to refactor get, update and delete to accomodtae contacts being an array
func (c *Client) GetHost(name string) (*Host, error) {
	var hostArray = []Host{}
	var host Host

	nagiosURL, err := c.buildURL("host", "GET", "host_name", name, "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("host_name", name)

	err = c.get(data, &hostArray, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error getting hostgroup from Nagios - %s", err.Error())
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

func (c *Client) UpdateHost(host *Host, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("host", "PUT", "host_name", host.Name, oldVal.(string))

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}

	// contactList := mapArrayToString(host.Contacts)
	// templatesList := mapArrayToString(host.Templates)
	// contactGroupsList := mapArrayToString(host.ContactGroups)
	// flapDetectionOptionsList := mapArrayToString(host.FlapDetectionOptions)
	// notificationOptionsList := mapArrayToString(host.NotificationOptions)
	// stalkingOptionsList := mapArrayToString(host.StalkingOptions)

	// TODO: Unsure if this should go to buildURL function or not. If we can find a way to pass it in through parameters via an interface
	// TODO: Need to build a function that could generate this dynamically based on passing in a struct
	// nagiosURL = nagiosURL + "&host_name=" + host.Name + "&alias=" + host.Alias + "&address=" + host.Address + "&max_check_attempts=" + host.MaxCheckAttempts +
	// 	"&check_period=" + host.CheckPeriod + "&notification_interval=" + host.NotificationInterval +
	// 	"&notification_period=" + host.NotificationPeriod + "&contacts=" + contactList + "&use=" + templatesList + "&contact_groups=" + contactGroupsList +
	// 	"&notes=" + host.Notes + "&notes_url=" + host.NotesURL + "&action_url=" + host.ActionURL + "&initial_state=" + host.InitialState +
	// 	"&retry_interval=" + host.RetryInterval + "&passive_checks_enabled=" + host.PassiveChecksEnabled + "&active_checks_enabled=" +
	// 	host.ActiveChecksEnabled + "&obsess_over_host=" + host.ObsessOverHost + "&event_handler=" + host.EventHandler +
	// 	"&event_handler_enabled=" + host.EventHandlerEnabled + "&flap_detection_enabled=" + host.FlapDetectionEnabled +
	// 	"&flap_detection_options=" + flapDetectionOptionsList + "&low_flap_threshold=" + host.LowFlapThreshold + "&high_flap_threshold=" + host.HighFlapThreshold +
	// 	"&process_perf_data=" + host.ProcessPerfData + "&retain_status_information=" + host.RetainStatusInformation + "&retain_nonstatus_information=" +
	// 	host.RetainNonstatusInformation + "&check_freshness=" + host.CheckFreshness + "&freshness_threshold=" + host.FreshnessThreshold +
	// 	"&first_notification_delay=" + host.FirstNotificationDelay + "&notification_options=" + notificationOptionsList + "&notifications_enabled=" +
	// 	host.NotificationsEnabled + "&stalking_options=" + stalkingOptionsList + "&icon_image=" + host.IconImage + "&icon_image_alt=" +
	// 	host.IconImageAlt + "&vrml_imag=" + host.VRMLImage + "&statusmap_image=" + host.StatusMapImage + "&2d_coords=" + host.TwoDCoords + "&3d_coords=" + host.ThreeDCoords

	nagiosURL = setUpdateURLParams(nagiosURL, host)

	// TODO: Want to dynamically set the URL encoded values based on passing the struct in to another func
	// For an update, we don't see to set any values in url.Values. We just pass in an empty one. The updates are handled through URL params defined above
	data := &url.Values{}

	_, err = c.put(data, nagiosURL)

	if err != nil {
		// If the error is this specific message, we want to "catch" it
		// and create a new host, then we can proceed on. Otherwise, we
		// can return the error and exit
		if strings.Contains(err.Error(), "Does the host exist?") {
			c.NewHost(host)
		} else {
			log.Printf("[ERROR] Error during HTTP PUT - %s", err.Error())
			return err
		}
	}

	err = c.applyConfig()

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteHost(name string) ([]byte, error) {
	nagiosURL, err := c.buildURL("host", "DELETE", "host_name", name, "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("host_name", name)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP DELETE - %s", err.Error())
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

func setUpdateURLParams(originalURL string, host *Host) string {
	var nagiosURL strings.Builder

	nagiosURL.WriteString(originalURL)
	nagiosURL.WriteString("&host_name=" + host.Name + "&alias=" + host.Alias + "&address=" + host.Address + "&max_check_attempts=" + host.MaxCheckAttempts +
		"&check_period=" + host.CheckPeriod + "&notification_interval=" + host.NotificationInterval +
		"&notification_period=" + host.NotificationPeriod + "&contacts=" + mapArrayToString(host.Contacts))

	// Optional attributes
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

	if host.NotificationOptions != nil {
		nagiosURL.WriteString("&notification_options=")
		nagiosURL.WriteString(mapArrayToString(host.NotificationOptions))
	}

	if host.NotificationsEnabled != "" {
		nagiosURL.WriteString("&notifications_enabled=")
		nagiosURL.WriteString(host.NotificationsEnabled)
	}

	if host.StalkingOptions != nil {
		nagiosURL.WriteString("&stalking_options=")
		nagiosURL.WriteString(mapArrayToString(host.StalkingOptions))
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

	nagiosURL.WriteString("&force=1")

	log.Printf("[DEBUG] Performing Update, URL - %s", nagiosURL.String())

	return nagiosURL.String()
}
