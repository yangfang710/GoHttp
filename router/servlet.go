package router

import (
	"github.com/gin-gonic/gin"
)

type Servlet interface {
	Example(c *gin.Context)

	AddCatProfile(c *gin.Context)
	GetCatProfile(c *gin.Context)

	GetTags(c *gin.Context)
	CreateMeowTag(c *gin.Context)
	CreateMeowTiezi(c *gin.Context)
	GetBbsList(c *gin.Context)
	GetTieziDetail(c *gin.Context)
	ClickGood(c *gin.Context)
	CreateMeowComment(c *gin.Context)
}
