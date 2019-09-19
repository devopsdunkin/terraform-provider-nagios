package nagios

import (
	"log"
	"net/url"
)

// TODO: Need to figure out how most of the funcs should be scoped. Thinking we don't need to expose most of these globally

func (c *Client) NewHost(host *Host) ([]byte, error) {
	nagiosURL, err := c.buildURL("host", "POST", "", "", "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	contactList := mapArrayToString(host.Contacts)
	templatesList := mapArrayToString(host.Templates)

	data := &url.Values{}
	data.Set("host_name", host.Name)
	data.Set("alias", host.Alias)
	data.Set("address", host.Address)
	data.Set("max_check_attempts", host.MaxCheckAttempts)
	data.Set("check_period", host.CheckPeriod)
	data.Set("notification_interval", host.NotificationInterval)
	data.Set("notification_period", host.NotificationPeriod)
	data.Set("contacts", contactList)
	data.Set("use", templatesList)

	body, err := c.post(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error occurred during HTTP POST - %s", err.Error())
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

	log.Printf("[DEBUG] Hostgroup Array - %s", hostArray)

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

	contactList := mapArrayToString(host.Contacts)
	templatesList := mapArrayToString(host.Templates)

	// TODO: Unsure if this should go to buildURL function or not. If we can find a way to pass it in through parameters via an interface
	nagiosURL = nagiosURL + "&host_name=" + host.Name + "&alias=" + host.Alias + "&address=" + host.Address + "&max_check_attempts=" + host.MaxCheckAttempts +
		"&check_period=" + host.CheckPeriod + "&notification_interval=" + host.NotificationInterval +
		"&notification_period=" + host.NotificationPeriod + "&contacts=" + contactList + "&use=" + templatesList

	data := &url.Values{}
	data.Set("host_name", host.Name)
	data.Set("alias", host.Alias)
	data.Set("address", host.Address)
	data.Set("max_check_attempts", host.MaxCheckAttempts)
	data.Set("check_period", host.CheckPeriod)
	data.Set("notification_interval", host.NotificationInterval)
	data.Set("notification_period", host.NotificationPeriod)
	data.Set("contacts", contactList)
	data.Set("use", templatesList)

	_, err = c.put(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP PUT - %s", err.Error())
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

	return body, nil
}
