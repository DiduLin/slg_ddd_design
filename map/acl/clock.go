package acl

import "time"

type Clock interface {
	Now() time.Time
}
