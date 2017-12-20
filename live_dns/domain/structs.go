package domain

import "github.com/google/uuid"

type DomainInfoBase struct {
	Fqdn              string `json:"fqdn,omitempty"`
	DomainRecordsHref string `json:"domain_records_href,omitempty"`
	DomainHref        string `json:"domain_href,omitempty"`
}

type DomainInfoExtra struct {
	ZoneUUID        *uuid.UUID `json:"zone_uuid,omitempty"`
	DomainKeysHref  string     `json:"domain_keys_href,omitempty"`
	ZoneHref        string     `json:"zone_href,omitempty"`
	ZoneRecordsHref string     `json:"zone_records_href,omitempty"`
}

type DomainInfo struct {
	*DomainInfoBase
	*DomainInfoExtra
}
