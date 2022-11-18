package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func redis_test() {

	n := 100000
	for i := 0; i < 3; i++ {
		n = n * 10
		runCommandWithoutKubearmor(n)
	}
	out, err := exec.Command("karmor", "install").Output()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out)
	fmt.Println("installing kubearmor", output)

	n2 := 100000
	for i := 0; i < 3; i++ {
		n2 = n2 * 10
		runCommandWithKubearmor(n)
	}
	out, err = exec.Command("kubectl", "create", "ns", "explorer").Output()
	if err != nil {
		log.Fatal(err)
	}
	output = string(out)
	fmt.Println(output)

	out, err = exec.Command("kubectl", "apply", "-f", "https://raw.githubusercontent.com/kubearmor/discovery-engine/dev/deployments/k8s/deployment.yaml", "-n", "explorer").Output()
	if err != nil {
		log.Fatal(err)
	}
	output = string(out)
	fmt.Println(output)

	fmt.Println("apply discovery engine", output)
	out, err = exec.Command("karmor", "discover", "--gRPC", ":9089", "|", "-a", "policy.yaml").Output()
	if err != nil {
		log.Fatal(err)
	}
	output = string(out)
	fmt.Println(output)

	out, err = exec.Command("kubectl", "apply", "-f", "policy.yaml").Output()
	if err != nil {
		log.Fatal(err)
	}
	output = string(out)
	fmt.Println(output)

	n3 := 100000

	for i := 0; i < 3; i++ {
		n3 = n3 * 10
		runCommandWithPolicy(n)
	}

}

func runCommandWithoutKubearmor(n int) {
	for i := 1; i < 11; i++ {

		out, err := exec.Command("kubectl", "exec", "-it", "redis-client", "--", "redis-benchmark", "-t", "GET", "-P", "1000", "-c", "1000", "-n", fmt.Sprint(n)).Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out)
		str := "test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
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
		str := "test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
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
		str := "test/withPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
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
