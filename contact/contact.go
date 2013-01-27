package contact

import (
    "github.com/kolo/xmlrpc"
    "github.com/prasmussen/gandi-api/client"
)

type Contact struct {
    *client.Client
}

func New(c *client.Client) *Contact {
    return &Contact{c}
}

// Get contact financial balance
func (self *Contact) Balance() (*BalanceInformation, error) {
    var res map[string]interface{}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key}}
    if err := self.Rpc.Call("contact.balance", params, &res); err != nil {
        return nil, err
    }
    return toBalanceInformation(res), nil
}

// Get contact information
func (self *Contact) Info(handle string) (*ContactInformation, error) {
    var res map[string]interface{}

    var params xmlrpc.Params
    if handle == "" {
        params = xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key}}
    } else {
        params = xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, handle}}
    }
    if err := self.Rpc.Call("contact.info", params, &res); err != nil {
        return nil, err
    }
    return toContactInformation(res), nil
}

// Create a contact
func (self *Contact) Create(opts ContactCreate) (*ContactInformation, error) {
    var res map[string]interface{}
    createArgs := xmlrpc.Struct{
        "given": opts.Firstname,
        "family": opts.Lastname,
        "email": opts.Email,
        "password": opts.Password,
        "streetaddr": opts.Address,
        "zip": opts.Zipcode,
        "city": opts.City,
        "country": opts.Country,
        "phone": opts.Phone,
        "type": opts.ContactType(),
    }

    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, createArgs}}
    if err := self.Rpc.Call("contact.create", params, &res); err != nil {
        return nil, err
    }
    return toContactInformation(res), nil
}

// Delete a contact
func (self *Contact) Delete(handle string) (bool, error) {
    var res bool

    var params xmlrpc.Params
    if handle == "" {
        params = xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key}}
    } else {
        params = xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, handle}}
    }
    if err := self.Rpc.Call("contact.delete", params, &res); err != nil {
        return false, err
    }
    return res, nil
}
