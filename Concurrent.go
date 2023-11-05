// Concurrency: Using goroutines to handle multiple tasks concurrently
func main() {
    go task1()
    go task2()
    time.Sleep(time.Second) // Wait for goroutines to finish
}

func task1() {
    // Do something
}

func task2() {
    // Do something else
}
