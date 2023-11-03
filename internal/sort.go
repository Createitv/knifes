// Copyright 2023 panghuang xfy150150@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.
// Package internal

package internal

import "sort"

// Sort sorts values (in-place) with respect to the given comparator.
//
// Uses Go's sort (hybrid of quicksort for large and then insertion sort for smaller slices).
func Sort[T any](values []T, comparator Comparator) {
	sort.Sort(sortable[T]{values, comparator})
}

type sortable[T any] struct {
	values     []T
	comparator Comparator
}

func (s sortable[T]) Len() int {
	return len(s.values)
}
func (s sortable[T]) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s sortable[T]) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}
