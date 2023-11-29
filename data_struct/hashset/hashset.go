package hashset

import "fmt"

type HashSet[T comparable] struct {
	items map[T]struct{}
}

var itemExists struct{}

func NewHashSet[T comparable](elements ...T) *HashSet[T] {
	hashSet := &HashSet[T]{
		items: make(map[T]struct{}),
	}
	hashSet.AddAll(elements...)
	return hashSet
}

func (hashSet *HashSet[T]) Add(element T) {
	hashSet.items[element] = itemExists
}

func (hashSet *HashSet[T]) AddAll(elements ...T) {
	for _, e := range elements {
		hashSet.Add(e)
	}
}

func (hashSet *HashSet[T]) Remove(element T) {
	delete(hashSet.items, element)
}

func (hashSet *HashSet[T]) RemoveAll(elements ...T) {
	for _, e := range elements {
		hashSet.Remove(e)
	}
}

func (hashSet *HashSet[T]) Clear() {
	hashSet.items = make(map[T]struct{})
}

func (hashSet *HashSet[T]) Size() int {
	return len(hashSet.items)
}

func (hashSet *HashSet[T]) IsEmpty() bool {
	return len(hashSet.items) == 0
}

func (hashSet *HashSet[T]) Contains(element T) bool {
	_, ok := hashSet.items[element]

	return ok
}

func (hashSet *HashSet[T]) Intersection(another *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	if hashSet.Size() <= another.Size() {
		for item := range hashSet.items {
			if _, contains := another.items[item]; contains {
				result.Add(item)
			}
		}

		return result
	}

	for item := range another.items {
		if _, contains := hashSet.items[item]; contains {
			result.Add(item)
		}
	}

	return result
}

func (hashSet *HashSet[T]) Union(another *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	for item := range hashSet.items {
		result.Add(item)
	}

	for item := range another.items {
		result.Add(item)
	}

	return result
}

func (hashSet *HashSet[T]) Difference(another *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	for item := range hashSet.items {
		if _, exist := another.items[item]; !exist {
			result.Add(item)
		}
	}
	return result
}

func (hashSet *HashSet[T]) ToSlice() []T {
	slices := make([]T, hashSet.Size())
	i := 0
	for val := range hashSet.items {
		slices[i] = val
		i++
	}

	return slices
}

func (set *HashSet[T]) ToString() string {
	s := ""
	for val := range set.items {
		if s == "" {
			s = fmt.Sprintf("%v", val)
		} else {
			s = fmt.Sprintf("%s,%v", s, val)
		}
	}

	return s
}
