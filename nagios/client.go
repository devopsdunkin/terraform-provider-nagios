package nagios

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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
	log.Printf("[DEBUG] Request body: [%p]", httpRequest)
	log.Printf("[DEBUG] HTTP Status code: %d", response.StatusCode)

	if err != nil {
		log.Printf("[ERROR] Error occurred completing HTTP request: %s", err.Error())
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	log.Printf("[DEBUG] HTTP response body: [%p]", body)

	if err != nil {
		log.Printf("[ERROR] Error occurred reading body: %s", err.Error())
		return nil, err
	}

	return body, nil
}

func (c *Client) get(objectURL string, resourceInfo interface{}) error {
	nagiosURL := c.url + objectURL + "?apikey=" + c.token + "&pretty=1"

	request, err := http.NewRequest(http.MethodGet, nagiosURL, nil)

	if err != nil {
		log.Printf("[ERROR] Error occurred during request: %s", err.Error())
		return err
	}

	body, err := c.sendRequest(request)

	if err != nil {
		log.Printf("[ERROR] Error occurred sending request: %s", err.Error())
		return err
	}

	return json.Unmarshal(body, resourceInfo)
}

func (c *Client) post(configURL string, requestBody interface{}) ([]byte, error) {
	nagiosURL := c.url + configURL + "?apikey=" + c.token + "&applyconfig=1"
	data, err := json.Marshal(requestBody)
	// log.Printf("[DEBUG] Request body: [%p]", requestBody)
	// log.Printf("[DEBUG] data variable value: [%p]", data)

	if err != nil {
		log.Printf("[ERROR] Error occurred: %s", err.Error())
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, nagiosURL, bytes.NewReader(data))
	log.Printf("[DEBUG] HTTP request body: [%p]", request.Body)

	if err != nil {
		log.Printf("[ERROR] Error occurred creating request: %s", err.Error())
		return nil, err
	}

	response, err := c.sendRequest(request)
	log.Printf("[DEBUG] Response from Nagios: [%p]", response)

	if err != nil {
		log.Printf("[ERROR] Error occurred sending request: %s", err.Error())
		return nil, err
	}

	return response, nil
}
