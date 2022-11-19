package main

import (
	"time"

	redisbenchmark "github.com/Shreyas220/Benchmarking/redis_benchmark"
)

func main() {
	redis()
}

func redis() {
	redisbenchmark.RunRedisBenchmark()
	time.Sleep(10 * time.Second)
	redisbenchmark.Filetxtthing()
}
