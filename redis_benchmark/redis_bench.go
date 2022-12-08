package redisbenchmark

import (
	"fmt"
	"log"
	"strconv"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/Shreyas220/Benchmarking/utils"
)

func RunRedisBenchmark() {

	fmt.Println("Starting")
 	n := 1000
	for i := 0; i < 1; i++ {
		n = n * 10
		fmt.Println("======WithoutKubearmor ", n, "======")
		runCommandWithoutKubearmor(n)
	}
 
/* 	utils.InstallKubearmor()

	n2 := 100000
	for i := 0; i < 3; i++ {
		n2 = n2 * 10
		fmt.Println("======WithKubearmor ", n2, "======")
		runCommandWithKubearmor(n2)
	}

	//apply discovery engine
	utils.ApplyDiscovery()
	utils.ApplyPolicy("redis_policy.yaml")
	n3 := 100000

	for i := 0; i < 3; i++ {
		fmt.Println("======WithPolicy-", n3, "======")
		n3 = n3 * 10
		runCommandWithPolicy(n3)
	}
 */
}


func runCommandWithoutKubearmor(n int) {
	name := "./redis_benchmark/runtime/timewithoutKubearmor0" + "_" + fmt.Sprint(n) + ".txt"
	num := getruntimens()
	saveNewRuntime(name,num)
	for i := 1; i < 11; i++ {
		command := "kubectl -n redis exec -it redis-0 -- redis-benchmark -a Hello@123 -t GET -n " + fmt.Sprintln(n)
		//out, err := exec.Command("kubectl","-n","redis" ,"exec", "-it", "redis-client", "--", "redis-benchmark", "-t", "GET", "-P", "1000", "-c", "1000", "-n", fmt.Sprint(n)).Output()
		out, err := utils.RunCommand(command)
		if err != nil {
			log.Fatal(err)
		}

		output := string(out)
		str := "./redis_benchmark/test/withoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		name := "./redis_benchmark/runtime/timewithoutKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		num := getruntimens()
		saveNewRuntime(name,num)
		f, err := os.Create(str)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(output)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("done without", i, n)
		time.Sleep(1 * time.Second)
	}

}

func runCommandWithKubearmor(n int) {
	for i := 1; i < 11; i++ {

		out, err := exec.Command("kubectl", "exec", "-it", "redis-client", "--", "redis-benchmark", "-t", "GET", "-P", "1000", "-c", "1000", "-n", fmt.Sprint(n)).Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out)
		str := "./redis_benchmark/test/withKubearmor" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		f, err := os.Create(str)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(output)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("done without", i, n)
		time.Sleep(1 * time.Second)

	}

}

func runCommandWithPolicy(n int) {
	for i := 1; i < 11; i++ {

		out, err := exec.Command("kubectl", "exec", "-it", "redis-client", "--", "redis-benchmark", "-t", "GET", "-P", "1000", "-c", "1000", "-n", fmt.Sprint(n)).Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out)
		str := "./redis_benchmark/test/withPolicy" + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
		f, err := os.Create(str)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(output)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("done without", i, n)
		time.Sleep(1 * time.Second)

	}

}

func saveNewRuntime(name string,num int64){
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	output:= fmt.Sprint(num)
	_, err2 := f.WriteString(output)

	if err2 != nil {
		log.Fatal(err2)
	}


}

func getruntimens() int64 {

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
				return 0
			}
			sum = sum + int2
		}
	}
	return sum
}
