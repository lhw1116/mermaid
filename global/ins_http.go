package global

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	api2 "mermaid/api"
	"mermaid/router"
)

var r *gin.Engine

func InitHttp(addr string) error {
	r = gin.New()

	// default is "debug/pprof"
	pprof.Register(r)
	router.RegisterRouter(r)
	api := r.Group("api")

	{
		api.GET("/articles", api2.GetArticles)
		//api.GET("/article/:id")
	}
	r.Run(addr)
	return nil
}

func GetRouter() *gin.Engine {
	return r
}
