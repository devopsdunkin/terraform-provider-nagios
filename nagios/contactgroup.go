package nagios

import (
	"net/url"
	"strings"
)

func (c *Client) newContactgroup(contactgroup *Contactgroup) ([]byte, error) {
	nagiosURL, err := c.buildURL("contactgroup", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	// data := setURLValuesFromContactgroup(contactgroup)
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

	err = c.get(data, &contactgroupArray, nagiosURL)

	if err != nil {
		return nil, err
	}

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

	// nagiosURL = setUpdateURLContactgroupParams(nagiosURL, contactgroup)
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

// func setURLValuesFromContactgroup(contactgroup *Contactgroup) *url.Values {
// 	data := &url.Values{}
// 	data.Set("contactgroup_name", contactgroup.ContactgroupName)
// 	data.Set("alias", contactgroup.Alias)

// 	// Optional attributes
// 	if contactgroup.Members != nil {
// 		data.Set("members", mapArrayToString(contactgroup.Members))
// 	}

// 	if contactgroup.ContactgroupMembers != nil {
// 		data.Set("contactgroup_members", mapArrayToString(contactgroup.ContactgroupMembers))
// 	}

// 	return data
// }

// func setUpdateURLContactgroupParams(originalURL string, contactgroup *Contactgroup) string {
// 	var nagiosURL strings.Builder

// 	nagiosURL.WriteString(originalURL)
// 	nagiosURL.WriteString(
// 		"&contactgroup_name=" + contactgroup.ContactgroupName +
// 			"&alias=" + contactgroup.Alias +
// 			"&members=" + mapArrayToString(contactgroup.Members) +
// 			"&contactgroup_members=" + mapArrayToString(contactgroup.ContactgroupMembers))

// 	return nagiosURL.String()
// }
