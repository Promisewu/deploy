package main

const (
	StatusWaiting  = 1
	StatusStarting = 2
	StatusEnd      = 3
)

var statusMap = map[uint]string{
	1: "未发布",
	2: "发布中",
	3: "已发布",
}

type Project struct {
	Id         uint
	Name       string
	Repository string
}

type deploy struct {
	Id        uint
	Name      string
	Relations []DepProRelation
}

type DepProRelation struct {
	ProjectId uint
	TagName   string
	Ordering  uint
}

type Env struct {
	Id        uint
	Name      string
	Config    string
	Namespace string
}

type Job struct {
	Id       uint
	DeployId uint
	EnvId    uint
	Status   uint
	Time     string
	Log      []string
}

type projectForm struct {
	Name       string `json:"name" binding:"required"`
	Repository string `json:"Repository" binding:"required"`
}

type DeployForm struct {
	Name      string               `json:"name" binding:"required"`
	Relations []DepProRelationForm `json:"relations"`
}

type DepProRelationForm struct {
	ProjectId uint   `json:"projectId" binding:"required"`
	TagName   string `json:"tagName" binding:"required"`
	Ordering  uint   `json:"ordering" binding:"required"`
}

type envForm struct {
	Name      string `json:"name" binding:"required"`
	Config    string `json:"config" binding:"required"`
	Namespace string `json:"namespace" binding:"required"`
}
