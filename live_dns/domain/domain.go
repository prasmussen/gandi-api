package domain

import (
	"fmt"

	"github.com/prasmussen/gandi-api/client"
	"github.com/prasmussen/gandi-api/live_dns/record"
)

// Domain holds the domain client stucture
type Domain struct {
	*client.Client
}

// New instanciates a new Domain client
func New(c *client.Client) *Domain {
	return &Domain{c}
}

// List domains associated to the contact represented by apikey
func (d *Domain) List() (domains []*DomainInfoBase, err error) {
	_, err = d.Get("/domains", &domains)
	return
}

// Info Gets domain information
func (d *Domain) Info(name string) (infos *DomainInfo, err error) {
	_, err = d.Get(fmt.Sprintf("/domains/%s", name), &infos)
	return
}

// Records Lists records for a given domain
func (d *Domain) Records(name string) (records []*record.RecordInfo, err error) {
	_, err = d.Get(fmt.Sprintf("/domains/%s/records", name), &records)
	return
}
