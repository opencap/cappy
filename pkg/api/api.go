package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opencap/opencap/pkg/messages"
	"net/http"
)

var (
	ErrNotFound = errors.New("not found")
)

const baseUrl = "http://%s:%d/v1"

type Client struct {
	client *http.Client
	host   string
	port   uint16
}

func NewClient(host string, port uint16) (*Client, error) {
	return &Client{
		client: &http.Client{},
		host:   host,
		port:   port,
	}, nil
}

func (c *Client) Lookup(domain, username string, typeId uint16) (subTypeId uint8, addrData []byte, extensions interface{}, err error) {
	var resp *http.Response
	resp, err = c.client.Get(fmt.Sprintf(baseUrl+"/domains/%s/users/%s/types/%d", c.host, c.port, domain, username, typeId))
	if err != nil {
		err = fmt.Errorf("request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = ErrNotFound
		return
	}

	var body messages.LookupResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		err = fmt.Errorf("json decoding failed: %v", err)
		return
	}

	subTypeId = body.SubType
	addrData, err = base64.StdEncoding.DecodeString(body.Address)
	if err != nil {
		err = fmt.Errorf("base64 decoding failed: %v", err)
		return
	}

	return
}

func (c *Client) AssociateDomain(domain string) (err error) {
	body := messages.AssociateDomainRequest{
		Domain: domain,
	}

	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(&body); err != nil {
		return
	}

	var resp *http.Response
	resp, err = c.client.Post(fmt.Sprintf(baseUrl+"/domains", c.host, c.port), "application/json", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return
}
