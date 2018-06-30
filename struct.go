package main

const (
	PRIVATE_TOKEN_NAME  = "PRIVATE-TOKEN"
	PRIVATE_TOKEN_VALUE = "i7TGyKz748BYigXqh9fB"
	GITLAB_ADDRESS      = "https://gitlab-wenba.xueba100.com:2443"
)

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
	Id uint
	Name string
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
	Name string `json:"name" binding:"required"`
}
