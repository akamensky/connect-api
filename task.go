package connect

import "fmt"

type GetTaskListRequest struct {
	ConnectorName string
}

type GetTaskListResponse struct {
	TaskList []*struct {
		Task *struct {
			Id            int64  `json:"task"`
			ConnectorName string `json:"connector"`
		} `json:"id"`
		ConnectorConfig map[string]string `json:"config"`
	} `json:"id"`
}

//path: /connectors/:name/tasks
func (c *Client) GetTaskList(req *GetTaskListRequest) (*GetTaskListResponse, error) {
	url := fmt.Sprintf("/connectors/%s/tasks", req.ConnectorName)

	result := new(GetTaskListResponse)

	err := c.get(url, &result.TaskList)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type GetTaskStatusRequest struct {
	ConnectorName string
	TaskId        int64
}

type GetTaskStatusResponse struct {
	TaskId     int64  `json:"id"`
	TaskState  string `json:"state"`
	WorkerId   string `json:"worker_id"`
	ErrorTrace string `json:"trace"`
}

//path: /connectors/:name/tasks/:taskId/status
func (c *Client) GetTaskStatus(req *GetTaskStatusRequest) (*GetTaskStatusResponse, error) {
	url := fmt.Sprintf("/connectors/%s/tasks/%d/status", req.ConnectorName, req.TaskId)

	result := new(GetTaskStatusResponse)

	err := c.get(url, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type RestartTaskRequest struct {
	ConnectorName string
	TaskId        int64
}

//path: /connectors/:name/tasks/:taskId/restart
func (c *Client) RestartTask(req *RestartTaskRequest) error {
	url := fmt.Sprintf("/connectors/%s/tasks/%d/restart", req.ConnectorName, req.TaskId)

	err := c.get(url, nil)
	if err != nil {
		return err
	}

	return nil
}
