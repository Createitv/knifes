// Copyright 2023 panghuang xfy150150@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package random implements some basic functions to generate random int and string.
package random

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TestRandInt(3, 4)", args{3, 4}, 0},
		{"TestRandInt(3, 100)", args{3, 100}, 0},
		{"TestRandInt(100, 100)", args{100, 100}, 0},
		{"TestRandInt(1000, 100)", args{1000, 100}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			if got := RandInt(tt.args.min, tt.args.max); tt.args.max > tt.args.min {
				assert.LessOrEqual(got, tt.args.max, fmt.Sprintf("%v should less than %v", got, tt.args.max))
				assert.GreaterOrEqual(got, tt.args.min, fmt.Sprintf("%v should greater or equal than %v", got, tt.args.min))
			}
			if got := RandInt(tt.args.min, tt.args.max); tt.args.max <= tt.args.min {
				assert.LessOrEqual(got, tt.args.min, fmt.Sprintf("%v should less than %v", got, tt.args.min))
				assert.GreaterOrEqual(got, tt.args.max, fmt.Sprintf("%v should greater or equal than %v", got, tt.args.max))
			}
		})
	}
}
