package main

import (
	"path/filepath"
	"strings"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func doDeploy() gin.HandlerFunc {
	return func(c *gin.Context) {

		deployId := getUintId(c, "deployId")
		deploy := deployMap[deployId]
		if deploy == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "没有该部署",
			})
			return
		}

		envId := getUintId(c, "envId")
		env := envMap[envId]
		if env == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "没有该环境",
			})
			return
		}

		go doJob(deployId, envId)

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "发布成功",
		})
		return
	}
}
func doJob(deployId uint, envId uint) {
	deploy := deployMap[deployId]
	env := envMap[envId]

	oldJob := jobMap[deployId][envId]
	oldJob.Status = StatusStarting
	oldJob.Log = append(oldJob.Log, dateTime()+"starting create client")

	namespace := env.Namespace
	projectList := deploy.Relations

	// create client
	clientset := createClient(envId)

	oldJob.Log = append(oldJob.Log, dateTime()+"had created client")

	//deploy project
	for _, val := range projectList {
		projectId := val.ProjectId
		tagName := val.TagName
		deployProject(clientset, namespace, projectId, tagName, &(oldJob.Log))
	}

	oldJob.Status = StatusEnd
}

func deployProject(clientset *kubernetes.Clientset, namespace string, projectId uint, tagName string, logList *[]string) {
	project := projectMap[projectId]
	projectName := project.Name
	repostory := project.Repository

	time.Sleep(1e9)
	*logList = append(*logList, dateTime()+"start to pull k8s file...")
	// 拉取 k8s.yml 文件
	fileString := pullK8sFile(projectName, repostory, tagName)

	*logList = append(*logList, dateTime()+"end to pull k8s file")

	time.Sleep(1e9)
	*logList = append(*logList, dateTime()+"start to pull split k8s file")
	// k8s文件 根据 "---" 拆分成多个string 进行部署
	k8sSlice := splitFile(fileString)
	*logList = append(*logList, dateTime()+"end to pull split k8s file")

	time.Sleep(1e9)
	*logList = append(*logList, dateTime()+"start to deploy "+projectName+"...")
	for _, val := range k8sSlice {
		decode := scheme.Codecs.UniversalDeserializer().Decode
		obj, versionKind, _ := decode([]byte(val), nil, nil)

		// 获取需要部署的类型
		kindType := versionKind.Kind

		// 根据部署类型选择资源部署
		switch kindType {
		case "Deployment":
			applyDeployment(clientset, namespace, projectName, tagName, obj, logList)
		case "Service":
			applyService(clientset, namespace, projectName, obj, logList)
		}
	}
	*logList = append(*logList, dateTime()+"end to deploy "+projectName+"...")
}

func createClient(envId uint) (clientset *kubernetes.Clientset) {
	var kubeconfig string
	var err error
	var config *rest.Config

	env := envMap[envId]
	configName := env.Config
	kubeconfig = filepath.Join("config", configName)
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func splitFile(fileString string) (k8sSlice []string) {
	k8sSlice = strings.Split(fileString, "---")
	return k8sSlice
}
