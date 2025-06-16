package view

import "github.com/gin-gonic/gin"

func Render(c *gin.Context, code int, name string, data gin.H){
	templateData, _ := c.Get("templateData")
	baseData, _ := templateData.(gin.H)

	for k, v := range data {
		baseData[k] = v
	}

	c.HTML(code,name,baseData)
}