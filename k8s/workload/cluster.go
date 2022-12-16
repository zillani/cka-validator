package workload

import (
	"k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func InitCluster() *kubernetes.Clientset {
	// Bootstrap k8s configuration from local Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Println("Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	// Create an rest client not targeting specific API version
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return clientset
}
