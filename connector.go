package connect

import (
	"fmt"
)

// ConnectorNameList returns list of connector names
// currently defined in KC cluster
func (c *Client) ConnectorNameList() ([]string, error) {
	url := "/connectors"

	result := make([]string, 0)

	err := c.get(url, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateOrUpdateConnector will create connector if it does not exist and will update configuration of existing one
func (c *Client) CreateOrUpdateConnector(name string, config map[string]string) error {
	url := fmt.Sprintf("/connectors/%s/config", name)

	err := c.put(url, config, nil)
	if err != nil {
		return err
	}

	return nil
}

// ConnectorConfig returns configuration map[string]string of connector
func (c *Client) ConnectorConfig(name string) (map[string]string, error) {
	url := fmt.Sprintf("/connectors/%s/config", name)

	result := make(map[string]string)

	err := c.get(url, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type ConnectorStatusResponse struct {
	ConnectorName   string `json:"name"`
	ConnectorStatus *struct {
		State    string
		WorkerId string
	} `json:"connector"`
	ConnectorTasks []*struct {
		TaskId     int64  `json:"id"`
		TaskState  string `json:"state"`
		WorkerId   string `json:"worker_id"`
		ErrorTrace string `json:"trace"`
	} `json:"tasks"`
}

// ConnectorStatus returns state of the connector and its tasks
func (c *Client) ConnectorStatus(name string) (*ConnectorStatusResponse, error) {
	url := fmt.Sprintf("/connectors/%s/status", name)

	result := new(ConnectorStatusResponse)

	err := c.get(url, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteConnector will delete connector from KafkaConnect cluster
func (c *Client) DeleteConnector(name string) error {
	url := fmt.Sprintf("/connectors/%s", name)

	err := c.delete(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
