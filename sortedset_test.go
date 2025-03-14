package sortedset

import (
	"testing"
)

func TestSortedSet_Int(t *testing.T) {
	// Test with integers
	set := New(func(i, j int) bool { return i < j })

	// Test Add and Contains
	t.Run("Add and Contains", func(t *testing.T) {
		set.Add(3)
		set.Add(1)
		set.Add(4)
		set.Add(1) // Duplicate

		if !set.Contains(1) || !set.Contains(3) || !set.Contains(4) {
			t.Error("Set should contain added elements")
		}
		if set.Contains(2) {
			t.Error("Set should not contain element that wasn't added")
		}
		if set.Len() != 3 {
			t.Errorf("Expected length 3, got %d", set.Len())
		}
	})

	// Test order
	t.Run("Order", func(t *testing.T) {
		elements := set.Elements()
		expected := []int{1, 3, 4}
		if len(elements) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(elements))
		}
		for i := range elements {
			if elements[i] != expected[i] {
				t.Errorf("Expected %d at position %d, got %d", expected[i], i, elements[i])
			}
		}
	})

	// Test Remove
	t.Run("Remove", func(t *testing.T) {
		set.Remove(3)
		if set.Contains(3) {
			t.Error("Set should not contain removed element")
		}
		if set.Len() != 2 {
			t.Errorf("Expected length 2, got %d", set.Len())
		}
	})

	// Test Pop operations
	t.Run("Pop operations", func(t *testing.T) {
		// Test PopLeft
		if val, ok := set.PopLeft(); !ok || val != 1 {
			t.Errorf("PopLeft expected 1, got %v", val)
		}

		// Test PopRight
		if val, ok := set.PopRight(); !ok || val != 4 {
			t.Errorf("PopRight expected 4, got %v", val)
		}

		if !set.IsEmpty() {
			t.Error("Set should be empty after pop operations")
		}
	})

	// Test bulk operations
	t.Run("Bulk operations", func(t *testing.T) {
		elements := []int{5, 2, 7, 1}
		set.AddAll(elements)

		if !set.ContainsAll(elements) {
			t.Error("Set should contain all added elements")
		}

		if !set.ContainsAny([]int{2, 9}) {
			t.Error("ContainsAny should return true for existing element")
		}

		popped := set.PopAll()
		expected := []int{1, 2, 5, 7}
		if len(popped) != len(expected) {
			t.Fatalf("PopAll: Expected length %d, got %d", len(expected), len(popped))
		}
		for i := range popped {
			if popped[i] != expected[i] {
				t.Errorf("PopAll: Expected %d at position %d, got %d", expected[i], i, popped[i])
			}
		}
	})

	// Test Clone
	t.Run("Clone", func(t *testing.T) {
		set.AddAll([]int{1, 2, 3})
		clone := set.Clone()

		if clone.Len() != set.Len() {
			t.Error("Clone should have same length as original")
		}

		original := set.Elements()
		cloned := clone.Elements()
		for i := range original {
			if original[i] != cloned[i] {
				t.Errorf("Clone elements mismatch at position %d", i)
			}
		}

		// Modify clone shouldn't affect original
		clone.Add(4)
		if set.Contains(4) {
			t.Error("Modifying clone should not affect original")
		}
	})

	// Test Clear
	t.Run("Clear", func(t *testing.T) {
		set.Clear()
		if !set.IsEmpty() {
			t.Error("Set should be empty after clear")
		}
	})
}

func TestSortedSet_Struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	set := New(func(a, b Person) bool {
		return a.Age < b.Age
	})

	set.Add(Person{Name: "Alice", Age: 30})
	set.Add(Person{Name: "Bob", Age: 25})
	set.Add(Person{Name: "Charlie", Age: 35})

	if !set.Contains(Person{Name: "Alice", Age: 30}) {
		t.Error("Set should contain Alice")
	}
	if set.Len() != 3 {
		t.Errorf("Expected length 3, got %d", set.Len())
	}

	elements := set.Elements()
	expected := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	for i := range elements {
		if elements[i].Name != expected[i].Name || elements[i].Age != expected[i].Age {
			t.Errorf("Expected %v at position %d, got %v", expected[i], i, elements[i])
		}
	}
}
