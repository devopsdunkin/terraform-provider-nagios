package nagios

import (
	"encoding/json"
	"net/url"
	"strings"
)

func (c *Client) newContactgroup(contactgroup *Contactgroup) ([]byte, error) {
	nagiosURL, err := c.buildURL("contactgroup", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	data := setURLParams(contactgroup)

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

func (c *Client) getContactgroup(name string) (*Contactgroup, error) {
	var contactgroupArray = []Contactgroup{}
	var contactgroup Contactgroup

	nagiosURL, err := c.buildURL("contactgroup", "GET", "contactgroup_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("contactgroup_name", name)

	body, err := c.get(data.Encode(), nagiosURL)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &contactgroupArray)

	for i, _ := range contactgroupArray {
		contactgroup.ContactgroupName = contactgroupArray[i].ContactgroupName
		contactgroup.Alias = contactgroupArray[i].Alias
		contactgroup.Members = contactgroupArray[i].Members
		contactgroup.ContactgroupMembers = contactgroupArray[i].ContactgroupMembers

		if i > 1 {
			break
		}
	}

	return &contactgroup, nil
}

func (c *Client) updateContactgroup(contactgroup *Contactgroup, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("contactgroup", "PUT", "contactgroup_name", contactgroup.ContactgroupName, oldVal.(string), "")

	if err != nil {
		return err
	}

	nagiosURL = nagiosURL + setURLParams(contactgroup).Encode()

	_, err = c.put(nagiosURL)

	if err != nil {
		// If the error is this specific message, we want to "catch" it
		// and create a new host, then we can proceed on. Otherwise, we
		// can return the error and exit
		if strings.Contains(err.Error(), "Does the contactgroup exist?") {
			c.newContactgroup(contactgroup)
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

func (c *Client) deleteContactgroup(name string) ([]byte, error) {
	nagiosURL, err := c.buildURL("contactgroup", "DELETE", "contactgroup_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("contactgroup_name", name)

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
