package main

func main() {
	main()
}


// When we call main inside main it is throwing goroutine stack exceeds 1000000000-byte limit

// Also it running 4 goroutines:
// goroutine 1 [running]
// goroutine 2 [force gc (idle)]
// goroutine 3 [GC sweep wait]
// goroutine 4 [GC scavenge wait]

// It seems after stack exceeds it started 4 goroutines to clean memory using Garbage collector.