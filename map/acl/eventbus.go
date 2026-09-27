package acl

type Eventbus interface {
	Publish(event any)
	Subscribe(cb func(event any))
}
