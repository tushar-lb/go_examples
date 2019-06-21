package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/tushar/terminal/file_server/server"
)

var (
	httpCache = make(map[string]*http.Client)
	cacheLock sync.Mutex
)

// Client is an HTTP REST wrapper. Use one of Get/Post/Put/Delete to get a request
// object.
type Client struct {
	BaseUrl    *url.URL
	HttpClient *http.Client
	UserAgent  string
}

type Resp struct {
}

// FileStats return file stats received from file server
func (c *Client) FileStats() (server.FileStatsInfo, error) {
	var fileStats server.FileStatsInfo
	req, err := c.newRequest("GET", "/stats", nil)
	if err != nil {
		return fileStats, err
	}

	_, err = c.do(req, &fileStats)
	return fileStats, err
}

func (c *Client) SendFileInfo(fileInfo server.FileStats) error {

	req, err := c.newRequest("POST", "/stats", &fileInfo)
	if err != nil {
		return err
	}

	var t Resp
	_, err = c.do(req, &t)
	return err
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseUrl.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
