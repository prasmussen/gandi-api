package record

import (
    "github.com/kolo/xmlrpc"
    "github.com/prasmussen/gandi-api/client"
)

type Record struct {
    *client.Client
}

func New(c *client.Client) *Record {
    return &Record{c}
}

// Count number of records for a given zone/version
func (self *Record) Count(zoneId, version int64) (int64, error) {
    var result int64
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, zoneId, version}}
    if err := self.Rpc.Call("domain.zone.record.count", params, &result); err != nil {
        return -1, err
    }
    return result, nil
}

// List records of a version of a DNS zone
func (self *Record) List(zoneId, version int64) ([]*RecordInfo, error) {
    var res []interface{}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, zoneId, version}}
    if err := self.Rpc.Call("domain.zone.record.list", params, &res); err != nil {
        return nil, err
    }

    records := make([]*RecordInfo, 0)
    for _, r := range res {
        record := ToRecordInfo(r.(xmlrpc.Struct))
        records = append(records, record)
    }
    return records, nil
}

// Add a new record to zone
func (self *Record) Add(args RecordAdd) (*RecordInfo, error) {
    var res map[string]interface{}
    createArgs := xmlrpc.Struct{
        "name": args.Name,
        "type": args.Type,
        "value": args.Value,
        "ttl": args.Ttl,
    }

    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, args.Zone, args.Version, createArgs}}
    if err := self.Rpc.Call("domain.zone.record.add", params, &res); err != nil {
        return nil, err
    }
    return ToRecordInfo(res), nil
}

// Remove a record from a zone/version
func (self *Record) Delete(zoneId, version, recordId int64) (bool, error) {
    var res int64
    deleteArgs := xmlrpc.Struct{"id": recordId}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, zoneId, version, deleteArgs}}
    if err := self.Rpc.Call("domain.zone.record.delete", params, &res); err != nil {
        return false, err
    }
    return (res == 1), nil
}

//// Set the current zone of a domain
//func (self *Record) Set(domainName string, id int64) (*domain.DomainInfo, error) {
//    var res map[string]interface{}
//    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, domainName, id}}
//    if err := self.Rpc.Call("domain.zone.set", params, &res); err != nil {
//        return nil, err
//    }
//    return domain.ToDomainInfo(res), nil
//}
