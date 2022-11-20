package nginxbenchmark

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shreyas220/Benchmarking/utils"
)

func RunNginxBench(config utils.Config) {

	setupNginx()
	external_ip := getExternalIp()
	runTest(config, external_ip, "withoutKubearmor")

	utils.InstallKubearmor()
	runTest(config, external_ip, "withKubearmor")

	utils.ApplyDiscovery()
	runTest(config, external_ip, "withPolicy")

}

func runTest(config utils.Config, external_ip string, status string) {
	for i := 0; i < 10; i++ {
		filename := "./nginx_benchmark/test/" + status + fmt.Sprint(i) + "_" + fmt.Sprint(config.N) + ".txt"
		s := "ab -q -m GET -n " + fmt.Sprint(config.N) + " -c " + fmt.Sprint(config.C) + " http://" + external_ip + "/hello-world-one"
		output, err := utils.RunCommand(s)
		if err != nil {
			fmt.Println(err, "\nError with filename ", filename)
		}
		utils.CreateTxtFile(filename, output)
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

	fmt.Println(external_ip)
	return external_ip
}
