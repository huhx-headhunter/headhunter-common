package model

import "github.com/aliyun/alibaba-cloud-sdk-go/services/sts"

type OssCredential struct {
	SecurityToken   string `json:"SecurityToken"`
	AccessKeyId     string `json:"AccessKeyId"`
	StatusCode      int    `json:"StatusCode"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Expiration      string `json:"Expiration"`
}

func NewOssCredential(response *sts.AssumeRoleResponse) *OssCredential {
	return &OssCredential{
		AccessKeyId:     response.Credentials.AccessKeyId,
		AccessKeySecret: response.Credentials.AccessKeySecret,
		SecurityToken:   response.Credentials.SecurityToken,
		StatusCode:      response.GetHttpStatus(),
		Expiration:      response.Credentials.Expiration,
	}
}
