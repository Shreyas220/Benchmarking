package main

import (
	"time"

	nginxbenchmark "github.com/Shreyas220/Benchmarking/nginx_benchmark"
	redisbenchmark "github.com/Shreyas220/Benchmarking/redis_benchmark"
)

func main() {
	//redis()
	nginxbenchmark.RunNginxBench()
}

func redis() {
	redisbenchmark.RunRedisBenchmark()
	time.Sleep(10 * time.Second)
	redisbenchmark.Filetxtthing()
}
