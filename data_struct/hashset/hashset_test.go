package hashset

import "testing"

func TestNewHashSet(t *testing.T) {
	hashSet := NewHashSet[int]()
	if hashSet == nil {
		t.Errorf("Hashset is nil")
		return
	}
}

func TestHashSet_Add(t *testing.T) {
	hashSet := NewHashSet[string]()
	hashSet.Add("test1")
	hashSet.Add("test2")
	hashSet.Add("test3")
	hashSet.Add("test4")

	if hashSet.Size() != 4 {
		t.Errorf("Expected size of set to be %d but got %d", 4, hashSet.Size())
	}
}

func TestHashSet_Remove(t *testing.T) {
	hashSet := NewHashSet[string]()
	hashSet.Add("test1")
	hashSet.Add("test2")
	hashSet.Remove("test2")
	hashSet.Add("test3")

	if hashSet.Size() != 2 {
		t.Errorf("Expected size of set to be %d but got %d", 2, hashSet.Size())
	}

	if hashSet.Contains("test") {
		t.Errorf("Expected set to not contain test")
	}
}

func TestHashSet_AddAll(t *testing.T) {
	hashSet := NewHashSet[string]()
	hashSet.AddAll("test1", "test2", "test3", "test4")

	if hashSet.Size() != 4 {
		t.Errorf("Expected size of set to be %d but got %d", 4, hashSet.Size())
	}
}

func TestHashSet_RemoveAll(t *testing.T) {
	hashSet := NewHashSet[string]()
	hashSet.AddAll("test1", "test2", "test3", "test4")
	hashSet.RemoveAll("test1", "test3")
	hashSet.RemoveAll("test1", "test2", "test4")

	if !hashSet.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}
}

func TestHashSet_Clear(t *testing.T) {
	hashSet := NewHashSet[string]()
	hashSet.AddAll("test1", "test2", "test3", "test4")
	hashSet.Clear()

	if !hashSet.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}
}

func TestHashSet_IntersectionFirstSetBiggerSize(t *testing.T) {
	hashSetA := NewHashSet[int](1, 2, 3, 4)
	hashSetB := NewHashSet[int](2, 5)

	newHashSet := hashSetA.Intersection(hashSetB)

	if newHashSet.Size() != 1 {
		t.Errorf("Expected intersection size to be %d but got %d", 1, newHashSet.Size())
	}

	if !newHashSet.Contains(2) {
		t.Errorf("Expected intersection to contain value 2")
	}

	if hashSetA.Size() != 4 {
		t.Errorf("Original set should remain unchanged after operation")
	}

	if hashSetB.Size() != 2 {
		t.Errorf("Original set should remain unchanged after operation")
	}
}

func TestHashSet_IntersectionFirstSetSmallerSize(t *testing.T) {
	hashSetA := NewHashSet[int](2, 5)
	hashSetB := NewHashSet[int](1, 2, 3, 4)

	newHashSet := hashSetA.Intersection(hashSetB)

	if newHashSet.Size() != 1 {
		t.Errorf("Expected intersection size to be %d but got %d", 1, newHashSet.Size())
	}

	if !newHashSet.Contains(2) {
		t.Errorf("Expected intersection to contain value 2")
	}

	if hashSetA.Size() != 2 {
		t.Errorf("Original set should remain unchanged after operation")
	}

	if hashSetB.Size() != 4 {
		t.Errorf("Original set should remain unchanged after operation")
	}
}

func TestHashSet_Union(t *testing.T) {
	hashSetA := NewHashSet[int](1, 2, 4)
	hashSetB := NewHashSet[int](2, 3)

	newHashSet := hashSetA.Union(hashSetB)

	if newHashSet.Size() != 4 {
		t.Errorf("Expected union size to be %d but got %d", 4, newHashSet.Size())
	}

	if !newHashSet.Contains(1) {
		t.Errorf("Expected union to contain value 1")
	}

	if !newHashSet.Contains(2) {
		t.Errorf("Expected union to contain value 2")
	}

	if !newHashSet.Contains(3) {
		t.Errorf("Expected union to contain value 3")
	}

	if !newHashSet.Contains(4) {
		t.Errorf("Expected union to contain value 4")
	}

	if hashSetA.Size() != 3 {
		t.Errorf("Original set should remain unchanged after operation")
	}

	if hashSetB.Size() != 2 {
		t.Errorf("Original set should remain unchanged after operation")
	}

}

func TestHashSet_Difference(t *testing.T) {
	hashSetA := NewHashSet[int](1, 2, 4)
	hashSetB := NewHashSet[int](2, 3)

	newHashSet := hashSetA.Difference(hashSetB)

	if newHashSet.Size() != 2 {
		t.Errorf("Expected difference size to be %d but got %d", 2, newHashSet.Size())
	}

	if !newHashSet.Contains(1) {
		t.Errorf("Expected union to contain value 1")
	}

	if !newHashSet.Contains(4) {
		t.Errorf("Expected union to contain value 4")
	}

	if hashSetA.Size() != 3 {
		t.Errorf("Original set should remain unchanged after operation")
	}

	if hashSetB.Size() != 2 {
		t.Errorf("Original set should remain unchanged after operation")
	}
}

func TestHashSet_ToString(t *testing.T) {
	hashSet := NewHashSet[string]("a", "b")
	expectedString1 := "a,b"
	expectedString2 := "b,a"

	result := hashSet.ToString()

	if result != expectedString1 && result != expectedString2 {
		t.Errorf("Expected string to be %s or %s but got %s", expectedString1, expectedString2, result)
	}
}

func TestHashSet_ToSlice(t *testing.T) {
	hashSet := NewHashSet[int](1, 2, 3)
	slice := hashSet.ToSlice()

	if len(slice) != hashSet.Size() {
		t.Errorf("Expected slice length to be %d but got %d", hashSet.Size(), len(slice))
	}
}
