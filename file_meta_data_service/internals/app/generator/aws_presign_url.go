package generator

import (
	"context"
	"filemetadata-service/internals/app/config"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsPresignUrl struct {
	client    *s3.PresignClient
	ctx       context.Context
	ObjectKey *string
	Bucket    *string
}

func NewAwsPresignUrl(objectKey, bucket *string) *AwsPresignUrl {

	// aws client config
	cfg := config.NewAwsCredentials()
	client := s3.NewFromConfig(aws.Config{
		Region:      os.Getenv("AWS_REGION"),
		Credentials: cfg,
	})

	presignClient := s3.NewPresignClient(client, func(po *s3.PresignOptions) {
		po.Expires = 168 * time.Hour
	})

	return &AwsPresignUrl{client: presignClient, ObjectKey: objectKey, Bucket: bucket, ctx: context.TODO()}
}

func (apu AwsPresignUrl) Generate() (string, error) {

	v4PresignedHttpReq, err := apu.client.PresignGetObject(apu.ctx, &s3.GetObjectInput{
		Bucket: apu.Bucket,
		Key:    apu.ObjectKey,
	})

	if err != nil {
		return "", err
	}

	return v4PresignedHttpReq.URL, nil
}
