package main

import (
	"flag"
	"fmt"
	//"time"

	nginxbenchmark "github.com/Shreyas220/Benchmarking/nginx_benchmark"
	redisbenchmark "github.com/Shreyas220/Benchmarking/redis_benchmark"
	"github.com/Shreyas220/Benchmarking/utils"

)

func main() {
	bench := flag.String("b", "", "# of iterations")
	num := flag.Int("n", 10000, "# of iterations")
	client := flag.Int("c", 100, "# of client")
	iteration := flag.Int("i", 2, "# of client")
	flag.Parse()

	config := utils.Config{}
	config.N = *num
	config.C = *client
	config.Iteration = *iteration
	config.Benchname = string(*bench)

	switch config.Benchname {
	case "redis":
		redis()
	case "nginx":
		nginxbenchmark.RunNginxBench(config)
	default:
		fmt.Println("correct flags not provided plesase try something like this \n->  go run main.go -b nginx -n 1000 -c 100 ")
	}

}

func redis() {
	//redisbenchmark.RunRedisBenchmark()
	//time.Sleep(10 * time.Second)
	redisbenchmark.NewFiletxtthing()
}


