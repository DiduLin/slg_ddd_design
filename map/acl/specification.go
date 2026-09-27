package acl

type Specification[T any] interface {
	IsSatisfiedBy(value T) bool
}
