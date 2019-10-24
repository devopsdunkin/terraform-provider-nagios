package nagios

import (
	"log"
	"net/url"
	"strings"
)

// TODO: Need to figure out how most of the funcs should be scoped. Thinking we don't need to expose most of these globally

func (c *Client) newService(service *Service) ([]byte, error) {
	nagiosURL, err := c.buildURL("service", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	data := setURLValuesFromService(service)

	body, err := c.post(data, nagiosURL)

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
func (c *Client) getService(name string) (*Service, error) {
	log.Printf("[DEBUG] getService, name: %s", name)
	var serviceArray = []Service{}
	var service Service

	nagiosURL, err := c.buildURL("service", "GET", "config_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("config_name", name)

	err = c.get(data, &serviceArray, nagiosURL)

	if err != nil {
		return nil, err
	}

	for i, _ := range serviceArray {
		service.ServiceName = serviceArray[i].ServiceName
		service.HostName = serviceArray[i].HostName
		service.Description = serviceArray[i].Description
		service.CheckCommand = serviceArray[i].CheckCommand
		service.MaxCheckAttempts = serviceArray[i].MaxCheckAttempts
		service.CheckInterval = serviceArray[i].CheckInterval
		service.RetryInterval = serviceArray[i].RetryInterval
		service.CheckPeriod = serviceArray[i].CheckPeriod
		service.NotificationInterval = serviceArray[i].NotificationInterval
		service.NotificationPeriod = serviceArray[i].NotificationPeriod
		service.Contacts = serviceArray[i].Contacts
		service.Templates = serviceArray[i].Templates
		service.IsVolatile = serviceArray[i].IsVolatile
		service.InitialState = serviceArray[i].InitialState
		service.ActiveChecksEnabled = serviceArray[i].ActiveChecksEnabled
		service.PassiveChecksEnabled = serviceArray[i].PassiveChecksEnabled
		service.ObsessOverService = serviceArray[i].ObsessOverService
		service.CheckFreshness = serviceArray[i].CheckFreshness
		service.FreshnessThreshold = serviceArray[i].FreshnessThreshold
		service.EventHandler = serviceArray[i].EventHandler
		service.EventHandlerEnabled = serviceArray[i].EventHandlerEnabled
		service.LowFlapThreshold = serviceArray[i].LowFlapThreshold
		service.HighFlapThreshold = serviceArray[i].HighFlapThreshold
		service.FlapDetectionEnabled = serviceArray[i].FlapDetectionEnabled
		service.FlapDetectionOptions = serviceArray[i].FlapDetectionOptions
		service.ProcessPerfData = serviceArray[i].ProcessPerfData
		service.RetainStatusInformation = serviceArray[i].RetainStatusInformation
		service.RetainNonStatusInformation = serviceArray[i].RetainNonStatusInformation
		service.FirstNotificationDelay = serviceArray[i].FirstNotificationDelay
		service.NotificationOptions = serviceArray[i].NotificationOptions
		service.NotificationsEnabled = serviceArray[i].NotificationsEnabled
		service.ContactGroups = serviceArray[i].ContactGroups
		service.Notes = serviceArray[i].Notes
		service.NotesURL = serviceArray[i].NotesURL
		service.ActionURL = serviceArray[i].ActionURL
		service.IconImage = serviceArray[i].IconImage
		service.IconImageAlt = serviceArray[i].IconImageAlt

		if i > 1 { // Nagios should only return 1 object during a GET with the way we are manipulating it. So only grab the first object and break if we have more than 1
			break
		}
	}

	return &service, nil
}

func (c *Client) updateService(service *Service, oldVal, oldDesc interface{}) error {
	nagiosURL, err := c.buildURL("service", "PUT", "config_name", service.ServiceName, oldVal.(string), oldDesc.(string))

	if err != nil {
		return err
	}

	nagiosURL = setUpdateURLServiceParams(nagiosURL, service)

	_, err = c.put(nagiosURL)

	if err != nil {
		// If the error is this specific message, we want to "catch" it
		// and create a new service, then we can proceed on. Otherwise, we
		// can return the error and exit
		if strings.Contains(err.Error(), "Does the service exist?") {
			c.newService(service)
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

func (c *Client) deleteService(hostName, serviceDescription string) ([]byte, error) {
	nagiosURL, err := c.buildURL("service", "DELETE", "host_name", hostName, "", serviceDescription)

	if err != nil {
		return nil, err
	}

	nagiosURL = nagiosURL + "&service_description=" + strings.Replace(serviceDescription, " ", "%20", -1)

	data := &url.Values{}
	data.Set("host_name", hostName)
	data.Set("service_description", serviceDescription)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func setURLValuesFromService(service *Service) *url.Values {
	// Required attributes
	data := &url.Values{}
	data.Set("config_name", service.ServiceName)
	data.Set("host_name", mapArrayToString(service.HostName))
	data.Set("service_description", service.Description)
	data.Set("max_check_attempts", service.MaxCheckAttempts)
	data.Set("check_interval", service.CheckInterval)
	data.Set("retry_interval", service.RetryInterval)
	data.Set("check_period", service.CheckPeriod)
	data.Set("notification_interval", service.NotificationInterval)
	data.Set("notification_period", service.NotificationPeriod)
	data.Set("contacts", mapArrayToString(service.Contacts))
	data.Set("templates", mapArrayToString(service.Templates))

	// optionsl attributes
	if service.CheckCommand != "" {
		data.Set("check_command", service.CheckCommand)
	}

	if service.Templates != nil {
		data.Set("use", mapArrayToString(service.Templates))
	}

	if service.IsVolatile != "" {
		data.Set("is_volatile", service.IsVolatile)
	}

	if service.InitialState != "" {
		data.Set("initial_state", service.InitialState)
	}

	if service.ActiveChecksEnabled != "" {
		data.Set("active_checks_enabled", service.ActiveChecksEnabled)
	}

	if service.PassiveChecksEnabled != "" {
		data.Set("passive_checks_enabled", service.PassiveChecksEnabled)
	}

	if service.ObsessOverService != "" {
		data.Set("obsess_over_service", service.ObsessOverService)
	}

	if service.CheckFreshness != "" {
		data.Set("check_freshness", service.CheckFreshness)
	}

	if service.FreshnessThreshold != "" {
		data.Set("freshness_threshold", service.FreshnessThreshold)
	}

	if service.EventHandler != "" {
		data.Set("event_handler", service.EventHandler)
	}

	if service.EventHandlerEnabled != "" {
		data.Set("event_handler_enabled", service.EventHandlerEnabled)
	}

	if service.LowFlapThreshold != "" {
		data.Set("low_flap_threshold", service.LowFlapThreshold)
	}

	if service.HighFlapThreshold != "" {
		data.Set("high_flap_threshold", service.HighFlapThreshold)
	}

	if service.FlapDetectionEnabled != "" {
		data.Set("flap_detection_enabled", service.FlapDetectionEnabled)
	}

	if service.FlapDetectionOptions != nil {
		data.Set("flap_detection_options", mapArrayToString(service.FlapDetectionOptions))
	}

	if service.ProcessPerfData != "" {
		data.Set("process_perf_data", service.ProcessPerfData)
	}

	if service.RetainStatusInformation != "" {
		data.Set("retain_status_information", service.RetainStatusInformation)
	}

	if service.RetainNonStatusInformation != "" {
		data.Set("retain_nonstatus_information", service.RetainNonStatusInformation)
	}

	if service.FirstNotificationDelay != "" {
		data.Set("first_notification_delay", service.FirstNotificationDelay)
	}

	if service.NotificationOptions != nil {
		data.Set("notification_options", mapArrayToString(service.NotificationOptions))
	}

	if service.NotificationsEnabled != "" {
		data.Set("notifications_enabled", service.NotificationsEnabled)
	}

	if service.ContactGroups != nil {
		data.Set("contact_groups", mapArrayToString(service.ContactGroups))
	}

	if service.Notes != "" {
		data.Set("notes", service.Notes)
	}

	if service.NotesURL != "" {
		data.Set("notes_url", service.NotesURL)
	}

	if service.ActionURL != "" {
		data.Set("action_url", service.ActionURL)
	}

	if service.IconImage != "" {
		data.Set("icon_image", service.IconImage)
	}

	if service.IconImageAlt != "" {
		data.Set("icon_image_alt", service.IconImageAlt)
	}

	return data
}

func setUpdateURLServiceParams(originalURL string, service *Service) string {
	var nagiosURL strings.Builder

	nagiosURL.WriteString(originalURL)
	nagiosURL.WriteString("&config_name=" + service.ServiceName + "&host_name=" + mapArrayToString(service.HostName) +
		"&max_check_attempts=" + service.MaxCheckAttempts + "&check_period=" + service.CheckPeriod + "&notification_interval=" + service.NotificationInterval +
		"&notification_period=" + service.NotificationPeriod + "&contacts=" + mapArrayToString(service.Contacts))

	// TODO: NEed to reorder these so they match the same order as the setURLValuesFromService func and the struct
	// Optional attributes
	if service.CheckCommand != "" {
		nagiosURL.WriteString("&check_command=")
		nagiosURL.WriteString(service.CheckCommand)
	}

	if service.Templates != nil {
		nagiosURL.WriteString("&use=")
		nagiosURL.WriteString(mapArrayToString(service.Templates))
	}

	if service.IsVolatile != "" {
		nagiosURL.WriteString("&is_volatile=")
		nagiosURL.WriteString(service.IsVolatile)
	}

	if service.InitialState != "" {
		nagiosURL.WriteString("&initial_state=")
		nagiosURL.WriteString(service.InitialState)
	}

	if service.ActiveChecksEnabled != "" {
		nagiosURL.WriteString("&active_checks_enabled=")
		nagiosURL.WriteString(service.ActiveChecksEnabled)
	}

	if service.PassiveChecksEnabled != "" {
		nagiosURL.WriteString("&passive_checks_enabled=")
		nagiosURL.WriteString(service.PassiveChecksEnabled)
	}

	if service.ObsessOverService != "" {
		nagiosURL.WriteString("&obsess_over_service=")
		nagiosURL.WriteString(service.ObsessOverService)
	}

	if service.CheckFreshness != "" {
		nagiosURL.WriteString("&check_freshness=")
		nagiosURL.WriteString(service.CheckFreshness)
	}

	if service.FreshnessThreshold != "" {
		nagiosURL.WriteString("&freshness_threshold=")
		nagiosURL.WriteString(service.FreshnessThreshold)
	}

	if service.EventHandler != "" {
		nagiosURL.WriteString("&event_handler=")
		nagiosURL.WriteString(service.EventHandler)
	}

	if service.EventHandlerEnabled != "" {
		nagiosURL.WriteString("&event_handler_enabled=")
		nagiosURL.WriteString(service.EventHandlerEnabled)
	}

	if service.LowFlapThreshold != "" {
		nagiosURL.WriteString("&low_flap_threshold=")
		nagiosURL.WriteString(service.LowFlapThreshold)
	}

	if service.HighFlapThreshold != "" {
		nagiosURL.WriteString("&high_flap_threshold=")
		nagiosURL.WriteString(service.HighFlapThreshold)
	}

	if service.FlapDetectionEnabled != "" {
		nagiosURL.WriteString("&flap_detection_enabled=")
		nagiosURL.WriteString(service.FlapDetectionEnabled)
	}

	if service.FlapDetectionOptions != nil {
		nagiosURL.WriteString("&flap_detection_options=")
		nagiosURL.WriteString(mapArrayToString(service.FlapDetectionOptions))
	}

	if service.ProcessPerfData != "" {
		nagiosURL.WriteString("&process_perf_data=")
		nagiosURL.WriteString(service.ProcessPerfData)
	}

	if service.RetainStatusInformation != "" {
		nagiosURL.WriteString("&retain_status_information=")
		nagiosURL.WriteString(service.RetainStatusInformation)
	}

	if service.RetainNonStatusInformation != "" {
		nagiosURL.WriteString("&retain_nonstatus_information")
		nagiosURL.WriteString(service.RetainNonStatusInformation)
	}

	if service.FirstNotificationDelay != "" {
		nagiosURL.WriteString("&first_notification_delay=")
		nagiosURL.WriteString(service.FirstNotificationDelay)
	}

	if service.NotificationOptions != nil {
		nagiosURL.WriteString("&notification_options=")
		nagiosURL.WriteString(mapArrayToString(service.NotificationOptions))
	}

	if service.NotificationsEnabled != "" {
		nagiosURL.WriteString("&notifications_enabled=")
		nagiosURL.WriteString(service.NotificationsEnabled)
	}

	if service.ContactGroups != nil {
		nagiosURL.WriteString("&contact_groups=")
		nagiosURL.WriteString(mapArrayToString(service.ContactGroups))
	}

	if service.Notes != "" {
		nagiosURL.WriteString("&notes=")
		nagiosURL.WriteString(service.Notes)
	}

	if service.NotesURL != "" {
		nagiosURL.WriteString("&notes_url=")
		nagiosURL.WriteString(service.NotesURL)
	}

	if service.ActionURL != "" {
		nagiosURL.WriteString("&action_url=")
		nagiosURL.WriteString(service.ActionURL)
	}

	if service.RetryInterval != "" {
		nagiosURL.WriteString("&retry_interval=")
		nagiosURL.WriteString(service.RetryInterval)
	}

	if service.IconImage != "" {
		nagiosURL.WriteString("&icon_image=")
		nagiosURL.WriteString(service.IconImage)
	}

	if service.IconImage != "" {
		nagiosURL.WriteString("&icon_image_alt=")
		nagiosURL.WriteString(service.IconImageAlt)
	}

	return nagiosURL.String()
}
