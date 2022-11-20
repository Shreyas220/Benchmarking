package main

import (
	"flag"
	"fmt"
	"time"

	nginxbenchmark "github.com/Shreyas220/Benchmarking/nginx_benchmark"
	redisbenchmark "github.com/Shreyas220/Benchmarking/redis_benchmark"
	"github.com/Shreyas220/Benchmarking/utils"
)

func main() {
	bench := flag.String("b", "redis", "# of iterations")
	num := flag.Int("n", 10000, "# of iterations")
	client := flag.Int("c", 100, "# of client")
	flag.Parse()
	config := utils.Config{}
	config.N = *num
	config.C = *client
	config.Benchname = string(*bench)

	switch config.Benchname {
	case "redis":
		redis()
	case "nginx":
		nginxbenchmark.RunNginxBench(config)
	default:
		fmt.Println("correct flags not provided plesase try something like this ")
	}
	//redis()

}

func redis() {
	redisbenchmark.RunRedisBenchmark()
	time.Sleep(10 * time.Second)
	redisbenchmark.Filetxtthing()
}
