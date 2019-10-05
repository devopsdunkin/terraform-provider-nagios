package nagios

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/devopsdunkin/terraform-provider-nagios/utils"
)

// TODO: Need to figure out how most of the funcs should be scoped. Thinking we don't need to expose most of these globally

func (c *Client) newService(service *Service) ([]byte, error) {
	nagiosURL, err := c.buildURL("service", "POST", "", "", "", "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	contactList := mapArrayToString(service.Contacts)
	templatesList := mapArrayToString(service.Templates)
	hostNameList := mapArrayToString(service.HostName)

	data := &url.Values{}
	data.Set("config_name", service.ServiceName)
	data.Set("host_name", hostNameList)
	data.Set("service_description", service.Description)
	data.Set("check_command", service.CheckCommand)
	data.Set("max_check_attempts", service.MaxCheckAttempts)
	data.Set("check_interval", service.CheckInterval)
	data.Set("retry_interval", service.RetryInterval)
	data.Set("check_period", service.CheckPeriod)
	data.Set("notification_interval", service.NotificationInterval)
	data.Set("notification_period", service.NotificationPeriod)
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
func (c *Client) getService(name string) (*Service, error) {
	var serviceArray = []Service{}
	var service Service

	nagiosURL, err := c.buildURL("service", "GET", "config_name", "", name, "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("config_name", name)

	log.Printf("[DEBUG] Right before c.get; URL encoded data - %s", data)

	err = c.get(data, &serviceArray, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error getting service from Nagios - %s", err.Error())
		return nil, err
	}

	log.Printf("[DEBUG] About to run for loop and grab data from serviceArray")

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

		if i > 1 { // Nagios should only return 1 object during a GET with the way we are manipulating it. So only grab the first object and break if we have more than 1
			break
		}
	}

	log.Printf("[DEBUG] Made it through getService")

	return &service, nil
}

func (c *Client) updateService(service *Service, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("service", "PUT", "config_name", service.Description, service.ServiceName, oldVal.(string))

	log.Printf("[DEBUG] updateService => nagios URL - %s", nagiosURL)

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}

	contactList := mapArrayToString(service.Contacts)
	templatesList := mapArrayToString(service.Templates)
	hostNameList := mapArrayToString(service.HostName)

	// TODO: Unsure if this should go to buildURL function or not. If we can find a way to pass it in through parameters via an interface
	// TODO: Nagios does their updates in a weird weay. Found that doing the URL encoded values doesn't do anything for
	// updates. Need to specify by generating the URL string. Need to determine if both are needed or just URL string as below
	nagiosURL = nagiosURL + "&config_name=" + service.ServiceName + "&host_name=" + hostNameList + "&service_description=" + service.Description + "&check_command=" + service.CheckCommand + "&max_check_attempts=" + service.MaxCheckAttempts +
		"&check_interval=" + service.CheckInterval + "&retry_interval=" + service.RetryInterval + "&check_period=" + service.CheckPeriod + "&notification_interval=" + service.NotificationInterval +
		"&notification_period=" + service.NotificationPeriod + "&contacts=" + contactList + "&use=" + templatesList

	log.Printf("[DEBUG] updateService => Nagios URL - %s", nagiosURL)

	data := &url.Values{}
	data.Set("config_name", service.ServiceName)
	data.Set("host_name", hostNameList)
	data.Set("service_description", service.Description)
	data.Set("check_command", service.CheckCommand)
	data.Set("max_check_attempts", service.MaxCheckAttempts)
	data.Set("check_interval", service.CheckInterval)
	data.Set("retry_interval", service.RetryInterval)
	data.Set("check_period", service.CheckPeriod)
	data.Set("notification_interval", service.NotificationInterval)
	data.Set("notification_period", service.NotificationPeriod)
	data.Set("contacts", contactList)
	data.Set("use", templatesList)

	log.Printf("[DEBUG] updateService => data - %s", data)

	updateBody, err := c.put(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP PUT - %s", err.Error())
		return err
	}

	log.Printf("[DEBUG] updateService => Right before calling updateServiceResponse")
	// updateServiceResponse := &UpdateServiceResponse{}
	updateServiceResponse := utils.ResponseStatus{}

	log.Printf("[DEBUG] updateService => Right before unmarshalling body into updateServiceResponse")
	json.Unmarshal(updateBody, &updateServiceResponse)

	// If Nagios returns an error and it contains the specific line of text, then the service was deleted outside of Terraform. We have to create the service again
	// Nagios API returns an error if the service does not exist and we attempt a PUT. Other Nagios objects don't seem vulnerable to this though
	if updateServiceResponse.StatusError != "" && strings.Contains(updateServiceResponse.StatusError, "Does the service exist?") {
		c.newService(service)
	}

	log.Printf("[DEBUG] Made it through updateService")

	return nil
}

func (c *Client) deleteService(hostName, serviceDescription string) ([]byte, error) {
	nagiosURL, err := c.buildURL("service", "DELETE", "host_name", "", hostName, "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	nagiosURL = nagiosURL + "&service_description=" + serviceDescription

	data := &url.Values{}
	data.Set("host_name", hostName)
	data.Set("service_description", serviceDescription)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP DELETE - %s", err.Error())
		return nil, err
	}

	return body, nil
}
