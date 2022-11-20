package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Config struct {
	N         int
	C         int
	Benchname string
	Iteration int
}

func RunCommand(s string) (string, error) {
	str := strings.Split(s, " ")
	out, err := exec.Command(str[0], str[1:]...).Output()
	output := string(out)
	if err != nil {
		return output, err
	}

	return output, nil

}

func CreateEmptyRecord(w csv.Writer) {
	defer w.Flush() // Using Write
	row := []string{}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

}

func InstallKubearmor() {
	fmt.Println("Sleeping for 1 minute")
	time.Sleep(60 * time.Second)

	fmt.Println("installing kubearmor now ")
	out, err := exec.Command("karmor", "install").Output()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out)
	fmt.Println("installed kubearmor  ", output)

	fmt.Println("Sleeping for 15 seconds")
	time.Sleep(15 * time.Second)

}

func CreateTxtFile(filename string, output string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(output)

	if err2 != nil {
		log.Fatal(err2)
	}

}

func ApplyDiscovery() {

	fmt.Println("Sleeping for 15 seconds")
	time.Sleep(15 * time.Second)

	out, err := RunCommand("kubectl create ns explorer")
	if err != nil {
		log.Fatal(err)
	}
	output := string(out)
	fmt.Println(output)

	fmt.Println("applying discovery engine")
	out, err = RunCommand("kubectl apply -f https://raw.githubusercontent.com/kubearmor/discovery-engine/dev/deployments/k8s/deployment.yaml -n explorer")
	if err != nil {
		log.Fatal(err)
	}
	output = string(out)
	fmt.Println(output)

	fmt.Println("Sleeping for 15 seconds")
	time.Sleep(15 * time.Second)

}

func ApplyPolicy(filename string) {
	fmt.Println("apply discovery engine")
	str := "kubectl apply -f " + filename
	out, err := RunCommand(str)
	if err != nil {
		log.Fatal(err)
		fmt.Println(out)
	}
	output := string(out)
	fmt.Println(output)

}
