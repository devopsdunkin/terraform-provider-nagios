package nagios

import (
	"net/url"
)

// NewServicegroup initiates the HTTP POST to the Nagios API to create a servicegroup
func (c *Client) newServicegroup(servicegroup *Servicegroup) ([]byte, error) {
	nagiosURL, err := c.buildURL("servicegroup", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("servicegroup_name", servicegroup.Name)
	data.Set("alias", servicegroup.Alias)
	data.Set("members", mapArrayToString(servicegroup.Members))

	body, err := c.post(data, nagiosURL)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) getServicegroup(name string) (*Servicegroup, error) {
	var servicegroupArray = []Servicegroup{}
	var servicegroup Servicegroup

	nagiosURL, err := c.buildURL("servicegroup", "GET", "servicegroup_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("servicegroup_name", name)

	err = c.get(data, &servicegroupArray, nagiosURL)

	if err != nil {
		return nil, err
	}

	for i, _ := range servicegroupArray {
		servicegroup.Name = servicegroupArray[i].Name
		servicegroup.Alias = servicegroupArray[i].Alias
		servicegroup.Members = servicegroupArray[i].Members
		if i > 1 { // Nagios should only return 1 object during a GET with the way we are manipulating it. So only grab the first object and break if we have more than 1
			break
		}
	}

	return &servicegroup, nil
}

func (c *Client) updateServicegroup(servicegroup *Servicegroup, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("servicegroup", "PUT", "servicegroup_name", servicegroup.Name, oldVal.(string), "")

	if err != nil {
		return err
	}

	// TODO: Needs migrated to buildURL func
	nagiosURL = nagiosURL + "&servicegroup_name=" + servicegroup.Name + "&alias=" + servicegroup.Alias

	data := &url.Values{}
	data.Set("servicegroup_name", servicegroup.Name)
	data.Set("alias", servicegroup.Alias)
	data.Set("members", mapArrayToString(servicegroup.Members))

	_, err = c.put(nagiosURL)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) deleteServicegroup(name string) ([]byte, error) {
	nagiosURL, err := c.buildURL("servicegroup", "DELETE", "servicegroup_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("servicegroup_name", name)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		return nil, err
	}

	return body, nil
}
