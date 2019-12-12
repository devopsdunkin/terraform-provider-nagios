package nagios

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
)

func (c *Client) newTimeperiod(timeperiod *Timeperiod) ([]byte, error) {
	nagiosURL, err := c.buildURL("timeperiod", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	data := setURLValuesFromTimeperiod(timeperiod)

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

func (c *Client) getTimeperiod(name string) (*Timeperiod, error) {
	var timeperiodArray = []Timeperiod{}
	var timeperiod Timeperiod

	nagiosURL, err := c.buildURL("timeperiod", "GET", "timeperiod_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("timeperiod_name", name)

	body, err := c.get(data.Encode(), nagiosURL)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &timeperiodArray)

	if err != nil {
		return nil, err
	}

	for i, _ := range timeperiodArray {
		timeperiod.TimeperiodName = timeperiodArray[i].TimeperiodName
		timeperiod.Alias = timeperiodArray[i].Alias
		timeperiod.Sunday = timeperiodArray[i].Sunday
		timeperiod.Monday = timeperiodArray[i].Monday
		timeperiod.Tuesday = timeperiodArray[i].Tuesday
		timeperiod.Wednesday = timeperiodArray[i].Wednesday
		timeperiod.Thursday = timeperiodArray[i].Thursday
		timeperiod.Friday = timeperiodArray[i].Friday
		timeperiod.Saturday = timeperiodArray[i].Saturday
		timeperiod.Exclude = timeperiodArray[i].Exclude

		if i > 1 {
			break
		}
	}

	return &timeperiod, nil
}

func (c *Client) updateTimeperiod(timeperiod *Timeperiod, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("timeperiod", "PUT", "timeperiod_name", timeperiod.TimeperiodName, oldVal.(string), "")

	if err != nil {
		return err
	}

	nagiosURL = setUpdateURLTimeperiodParams(nagiosURL, timeperiod)

	log.Printf("[DEBUG] Timeperiod update: Nagios URL - %s", nagiosURL)

	_, err = c.put(nagiosURL)

	if err != nil {
		// If the error is this specific message, we want to "catch" it
		// and create a new host, then we can proceed on. Otherwise, we
		// can return the error and exit
		if strings.Contains(err.Error(), "Does the timeperiod exist?") {
			c.newTimeperiod(timeperiod)
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

func (c *Client) deleteTimeperiod(name string) ([]byte, error) {
	nagiosURL, err := c.buildURL("timeperiod", "DELETE", "timeperiod_name", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("timeperiod_name", name)

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

func setURLValuesFromTimeperiod(timeperiod *Timeperiod) *url.Values {
	data := &url.Values{}
	data.Set("timeperiod_name", timeperiod.TimeperiodName)
	data.Set("alias", timeperiod.Alias)

	// Optional attributes
	if timeperiod.Sunday != "" {
		data.Set("sunday", timeperiod.Sunday)
	}

	if timeperiod.Monday != "" {
		data.Set("monday", timeperiod.Monday)
	}

	if timeperiod.Tuesday != "" {
		data.Set("tuesday", timeperiod.Tuesday)
	}

	if timeperiod.Wednesday != "" {
		data.Set("wednesday", timeperiod.Wednesday)
	}

	if timeperiod.Thursday != "" {
		data.Set("thursday", timeperiod.Thursday)
	}

	if timeperiod.Friday != "" {
		data.Set("friday", timeperiod.Friday)
	}

	if timeperiod.Saturday != "" {
		data.Set("saturday", timeperiod.Saturday)
	}

	if timeperiod.Exclude != nil {
		data.Set("exclude", mapArrayToString(timeperiod.Exclude))
	}

	return data
}

// Function is being deprecated and will be removed in a future release
func setUpdateURLTimeperiodParams(originalURL string, timeperiod *Timeperiod) string {
	var nagiosURL strings.Builder

	nagiosURL.WriteString(originalURL)
	nagiosURL.WriteString("&timeperiod_name=" + timeperiod.TimeperiodName)
	nagiosURL.WriteString("&alias=" + timeperiod.Alias)

	if timeperiod.Sunday != "" {
		nagiosURL.WriteString("&sunday=" + timeperiod.Sunday)
	}

	if timeperiod.Monday != "" {
		nagiosURL.WriteString("&monday=" + timeperiod.Monday)
	}

	if timeperiod.Tuesday != "" {
		nagiosURL.WriteString("&tuesday=" + timeperiod.Tuesday)
	}

	if timeperiod.Wednesday != "" {
		nagiosURL.WriteString("&wednesday=" + timeperiod.Wednesday)
	}

	if timeperiod.Thursday != "" {
		nagiosURL.WriteString("&thursday=" + timeperiod.Thursday)
	}

	if timeperiod.Friday != "" {
		nagiosURL.WriteString("&friday=" + timeperiod.Friday)
	}

	if timeperiod.Saturday != "" {
		nagiosURL.WriteString("&saturday=" + timeperiod.Saturday)
	}

	return nagiosURL.String()
}
