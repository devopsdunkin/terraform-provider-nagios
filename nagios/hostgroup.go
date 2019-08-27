package nagios

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// NewHostgroup initiates the HTTP POST to the Nagios API to create a hostgroup
func (c *Client) NewHostgroup(hostgroup *Hostgroup) ([]byte, error) {
	nagiosURL := c.url + "/config/hostgroup?apikey=" + c.token + "&applyconfig=1"

	data := url.Values{}
	data.Set("hostgroup_name", hostgroup.Name)
	data.Set("alias", hostgroup.Alias)
	// log.Printf("[DEBUG] Data to be sent to Nagios - %s", string(data))

	// if err != nil {
	// 	log.Printf("[ERROR] Error occurred converting hostgroup struct")
	// 	return nil, err
	// }

	request, err := http.NewRequest(http.MethodPost, nagiosURL, strings.NewReader(data.Encode()))

	if err != nil {
		log.Printf("[ERROR] Error creating HTTP request - %s", err.Error())
		return nil, err
	}

	// Add headers
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "/")

	// Perform HTTP request
	response, err := c.httpClient.Do(request)

	if err != nil {
		log.Printf("[ERROR] Error occurred sending request: %s", err.Error())
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	log.Printf("[DEBUG] Response - %s", string(body))

	if err != nil {
		log.Printf("[ERROR] Error processing response.Body: %s", err.Error())
		return nil, err
	}

	return body, nil
}

func (c *Client) GetHostgroup(name string) (*Hostgroup, error) {
	var hostgroup Hostgroup

	objectURL := "/objects/hostgroup"

	err := c.get(objectURL, &hostgroup)

	if err != nil {
		return nil, err
	}

	return &hostgroup, nil
}
