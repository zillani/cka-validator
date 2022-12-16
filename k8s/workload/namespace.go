package workload

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"strings"
)

var clientset *kubernetes.Clientset

func init() {
	log.Println("ClientSet initialized!")
	clientset = InitCluster()
}

func GetNamespaces() *v1.NamespaceList {
	ns, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("error ", err)
	}
	return ns
}

func GetNamespace(namespace string) string {
	ns, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("error ", err)
	}
	for _, item := range ns.Items {
		if strings.Contains(item.Name, namespace) {
			return namespace
		}
	}
	return ""
}
