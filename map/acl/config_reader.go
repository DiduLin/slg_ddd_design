package acl

type ConfigReader[K comparable, T any] interface {
	GetConfig(key K) (T, error)
	GetAllConfigs() ([]T, error)
	GetConfigBySpecification(specification Specification[T]) (T, error)
	ListBySpecification(specification Specification[T]) ([]T, error)
}
