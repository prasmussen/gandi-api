package record

const (
	A     = "A"
	AAAA  = "AAAA"
	CAA   = "CAA"
	CDS   = "CDS"
	CNAME = "CNAME"
	DNAME = "DNAME"
	DS    = "DS"
	LOC   = "LOC"
	MX    = "MX"
	NS    = "NS"
	PTR   = "PTR"
	SPF   = "SPF"
	SRV   = "SRV"
	SSHFP = "SSHFP"
	TLSA  = "TLSA"
	TXT   = "TXT"
	WKS   = "WKS"
)

type RecordInfo struct {
	Href   string   `json:"rrset_href,omitempty"`
	Name   string   `json:"rrset_name,omitempty"`
	Ttl    int64    `json:"rrset_ttl,omitempty"`
	Type   string   `json:"rrset_type,omitempty"`
	Values []string `json:"rrset_values,omitempty"`
}

type Status struct {
	Message string `json:"message"`
}