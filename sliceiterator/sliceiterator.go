package sliceiterator

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type SliceIterator[T any] struct {
	slice []T	
	index int
}

func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{slice: slice, index: 0}
}

func (it *SliceIterator[T]) HasNext() bool {
	return it.index < len(it.slice)
}

func (it *SliceIterator[T]) Next() T {
	if !it.HasNext() {
		var zero T
		return zero
	}

	element := it.slice[it.index]
	it.index++
	return element
}





