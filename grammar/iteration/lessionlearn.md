## go only have "for" statement
## benchmark test

### first benchmark test

```
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
```

explain :

1. function name starts with "Benchmark"

2. The `testing.B` gives you access to the loop function.

3. how to execute it ?

   ```
   go test -bench=.
   ```

   > ➜  iteration git:(main) ✗ go test -bench=.
   > goos: darwin
   > goarch: arm64
   > pkg: github.com/huan11/learn-go-by-tdd/grammar/iteration
   > cpu: Apple M4 Pro
   >
   > BenchmarkRepeat-12      37569489                31.62 ns/op
   > PASS
   > ok      github.com/huan11/learn-go-by-tdd/grammar/iteration     1.337s

   **op**: This stands for "operation." 

   **ns**: This stands for nanoseconds.





### benchmark structure

A typical benchmark is structured like:

```go
func Benchmark(b *testing.B) {
	//... setup ...
	for b.Loop() {
		//... code to measure ...
	}
	//... cleanup ...
}
```



explain:

1. Only the body of the loop is timed; 
2. it automatically excludes setup and cleanup code from benchmark timing.