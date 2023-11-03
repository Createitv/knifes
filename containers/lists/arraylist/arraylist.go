// Copyright 2023 panghuang xfy150150@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.
// Package arraylist

package arraylist

import (
	"fmt"
	"github.com/createitv/knifes/containers/lists"
	"github.com/createitv/knifes/internal"
	"strings"
)

var _ lists.List = (lists.List)(nil)

// List holds the elements in a slice
type List[T any] struct {
	elements []T
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New[T any](values ...T) *List[T] {
	list := &List[T]{
		elements: nil,
		size:     0,
	}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value at the end of the list
func (list *List[T]) Add(values ...T) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List[T]) Get(index int) (*T, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	return &list.elements[index], true
}

// Remove removes the element at the given index from the list.
func (list *List[T]) Remove(index int) {

	if !list.withinRange(index) {
		return
	}
	for i := 0; i < len(list.elements); i++ {
		if i == index {
			list.elements = append(list.elements[:i], list.elements[i+1:]...)
			list.size--
			list.shrink()

		}
	}
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[T]) Contains(values ...T) bool {

	for _, searchValue := range values {
		found := false
		for index := 0; index < list.size; index++ {
			if list.elements[index] == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values returns all elements in the list.
func (list *List[T]) Values() []T {
	newElements := make([]T, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// IndexOf returns index of provided element
func (list *List[T]) IndexOf(value T) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if element == value {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any elements.
func (list *List[T]) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *List[T]) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.size = 0
	list.elements = []T{}
}

// Sort sorts values (in-place) using.
func (list *List[T]) Sort(comparator internal.Comparator) {
	if len(list.elements) < 2 {
		return
	}
	internal.Sort(list.elements[:list.size], comparator)
}

// Swap swaps the two values at the specified positions.
func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Insert(index int, values ...T) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

// Set the value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Set(index int, value interface{}) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(value)
		}
		return
	}

	list.elements[index] = value
}

// String returns a string representation of container
func (list *List[T]) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List[T]) resize(cap int) {
	newElements := make([]T, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *List[T]) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *List[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}