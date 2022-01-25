package connect

import "net/http"

type NodeStatus struct {
	Host           string `json:"host"`
	Version        string `json:"version"`
	Commit         string `json:"commit"`
	KafkaClusterId string `json:"kafka_cluster_id"`
	Error          error  `json:"error"`
}

// ClusterInfo returns information about cluster from each node of the cluster
func (c *Client) ClusterInfo() []*NodeStatus {
	result := make([]*NodeStatus, 0)

	for _, host := range c.hosts {
		url := "/"
		info := new(NodeStatus)

		err := c.do(http.MethodGet, host, url, nil, info)
		if err != nil {
			info.Error = err
		}
		info.Host = host

		result = append(result, info)
	}

	return result
}
