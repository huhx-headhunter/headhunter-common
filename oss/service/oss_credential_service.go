package service

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/huhx-headhunter/headhunter-common/db"
	"github.com/huhx-headhunter/headhunter-common/oss/model"
	"time"
)

var credentialCtx = context.Background()

const OssCredentialKey = "oss_credentials"

func SaveOssCredential(credential *model.OssCredential) error {
	expireDate, err := time.Parse(time.RFC3339, credential.Expiration)
	if err != nil {
		return fmt.Errorf("error parsing expiration date: %v", err)
	}
	seconds := int(expireDate.Sub(time.Now()).Seconds())

	credentialsJSON, err := json.Marshal(credential)
	if err != nil {
		return fmt.Errorf("error marshalling credential: %v", err)
	}

	err = db.Redis.Set(credentialCtx, OssCredentialKey, credentialsJSON, time.Duration(seconds)*time.Second).Err()
	if err != nil {
		return fmt.Errorf("error saving credential to Redis: %v", err)
	}

	return nil
}

func GetOssCredential() (*model.OssCredential, error) {
	var credentials model.OssCredential
	credentialsJSON, err := db.Redis.Get(credentialCtx, OssCredentialKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting credentials from Redis: %v", err)
	}

	if err = json.Unmarshal([]byte(credentialsJSON), &credentials); err != nil {
		return nil, fmt.Errorf("error unmarshalling credentials: %v", err)
	}

	return &credentials, nil
}
