package main

import (
	"flag"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/homedir"
	"io/ioutil"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
	"strings"
)

func test111() {
	// create the clientset
	config := getConfig()

	//create client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// separate k8s to deployment service
	splitFile()

	// create deployment
	createDeployment(clientset)

	// create service
	createService(clientset)
}

func splitFile() {
	filecontent, _ := ioutil.ReadFile("./aixue-homework/k8s.yml")
	filestring := string(filecontent)
	arr := strings.Split(filestring, "---")
	deploymentData := []byte(arr[0])
	_ = ioutil.WriteFile("./aixue-homework/k8s.deployment.yml", deploymentData, 0644)

	serviceData := []byte(arr[1])
	_ = ioutil.WriteFile("./aixue-homework/k8s.service.yml", serviceData, 0644)
}

func createService(clientset *(kubernetes.Clientset)) {
	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	filecontent, _ := ioutil.ReadFile("./aixue-homework/k8s.service.yml")
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, _ := decode(filecontent, nil, nil)
	service := obj.(*apiv1.Service)

	fmt.Println("Creating service...")

	result, err := serviceClient.Create(service)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created service %q.\n", result.GetObjectMeta().GetName())
}

func createDeployment(clientset *(kubernetes.Clientset)) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	filecontent, _ := ioutil.ReadFile("./aixue-homework/k8s.deployment.yml")
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, _ := decode(filecontent, nil, nil)
	deployment := obj.(*appsv1.Deployment)

	fmt.Println("Creating deployment...")

	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func getConfig() (config *rest.Config) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	return config
}
