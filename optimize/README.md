# Optimizer Benchmark
## Get Started

```
$ go test -bench . -benchmem
goos: darwin
goarch: arm64
pkg: github.com/jun06t/cel-sample/optimize
BenchmarkNewProgramOptimizeTrue-10       1818870               657.3 ns/op           144 B/op          9 allocs/op
BenchmarkNewProgramOptimizeFalse-10       836688              1329 ns/op             896 B/op         33 allocs/op
BenchmarkRawCode-10                     589481318                2.036 ns/op           0 B/op          0 allocs/op
PASS
```
