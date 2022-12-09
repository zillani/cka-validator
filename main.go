package main

import (
	"fmt"
	"github.com/zillani/cka-validator/k8s/workload"
)

func main() {
	ns := workload.GetNamespace()
	fmt.Println(ns)
}
