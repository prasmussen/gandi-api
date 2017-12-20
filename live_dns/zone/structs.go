package zone

import "github.com/google/uuid"

type ZoneInfo struct {
	Retry           int        `json:"retry,omitempty"`
	UUID            *uuid.UUID `json:"uuid,omitempty"`
	Minimum         int        `json:"minimum,omitempty"`
	Refresh         int        `json:"refresh,omitempty"`
	Expire          int64      `json:"expire,omitempty"`
	SharingID       *uuid.UUID `json:"sharing_id,omitempty"`
	Serial          int        `json:"serial,omitempty"`
	Email           string     `json:"email,omitempty"`
	PrimaryNS       string     `json:"primary_ns,omitempty"`
	Name            string     `json:"name,omitempty"`
	DomainsHref     string     `json:"domains_href,omitempty"`
	ZoneHref        string     `json:"zone_href,omitempty"`
	ZoneRecordsHref string     `json:"zone_records_href,omitempty"`
}

type Status struct {
	Message string `json:"message"`
}

type CreateStatus struct {
	*Status
	UUID *uuid.UUID `json:"uuid"`
}
