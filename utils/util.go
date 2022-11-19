package utils

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func RunCommand(s string) (string, error) {
	str := strings.Split(s, " ")
	out, err := exec.Command(str[0], str[1:]...).Output()
	output := string(out)
	if err != nil {
		return output, err
	}

	return output, nil

}

func CreateFile(name string) *csv.Writer {
	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)

	return w

}

func CreateRecord(w csv.Writer, filename string, status string) {
	thr, avg := FindValues(filename)
	defer w.Flush() // Using Write
	row := []string{status, thr, avg}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

}

func CreateEmptyRecord(w csv.Writer) {
	defer w.Flush() // Using Write
	row := []string{}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

}

func FindValues(fileName string) (string, string) {

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sliceData := strings.Split(string(fileBytes), "\n")
	length := len(sliceData)

	fmt.Println(strings.Replace(sliceData[length-3][0:14], " ", "", -1))
	avg := strings.Replace(sliceData[length-3][0:14], " ", "", -1)

	fmt.Println(sliceData[length-6])

	str1 := strings.Replace(sliceData[length-6][22:], "requests per second", "", -1)
	throughput := strings.ReplaceAll(str1, " ", "")

	return throughput, avg
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

	fmt.Println("Sleeping for 1 minute")
	time.Sleep(15 * time.Second)

}
