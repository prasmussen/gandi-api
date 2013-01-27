package record


type RecordInfo struct {
    Id int
    Name string
    Ttl int
    Type string
    Value string
}

type RecordAdd struct {
    Id int `goptions:"-i, --id, obligatory, description='Zone id'"`
    Version int `goptions:"-v, --version, obligatory, description='Zone version'"`
    Name string `goptions:"-n, --name, obligatory, description='Record name. Relative name, may contain leading wildcard. @ for empty name'"`
    Type string `goptions:"-t, --type, obligatory, description='Record type'"`
    Value string `goptions:"-V, --value, obligatory, description='Value for record. Semantics depends on the record type.'"`
    Ttl int `goptions:"-T, --ttl, description='Time to live, in seconds, between 5 minutes and 30 days'"`
}

