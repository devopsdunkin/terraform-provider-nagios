package nagios

import (
	"encoding/json"
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

	log.Printf("[DEBUG] UpdateHost => Nagios URL - %s", nagiosURL)

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

	updateBody, err := c.put(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP PUT - %s", err.Error())
		return err
	}

	updateResponse := &UpdateResponse{}

	json.Unmarshal(updateBody, &updateResponse)

	// If Nagios returns an error and it contains the specific line of text, then the host was deleted outside of Terraform. We have to create the host again
	// Nagios API returns an error if the host does not exist and we attempt a PUT. Other Nagios objects don't seem vulnerable to this though
	if updateResponse.StatusError != "" && strings.Contains(updateResponse.StatusError, "Does the host exist?") {
		c.NewHost(host)
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
