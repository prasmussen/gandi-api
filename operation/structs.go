package operation

import (
    "time"
)

type OperationInfo struct {
    DateCreated time.Time
    DateStart time.Time
    DateUpdated time.Time
    Eta string
    Id int
    LastError string
    SessionId int
    Source string
    Step string
    Type string
    Params map[string]interface{}
    OperationDetails *OperationDetails
}

type OperationDetails struct {
    Id string
    Label string
    ProductAction string
    ProductName string
    ProductType string
    Quantity int
}
