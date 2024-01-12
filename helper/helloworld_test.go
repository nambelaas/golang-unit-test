package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testMain hanya berjalan pada 1 package, jika ingin digunakan pada package lain maka perlu di define lagi
func TestMain(m *testing.M) {
	// Sebelum
	fmt.Println("Sebelum Run Testing")
	m.Run()
	// Sesudah
	fmt.Println("Setelah Run Testing")
}

// untuk membatalkan test dapat menggunakan Test()
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Darwin")
	}

	result := HelloWorld("Salman")

	require.Equal(t, "Hello Salman", result, "Result is not 'Hello Salman'")
}

// good practice testing using assert/require
// akan menggagalkan test dengan menggunakan FailNow
func TestHelloWorldRequires(t *testing.T) {
	result := HelloWorld("Salman")

	require.Equal(t, "Hello Salman", result, "Result is not 'Hello Salman'")

	fmt.Println("Test Hello World Require Is Done")
}

// akan menggagalkan test dengan menggunakan Fail
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Salman")

	assert.Equal(t, "Hello Salman", result, "Result is not 'Hello Salman'")

	fmt.Println("Test Hello World Assert Is Done")
}

// bad practice testing using if
func TestHelloWorldSalman(t *testing.T) {
	result := HelloWorld("Salman")

	if result != "Hello Salman" {
		// unit test failed
		t.Error("Result Must be Hello Salman")
	}

	fmt.Println("Test Hello World Salman")
}

func TestHelloWorldSeif(t *testing.T) {
	result := HelloWorld("Seif")

	if result != "Hello Seif" {
		// unit test failed
		t.Fatal("Result Must be Hello Seif")
	}

	fmt.Println("Test Hello World Salman")
}

func TestSubTest(t *testing.T) {
	t.Run("Salman", func(t *testing.T) {
		result := HelloWorld("Salman")

		require.Equal(t, "Hello Salman", result, "Result is not 'Hello Salman'")
	})
	t.Run("Seif", func(t *testing.T) {
		result := HelloWorld("Seif")

		require.Equal(t, "Hello Seif", result, "Result is not 'Hello Seif'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Salman",
			request:  "Salman",
			expected: "Hello Salman",
		},
		{
			name:     "Seif",
			request:  "Seif",
			expected: "Hello Seif",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Salman")
	}
}

func BenchmarkHelloWordSeif(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Seif")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Salman", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Salman")
		}
	})
	b.Run("Seif", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Seif")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Salman",
			request: "Salman",
		},
		{
			name:    "Seif",
			request: "Seif",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

/*
	Command run benchmark
	go test -v -bench=. 					-> menjalankan semua benchmark pada module

	go test -v -bench=NamaBenchmark 		-> menjalankan benchmark yang disebut pada module

	go test -v -run=Test -bench=. 			-> menjalankan semua benchmark tanpa menjalankan unit test pada module

	go test -v -run=Test -bench=. ./...		-> menjalankan benchmark pada semua module dan package tanpa unit testnya
*/
