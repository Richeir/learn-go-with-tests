package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("excepted %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

	// $ go test -bench=.
	// goos: windows
	// goarch: amd64
	// pkg: github.com/richeir/learn-go-with-tests/test/iteration
	// cpu: Intel(R) Xeon(R) CPU E31280 @ 3.50GHz
	// BenchmarkRepeat-8        6852866               174.0 ns/op
	// PASS
	// ok      github.com/richeir/learn-go-with-tests/test/iteration   2.419s

	// What 136 ns/op means is our function takes on average 136 nanoseconds to run (on my computer).
	// Which is pretty ok! To test this it ran it 10000000 times.
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)

	// Output: aaaaa
}
