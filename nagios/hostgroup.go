package nagios

import (
	"net/url"
)

// NewHostgroup initiates the HTTP POST to the Nagios API to create a hostgroup
func (c *Client) newHostgroup(hostgroup *Hostgroup) ([]byte, error) {
	nagiosURL, err := c.buildURL("hostgroup", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	hostGroupMemberList := mapArrayToString(hostgroup.Members)

	data := &url.Values{}
	data.Set("hostgroup_name", hostgroup.Name)
	data.Set("alias", hostgroup.Alias)
	data.Set("members", hostGroupMemberList)

	body, err := c.post(data, nagiosURL)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) getHostgroup(name string) (*Hostgroup, error) {
	var hostgroupArray = []Hostgroup{}
	var hostgroup Hostgroup

	nagiosURL, err := c.buildURL("hostgroup", "GET", "hostgroup_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("hostgroup_name", name)

	err = c.get(data, &hostgroupArray, nagiosURL)

	if err != nil {
		return nil, err
	}

	for i, _ := range hostgroupArray {
		hostgroup.Name = hostgroupArray[i].Name
		hostgroup.Alias = hostgroupArray[i].Alias
		hostgroup.Members = hostgroupArray[i].Members
		if i > 1 { // Nagios should only return 1 object during a GET with the way we are manipulating it. So only grab the first object and break if we have more than 1
			break
		}
	}

	return &hostgroup, nil
}

func (c *Client) updateHostgroup(hostgroup *Hostgroup, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("hostgroup", "PUT", "hostgroup_name", hostgroup.Name, oldVal.(string), "")

	if err != nil {
		return err
	}

	hostGroupMemberList := mapArrayToString(hostgroup.Members)

	// TODO: Needs migrated to buildURL func
	nagiosURL = nagiosURL + "&hostgroup_name=" + hostgroup.Name + "&alias=" + hostgroup.Alias + "&members=" + hostGroupMemberList

	data := &url.Values{}
	data.Set("hostgroup_name", hostgroup.Name)
	data.Set("alias", hostgroup.Alias)
	data.Set("members", hostGroupMemberList)

	_, err = c.put(nagiosURL)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) deleteHostgroup(name string) ([]byte, error) {
	nagiosURL, err := c.buildURL("hostgroup", "DELETE", "hostgroup_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("hostgroup_name", name)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		return nil, err
	}

	return body, nil
}
