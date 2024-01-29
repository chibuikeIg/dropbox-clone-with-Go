package repository

import (
	"context"
	"file-upload-service/internals/app/config"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type AwsS3 struct {
	client *s3.Client
	ctx    context.Context
	Bucket *string
}

func NewAwsS3(bucket *string) *AwsS3 {

	// aws client config
	cfg := config.NewAwsCredentials()
	client := s3.NewFromConfig(aws.Config{
		Region:      os.Getenv("AWS_REGION"),
		Credentials: cfg,
	})

	ctx := context.TODO()

	return &AwsS3{client: client, ctx: ctx, Bucket: bucket}
}

func (awsS3 *AwsS3) CreateMultipartUpload(data map[string]*string) (string, error) {
	uploadOuput, err := awsS3.client.CreateMultipartUpload(awsS3.ctx, &s3.CreateMultipartUploadInput{
		Bucket:      awsS3.Bucket,
		Key:         data["object_key"],
		ContentType: data["content_type"],
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	return *uploadOuput.UploadId, nil
}

func (awsS3 *AwsS3) UploadPart(data map[string]any) (string, error) {

	uploadOutput, err := awsS3.client.UploadPart(awsS3.ctx, &s3.UploadPartInput{
		Bucket:     awsS3.Bucket,
		Key:        data["object_key"].(*string),
		PartNumber: data["part_number"].(*int32),
		UploadId:   data["upload_id"].(*string),
		Body:       data["requestBody"].(io.Reader),
	})

	if err != nil {
		return "", err
	}

	return *uploadOutput.ETag, nil
}

func (awsS3 *AwsS3) CompleteMultipartUpload(data map[string]any) (string, error) {

	var completedParts []types.CompletedPart

	for _, part := range data["completed_parts"].([]map[string]any) {

		if len(part) > 0 {
			completedPart := types.CompletedPart{
				ETag:       part["ETag"].(*string),
				PartNumber: part["PartNumber"].(*int32),
			}
			completedParts = append(completedParts, completedPart)
		}

	}

	completedMultiPartUpload := &types.CompletedMultipartUpload{
		Parts: completedParts,
	}

	response, err := awsS3.client.CompleteMultipartUpload(awsS3.ctx, &s3.CompleteMultipartUploadInput{
		Bucket:          awsS3.Bucket,
		Key:             data["object_key"].(*string),
		UploadId:        data["upload_id"].(*string),
		MultipartUpload: completedMultiPartUpload,
	})

	if err != nil {
		return "", err
	}

	return *response.Key, nil
}

func (awsS3 *AwsS3) AbortMultipartUpload(data map[string]any) (bool, error) {

	_, err := awsS3.client.AbortMultipartUpload(awsS3.ctx, &s3.AbortMultipartUploadInput{
		Bucket:   awsS3.Bucket,
		Key:      data["object_key"].(*string),
		UploadId: data["upload_id"].(*string),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
