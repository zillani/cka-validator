package workload

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func GetNamespace() string {
	clientset := InitCluster()
	ns, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("error ", err)
	}
	return ns.String()
}
