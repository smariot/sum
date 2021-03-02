package sum

import (
	"math"
	"sync"
)

// modified from the container/heap package.
func down(values []float64, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			return
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && math.Abs(values[j2]) < math.Abs(values[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if math.Abs(values[j]) >= math.Abs(values[i]) {
			return
		}
		values[i], values[j] = values[j], values[i]
		i = j
	}
}

// SliceDestructive returns the sum of the values in the provided slice,
// by adding the value with the second smallest magnitude
// to the value with the smallest magnitude, and
// repeating this process until only one value remains.
//
// The contents of the slice is not preserved, and will contain
// contain garbage after this call.
//
// If the provided slice is empty, the returned sum is 0.
func SliceDestructive(values []float64) float64 {
	n := len(values)

	switch n {
	case 0:
		return 0
	case 1:
		return values[0]
	case 2:
		return values[0] + values[1]
	}

	// initialize heap
	for i := n/2 - 1; i >= 0; i-- {
		down(values, i, n)
	}

	for {
		// remove the value at the root of the heap
		n--
		tmp := values[0]
		values[0] = values[n]
		down(values, 0, n)

		// add the removed value to the new root
		values[0] += tmp

		// if the heap only contains two items, we're done.
		if n == 2 {
			return values[0] + values[1]
		}

		// otherwise, fix the root and repeat.
		down(values, 0, n)
	}
}

const poolItemLen = 8192

var pool = sync.Pool{New: func() interface{} {
	return make([]float64, poolItemLen)
}}

// Slice copies the passed slice and then calls SliceDestructive on the copy.
func Slice(values []float64) float64 {
	switch n := len(values); true {
	case n == 0:
		return 0
	case n == 1:
		return values[0]
	case n == 2:
		return values[0] + values[1]
	case n <= poolItemLen:
		tmp := pool.Get().([]float64)
		defer pool.Put(tmp)
		values = tmp[:copy(tmp, values)]
	default:
		tmp := make([]float64, len(values))
		values = tmp[:copy(tmp, values)]
	}
	return SliceDestructive(values)
}
