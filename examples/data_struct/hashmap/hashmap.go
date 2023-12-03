package main

import (
	"fmt"

	"github.com/Mahes2/go-libs/data_struct/hashmap"
)

func main() {
	hashMap := hashmap.NewHashMap[int, string]()
	if hashMap.IsEmpty() {
		fmt.Println("stack is empty")
	}

	hashMap.Put(1, "a")
	hashMap.Put(2, "b")
	hashMap.Put(3, "c")

	el := hashMap.Get(2)
	fmt.Println("value for key 2: ", el)

	size := hashMap.Size()
	fmt.Println("size of hashmap: ", size)

	hashMap.Remove(2)
	size = hashMap.Size()
	fmt.Println("size of hashmap after removing key 2: ", size)

	hashMap.Clear()
	size = hashMap.Size()
	fmt.Println("size of hashmap after clearing: ", size)

	if hashMap.IsEmpty() {
		fmt.Println("stack is empty")
	}
}
