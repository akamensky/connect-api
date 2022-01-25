package connect

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	c := NewClient(&Configuration{
		Hosts:   []string{"http://localhost:8183", "http://localhost:8283", "http://localhost:8383"},
		Timeout: 1 * time.Second,
	})
	testConnectorName := "test-connector"
	testConnectorConfig := map[string]string{"name": testConnectorName, "connector.class": "org.apache.kafka.connect.tools.MockSinkConnector", "tasks.max": "1", "topics": "test", "connection.url": "jdbc:sqlite:test.db", "auto.create": "true"}
	testConnectorUpdateConfig := map[string]string{"name": testConnectorName, "connector.class": "org.apache.kafka.connect.tools.MockSinkConnector", "tasks.max": "1", "topics": "test2", "connection.url": "jdbc:sqlite:test.db", "auto.create": "true"}

	t.Run("TestClient_ClusterInfo", getClusterInfoTest())
	t.Run("TestClient_CreateConnector", getCreateOrUpdateConnectorTest(c, testConnectorName, testConnectorConfig))
	t.Run("TestClient_ConnectorNameList_NonEmpty", getConnectorNameListTest(c, []string{testConnectorName}))
	t.Run("TestClient_ConnectorConfig", getConnectorConfigTest(c, testConnectorName, testConnectorConfig))
	t.Run("TestClient_UpdateConnector", getCreateOrUpdateConnectorTest(c, testConnectorName, testConnectorUpdateConfig))
	t.Run("TestClient_ConnectorUpdateConfig", getConnectorConfigTest(c, testConnectorName, testConnectorUpdateConfig))
	t.Run("TestClient_DeleteConnector", getDeleteConnectorTest(c, testConnectorName))
	t.Run("TestClient_ConnectorNameList_Empty", getConnectorNameListTest(c, []string{}))
}

func getClusterInfoTest() func(t *testing.T) {
	return func(t *testing.T) {
		c := NewClient(&Configuration{
			Hosts:   []string{"http://dontexist:1235", "http://dontexist2:1235", "http://dontexist3:1235"},
			Timeout: 1 * time.Second,
		})

		info := c.ClusterInfo()
		for _, nodeInfo := range info {
			if nodeInfo.Error == nil {
				b, _ := json.Marshal(nodeInfo)
				t.Errorf("expected dns error, but got something else: %s", string(b))
			}
		}

		c = NewClient(&Configuration{
			Hosts:   []string{"http://localhost:8183", "http://localhost:8283", "http://localhost:8383"},
			Timeout: 1 * time.Second,
		})

		info = c.ClusterInfo()
		var nInfo *NodeStatus
		for i, nodeInfo := range info {
			if nInfo == nil {
				nInfo = nodeInfo
			}
			if nodeInfo.Error != nil {
				b, _ := json.Marshal(nodeInfo)
				t.Errorf("expected no error, but got something else: %s", string(b))
			}
			if nodeInfo.Host != fmt.Sprintf("http://localhost:8%d83", i+1) {
				t.Errorf("expected hostname [%s], but got [%s]", fmt.Sprintf("http://localhost:8%d83", i+1), nodeInfo.Host)
			}
			if nodeInfo.Version != nInfo.Version {
				t.Errorf("expected version [%s], but got [%s]", nInfo.Version, nodeInfo.Version)
			}
			if nodeInfo.Commit != nInfo.Commit {
				t.Errorf("expected commit [%s], but got [%s]", nInfo.Commit, nodeInfo.Commit)
			}
			if nodeInfo.KafkaClusterId != nInfo.KafkaClusterId {
				t.Errorf("expected cluster id [%s], but got [%s]", nInfo.KafkaClusterId, nodeInfo.KafkaClusterId)
			}
		}

		c = NewClient(&Configuration{
			Hosts:   []string{"http://localhost:1235", "http://dontexist2:1235", "http://dontexist3:1235", "http://localhost:8183"},
			Timeout: 1 * time.Second,
		})

		info = c.ClusterInfo()
		for i, nodeInfo := range info {
			b, _ := json.Marshal(nodeInfo)
			if i != 3 {
				if nodeInfo.Error == nil {
					t.Errorf("expected error, but got something else: %s", string(b))
				}
			} else if i == 3 {
				if nodeInfo.Error != nil {
					t.Errorf("expected NO error, but got something else: %s", string(b))
				}
			}
		}
	}
}

func getConnectorNameListTest(c *Client, expected []string) func(t *testing.T) {
	return func(t *testing.T) {
		connectorNames, err := c.ConnectorNameList()
		if err != nil {
			t.Error(err)
		} else {
			if !reflect.DeepEqual(connectorNames, expected) {
				t.Errorf("expected %v, but got %v", expected, connectorNames)
			}
		}
	}
}

func getCreateOrUpdateConnectorTest(c *Client, name string, config map[string]string) func(t *testing.T) {
	return func(t *testing.T) {
		err := c.CreateOrUpdateConnector(name, config)
		if err != nil {
			t.Error(err)
		}
	}
}

func getDeleteConnectorTest(c *Client, name string) func(t *testing.T) {
	return func(t *testing.T) {
		err := c.DeleteConnector(name)
		if err != nil {
			t.Error(err)
		}
	}
}

func getConnectorConfigTest(c *Client, name string, expected map[string]string) func(t *testing.T) {
	return func(t *testing.T) {
		conf, err := c.ConnectorConfig(name)
		if err != nil {
			t.Error(err)
		} else {
			if !reflect.DeepEqual(expected, conf) {
				t.Errorf("expected %v, but got %v", expected, conf)
			}
		}
	}
}
