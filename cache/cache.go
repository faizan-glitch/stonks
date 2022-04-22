package cache

// All Caches must implement the following interface
type Cache interface {
	Get(k string) (interface{}, bool)
	Set(k string, v interface{})
	Delete(k string)
	Clear()
}
