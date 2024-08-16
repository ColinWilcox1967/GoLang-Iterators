package main

import (
	"fmt"

	sliceiterator "./sliceiterator"
	mapiterator "./mapiterator"
)

func main () {

	numbers := []int{1,2,3,4,5}
	numericIterator := sliceiterator.NewSliceIterator(numbers)
	for numericIterator.HasNext() {
		fmt.Println(numericIterator.Next())
	}

	strings := []string{"a","b","c","d"}
	stringIterator := sliceiterator.NewSliceIterator(strings)
	for stringIterator.HasNext() {
		fmt.Println(stringIterator.Next())
	}

	m := map[string]int{"first":1, "second":2, "third":3}
	mapIterator := mapiterator.NewMapIterator(m)

	for mapIterator.HasNext() {
		key, value := mapIterator.Next()
		fmt.Printf("Key: %s Value:%d\n", key, value)
	}


}