package version

import (
    "github.com/kolo/xmlrpc"
    "github.com/prasmussen/gandi-api/client"
)

type Version struct {
    *client.Client
}

func New(c *client.Client) *Version {
    return &Version{c}
}

// Count this zone versions
func (self *Version) Count(zoneId int) (int, error) {
    var result int64
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, zoneId}}
    if err := self.Rpc.Call("domain.zone.version.count", params, &result); err != nil {
        return -1, err
    }
    return int(result), nil
}

// List this zone versions, with their creation date
func (self *Version) List(zoneId int) ([]*VersionInfo, error) {
    var res []interface{}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, zoneId}}
    if err := self.Rpc.Call("domain.zone.version.list", params, &res); err != nil {
        return nil, err
    }

    versions := make([]*VersionInfo, 0)
    for _, r := range res {
        version := ToVersionInfo(r.(xmlrpc.Struct))
        versions = append(versions, version)
    }
    return versions, nil
}

// Create a new version from another version. This will duplicate the version’s records
func (self *Version) New(zoneId, version int) (int, error) {
    var res int64

    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, zoneId, version}}
    if err := self.Rpc.Call("domain.zone.version.new", params, &res); err != nil {
        return -1, err
    }
    return int(res), nil
}

// Delete a specific version
func (self *Version) Delete(zoneId, version int) (bool, error) {
    var res bool
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, zoneId, version}}
    if err := self.Rpc.Call("domain.zone.version.delete", params, &res); err != nil {
        return false, err
    }
    return res, nil
}

// Set the active version of a zone
func (self *Version) Set(zoneId, version int) (bool, error) {
    var res bool
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, zoneId, version}}
    if err := self.Rpc.Call("domain.zone.version.set", params, &res); err != nil {
        return false, err
    }
    return res, nil
}