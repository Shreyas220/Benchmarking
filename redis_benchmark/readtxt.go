package redisbenchmark

import (
	"fmt"
	"time"

	"github.com/Shreyas220/Benchmarking/utils"
)

func Filetxtthing() {
	w := utils.CreateFile("records.csv")

	n := 1000000
	for i := 0; i < 3; i++ {
		for i := 1; i < 11; i++ {
			str := "./test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			utils.CreateRecord(*w, str, "00")
		}
		utils.CreateEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			utils.CreateRecord(*w, str, "10")
		}
		utils.CreateEmptyRecord(*w)
		time.Sleep(1 * time.Second)

		for i := 1; i < 11; i++ {
			str := "./test/withPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			utils.CreateRecord(*w, str, "11")
		}

		utils.CreateEmptyRecord(*w)
		utils.CreateEmptyRecord(*w)

		time.Sleep(5 * time.Second)
		n = n * 10
	}
}
