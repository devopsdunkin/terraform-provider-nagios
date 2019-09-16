package nagios

import (
	"log"
	"net/url"
)

// NewHostgroup initiates the HTTP POST to the Nagios API to create a hostgroup
func (c *Client) NewHostgroup(hostgroup *Hostgroup) ([]byte, error) {
	nagiosURL, err := c.buildURL("hostgroup", "POST", "", "", "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("hostgroup_name", hostgroup.Name)
	data.Set("alias", hostgroup.Alias)

	body, err := c.post(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error occurred during HTTP POST - %s", err.Error())
		return nil, err
	}

	return body, nil
}

func (c *Client) GetHostgroup(name string) (*Hostgroup, error) {
	var hostgroupArray = []Hostgroup{}
	var hostgroup Hostgroup

	nagiosURL, err := c.buildURL("hostgroup", "GET", "hostgroup_name", name, "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("hostgroup_name", name)

	err = c.get(data, &hostgroupArray, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error getting hostgroup from Nagios - %s", err.Error())
		return nil, err
	}

	log.Printf("[DEBUG] Hostgroup Array - %s", hostgroupArray)

	for i, _ := range hostgroupArray {
		hostgroup.Name = hostgroupArray[i].Name
		hostgroup.Alias = hostgroupArray[i].Alias
		if i > 1 { // Nagios should only return 1 object during a GET with the way we are manipulating it. So only grab the first object and break if we have more than 1
			break
		}
	}
	log.Printf("[DEBUG] GetHostgroup func: hostgroup.Name - %s", hostgroup.Name)
	log.Printf("[DEBUG] GetHostgroup func: hostgroup.Alias - %s", hostgroup.Alias)
	return &hostgroup, nil
}

func (c *Client) UpdateHostgroup(hostgroup *Hostgroup, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("hostgroup", "PUT", "hostgroup_name", hostgroup.Name, oldVal.(string))

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}

	// TODO: Needs migrated to buildURL func
	nagiosURL = nagiosURL + "&hostgroup_name=" + hostgroup.Name + "&alias=" + hostgroup.Alias

	data := &url.Values{}
	data.Set("hostgroup_name", hostgroup.Name)
	data.Set("alias", hostgroup.Alias)

	log.Printf("[DEBUG] hostgroup.Name in UpdateHostgroup func - %s", hostgroup.Name) // TODO: Clean up logging and make it more consistent
	log.Printf("[DEBUG] Value of url.Values (data) - %s", data)

	_, err = c.put(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP PUT - %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) DeleteHostgroup(name string) ([]byte, error) {
	// TODO: Come back to this func. Not sure if implementing correctly
	// Not sure if we should be creating a pointer to hostgroup when deleting
	// Or do we just pass in the name of the hostgroup to delete since it no longer exists?
	// hostgroup := &Hostgroup{}
	nagiosURL, err := c.buildURL("hostgroup", "DELETE", "hostgroup_name", name, "")

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	data := &url.Values{}
	data.Set("hostgroup_name", name)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		log.Printf("[ERROR] Error during HTTP DELETE - %s", err.Error())
		return nil, err
	}

	return body, nil
}
