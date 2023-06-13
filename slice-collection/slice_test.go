package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestIntersect(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	expected := []int{4, 5}

	result := UniqueSliceIntersect(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIntersectEmpty(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{}
	expected := []int{}

	result := UniqueSliceIntersect(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIntersectNoMatches(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	expected := []int{}

	result := UniqueSliceIntersect(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIntersectDuplicates(t *testing.T) {
	slice1 := []int{1, 2, 2, 3, 4, 5, 5}
	slice2 := []int{2, 3, 3, 4, 5, 6}
	expected := []int{2, 3, 4, 5}

	result := UniqueSliceIntersect(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestUnion(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := UniqueSliceUnion(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestUnionEmpty(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{}
	expected := []int{1, 2, 3}

	result := UniqueSliceUnion(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestUnionNoDuplicates(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6}

	result := UniqueSliceUnion(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestUnionDuplicates(t *testing.T) {
	slice1 := []int{1, 2, 2, 3, 4, 5, 5}
	slice2 := []int{2, 3, 3, 4, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6}

	result := UniqueSliceUnion(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestComplement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 3, 4, 6}
	remove := []int{2, 4, 6}
	expected := []int{1, 3, 5}

	result := UniqueSliceComplement(slice, remove)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestComplementEmpty(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	remove := []int{}
	expected := []int{1, 2, 3, 4, 5, 6}

	result := UniqueSliceComplement(slice, remove)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestComplementNoMatches(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	remove := []int{7, 8, 9}
	expected := []int{1, 2, 3, 4, 5, 6}

	result := UniqueSliceComplement(slice, remove)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestComplementEmptySlice(t *testing.T) {
	slice := []int{}
	remove := []int{2, 4, 6}
	expected := []int{}

	result := UniqueSliceComplement(slice, remove)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestComplementEmptyBoth(t *testing.T) {
	slice := []int{}
	remove := []int{}
	expected := []int{}

	result := UniqueSliceComplement(slice, remove)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestUniqueSliceIntersect_Concurrent(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	var wg sync.WaitGroup
	results := make(chan []int, 2)

	// 启动两个 goroutine 并发调用 UniqueSliceIntersect
	wg.Add(2)
	go func() {
		defer wg.Done()
		result := UniqueSliceIntersect(slice1, slice2)
		results <- result
	}()
	go func() {
		defer wg.Done()
		result := UniqueSliceIntersect(slice1, slice2)
		results <- result
	}()

	wg.Wait()
	close(results)

	// 检查结果
	for result := range results {
		expected := []int{4, 5}
		if !equalSlice(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	}
}

func TestUniqueSliceUnion_Concurrent(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	concurrentCalls := 100
	var wg sync.WaitGroup
	results := make(chan []int, concurrentCalls)

	// 启动多个 goroutine 并发调用 UniqueSliceUnion
	for i := 0; i < concurrentCalls; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := UniqueSliceUnion(slice1, slice2)
			results <- result
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集并检查结果
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	count := 0
	for result := range results {
		count++
		if !equalSlice(result, expected) {
			t.Errorf("Unexpected result: %v", result)
		}
	}

	if count != concurrentCalls {
		t.Errorf("Expected %d results, but got %d", concurrentCalls, count)
	}
}

func TestUniqueSliceComplement_Concurrent(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	remove := []int{3, 5, 7}

	concurrentCalls := 100
	var wg sync.WaitGroup
	results := make(chan []int, concurrentCalls)

	// 启动多个 goroutine 并发调用 UniqueSliceComplement
	for i := 0; i < concurrentCalls; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := UniqueSliceComplement(slice, remove)
			results <- result
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集并检查结果
	expected := []int{1, 2, 4, 6, 8}
	count := 0
	for result := range results {
		count++
		if !equalSlice(result, expected) {
			t.Errorf("Unexpected result: %v", result)
		}
	}

	if count != concurrentCalls {
		t.Errorf("Expected %d results, but got %d", concurrentCalls, count)
	}
}

func equalSlice(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}
