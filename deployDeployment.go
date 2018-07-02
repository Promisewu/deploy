package main

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1Type "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/runtime"
	"strings"
)

func applyDeployment(clientset *kubernetes.Clientset, namespace string, projectName string, tag string, obj runtime.Object, logList *[]string) {

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	_, err := deploymentsClient.Get(projectName, metav1.GetOptions{})

	deployment := obj.(*appsv1.Deployment)
	image := deployment.Spec.Template.Spec.Containers[0].Image
	slice := strings.Split(image, ":")
	slice[1] = tag
	deployment.Spec.Template.Spec.Containers[0].Image = strings.Join(slice, ":")

	if err == nil {
		updateDeployment(deployment, deploymentsClient, projectName, logList)
	} else {
		createDeployment(deployment, deploymentsClient, projectName, logList)
	}
	return
}

func updateDeployment(deployment *appsv1.Deployment, deploymentsClient appsv1Type.DeploymentInterface, projectName string, logList *[]string) {

	*logList= append(*logList,dateTime()+ "Updating deployment "+projectName+"...")

	fmt.Println("Updating deployment " + projectName + "...")
	deployment.Name = projectName

	result, err := deploymentsClient.Update(deployment)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Updated deployment %q.\n", result.GetObjectMeta().GetName())
	*logList= append(*logList,dateTime() + "Updated deployment "+projectName+"...")

	return
}

func createDeployment(deployment *appsv1.Deployment, deploymentsClient appsv1Type.DeploymentInterface, projectName string, logList *[]string) {

	*logList= append(*logList, dateTime() + "Creating deployment "+projectName+"...")

	fmt.Println("Creating deployment " + projectName + "...")

	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	*logList= append(*logList,dateTime() + "Created deployment "+projectName+"...")
	return
}
