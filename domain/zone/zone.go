package zone

import (
    "github.com/kolo/xmlrpc"
    "github.com/prasmussen/gandi-api/domain"
    "github.com/prasmussen/gandi-api/client"
)

type Zone struct {
    *client.Client
}

func New(c *client.Client) *Zone {
    return &Zone{c}
}

// Counts accessible zones
func (self *Zone) Count() (int64, error) {
    var result int64
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key}}
    if err := self.Call("domain.zone.count", params, &result); err != nil {
        return -1, err
    }
    return result, nil
}

// Get zone information
func (self *Zone) Info(id int64) (*ZoneInfo, error) {
    var res map[string]interface{}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, id}}
    if err := self.Call("domain.zone.info", params, &res); err != nil {
        return nil, err
    }
    return ToZoneInfo(res), nil
}

// List accessible DNS zones.
func (self *Zone) List() ([]*ZoneInfoBase, error) {
    var res []interface{}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key}}
    if err := self.Call("domain.zone.list", params, &res); err != nil {
        return nil, err
    }

    zones := make([]*ZoneInfoBase, 0)
    for _, r := range res {
        zone := ToZoneInfoBase(r.(xmlrpc.Struct))
        zones = append(zones, zone)
    }
    return zones, nil
}

// Create a zone
func (self *Zone) Create(name string) (*ZoneInfo, error) {
    var res map[string]interface{}
    createArgs := xmlrpc.Struct{"name": name}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, createArgs}}
    if err := self.Call("domain.zone.create", params, &res); err != nil {
        return nil, err
    }
    return ToZoneInfo(res), nil
}

// Delete a zone
func (self *Zone) Delete(id int64) (bool, error) {
    var res bool
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, id}}
    if err := self.Call("domain.zone.delete", params, &res); err != nil {
        return false, err
    }
    return res, nil
}

// Set the current zone of a domain
func (self *Zone) Set(domainName string, id int64) (*domain.DomainInfo, error) {
    var res map[string]interface{}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, domainName, id}}
    if err := self.Call("domain.zone.set", params, &res); err != nil {
        return nil, err
    }
    return domain.ToDomainInfo(res), nil
}
