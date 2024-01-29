package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Credentials struct{}

func NewCredentials() *Credentials {
	return &Credentials{}
}

func (c *Credentials) Retrieve(ctx context.Context) (aws.Credentials, error) {

	return aws.Credentials{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
	}, nil
}
