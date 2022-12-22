package workload

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"strings"
)

func GetPods() {
	clientset := InitCluster()
	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatal("error ", err)
		}
		fmt.Printf("pod aree  ***** ", pods)
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			log.Fatal("error ", err)
		} else {
			fmt.Printf("Found example-xxxxx pod in default namespace\n")
		}
	}
}

func GetDeployment(deploymentName, namespace string) (string, int32, v1.Deployment) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("error ", err)
	}
	for _, item := range deployments.Items {
		if strings.Contains(item.Name, deploymentName) {
			return deploymentName, item.Status.Replicas, item
		}
	}
	return "", 0, v1.Deployment{}
}

func GetDeployments(namespace string) []string {
	var deploymentList []string
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("error ", err)
	}
	for _, item := range deployments.Items {
		deploymentList = append(deploymentList, item.Name)
	}
	return deploymentList
}
