package api

import (
	"github.com/gin-gonic/gin"
	"mermaid/api/ginreturn"
	"mermaid/service"
	"net/http"
)

func GetArticles(c *gin.Context) {
	data, err := service.DefaultArticle.GetAllArticles()
	if err != nil {
		ginreturn.GinReturn(c, http.StatusOK, nil, err)
		return
	}
	ginreturn.GinReturn(c, http.StatusOK, data, nil)
}
