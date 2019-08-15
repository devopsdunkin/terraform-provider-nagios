package nagios

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Client used to store info required to communicate with Nagios
type Client struct {
	url        string
	token      string
	httpClient *http.Client
}

// NewClient creates a pointer to the client that will be used to send requests to Nagios
func NewClient(url string, token string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	return &Client{
		url,
		token,
		httpClient,
	}
}

func (c *Client) sendRequest(httpRequest *http.Request) ([]byte, error) {
	response, err := c.httpClient.Do(httpRequest)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) get() {
	// placeholder
}

func (c *Client) post(configURL string, requestBody interface{}) ([]byte, error) {
	nagiosURL := c.url + configURL + "?apikey=" + c.token + "&pretty=1"

	data, err := json.Marshal(requestBody)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, nagiosURL, bytes.NewReader(data))

	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest(request)

	if err != nil {
		return nil, err
	}

	return response, nil

	// curl -XPOST "https://nagiosdev1-v.hy-vee.net/nagiosxi/api/v1/config/hostgroup?apikey=f0RWe3VY5prUC58FmuVQAYlbevA0iujRvbN3tYhSHMMf7abMH6Ct4eIC3uSF2bsQ&pretty=1"
	// -d "hostgroup_name=testapihostgroup&alias=HostGroup&applyconfig=1"
}
