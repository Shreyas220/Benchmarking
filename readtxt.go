package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Filetxtthing() {
	file, err := os.Create("records2.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	n := 1000000
	w := csv.NewWriter(file)
	for i := 0; i < 2; i++ {
		for i := 1; i < 11; i++ {
			str := "./test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			createRecord(*w, str, "00")
		}
		createEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			createRecord(*w, str, "10")
		}
		createEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./test/withPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			createRecord(*w, str, "11")
		}

		createEmptyRecord(*w)
		createEmptyRecord(*w)

		time.Sleep(5 * time.Second)
		n = n * 10
	}
}

func createRecord(w csv.Writer, filename string, status string) {
	thr, avg := FindValues(filename)
	defer w.Flush() // Using Write
	row := []string{status, thr, avg}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

}

func createEmptyRecord(w csv.Writer) {
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
	f, err := strconv.ParseFloat(avg, 8)
	f = f + 1

	fmt.Println(sliceData[length-6])

	str1 := strings.Replace(sliceData[length-6][22:], "requests per second", "", -1)
	throughput := strings.ReplaceAll(str1, " ", "")

	return throughput, avg
}
