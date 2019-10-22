package nagios

import (
	"log"
	"net/url"
)

// NewServicegroup initiates the HTTP POST to the Nagios API to create a servicegroup
func (c *Client) NewServicegroup(servicegroup *Servicegroup) ([]byte, error) {
	nagiosURL, err := c.buildURL("servicegroup", "POST", "", "", "", "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("servicegroup_name", servicegroup.Name)
	data.Set("alias", servicegroup.Alias)

	body, err := c.post(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error occurred during HTTP POST - %s", err.Error())
		return nil, err
	}

	return body, nil
}

func (c *Client) GetServicegroup(name string) (*Servicegroup, error) {
	var servicegroupArray = []Servicegroup{}
	var servicegroup Servicegroup

	nagiosURL, err := c.buildURL("servicegroup", "GET", "servicegroup_name", name, "", "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("servicegroup_name", name)

	err = c.get(data, &servicegroupArray, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error getting servicegroup from Nagios - %s", err.Error())
		return nil, err
	}

	log.Printf("[DEBUG] Hostgroup Array - %s", servicegroupArray)

	for i, _ := range servicegroupArray {
		servicegroup.Name = servicegroupArray[i].Name
		servicegroup.Alias = servicegroupArray[i].Alias
		if i > 1 { // Nagios should only return 1 object during a GET with the way we are manipulating it. So only grab the first object and break if we have more than 1
			break
		}
	}
	log.Printf("[DEBUG] GetHostgroup func: servicegroup.Name - %s", servicegroup.Name)
	log.Printf("[DEBUG] GetHostgroup func: servicegroup.Alias - %s", servicegroup.Alias)
	return &servicegroup, nil
}

func (c *Client) UpdateServicegroup(servicegroup *Servicegroup, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("servicegroup", "PUT", "servicegroup_name", servicegroup.Name, oldVal.(string), "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}

	// TODO: Needs migrated to buildURL func
	nagiosURL = nagiosURL + "&servicegroup_name=" + servicegroup.Name + "&alias=" + servicegroup.Alias

	data := &url.Values{}
	data.Set("servicegroup_name", servicegroup.Name)
	data.Set("alias", servicegroup.Alias)

	log.Printf("[DEBUG] servicegroup.Name in UpdateServicegroup func - %s", servicegroup.Name) // TODO: Clean up logging and make it more consistent
	log.Printf("[DEBUG] Value of url.Values (data) - %s", data)

	_, err = c.put(nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP PUT - %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) DeleteServicegroup(name string) ([]byte, error) {
	// TODO: Come back to this func. Not sure if implementing correctly
	// Not sure if we should be creating a pointer to servicegroup when deleting
	// Or do we just pass in the name of the servicegroup to delete since it no longer exists?
	// servicegroup := &Servicegroup{}
	nagiosURL, err := c.buildURL("servicegroup", "DELETE", "servicegroup_name", name, "", "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("servicegroup_name", name)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP DELETE - %s", err.Error())
		return nil, err
	}

	return body, nil
}
