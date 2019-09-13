package nagios

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client used to store info required to communicate with Nagios
type Client struct {
	url        string
	token      string
	httpClient *http.Client
}

// NewClient creates a pointer to the client that will be used to send requests to Nagios
func NewClient(url, token string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	nagiosClient := &Client{
		url:        url,
		token:      token,
		httpClient: httpClient,
	}

	return nagiosClient
}

func (c *Client) sendRequest(httpRequest *http.Request) ([]byte, error) {
	c.addRequestHeaders(httpRequest)

	response, err := c.httpClient.Do(httpRequest)

	if err != nil {
		log.Printf("[ERROR] Error occurred completing HTTP request: %s", err.Error())
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("[ERROR] Error occurred reading body: %s", err.Error())
		return nil, err
	}

	return body, nil
}

func (c *Client) buildURL(objectType, method, objectName, name string) (string, error) {
	var nagiosURL strings.Builder

	var apiURL string
	if strings.HasSuffix(c.url, "/") {
		apiURL = "api/v1/config/"
	} else {
		apiURL = "/api/v1/config/"
	}

	// All of this creates the nagiosURL to get the object
	// that has been supplied through parameters
	nagiosURL.WriteString(c.url)
	nagiosURL.WriteString(apiURL)
	nagiosURL.WriteString(objectType)

	// If we are doing a GET, PUT or DELETE, we need to provide the name of the object
	// and type to filter results to only that. Otherwise, Nagios
	// will return all results for that particular object type
	// TODO: This is getting messy. Need to figure out a more streamlined way to handle all of this
	if method == "GET" || method == "DELETE" {
		nagiosURL.WriteString("?apikey=")
		nagiosURL.WriteString(c.token)
		nagiosURL.WriteString("&")
		nagiosURL.WriteString(objectName)
		nagiosURL.WriteString("=")

		if name == "" {
			errMsg := "Name must be provided when using the " + method + " method"
			log.Printf("[ERROR] %s", errMsg)
			return "", errors.New(errMsg)
		} else {
			nagiosURL.WriteString(name)
		}
	} else if method == "PUT" {
		nagiosURL.WriteString("/")
		nagiosURL.WriteString(name)
		nagiosURL.WriteString("?apikey=")
		nagiosURL.WriteString(c.token)
		nagiosURL.WriteString("&pretty=1&applyconfig=1")
		// nagiosURL.WriteString()
	} else if method == "POST" {
		nagiosURL.WriteString("?apikey=")
		nagiosURL.WriteString(c.token)
		nagiosURL.WriteString("&applyconfig=1")
	}

	if method != "GET" {
		nagiosURL.WriteString("&applyconfig=1") // During a POST/PUT/DELETE, we want to tell Nagios to update the config
	} else {
		nagiosURL.WriteString("&pretty=1") // Otherwise, if just doing a GET, we want it returned in a pretty format (JSON)
	}

	log.Printf("[DEBUG] Nagios URL - %s", nagiosURL.String()) // TODO: Need to scrub API key from logs

	return nagiosURL.String(), nil
}

func (c *Client) addRequestHeaders(request *http.Request) {
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "/")

	return
}

func (c *Client) get(data *url.Values, resourceInfo interface{}, nagiosURL string) error {
	request, err := http.NewRequest(http.MethodGet, nagiosURL, strings.NewReader(data.Encode()))

	if err != nil {
		log.Printf("[ERROR] Error occurred during request: %s", err.Error())
		return err
	}

	body, err := c.sendRequest(request)
	// log.Printf("[DEBUG] Body - %s", string(body))

	if err != nil {
		log.Printf("[ERROR] Error occurred sending request: %s", err.Error())
		return err
	}

	if err != nil {
		log.Printf("[ERROR] Error unmarshaling JSON data from Nagios - %s", err.Error())
		return err
	}

	test := body
	log.Printf("[DEBUG] Test value - %s", string(test))

	return json.Unmarshal(body, resourceInfo)
}

func (c *Client) post(data *url.Values, nagiosURL string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, nagiosURL, strings.NewReader(data.Encode()))

	if err != nil {
		log.Printf("[ERROR] Error creating HTTP request - %s", err.Error())
		return nil, err
	}

	body, err := c.sendRequest(request)
	// log.Printf("[DEBUG] Response from Nagios - %s", string(body))

	if err != nil {
		log.Printf("[ERROR] Error sending request: %s", err.Error())
		return nil, err
	}

	return body, nil
}

func (c *Client) put(data *url.Values, nagiosURL string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPut, nagiosURL, strings.NewReader(data.Encode()))

	if err != nil {
		log.Printf("[ERROR] Error creating HTTP request - %s", err.Error())
		return nil, err
	}

	body, err := c.sendRequest(request)

	if err != nil {
		log.Printf("[ERROR] Error sending request - %s", err.Error())
		return nil, err
	}

	return body, nil
}

func (c *Client) delete(data *url.Values, nagiosURL string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodDelete, nagiosURL, strings.NewReader(data.Encode()))

	if err != nil {
		log.Printf("[ERROR] Error creating HTTP request - %s", err.Error())
		return nil, err
	}

	body, err := c.sendRequest(request)

	if err != nil {
		log.Printf("[ERROR] Error sending request - %s", err.Error())
		return nil, err
	}

	return body, nil
}
