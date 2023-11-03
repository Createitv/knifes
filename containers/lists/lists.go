// Copyright 2023 panghuang xfy150150@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.
// Package lists

package lists

import (
	"github.com/createitv/knifes/containers"
	"github.com/createitv/knifes/internal"
)

// List interface that all lists implement
type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(values ...interface{}) bool
	Sort(comparator internal.Comparator)
	Swap(index1, index2 int)
	Insert(index int, values ...interface{})
	Set(index int, value interface{})

	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}
