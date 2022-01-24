package connect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultTimeout = 10 * time.Second
)

type Client struct {
	hosts      []string
	httpClient *http.Client
}

type Configuration struct {
	Hosts   []string
	Timeout time.Duration
}

func NewClient(conf *Configuration) *Client {
	c := &Client{
		hosts: conf.Hosts,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	if conf.Timeout != 0 {
		c.httpClient.Timeout = conf.Timeout
	}
	return c
}

func (c *Client) makeUrl(host, url string) string {
	return fmt.Sprintf("%s%s", host, url)
}

func (c *Client) get(url string, resp interface{}) error {
	return c.doWithFailover(http.MethodGet, url, nil, resp)
}

func (c *Client) post(url string, reqData interface{}, resp interface{}) error {
	return c.doWithFailover(http.MethodPost, url, reqData, resp)
}

func (c *Client) put(url string, reqData interface{}, resp interface{}) error {
	return c.doWithFailover(http.MethodPut, url, reqData, resp)
}

func (c *Client) delete(url string, reqData interface{}, resp interface{}) error {
	return c.doWithFailover(http.MethodDelete, url, reqData, resp)
}

func (c *Client) doWithFailover(method, url string, reqData interface{}, ret interface{}) error {
	errStack := make([]error, 0)
	for _, host := range c.hosts {
		err := c.do(method, host, url, reqData, ret)
		if err != nil {
			// If err is API err, then we just return, if it is actual error, we try until the end
			if _, ok := err.(Err); !ok {
				errStack = append(errStack, fmt.Errorf("host %s: %w", host, err))
				continue
			} else {
				return err
			}
		}
	}
	result := fmt.Errorf("all hosts in cluster failed:\n")
	for _, err := range errStack {
		result = fmt.Errorf("    %w", err)
	}
	return result
}

func (c *Client) do(method, host, url string, reqData interface{}, ret interface{}) error {
	var reqBody io.Reader
	if reqData != nil {
		bodyBytes, err := json.Marshal(reqData)
		if err != nil {
			return err
		}
		reqBody = bytes.NewReader(bodyBytes)
	}

	r, err := http.NewRequest(method, c.makeUrl(host, url), reqBody)
	if err != nil {
		return err
	}

	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// check error
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		// try to parse body as error
		e := Err{}
		err := json.Unmarshal(body, &e)
		if err != nil {
			return fmt.Errorf("failed to parse response body as error, HTTP status code was [%d]", resp.StatusCode)
		}

		return e
	}

	if ret != nil {
		err = json.Unmarshal(body, ret)
		if err != nil {
			return fmt.Errorf("failed to parse response body as error, HTTP status code was [%d]", resp.StatusCode)
		}
	}

	return nil
}
