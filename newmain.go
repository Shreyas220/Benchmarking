package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/Shreyas220/Benchmarking/utils"
)

func notmain() {
	var sum int64 = 0
	for i:= 546;i<591;i++{
		str := "sudo bpftool prog show id " + fmt.Sprint(i)
		out, err := utils.RunCommand(str)
		if err != nil {
			fmt.Println("err", err)
		}
		newstr := strings.Split(out, "\n")
		newstr2 := strings.Split(newstr[0], " ")
		if strings.Contains(newstr[0], "run_time_ns"){
			int2, err := strconv.ParseInt(newstr2[11], 10, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			sum = sum + int2
		}
	}
	fmt.Println(sum)
	
}
