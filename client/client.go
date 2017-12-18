package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kolo/xmlrpc"
)

const (
	Production SystemType = iota
	Testing
	LiveDNS
)

type SystemType int

func (self SystemType) Url() string {
	if self == Production {
		return "https://rpc.gandi.net/xmlrpc/"
	}
	if self == LiveDNS {
		return "https://dns.api.gandi.net/api/v5/"
	}
	return "https://rpc.ote.gandi.net/xmlrpc/"
}

type Client struct {
	Key string
	Url string
}

func New(apiKey string, system SystemType) *Client {
	return &Client{
		Key: apiKey,
		Url: system.Url(),
	}
}

func (self *Client) Call(serviceMethod string, args []interface{}, reply interface{}) error {
	rpc, err := xmlrpc.NewClient(self.Url, nil)
	if err != nil {
		return err
	}
	return rpc.Call(serviceMethod, args, reply)
}

func (self *Client) DoRest(req *http.Request, decoded interface{}) (*http.Response, error) {
	req.Header.Set("X-Api-Key", self.Key)
	req.Header.Set("Accept", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if decoded != nil {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if len(b) > 0 {
			err = json.Unmarshal(b, decoded)
			if err != nil {
				return nil, err
			}
		}
		resp.Body = ioutil.NopCloser(bytes.NewReader(b))
	}
	return resp, err
}

func (self *Client) NewJsonRequest(method string, url string, data interface{}) (*http.Request, error) {
	var reader io.Reader
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", strings.TrimRight(self.Url, "/"), strings.TrimLeft(url, "/")), reader)
	if err != nil {
		return nil, err
	}
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (self *Client) Get(Uri string, decoded interface{}) (*http.Response, error) {
	req, err := self.NewJsonRequest("GET", Uri, nil)
	if err != nil {
		return nil, err
	}
	resp, err := self.DoRest(req, decoded)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected http code %d on URL %v. expecting %d", resp.StatusCode, resp.Request.URL, http.StatusOK)
	}
	return resp, err
}

func (self *Client) Delete(Uri string, decoded interface{}) (*http.Response, error) {
	req, err := self.NewJsonRequest("DELETE", Uri, nil)
	if err != nil {
		return nil, err
	}
	resp, err := self.DoRest(req, decoded)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("Unexpected http code %d on URL %v. expecting %d", resp.StatusCode, resp.Request.URL, http.StatusNoContent)
	}
	return resp, err
}

func (self *Client) Post(Uri string, data interface{}, decoded interface{}) (*http.Response, error) {
	req, err := self.NewJsonRequest("POST", Uri, data)
	if err != nil {
		return nil, err
	}
	resp, err := self.DoRest(req, decoded)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Unexpected http code %d on URL %v. expecting %d", resp.StatusCode, resp.Request.URL, http.StatusCreated)
	}
	return resp, err
}

func (self *Client) Put(Uri string, data interface{}, decoded interface{}) (*http.Response, error) {
	req, err := self.NewJsonRequest("PUT", Uri, data)
	if err != nil {
		return nil, err
	}
	return self.DoRest(req, decoded)
}

func (self *Client) Patch(Uri string, data interface{}, decoded interface{}) (*http.Response, error) {
	req, err := self.NewJsonRequest("PATCH", Uri, data)
	if err != nil {
		return nil, err
	}
	resp, err := self.DoRest(req, decoded)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("Unexpected http code %d on URL %v. expecting %d", resp.StatusCode, resp.Request.URL, http.StatusAccepted)
	}
	return resp, err
}
