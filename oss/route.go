package oss

import (
	"github.com/gin-gonic/gin"
	"github.com/huhx-headhunter/headhunter-common/oss/controller"
	"github.com/huhx-headhunter/headhunter-common/oss/service"
)

func Routes(router *gin.RouterGroup) {
	ossService := service.NewOssService()
	ossController := controller.NewOssController(ossService)

	group := router.Group("/oss")
	{
		group.GET("/sts", ossController.GetOssCredential)
	}
}
