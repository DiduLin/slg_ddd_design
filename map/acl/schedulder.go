package acl

import "time"

type Scheduler interface {
	ExecuteOnceAfter(delay time.Duration, cb func(time.Time))
	Stop()
}
