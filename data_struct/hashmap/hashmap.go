package hashmap

type HashMap[K, V comparable] struct {
	items map[K]V
}

func NewHashMap[K, V comparable]() *HashMap[K, V] {
	hashMap := &HashMap[K, V]{
		items: make(map[K]V),
	}
	return hashMap
}

func (hashMap *HashMap[K, V]) Put(key K, value V) {
	hashMap.items[key] = value
}

func (hashMap *HashMap[K, V]) Get(key K) (value V) {
	return hashMap.items[key]
}

func (hashMap *HashMap[K, V]) Remove(key K) {
	delete(hashMap.items, key)
}

func (hashMap *HashMap[K, V]) Clear() {
	hashMap.items = make(map[K]V)
}

func (hashMap *HashMap[K, V]) Size() int {
	return len(hashMap.items)
}

func (hashMap *HashMap[K, V]) IsEmpty() bool {
	return len(hashMap.items) == 0
}

func (hashMap *HashMap[K, V]) ContainsKey(key K) bool {
	_, ok := hashMap.items[key]

	return ok
}

func (hashMap *HashMap[K, V]) KeySet() []K {
	keys := make([]K, hashMap.Size())
	i := 0
	for key := range hashMap.items {
		keys[i] = key
	}

	return keys
}
