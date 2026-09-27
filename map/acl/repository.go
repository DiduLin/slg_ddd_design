package acl

type Repository[K comparable, V any] interface {
	Get(key K) (V, error)
	Save(value V) error
	FindBySpecification(specification Specification[V]) ([]V, error)
}
