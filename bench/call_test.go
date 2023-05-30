package bench

import (
	"testing"
)

type base struct {
	i int64
}

func (b *base) Inc() {
	b.i++
}

type derived struct {
	base
}

type counter interface {
	Inc()
}

func BenchmarkEmbed(b *testing.B) {
	// go test -gcflags=-N -v -run='^$' -bench='Embed' -benchmem -benchtime=10s
	// goos: darwin
	// goarch: amd64
	// pkg: github.com/telnet2/gopkg/bench
	// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	// BenchmarkEmbed
	// BenchmarkEmbed-12       1000000000               1.871 ns/op           0 B/op          0 allocs/op
	// PASS
	// ok      github.com/telnet2/gopkg/bench  2.183s
	derived := &derived{}
	for i := 0; i < b.N; i++ {
		derived.Inc()
	}
}

func BenchmarkInterface(b *testing.B) {
	// go test -gcflags=-N -v -run='^$' -bench='Interface' -benchmem -benchtime=10s
	// goos: darwin
	// goarch: amd64
	// pkg: github.com/telnet2/gopkg/bench
	// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	// BenchmarkInterface
	// BenchmarkInterface-12           1000000000               2.979 ns/op           0 B/op          0 allocs/op
	// PASS
	// ok      github.com/telnet2/gopkg/bench  3.526s
	var c counter = &base{}
	for i := 0; i < b.N; i++ {
		c.Inc()
	}
}

func BenchmarkBase(b *testing.B) {
	// go test -gcflags=-N -v -run='^$' -bench='Base' -benchmem -benchtime=10s
	// goos: darwin
	// goarch: amd64
	// pkg: github.com/telnet2/gopkg/bench
	// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	// BenchmarkBase
	// BenchmarkBase-12        1000000000               1.858 ns/op           0 B/op          0 allocs/op
	// PASS
	// ok      github.com/telnet2/gopkg/bench  2.174s
	var c = &base{}
	for i := 0; i < b.N; i++ {
		c.Inc()
	}
}
