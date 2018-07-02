package main

import (
	"os/exec"
	"io/ioutil"
)

func pullK8sFile(projectName string, repostory string, tag string) (filestring string) {

	cmd := exec.Command("/bin/bash", "-c", "mkdir "+projectName)
	cmd.Run()

	cmd2 := exec.Command("/bin/bash", "-c", "git init")
	cmd2.Dir = "./" + projectName
	cmd2.Run()

	cmd3 := exec.Command("/bin/bash", "-c", "git config core.sparsecheckout true")
	cmd3.Dir = "./" + projectName
	cmd3.Run()

	cmd4 := exec.Command("/bin/bash", "-c", "echo 'k8s.yml' >> .git/info/sparse-checkout")
	cmd4.Dir = "./" + projectName
	cmd4.Run()

	cmd5 := exec.Command("/bin/bash", "-c", "git remote add origin "+repostory)
	cmd5.Dir = "./" + projectName
	cmd5.Run()

	//todo
	//cmd6 := exec.Command("/bin/bash", "-c", "git pull origin "+tag)
	cmd6 := exec.Command("/bin/bash", "-c", "git pull origin k8s")

	cmd6.Dir = "./" + projectName
	cmd6.Run()

	filecontent, _ := ioutil.ReadFile("./" + projectName + "/k8s.yml")
	filestring = string(filecontent)
	return filestring
}
