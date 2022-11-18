package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	runCommand()
}

func runCommand() {
	out, err := exec.Command("kubectl", "exec", "-it", "redis-client", "--", "redis-benchmark", "-t", "GET", "-n", "1000").Output()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out)

	fmt.Println(output)
	f, err := os.Create("test/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(output)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")

}
