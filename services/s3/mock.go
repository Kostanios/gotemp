package s3service

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var ResponseData []interface{}
var responseDataIndex int

type S3Mock struct {
	s3.Client
}

func Mock() Interface {
	return S3Mock{}
}

func Reset() {
	ResponseData = nil
	responseDataIndex = 0
}

func (mock S3Mock) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {

	if params.Bucket == nil {
		panic("This should not happen bro...")
	}

	if params.Key == nil {
		panic("This should not happen bro...")
	}

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err,ok := result.(error); ok {
		return nil,err
	}

	if result == nil {
		return nil,nil
	}

	if v, ok := result.(s3.PutObjectOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}

func (mock S3Mock) 	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {

	if params.Bucket == nil {
		panic("This should not happen bro...")
	}

	if params.Key == nil {
		panic("This should not happen bro...")
	}

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err,ok := result.(error); ok {
		return nil,err
	}

	if result == nil {
		return nil,nil
	}

	if v, ok := result.(s3.DeleteObjectOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}

func (mock S3Mock) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {

	if params.Bucket == nil {
		panic("This should not happen bro...")
	}

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err,ok := result.(error); ok {
		return nil,err
	}

	if result == nil {
		return nil,nil
	}

	if v, ok := result.(s3.ListObjectsV2Output); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}