package domain

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/prasmussen/gandi-api/client"
	"github.com/prasmussen/gandi-api/live_dns/record"
	"github.com/prasmussen/gandi-api/live_dns/test_helpers"
	"github.com/stretchr/testify/assert"
)

func RunTest(t testing.TB, method, uri, requestBody, responseBody string, code int, call func(t testing.TB, d *Domain)) {
	test_helpers.RunTest(t, method, uri, requestBody, responseBody, code, func(t testing.TB, c *client.Client) {
		call(t, New(c))
	})
}

func TestList(t *testing.T) {
	RunTest(t,
		"GET", "/api/v5/domains",
		``,
		`[
			{
			  "fqdn": "example.com",
			  "domain_records_href": "https://dns.api.gandi.net/api/v5/domains/example.com/records",
			  "domain_href": "https://dns.api.gandi.net/api/v5/domains/example.com"
			},
			{
			  "fqdn": "example.fr",
			  "domain_records_href": "https://dns.api.gandi.net/api/v5/domains/example.fr/records",
			  "domain_href": "https://dns.api.gandi.net/api/v5/domains/example.fr"
			},
			{
			  "fqdn": "example.cat",
			  "domain_records_href": "https://dns.api.gandi.net/api/v5/domains/example.cat/records",
			  "domain_href": "https://dns.api.gandi.net/api/v5/domains/example.cat"
			}
		  ]`,
		http.StatusOK,
		func(t testing.TB, d *Domain) {
			info, err := d.List()
			assert.NoError(t, err)
			assert.Equal(t, []*DomainInfoBase{
				&DomainInfoBase{
					Fqdn:              "example.com",
					DomainRecordsHref: "https://dns.api.gandi.net/api/v5/domains/example.com/records",
					DomainHref:        "https://dns.api.gandi.net/api/v5/domains/example.com",
				},
				&DomainInfoBase{
					Fqdn:              "example.fr",
					DomainRecordsHref: "https://dns.api.gandi.net/api/v5/domains/example.fr/records",
					DomainHref:        "https://dns.api.gandi.net/api/v5/domains/example.fr",
				},
				&DomainInfoBase{
					Fqdn:              "example.cat",
					DomainRecordsHref: "https://dns.api.gandi.net/api/v5/domains/example.cat/records",
					DomainHref:        "https://dns.api.gandi.net/api/v5/domains/example.cat",
				},
			}, info)
		},
	)
}

func TestInfo(t *testing.T) {
	RunTest(t,
		"GET", "/api/v5/domains/example.com",
		``,
		`{
			"zone_uuid": "f05ac8b8-e447-11e7-8e33-00163ec31f40",
			"domain_keys_href": "https://dns.api.gandi.net/api/v5/domains/example.com/keys",
			"fqdn": "example.com",
			"zone_href": "https://dns.api.gandi.net/api/v5/zones/f05ac8b8-e447-11e7-8e33-00163ec31f40",
			"zone_records_href": "https://dns.api.gandi.net/api/v5/zones/f05ac8b8-e447-11e7-8e33-00163ec31f40/records",
			"domain_records_href": "https://dns.api.gandi.net/api/v5/domains/example.com/records",
			"domain_href": "https://dns.api.gandi.net/api/v5/domains/example.com"
		  }`,
		http.StatusOK,
		func(t testing.TB, d *Domain) {
			info, err := d.Info("example.com")
			assert.NoError(t, err)
			id, err := uuid.Parse("f05ac8b8-e447-11e7-8e33-00163ec31f40")
			assert.NoError(t, err)
			assert.Equal(t, &DomainInfo{
				&DomainInfoBase{
					Fqdn:              "example.com",
					DomainRecordsHref: "https://dns.api.gandi.net/api/v5/domains/example.com/records",
					DomainHref:        "https://dns.api.gandi.net/api/v5/domains/example.com",
				},
				&DomainInfoExtra{
					ZoneUUID:        &id,
					ZoneHref:        "https://dns.api.gandi.net/api/v5/zones/f05ac8b8-e447-11e7-8e33-00163ec31f40",
					ZoneRecordsHref: "https://dns.api.gandi.net/api/v5/zones/f05ac8b8-e447-11e7-8e33-00163ec31f40/records",
					DomainKeysHref:  "https://dns.api.gandi.net/api/v5/domains/example.com/keys",
				},
			}, info)
		},
	)
}

func TestRecords(t *testing.T) {
	RunTest(t,
		"GET", "/api/v5/domains/example.com/records",
		``,
		`[
			{
			  "rrset_type": "MX",
			  "rrset_ttl": 10800,
			  "rrset_name": "@",
			  "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12bb7678-e43e-11e7-80c1-00163e6dc886/records/%40/MX",
			  "rrset_values": [
				"10 spool.mail.gandi.net.",
				"50 fb.mail.gandi.net."
			  ]
			},
			{
			  "rrset_type": "CNAME",
			  "rrset_ttl": 10800,
			  "rrset_name": "example",
			  "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12bb7678-e43e-11e7-80c1-00163e6dc886/records/example/CNAME",
			  "rrset_values": [
				"example.com."
			  ]
			}
		  ]`,
		http.StatusOK,
		func(t testing.TB, d *Domain) {
			records, err := d.Records("example.com")
			assert.NoError(t, err)
			assert.Equal(t, []*record.RecordInfo{
				&record.RecordInfo{
					Type: record.MX,
					Ttl:  10800,
					Name: "@",
					Href: "https://dns.api.gandi.net/api/v5/zones/12bb7678-e43e-11e7-80c1-00163e6dc886/records/%40/MX",
					Values: []string{
						"10 spool.mail.gandi.net.",
						"50 fb.mail.gandi.net.",
					},
				},
				&record.RecordInfo{
					Type: record.CNAME,
					Ttl:  10800,
					Name: "example",
					Href: "https://dns.api.gandi.net/api/v5/zones/12bb7678-e43e-11e7-80c1-00163e6dc886/records/example/CNAME",
					Values: []string{
						"example.com.",
					},
				},
			}, records)
		},
	)
}
