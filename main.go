package main

import (
	"fmt"
	"github.com/zillani/cka-validator/api"
	"github.com/zillani/cka-validator/k8s/workload"
)

func main() {
	ns := workload.GetNamespace("default")
	deploy, active, _ := workload.GetDeployment("webapp", "default")
	fmt.Printf("deployment name %s and %d pods running\n", deploy, active)
	fmt.Println("namespace ", ns)
	api.Server()
}
