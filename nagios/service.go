package nagios

import (
	"encoding/json"
	"net/url"
	"strings"
)

// TODO: Need to figure out how most of the funcs should be scoped. Thinking we don't need to expose most of these globally

func (c *Client) newService(service *Service) ([]byte, error) {
	nagiosURL, err := c.buildURL("service", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	data := setURLParams(service)

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
	var serviceArray = []Service{}
	var service Service

	nagiosURL, err := c.buildURL("service", "GET", "config_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("config_name", name)

	body, err := c.get(data.Encode(), nagiosURL)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &serviceArray)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &service.FreeVariables)

	for i := range serviceArray {
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
		service.Register = serviceArray[i].Register
		service.FreeVariables = serviceArray[i].FreeVariables

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

	nagiosURL = nagiosURL + setURLParams(service).Encode()

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
