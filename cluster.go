package connect

import "net/http"

type ClusterInfoResponse struct {
	Host           string `json:"host"`
	Version        string `json:"version"`
	Commit         string `json:"commit"`
	KafkaClusterId string `json:"kafka_cluster_id"`
	Error          string `json:"error"`
}

// ClusterInfo returns information about cluster from each node of the cluster
func (c *Client) ClusterInfo() []*ClusterInfoResponse {
	result := make([]*ClusterInfoResponse, 0)

	for _, host := range c.hosts {
		url := "/"
		info := new(ClusterInfoResponse)

		err := c.do(http.MethodGet, host, url, nil, info)
		if err != nil {
			info.Error = err.Error()
		}
		info.Host = host

		result = append(result, info)
	}

	return result
}
