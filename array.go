package numgo

import (
	"fmt"
)

// NumgoArray represents a multidimensional array
type NumgoArray struct {
	Data  []float64
	Shape []int
}

// NewNumgoArray creates a new NumgoArray with the given shape and initializes it with zeros
func NewNumgoArray(shape []int) *NumgoArray {
	size := 1
	for _, dim := range shape {
		size *= dim
	}
	data := make([]float64, size)
	return &NumgoArray{Data: data, Shape: shape}
}

// NewNumgoArrayFromSlice creates a new NumgoArray from a slice of float64 values and the given shape
func NewNumgoArrayFromSlice(data []float64, shape []int) *NumgoArray {
	if len(data) != shape[0]*shape[1] {
		panic("invalid shape")
	}
	return &NumgoArray{Data: data, Shape: shape}
}

// Set sets the value of an element at the given indices
func (a *NumgoArray) Set(value float64, indices ...int) {
	idx := a.getIndex(indices...)
	a.Data[idx] = value
}

// Get returns the value of an element at the given indices
func (a *NumgoArray) Get(indices ...int) float64 {
	idx := a.getIndex(indices...)
	return a.Data[idx]
}

// getIndex calculates the index of an element based on its indices
func (a *NumgoArray) getIndex(indices ...int) int {
	if len(indices) != len(a.Shape) {
		panic("invalid number of indices")
	}
	idx := 0
	stride := 1
	for i := len(indices) - 1; i >= 0; i-- {
		if indices[i] < 0 || indices[i] >= a.Shape[i] {
			panic(fmt.Sprintf("index out of bounds: %v", indices))
		}
		idx += indices[i] * stride
		stride *= a.Shape[i]
	}
	return idx
}

// Array creates a new NumgoArray from a slice of float64 values
func Array(data []float64) *NumgoArray {
	return NewNumgoArrayFromSlice(data, []int{len(data)})
}
