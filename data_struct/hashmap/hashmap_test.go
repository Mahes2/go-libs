package hashmap

import (
	"testing"
)

func TestNewHashMap(t *testing.T) {
	hashMap := NewHashMap[int, string]()
	if hashMap == nil {
		t.Errorf("Hashmap is nil")
		return
	}
}

func TestPut(t *testing.T) {
	hashMap := NewHashMap[int, string]()
	if hashMap == nil {
		t.Errorf("Hashmap is nil")
		return
	}

	hashMap.Put(1, "test")
	hashMap.Put(2, "test 2")

	if hashMap.Size() != 2 {
		t.Errorf("Expected size of map to be %d but got %d", 2, hashMap.Size())
	}

	if !hashMap.ContainsKey(1) {
		t.Errorf("Map should contain key: %v", 1)
	}

	if !hashMap.ContainsKey(2) {
		t.Errorf("Map should contain key: %v", 2)
	}
}

func TestGet(t *testing.T) {
	hashMap := NewHashMap[int, string]()
	if hashMap == nil {
		t.Errorf("Hashmap is nil")
		return
	}

	hashMap.Put(1, "test")
	hashMap.Put(2, "test 2")

	if key := hashMap.Get(1); key != "test" {
		t.Errorf("Expected value for key: %v to be 'test' but got '%s'", 1, key)
	}

	if key := hashMap.Get(2); key != "test 2" {
		t.Errorf("Expected value for key: %v to be 'test 2' but got '%s'", 2, key)
	}
}

func TestRemove(t *testing.T) {
	hashMap := NewHashMap[int, string]()
	if hashMap == nil {
		t.Errorf("Hashmap is nil")
		return
	}

	hashMap.Put(1, "test")
	hashMap.Put(2, "test 2")

	hashMap.Remove(1)

	if _, ok := hashMap.items[1]; ok {
		t.Errorf("Removed key still exists in the hashmap")
	}

	if key := hashMap.Get(2); key != "test 2" {
		t.Errorf("Expected value for key: %v to be 'test 2' but got '%s'", 2, key)
	}
}

func TestClear(t *testing.T) {
	hashMap := NewHashMap[int, string]()
	if hashMap == nil {
		t.Errorf("Hashmap is nil")
		return
	}

	hashMap.Put(1, "test")
	hashMap.Put(2, "test 2")

	hashMap.Clear()

	if res := hashMap.IsEmpty(); !res {
		t.Errorf("Expected hashmap to be empty after clearing it")
	}
}

func TestKeySet(t *testing.T) {
	hashMap := NewHashMap[int, string]()
	if hashMap == nil {
		t.Errorf("Hashmap is nil")
		return
	}

	hashMap.Put(1, "test")
	hashMap.Put(2, "test 2")

	keys := hashMap.KeySet()

	if len(keys) != 2 {
		t.Errorf("Expected key set to contain 2 keys but got %d", len(keys))
	}
}

func BenchmarkHashMapPut100(b *testing.B) {
	hashMap := NewHashMap[int, string]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			hashMap.Put(j, "test")
		}
	}
}

func BenchmarkHashMapPut10000(b *testing.B) {
	hashMap := NewHashMap[int, string]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			hashMap.Put(j, "test")
		}
	}
}

func BenchmarkHashMapPut1000000(b *testing.B) {
	hashMap := NewHashMap[int, string]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			hashMap.Put(j, "test")
		}
	}
}

func BenchmarkHashMapRemove100(b *testing.B) {
	hashMap := NewHashMap[int, string]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 100; j++ {
			hashMap.Put(j, "test")
		}
		b.StartTimer()
		for j := 0; j < 100; j++ {
			hashMap.Remove(j)
		}
	}
}

func BenchmarkHashMapRemove10000(b *testing.B) {
	hashMap := NewHashMap[int, string]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 10000; j++ {
			hashMap.Put(j, "test")
		}
		b.StartTimer()
		for j := 0; j < 10000; j++ {
			hashMap.Remove(j)
		}
	}
}

func BenchmarkHashMapRemove1000000(b *testing.B) {
	hashMap := NewHashMap[int, string]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 1000000; j++ {
			hashMap.Put(j, "test")
		}
		b.StartTimer()
		for j := 0; j < 1000000; j++ {
			hashMap.Remove(j)
		}
	}
}

func BenchmarkHashMapContains100(b *testing.B) {
	hashMap := NewHashMap[int, string]()
	for j := 0; j < 100; j++ {
		hashMap.Put(j, "test")
	}
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = hashMap.ContainsKey(50)
	}
}

func BenchmarkHashMapContains10000(b *testing.B) {
	hashMap := NewHashMap[int, string]()
	for j := 0; j < 10000; j++ {
		hashMap.Put(j, "test")
	}
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = hashMap.ContainsKey(5000)
	}
}

func BenchmarkHashMapContains1000000(b *testing.B) {
	hashMap := NewHashMap[int, string]()
	for j := 0; j < 1000000; j++ {
		hashMap.Put(j, "test")
	}
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = hashMap.ContainsKey(500000)
	}
}
