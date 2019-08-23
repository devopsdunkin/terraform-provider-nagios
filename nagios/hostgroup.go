package nagios

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// NewHostgroup initiates the HTTP POST to the Nagios API to create a hostgroup
func (c *Client) NewHostgroup(hostgroup *Hostgroup) error {
	nagiosURL := c.url + "/config/hostgroup?apikey=" + c.token + "&applyconfig=1"
	data, err := json.Marshal(hostgroup)

	if err != nil {
		log.Printf("[ERROR] Error occurred converting hostgroup struct")
		return err
	}

	request, err := http.NewRequest(http.MethodPost, nagiosURL, bytes.NewReader(data))
	log.Printf("[DEBUG] HTTP request body: [%p]", request.Body)

	if err != nil {
		log.Printf("[ERROR] Error occurred creating request: %s", err.Error())
		return err
	}

	response, err := c.httpClient.Do(request)

	log.Printf("[DEBUG] Response from Nagios: [%p]", response)

	if err != nil {
		log.Printf("[ERROR] Error occurred sending request: %s", err.Error())
		return err
	}

	defer response.Body.Close()

	return nil
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
