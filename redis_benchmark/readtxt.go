package redisbenchmark

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"strconv"

	"github.com/Shreyas220/Benchmarking/utils"
)

func Filetxtthing() {
	file, err := os.Create("records.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)

	n := 1000000
	for i := 0; i < 3; i++ {
		for i := 1; i < 11; i++ {
			str := "./redis_benchmark/test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordRedis(*w, str, "withoutkarmor", fmt.Sprint(n))
		}
		utils.CreateEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./redis_benchmark/test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordRedis(*w, str, "withkarmor", fmt.Sprint(n))
		}
		utils.CreateEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./redis_benchmark/test/withPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordRedis(*w, str, "withPolicy", fmt.Sprint(n))
		}

		utils.CreateEmptyRecord(*w)
		utils.CreateEmptyRecord(*w)

		time.Sleep(5 * time.Second)
		n = n * 10
	}
}

func NewFiletxtthing() {
	file, err := os.Create("newrecord.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)

	n := 1000
	for i := 0; i < 11; i++ {
		str := "./redis_benchmark/runtime/timewithoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		CreatenewRecordRedis(*w, str, "withoutkarmor", fmt.Sprint(n))
	}
	utils.CreateEmptyRecord(*w)
	time.Sleep(1 * time.Second)

	time.Sleep(5 * time.Second)
	n = n * 10
	
}


func CreatenewRecordRedis(w csv.Writer, filename string, status string, n string) {
	thr, avg := Readruntimefile(filename)
	defer w.Flush() // Using Write
	row := []string{avg}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

}

func Readruntimefile() string{
	fileBytes, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(fileBytes)
}
