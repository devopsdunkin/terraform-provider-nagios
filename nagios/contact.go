package nagios

import (
	"net/url"
	"strings"
)

func (c *Client) newContact(contact *Contact) ([]byte, error) {
	nagiosURL, err := c.buildURL("contact", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	// data := setURLValuesFromContact(contact)
	data := setURLParams(contact)

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

func (c *Client) getContact(name string) (*Contact, error) {
	var contactArray = []Contact{}
	var contact Contact

	nagiosURL, err := c.buildURL("contact", "GET", "contact_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("contact_name", name)

	err = c.get(data, &contactArray, nagiosURL)

	if err != nil {
		return nil, err
	}

	for i, _ := range contactArray {
		contact.ContactName = contactArray[i].ContactName
		contact.HostNotificationsEnabled = contactArray[i].HostNotificationsEnabled
		contact.ServiceNotificationsEnabled = contactArray[i].ServiceNotificationsEnabled
		contact.HostNotificationPeriod = contactArray[i].HostNotificationPeriod
		contact.ServiceNotificationPeriod = contactArray[i].ServiceNotificationPeriod
		contact.HostNotificationOptions = contactArray[i].HostNotificationOptions
		contact.ServiceNotificationOptions = contactArray[i].ServiceNotificationOptions
		contact.HostNotificationCommands = contactArray[i].HostNotificationCommands
		contact.ServiceNotificationCommands = contactArray[i].ServiceNotificationCommands
		contact.Alias = contactArray[i].Alias
		contact.ContactGroups = contactArray[i].ContactGroups
		contact.Templates = contactArray[i].Templates
		contact.Email = contactArray[i].Email
		contact.Pager = contactArray[i].Pager
		contact.Address1 = contactArray[i].Address1
		contact.Address2 = contactArray[i].Address2
		contact.Address3 = contactArray[i].Address3
		contact.CanSubmitCommands = contactArray[i].CanSubmitCommands
		contact.RetainStatusInformation = contactArray[i].RetainStatusInformation
		contact.RetainNonstatusInformation = contactArray[i].RetainNonstatusInformation

		if i > 1 {
			break
		}
	}

	return &contact, nil
}

func (c *Client) updateContact(contact *Contact, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("contact", "PUT", "contact_name", contact.ContactName, oldVal.(string), "")

	if err != nil {
		return err
	}

	// nagiosURL = nagiosURL + setUpdateURLContactParams(nagiosURL, contact)
	nagiosURL = nagiosURL + setURLParams(contact).Encode()

	_, err = c.put(nagiosURL)

	if err != nil {
		// If the error is this specific message, we want to "catch" it
		// and create a new host, then we can proceed on. Otherwise, we
		// can return the error and exit
		if strings.Contains(err.Error(), "Does the contact exist?") {
			c.newContact(contact)
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

func (c *Client) deleteContact(name string) ([]byte, error) {
	nagiosURL, err := c.buildURL("contact", "DELETE", "contact_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("contact_name", name)

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

// func setURLValuesFromContact(contact *Contact) *url.Values {
// 	data := &url.Values{}
// 	data.Set("contact_name", contact.ContactName)
// 	data.Set("host_notifications_enabled", contact.HostNotificationsEnabled)
// 	data.Set("service_notifications_enabled", contact.ServiceNotificationsEnabled)
// 	data.Set("host_notification_period", contact.HostNotificationPeriod)
// 	data.Set("service_notification_period", contact.ServiceNotificationPeriod)
// 	data.Set("host_notification_options", contact.HostNotificationOptions)
// 	data.Set("service_notification_options", contact.ServiceNotificationOptions)
// 	data.Set("host_notification_commands", mapArrayToString(contact.HostNotificationCommands))
// 	data.Set("service_notification_commands", mapArrayToString(contact.ServiceNotificationCommands))

// 	// Optional attributes
// 	if contact.Alias != "" {
// 		data.Set("alias", contact.Alias)
// 	}

// 	if contact.ContactGroups != nil {
// 		data.Set("contact_groups", mapArrayToString(contact.ContactGroups))
// 	}

// 	if contact.Templates != nil {
// 		data.Set("use", mapArrayToString(contact.Templates))
// 	}

// 	if contact.Email != "" {
// 		data.Set("email", contact.Email)
// 	}

// 	if contact.Pager != "" {
// 		data.Set("pager", contact.Pager)
// 	}

// 	if contact.Address1 != "" {
// 		data.Set("address1", contact.Address1)
// 	}

// 	if contact.Address2 != "" {
// 		data.Set("address2", contact.Address2)
// 	}

// 	if contact.Address3 != "" {
// 		data.Set("address3", contact.Address3)
// 	}

// 	if contact.CanSubmitCommands != "" {
// 		data.Set("can_submit_commands", contact.CanSubmitCommands)
// 	}

// 	if contact.RetainStatusInformation != "" {
// 		data.Set("retain_status_information", contact.RetainStatusInformation)
// 	}

// 	if contact.RetainNonstatusInformation != "" {
// 		data.Set("retain_nonstatus_information", contact.RetainNonstatusInformation)
// 	}

// 	return data
// }

// func setUpdateURLContactParams(originalURL string, contact *Contact) string {
// 	var nagiosURL strings.Builder

// 	nagiosURL.WriteString(originalURL)
// 	nagiosURL.WriteString(
// 		"&contact_name=" + contact.ContactName +
// 			"&host_notifications_enabled=" + contact.HostNotificationsEnabled +
// 			"&service_notifications_enabled=" + contact.ServiceNotificationsEnabled +
// 			"&host_notification_period=" + contact.HostNotificationPeriod +
// 			"&service_notification_period=" + contact.ServiceNotificationPeriod +
// 			"&host_notification_options=" + contact.HostNotificationOptions +
// 			"&service_notification_options=" + contact.ServiceNotificationOptions +
// 			"&host_notification_commands=" + mapArrayToString(contact.HostNotificationCommands) +
// 			"&service_notification_commands=" + mapArrayToString(contact.ServiceNotificationCommands))

// 	if contact.Alias != "" {
// 		nagiosURL.WriteString("&alias=" + contact.Alias)
// 	}

// 	if contact.ContactGroups != nil {
// 		nagiosURL.WriteString("&contact_groups=" + mapArrayToString(contact.ContactGroups))
// 	}

// 	if contact.Templates != nil {
// 		nagiosURL.WriteString("&use=" + mapArrayToString(contact.Templates))
// 	}

// 	if contact.Email != "" {
// 		nagiosURL.WriteString("&email=" + contact.Email)
// 	}

// 	if contact.Pager != "" {
// 		nagiosURL.WriteString("&pager=" + contact.Pager)
// 	}

// 	if contact.Address1 != "" {
// 		nagiosURL.WriteString("&address1=" + contact.Address1)
// 	}

// 	if contact.Address2 != "" {
// 		nagiosURL.WriteString("&address2=" + contact.Address2)
// 	}

// 	if contact.Address3 != "" {
// 		nagiosURL.WriteString("&address3=" + contact.Address3)
// 	}

// 	if contact.CanSubmitCommands != "" {
// 		nagiosURL.WriteString("&can_submit_commands=" + contact.CanSubmitCommands)
// 	}

// 	if contact.RetainStatusInformation != "" {
// 		nagiosURL.WriteString("&retain_status_information=" + contact.RetainStatusInformation)
// 	}

// 	if contact.RetainNonstatusInformation != "" {
// 		nagiosURL.WriteString("&retain_nonstatus_information=" + contact.RetainNonstatusInformation)
// 	}

// 	return nagiosURL.String()
// }
