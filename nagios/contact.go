package nagios

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
)

func (c *Client) newContact(contact *Contact) ([]byte, error) {
	nagiosURL, err := c.buildURL("contact", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

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

	body, err := c.get(data.Encode(), nagiosURL)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &contactArray)

	if err != nil {
		return nil, err
	}

	for i := range contactArray {
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
	log.Printf("[DEBUG] contact name: %s; oldVal: %s", contact.ContactName, oldVal.(string))
	nagiosURL, err := c.buildURL("contact", "PUT", "contact_name", contact.ContactName, oldVal.(string), "")

	if err != nil {
		log.Printf("[ERROR] Error occurred: %s", err.Error())
		return err
	}

	nagiosURL = nagiosURL + setURLParams(contact).Encode()

	_, err = c.put(nagiosURL)

	if err != nil {
		// If the error is this specific message, we want to "catch" it
		// and create a new host, then we can proceed on. Otherwise, we
		// can return the error and exit
		if strings.Contains(err.Error(), "Does the contact exist?") {
			c.newContact(contact)
		} else {
			log.Printf("[ERROR] Error occurred during put. %s", err.Error())
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
