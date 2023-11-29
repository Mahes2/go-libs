package main

import (
	"fmt"
	"go-libs/data_struct/hashset"
)

func main() {
	hashSet := hashset.NewHashSet[int]()
	if hashSet.IsEmpty() {
		fmt.Println("hashset is empty")
	}

	hashSet.Add(1)
	hashSet.Add(2)
	hashSet.Add(3)

	if hashSet.Contains(2) {
		fmt.Println("hashset contains value 2")
	}

	hashSet.Remove(2)

	if !hashSet.Contains(2) {
		fmt.Println("hashset does not contain value 2")
	}

	fmt.Println("size of hashset is: ", hashSet.Size())

	anotherHashSet := hashset.NewHashSet[int](3, 6)
	fmt.Println("another hashset is: ", anotherHashSet.ToString())

	intersectionHashSet := hashSet.Intersection(anotherHashSet)
	fmt.Println("intersection hashset is: ", intersectionHashSet.ToString())

	differenceHashSet := hashSet.Difference(anotherHashSet)
	fmt.Println("difference hashset is: ", differenceHashSet.ToString())

	unionHashSet := hashSet.Union(anotherHashSet)
	fmt.Println("union hashset is: ", unionHashSet.ToString())
}
