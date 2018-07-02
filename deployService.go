package main

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1Type "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/runtime"
)

func applyService(clientset *kubernetes.Clientset, namespace string, projectName string, obj runtime.Object, logList *[]string) {
	serviceClient := clientset.CoreV1().Services(namespace)
	result, err := serviceClient.Get(projectName, metav1.GetOptions{})

	service := obj.(*apiv1.Service)

	if err == nil {
		updateService(service, serviceClient, projectName, result, logList)
	} else {
		createService(service, serviceClient, projectName, logList)
	}
	return
}

func updateService(service *apiv1.Service, serviceClient apiv1Type.ServiceInterface, projectName string, result *apiv1.Service, logList *[]string) {

	*logList= append(*logList,dateTime()+ "Updating service "+projectName+"...")
	fmt.Println("Updating service" + projectName + "...")
	service.Name = projectName
	service.ObjectMeta.ResourceVersion = result.ObjectMeta.ResourceVersion
	service.Spec.ClusterIP = result.Spec.ClusterIP

	result, err := serviceClient.Update(service)
	if err != nil {
		panic(err)
	}
	*logList= append(*logList,dateTime()+ "Updated service "+projectName+"...")
	fmt.Printf("Updated service %q.\n", result.GetObjectMeta().GetName())
}

func createService(service *apiv1.Service, serviceClient apiv1Type.ServiceInterface, projectName string, logList *[]string) {

	*logList= append(*logList,dateTime()+ "Updating service "+projectName+"...")
	fmt.Println("Creating service" + projectName + "...")

	result, err := serviceClient.Create(service)
	if err != nil {
		panic(err)
	}

	*logList= append(*logList,dateTime()+ "Created service "+projectName+"...")
	fmt.Printf("Created service %q.\n", result.GetObjectMeta().GetName())
}
