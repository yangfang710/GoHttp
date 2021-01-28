package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"pkg/mod/github.com/hashicorp/go-multierror@v1.1.0"

	"GoHttp/env"
	"GoHttp/env/global"
	"GoHttp/model"
)

const (
	prefixRoute = "/api"
)

type Router struct {
	GinEngine *gin.Engine
}

func NewRouter(s Servlet) *Router {

	engine := gin.New()

	if global.Config.Metrics.Addr != "" {
		p := ginprometheus.NewPrometheus("http")
		p.ReqCntURLLabelMappingFn = (*gin.Context).HandlerName
		p.SetListenAddress(global.Config.Metrics.Addr)
		p.Use(engine)
	}

	engine.Use(
		cors.New(env.CorsConfig()),
	)

	engine.Static("/apidoc", "./apidoc")

	group := engine.Group(prefixRoute)
	group.GET("/example", s.Example)

	{
		catProfileGroup := group.Group("/cat")
		catProfileGroup.GET("/profile", s.GetCatProfile)
		catProfileGroup.POST("/profile", s.AddCatProfile)
	}

	{
		bbsGroup := group.Group("/bbs")

		bbsGroup.GET("/tags", s.GetTags)
		bbsGroup.POST("/tag", s.CreateMeowTag)

		bbsGroup.POST("/tiezi", s.CreateMeowTiezi)
		bbsGroup.GET("/list", s.GetBbsList)
		bbsGroup.GET("/tiezi", s.GetTieziDetail)
		bbsGroup.GET("/like", s.ClickGood)
		bbsGroup.POST("/comment", s.CreateMeowComment)
	}

	return &Router{GinEngine: engine}
}

func (r *Router) Close() error {
	var dbErr error
	dbEngine := model.GetEngine()
	if dbEngine != nil {
		dbErr = dbEngine.Close()
	}
	return multierror.Append(dbErr).ErrorOrNil()
}
