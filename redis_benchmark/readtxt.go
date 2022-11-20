package redisbenchmark

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shreyas220/Benchmarking/utils"
)

func Filetxtthing() {
	w := utils.CreateFile("records.csv")

	n := 1000000
	for i := 0; i < 3; i++ {
		for i := 1; i < 11; i++ {
			str := "./redis_benchmark/test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordRedis(*w, str, "withoutPolicy", fmt.Sprint(n))
		}
		utils.CreateEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./redis_benchmark/test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordRedis(*w, str, "withoutPolicy", fmt.Sprint(n))
		}
		utils.CreateEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./redis_benchmark/test/withPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordRedis(*w, str, "withoutPolicy", fmt.Sprint(n))
		}

		utils.CreateEmptyRecord(*w)
		utils.CreateEmptyRecord(*w)

		time.Sleep(5 * time.Second)
		n = n * 10
	}
}

func CreateRecordRedis(w csv.Writer, filename string, status string, n string) {
	thr, avg := FindValues(filename)
	defer w.Flush() // Using Write
	row := []string{status, n, thr, avg}
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
