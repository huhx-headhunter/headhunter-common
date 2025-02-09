package config

import "os"

type OssProperties struct {
	Region      string `json:"region"`
	AccessKey   string `json:"accessKey"`
	SecretKey   string `json:"secretKey"`
	RoleArn     string `json:"roleArn"`
	SessionName string `json:"sessionName"`
	Host        string `json:"host"`
}

func NewOssProperties() *OssProperties {
	return &OssProperties{
		Region:      os.Getenv("oss_region"),
		AccessKey:   os.Getenv("oss_access_key"),
		SecretKey:   os.Getenv("oss_secret_key"),
		RoleArn:     os.Getenv("oss_role_arn"),
		SessionName: os.Getenv("oss_session_name"),
		Host:        os.Getenv("oss_host"),
	}
}
