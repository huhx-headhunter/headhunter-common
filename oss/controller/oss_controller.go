package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/huhx-headhunter/headhunter-common/oss/service"
	"github.com/huhx/common-go/exception"
	"net/http"
)

type OssController struct {
	service *service.OssService
}

func NewOssController(service *service.OssService) *OssController {
	return &OssController{service}
}

// GetOssCredential godoc
//
//	@Tags			Oss Resource
//	@Summary		Get OSS credential
//	@Description	Get OSS credential for uploading files
//	@Success		200	{object}	model.OssCredential
//	@Failure		500	{object}	swag.SystemError
//	@Router			/oss/credential [get]
//	@Security		BearerAuth
func (uc *OssController) GetOssCredential(c *gin.Context) {
	credential, err := uc.service.GetOssCredential()
	if err != nil {
		panic(exception.System{Content: err.Error()})
	}

	c.JSON(http.StatusOK, credential)
}
