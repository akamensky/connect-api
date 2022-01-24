package connect

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestClient_ClusterInfo(t *testing.T) {
	c := NewClient(&Configuration{
		Hosts:   []string{"http://dontexist:1235", "http://dontexist2:1235", "http://dontexist3:1235"},
		Timeout: 1 * time.Second,
	})

	info := c.ClusterInfo()
	b, _ := json.MarshalIndent(info, "", "    ")
	log.Print(string(b))
}
