package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/huhx-headhunter/headhunter-common/oss/config"
	"github.com/huhx-headhunter/headhunter-common/oss/model"
)

type OssService struct {
	properties *config.OssProperties
}

func NewOssService() *OssService {
	return &OssService{config.NewOssProperties()}
}

func (uc *OssService) GetOssCredential() (*model.OssCredential, error) {
	if credential, _ := GetOssCredential(); credential != nil {
		return credential, nil
	}

	client, err := sts.NewClientWithAccessKey(uc.properties.Region, uc.properties.AccessKey, uc.properties.SecretKey)
	if err != nil {
		fmt.Println("Error creating STS client:", err)
		return nil, err
	}
	response, err := client.AssumeRole(uc.createRequest())
	if err != nil {
		fmt.Println("Error assuming role:", err)
		return nil, err
	}

	ossCredential := model.NewOssCredential(response)
	if saveErr := SaveOssCredential(ossCredential); saveErr != nil {
		return nil, saveErr
	}
	return ossCredential, nil
}

func (uc *OssService) createRequest() *sts.AssumeRoleRequest {
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.RoleArn = uc.properties.RoleArn
	request.RoleSessionName = uc.properties.SessionName
	return request
}
