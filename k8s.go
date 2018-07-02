package main

import "os/exec"

func pullK8sFile() {

	repostory := "ssh://git@gitlab-wenba.xueba100.com:9922/aixue-open/aixue-exam-new.git"
	tag := "k8s"
	cmd := exec.Command("/bin/bash", "-c", "mkdir aixue-homework")
	cmd.Run()

	cmd2 := exec.Command("/bin/bash", "-c", "git init")
	cmd2.Dir = "./aixue-homework"
	cmd2.Run()

	cmd3 := exec.Command("/bin/bash", "-c", "git config core.sparsecheckout true")
	cmd3.Dir = "./aixue-homework"
	cmd3.Run()

	cmd4 := exec.Command("/bin/bash", "-c", "echo 'k8s.yml' >> .git/info/sparse-checkout")
	cmd4.Dir = "./aixue-homework"
	cmd4.Run()

	cmd5 := exec.Command("/bin/bash", "-c", "git remote add origin "+repostory)
	cmd5.Dir = "./aixue-homework"
	cmd5.Run()

	cmd6 := exec.Command("/bin/bash", "-c", "git pull origin "+tag)
	cmd6.Dir = "./aixue-homework"
	cmd6.Run()
}
