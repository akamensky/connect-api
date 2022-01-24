package connect

import (
	"fmt"
)

type GetConnectorNameListResponse struct {
	ConnectorNameList []string
}

//path: /connectors
func (c *Client) GetConnectorNameList() (*GetConnectorNameListResponse, error) {
	url := "/connectors"

	result := new(GetConnectorNameListResponse)

	err := c.get(url, result.ConnectorNameList)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type CreateConnectorRequest struct {
	ConnectorName   string            `json:"name"`
	ConnectorConfig map[string]string `json:"config"`
}

type CreateConnectorResponse struct {
	ConnectorName   string            `json:"name"`
	ConnectorConfig map[string]string `json:"config"`
	ConnectorTasks  []*struct {
		TaskId        int64  `json:"task"`
		ConnectorName string `json:"connector"`
	} `json:"tasks"`
}

//path: /connectors
func (c *Client) CreateConnector(req *CreateConnectorRequest) (*CreateConnectorResponse, error) {
	url := "/connectors"

	result := new(CreateConnectorResponse)

	err := c.post(url, req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type GetConnectorRequest struct {
	ConnectorName string `json:"name"`
}

type GetConnectorResponse struct {
	ConnectorName   string            `json:"name"`
	ConnectorConfig map[string]string `json:"config"`
	ConnectorTasks  []*struct {
		TaskId        int64  `json:"task"`
		ConnectorName string `json:"connector"`
	} `json:"tasks"`
}

//path: /connectors/:name
func (c *Client) GetConnector(req GetConnectorRequest) (*GetConnectorResponse, error) {
	url := fmt.Sprintf("/connectors/%s", req.ConnectorName)

	result := new(GetConnectorResponse)

	err := c.get(url, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type GetConnectorConfigRequest struct {
	ConnectorName string `json:"name"`
}

type GetConnectorConfigResponse struct {
	ConnectorConfig map[string]string
}

//path: /connectors/:name/config
func (c *Client) GetConnectorConfig(req *GetConnectorConfigRequest) (*GetConnectorConfigResponse, error) {
	url := fmt.Sprintf("/connectors/%s/config", req.ConnectorName)

	result := new(GetConnectorConfigResponse)

	err := c.get(url, result.ConnectorConfig)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type UpdateConnectorConfigRequest struct {
	ConnectorName   string            `json:"name"`
	ConnectorConfig map[string]string `json:"config"`
}

type UpdateConnectorConfigResponse struct {
	ConnectorName   string            `json:"name"`
	ConnectorConfig map[string]string `json:"config"`
	ConnectorTasks  []*struct {
		TaskId        int64  `json:"task"`
		ConnectorName string `json:"connector"`
	} `json:"tasks"`
}

//path: /connectors/:name/config
func (c *Client) UpdateConnectorConfig(req *UpdateConnectorConfigRequest) (*UpdateConnectorConfigResponse, error) {
	url := fmt.Sprintf("/connectors/%s/config", req.ConnectorName)

	result := new(UpdateConnectorConfigResponse)

	err := c.put(url, req.ConnectorConfig, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type GetConnectorStatusRequest struct {
	ConnectorName string
}

type GetConnectorStatusResponse struct {
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

//path: /connectors/:name/status
func (c *Client) GetConnectorStatus(req *GetConnectorStatusRequest) (*GetConnectorStatusResponse, error) {
	url := fmt.Sprintf("/connectors/%s/status", req.ConnectorName)

	result := new(GetConnectorStatusResponse)

	err := c.get(url, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type RestartConnectorRequest struct {
	ConnectorName string
}

//path: /connectors/:name/restart
func (c *Client) RestartConnector(req *RestartConnectorRequest) error {
	url := fmt.Sprintf("/connectors/%s/restart", req.ConnectorName)

	err := c.post(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

type PauseConnectorRequest struct {
	ConnectorName string
}

//path: /connectors/:name/pause
func (c *Client) PauseConnector(req *PauseConnectorRequest) error {
	url := fmt.Sprintf("/connectors/%s/pause", req.ConnectorName)

	err := c.put(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

type ResumeConnectorRequest struct {
	ConnectorName string
}

//path: /connectors/:name/resume
func (c *Client) ResumeConnector(req *ResumeConnectorRequest) error {
	url := fmt.Sprintf("/connectors/%s/resume", req.ConnectorName)

	err := c.put(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

type DeleteConnectorRequest struct {
	ConnectorName string
}

//path: /connectors/:name/delete
func (c *Client) DeleteConnector(req *DeleteConnectorRequest) error {
	url := fmt.Sprintf("/connectors/%s/delete", req.ConnectorName)

	err := c.delete(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
