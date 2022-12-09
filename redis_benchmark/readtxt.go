package redisbenchmark

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	"strings"
	"strconv"
	"github.com/Shreyas220/Benchmarking/utils"
)

var num int = 1
var prev int64 

func Filetxtthing() {
	csvfilename := "./records/records" + fmt.Sprint(num) + ".csv"
	file, err := os.Create(csvfilename)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)

	n := 1000000
	for i := 0; i < 1; i++ {
/* 		for i := 1; i < 11; i++ {
			str := "./redis_benchmark/test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordRedis(*w, str, "withoutKubearmor", fmt.Sprint(n))
		} */
		// utils.CreateEmptyRecord(*w)
		// time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./redis_benchmark/test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordRedis(*w, str, "withKubearmor", fmt.Sprint(n))
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

func CreateRecordRedis(w csv.Writer, filename string, status string, n string) {
	thr, avg,time := FindValues(filename)
	defer w.Flush() // Using Write
	row := []string{status, n, thr, avg, time}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}
}

func FindValues(fileName string) (string, string,string) {

	fileBytes, err := ioutil.ReadFile(fileName)
	var timefortest string
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sliceData := strings.Split(string(fileBytes), "\n")
	length := len(sliceData)

	for _, line := range sliceData {
		if strings.Contains(line, "requests completed in") {
			sliceData1 := strings.Split(line, " ")
			timefortest = sliceData1[6]
		}
	}

	fmt.Println(strings.Replace(sliceData[length-3][0:14], " ", "", -1))
	avg := strings.Replace(sliceData[length-3][0:14], " ", "", -1)

	fmt.Println(sliceData[length-6])

	str1 := strings.Replace(sliceData[length-6][22:], "requests per second", "", -1)
	throughput := strings.ReplaceAll(str1, " ", "")

	return throughput, avg, timefortest
}

func NewFiletxtthing() {
	csvfilename := "./records/newrecordKube" + fmt.Sprint(num) + ".csv"
	file, err := os.Create(csvfilename)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)

	n := 1000000
	assign_prev()
	for i := 1; i < 11; i++ {
		str := "./redis_benchmark/runtime/timewithKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt" 
		CreatenewRecordRedis(*w, str, "withoutkarmor", fmt.Sprint(n))
	}
	utils.CreateEmptyRecord(*w)
	time.Sleep(1 * time.Second)

	time.Sleep(5 * time.Second)
	n = n * 10
	
}

func NewFiletxtthing2() {
	csvfilename := "./records/newrecordPolicy" + fmt.Sprint(num) + ".csv"
	file, err := os.Create(csvfilename)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)

	n := 1000000
	assign_prev2()
	for i := 1; i < 11; i++ {
		str := "./redis_benchmark/runtime/timewithPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		CreatenewRecordRedis(*w, str, "withoutkarmor", fmt.Sprint(n))
	}
	utils.CreateEmptyRecord(*w)
	time.Sleep(1 * time.Second)

	time.Sleep(5 * time.Second)
	n = n * 10

	num = num +1 
	
}


func CreatenewRecordRedis(w csv.Writer, filename string, status string, n string) {
	diff := Readnewruntimefile(filename)
	defer w.Flush() // Using Write
	row := []string{diff}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

}

func assign_prev() {
	fileBytes, err := ioutil.ReadFile("./redis_benchmark/runtime/timewithKubearmor0_1000000.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	int2, err := strconv.ParseInt(string(fileBytes), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	prev = 0
	fmt.Println("prev is ", int2)

}

func assign_prev2() {
	fileBytes, err := ioutil.ReadFile("./redis_benchmark/runtime/timewithPolicy0_1000000.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	int2, err := strconv.ParseInt(string(fileBytes), 10, 64)
	if err != nil {
		fmt.Println(err)
		return 
	}
	prev = 0
	fmt.Println("prev is ", int2)
}



func Readnewruntimefile(name string) string {
	fileBytes, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	int2, err := strconv.ParseInt(string(fileBytes), 10, 64)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	result := int2 - prev
	prev = int2
	var y float64 = float64(result)
	var result2 float64 = y / 1000000000
	s := fmt.Sprintf("%f", result2) 
	fmt.Println("Result is ", s)
	return s
}
