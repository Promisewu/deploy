package main

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


