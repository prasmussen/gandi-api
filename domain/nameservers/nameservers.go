package zone

import (
	"fmt"
	"github.com/prasmussen/gandi-api/client"
	"github.com/prasmussen/gandi-api/domain"
)

type Nameservers struct {
	*client.Client
}

func New(c *client.Client) *Nameservers {
	return &Nameservers{c}
}

// Set the current zone of a domain
func (self *Nameservers) Set(domainName string, nameservers []string) (*domain.DomainInfo, error) {
	var res map[string]interface{}
	fmt.Println(nameservers)
	params := []interface{}{self.Key, domainName, nameservers}
	if err := self.Call("domain.nameservers.set", params, &res); err != nil {
		return nil, err
	}
	return domain.ToDomainInfo(res), nil
}
