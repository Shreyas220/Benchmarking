package nginxbenchmark

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shreyas220/Benchmarking/utils"
)

func RunNginxBench(config utils.Config) {
	fmt.Println("======Starting Nginx Test======")

	setupNginx()
	time.Sleep(15 * time.Second)

	external_ip := getExternalIp()

	fmt.Println("---Running test without kubearmor---")
	runTest(config, external_ip, "withoutKubearmor")

	fmt.Println("---Running test with kubearmor---")
	utils.InstallKubearmor()
	runTest(config, external_ip, "withKubearmor")

	fmt.Println("---Running test with Policy---")
	utils.ApplyDiscovery()
	utils.ApplyPolicy("nginx_policy.yaml")
	runTest(config, external_ip, "withPolicy")

	Readtxt(config)

}

func runTest(config utils.Config, external_ip string, status string) {
	n := config.N
	for j := 0; j < config.Iteration; j++ {
		fmt.Println("======", status, "-", n, "======")
		for i := 1; i < 11; i++ {
			filename := "./nginx_benchmark/test/" + status + fmt.Sprint(i) + "_" + fmt.Sprint(n) + ".txt"
			s := "ab -q -m GET -n " + fmt.Sprint(config.N) + " -c " + fmt.Sprint(config.C) + " http://" + external_ip + "/hello-world-one"
			output, err := utils.RunCommand(s)
			if err != nil {
				fmt.Println(err, "\nError with filename ", filename)
			}
			utils.CreateTxtFile(filename, output)
			fmt.Println("Test ", status, " done ", i)
		}
		n = n * 10
	}
}

func setupNginx() {

	s := "kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.3.0/deploy/static/provider/cloud/deploy.yaml"
	str, err := utils.RunCommand(s)
	if err != nil {
		fmt.Println(err, str)
	}

	s = "kubectl apply -f hello.yaml -n ingress-nginx"
	str, err = utils.RunCommand(s)
	if err != nil {
		fmt.Println(err, str)
	}

	s = "kubectl apply -f hello2.yaml -n ingress-nginx"
	str, err = utils.RunCommand(s)
	if err != nil {
		fmt.Println(err, str)
	}
	fmt.Println(str)

	s = "kubectl apply -f ingress_policy.yaml -n ingress-nginx"
	str, err = utils.RunCommand(s)
	if err != nil {
		fmt.Println(err, str)
	}
	fmt.Println(str)

}

func getExternalIp() string {
	var external_ip string
	s := "kubectl get service ingress-nginx-controller --namespace=ingress-nginx"
	str, err := utils.RunCommand(s)
	if err != nil {
		fmt.Println(err, str)
	}
	fmt.Println(str)

	if !strings.Contains(str, "No resources found") {
		for i := 0; i < 10; i++ {
			newstr := strings.Split(str, "\n")
			newstr = strings.Split(newstr[1], " ")
			external_ip = newstr[9]
			if external_ip == "<pending>" {
				time.Sleep(10 * time.Second)
				continue
			}
			break
		}
	}

	return external_ip
}
