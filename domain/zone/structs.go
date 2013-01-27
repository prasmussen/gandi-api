package zone

import (
    "time"
)

type ZoneInfoBase struct {
    DateUpdated time.Time
    Id int
    Name string
    Public bool
    Version int
}

type ZoneInfoExtra struct {
    Domains int
    Owner string
    Versions []int
}

type ZoneInfo struct {
    *ZoneInfoBase
    *ZoneInfoExtra
}
