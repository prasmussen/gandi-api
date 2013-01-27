package client

import (
    "github.com/kolo/xmlrpc"
)

const (
    Production SystemType = iota
    Testing
)

type SystemType int

func (self SystemType) Url() string {
    if self == Production {
        return "https://rpc.gandi.net/xmlrpc/"
    }
    return "https://rpc.ote.gandi.net/xmlrpc/"
}

type Client struct {
    Key string
    Rpc *xmlrpc.Client
}

func New(apiKey string, system SystemType) (*Client, error) {
    rpc, err := xmlrpc.NewClient(system.Url(), nil)
    if err != nil {
        return nil, err
    }

    return &Client {
        Key: apiKey,
        Rpc: rpc,
    }, nil
}
