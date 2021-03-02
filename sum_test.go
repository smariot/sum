package sum_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/smariot/sum"
)

func ExampleSliceDestructive() {
	const n = 1000000
	repeatedAddition := 0.
	var values []float64

	// Going from largest to smallest to make this the worst case scenario
	// for repeated addition.
	for i := n; i > 0; i-- {
		repeatedAddition += float64(i * i)
		values = append(values, float64(i*i))
	}

	const exact = n * (n + 1) * (2*n + 1) / 6
	sumSliceDestructive := sum.SliceDestructive(values)
	fmt.Printf("sum of n² for 1 <= n <= %d:\n", n)
	fmt.Printf("  exact:                %d\n", exact)
	fmt.Printf("  repeated addition:    %g (difference=%g)\n", repeatedAddition, exact-repeatedAddition)
	fmt.Printf("  sum.SliceDestructive: %g (difference=%g)\n", sumSliceDestructive, exact-sumSliceDestructive)

	// output:
	// sum of n² for 1 <= n <= 1000000:
	//   exact:                333333833333500000
	//   repeated addition:    3.333338333342835e+17 (difference=-783488)
	//   sum.SliceDestructive: 3.3333383333349997e+17 (difference=64)
}

func TestSliceDestructive(t *testing.T) {
	values := make([]float64, 262144)

	for _, i := range []int{0, 1, 2, 3, 4, 64, 4096, 262144} {
		for j := 0; j < i; j++ {
			values[j] = float64(j + 1)
		}

		expected := float64(i * (i + 1) / 2)
		sum := sum.SliceDestructive(values[:i])
		if expected != sum {
			t.Errorf("sum.SliceDestructive([0 < n < %d]) = %v, want %v", i, sum, expected)
		}
	}
}

func TestSlice(t *testing.T) {
	values := make([]float64, 262144)
	for i := range values {
		values[i] = float64(1 + i)
	}

	for _, i := range []int{0, 1, 2, 3, 4, 64, 4096, 262144} {
		expected := float64(i * (i + 1) / 2)
		sum := sum.Slice(values[:i])
		if expected != sum {
			t.Errorf("sum.Slice([0 < n < %d]) = %v, want %v", i, sum, expected)
		}
	}
}

// assigned to to prevent optimizer from removing the code we're trying to benchmark.
var unused float64

func benchmarkRepeatedAdditionN(b *testing.B, sz int) {
	r := rand.New(rand.NewSource(42))
	slice := make([]float64, sz)
	sliceOrig := make([]float64, sz)

	for i := range slice {
		sliceOrig[i] = r.NormFloat64()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copy(slice, sliceOrig)
		b.StartTimer()
		total := 0.
		for _, value := range slice {
			total += value
		}
		unused = total
	}
}

func benchmarkSliceDestructiveN(b *testing.B, sz int) {
	r := rand.New(rand.NewSource(42))
	slice := make([]float64, sz)
	sliceOrig := make([]float64, sz)
	for i := range slice {
		sliceOrig[i] = r.NormFloat64()
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copy(slice, sliceOrig)
		b.StartTimer()
		unused = sum.SliceDestructive(slice)
	}
}

func BenchmarkRepeatedAddition1000(b *testing.B) {
	benchmarkRepeatedAdditionN(b, 1000)
}

func BenchmarkRepeatedAddition1000000(b *testing.B) {
	benchmarkRepeatedAdditionN(b, 1000000)
}

func BenchmarkSliceDestructive1000(b *testing.B) {
	benchmarkSliceDestructiveN(b, 1000)
}

func BenchmarkSliceDestructive1000000(b *testing.B) {
	benchmarkSliceDestructiveN(b, 1000000)
}
