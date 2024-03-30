# Optimizer Benchmark
## Get Started

```
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/jun06t/cel-sample/optimize
BenchmarkNewProgramOptimizeTrue-10       2292873               509.6 ns/op
BenchmarkNewProgramOptimizeFalse-10      1341428               911.4 ns/op
PASS
```