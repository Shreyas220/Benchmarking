package nginxbenchmark

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

func Readtxt(config utils.Config) {
	file, err := os.Create("nginxBenchmark.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)

	n := config.N
	for i := 0; i < config.Iteration; i++ {
		for i := 1; i < 11; i++ {
			str := "./nginx_benchmark/test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordNginx(*w, str, "withoutKubearmor", fmt.Sprint(n))
		}
		utils.CreateEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./nginx_benchmark/test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordNginx(*w, str, "withKubearmor", fmt.Sprint(n))
		}
		utils.CreateEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./nginx_benchmark/test/withPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			CreateRecordNginx(*w, str, "withoutPolicy", fmt.Sprint(n))
		}

		utils.CreateEmptyRecord(*w)
		utils.CreateEmptyRecord(*w)

		time.Sleep(5 * time.Second)
		n = n * 10
	}

}

func CreateRecordNginx(w csv.Writer, filename string, status string, n string) {
	throughput, rps := FindValues(filename)
	defer w.Flush() // Using Write
	row := []string{status, n, throughput, rps}
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
	fmt.Println(length)

	for i, line := range sliceData {
		fmt.Println(i, line)
	}

	str := fmt.Sprintln("this", sliceData[21])
	str2 := fmt.Sprintln("this", sliceData[22])

	fmt.Println(str)
	sliceData1 := strings.Split(str, " ")
	sliceData2 := strings.Split(str2, " ")

	for i, line := range sliceData1 {
		fmt.Println("slice 1 ", i, line)
	}
	for i, line := range sliceData2 {
		fmt.Println("slice 2 ", i, line)
	}

	fmt.Println(sliceData1[7], sliceData2[10])

	return sliceData1[7], sliceData2[10]

}
