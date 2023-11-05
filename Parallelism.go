// Parallelism: Utilizing multiple cores to execute tasks in parallel
func main() {
    runtime.GOMAXPROCS(2) // Set GOMAXPROCS to utilize 2 cores
    go task1()
    go task2()
    time.Sleep(time.Second) // Wait for goroutines to finish
}
