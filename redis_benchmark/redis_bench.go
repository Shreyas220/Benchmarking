package redisbenchmark

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Shreyas220/Benchmarking/utils"
)

func RunRedisBenchmark() {

	fmt.Println("Starting")
	n := 100000
	for i := 0; i < 3; i++ {
		n = n * 10
		fmt.Println("======WithoutKubearmor ", n, "======")
		runCommandWithoutKubearmor(n)
	}

	utils.InstallKubearmor()

	n2 := 100000
	for i := 0; i < 3; i++ {
		n2 = n2 * 10
		fmt.Println("======WithKubearmor ", n2, "======")
		runCommandWithKubearmor(n2)
	}

	//apply discovery engine
	utils.ApplyDiscovery()
	utils.ApplyPolicy("redis_policy.yaml")
	n3 := 100000

	for i := 0; i < 3; i++ {
		fmt.Println("======WithPolicy-", n3, "======")
		n3 = n3 * 10
		runCommandWithPolicy(n3)
	}

}

func runCommandWithoutKubearmor(n int) {
	for i := 1; i < 11; i++ {

		out, err := exec.Command("kubectl", "exec", "-it", "redis-client", "--", "redis-benchmark", "-t", "GET", "-P", "1000", "-c", "1000", "-n", fmt.Sprint(n)).Output()
		if err != nil {
			log.Fatal(err)
		}

		output := string(out)
		str := "./redis_benchmark/test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		f, err := os.Create(str)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(output)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("done without", i, n)
		time.Sleep(1 * time.Second)
	}

}

func runCommandWithKubearmor(n int) {
	for i := 1; i < 11; i++ {

		out, err := exec.Command("kubectl", "exec", "-it", "redis-client", "--", "redis-benchmark", "-t", "GET", "-P", "1000", "-c", "1000", "-n", fmt.Sprint(n)).Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out)
		str := "./redis_benchmark/test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		f, err := os.Create(str)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(output)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("done without", i, n)
		time.Sleep(1 * time.Second)

	}

}

func runCommandWithPolicy(n int) {
	for i := 1; i < 11; i++ {

		out, err := exec.Command("kubectl", "exec", "-it", "redis-client", "--", "redis-benchmark", "-t", "GET", "-P", "1000", "-c", "1000", "-n", fmt.Sprint(n)).Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out)
		str := "./redis_benchmark/test/withPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		f, err := os.Create(str)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(output)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("done without", i, n)
		time.Sleep(1 * time.Second)

	}

}
