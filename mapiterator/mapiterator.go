package mapiterator

type MapIterator[K comparable, V any] struct {
	keys []K
	values map[K]V 
	index int
}

func NewMapIterator[K comparable, V any](m map[K]V) *MapIterator[K,V] {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return &MapIterator[K,V]{keys: keys, values:m, index:0}
}

func (it *MapIterator[K,V])HasNext() bool {
	return it.index < len(it.keys)
}

func (it *MapIterator[K,V])Next() (K,V) {
	if !it.HasNext() {
		var zeroK K
		var zeroV V 
		return zeroK, zeroV
	}

	key := it.keys[it.index]
	value := it.values[key]

	it.index++
	return key, value

}