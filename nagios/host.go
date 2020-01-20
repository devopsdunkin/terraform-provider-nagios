package nagios

import (
	"encoding/json"
	"net/url"
	"strings"
)

// TODO: Need to figure out how most of the funcs should be scoped. Thinking we don't need to expose most of these globally
func (c *Client) newHost(host *Host) ([]byte, error) {
	nagiosURL, err := c.buildURL("host", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	data := setURLParams(host)

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
func (c *Client) getHost(name string) (*Host, error) {
	var hostArray = []Host{}
	var host Host

	nagiosURL, err := c.buildURL("host", "GET", "host_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("host_name", name)

	body, err := c.get(data.Encode(), nagiosURL)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &hostArray)

	if err != nil {
		return nil, err
	}

	// We are not capturing errors here because free vars may not be provided
	// It will throw an error if it isn't set on the Nagios host and it attempts to Unmarshal here
	// TODO: We need to find a better way of checking for this as an optional field
	json.Unmarshal(body, &host.FreeVariables)

	for i := range hostArray {
		host.HostName = hostArray[i].HostName
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
		host.Register = hostArray[i].Register

		if i > 1 { // Nagios should only return 1 object during a GET with the way we are manipulating it. So only grab the first object and break if we have more than 1
			break
		}
	}

	return &host, nil
}

func (c *Client) updateHost(host *Host, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("host", "PUT", "host_name", host.HostName, oldVal.(string), "")

	if err != nil {
		return err
	}

	nagiosURL = nagiosURL + setURLParams(host).Encode()

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
