package dynamodbservice

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var ResponseData []interface{}
var responseDataIndex int

type DynamodbMock struct {
	dynamodb.Client
}

func Mock() Interface {
	return DynamodbMock{}
}

func Reset() {
	ResponseData = nil
	responseDataIndex = 0
}

func (mock DynamodbMock) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(options *dynamodb.Options)) (*dynamodb.GetItemOutput, error) {

	if params.TableName == nil {
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

	if v, ok := result.(dynamodb.GetItemOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}

func (mock DynamodbMock) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {

	if params.TableName == nil {
		panic("This should not happen bro...")
	}

	if params.Item == nil {
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

	if v, ok := result.(dynamodb.PutItemOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}

func (mock DynamodbMock) Scan(ctx context.Context, input *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {

	if input.TableName == nil {
		panic("This should not happen bro...")
	}

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err,ok := result.(error); ok {
		return nil,err
	}

	if result == nil {
		return &dynamodb.ScanOutput{
			Items:[]map[string]types.AttributeValue{},
		},nil
	}

	if v, ok := result.(dynamodb.ScanOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}

func (mock DynamodbMock) Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {

	if input.TableName == nil {
		panic("This should not happen bro...")
	}

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err,ok := result.(error); ok {
		return nil,err
	}

	if result == nil {
		return &dynamodb.QueryOutput{
			Items:[]map[string]types.AttributeValue{},
		},nil
	}

	if v, ok := result.(dynamodb.QueryOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}

func (mock DynamodbMock) UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {

	if input.TableName == nil {
		panic("This should not happen bro...")
	}

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err,ok := result.(error); ok {
		return nil,err
	}

	if result == nil {
		return &dynamodb.UpdateItemOutput{},nil
	}

	if v, ok := result.(dynamodb.UpdateItemOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}

func (mock DynamodbMock) DeleteItem(ctx context.Context, input *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {

	if input.TableName == nil {
		panic("This should not happen bro...")
	}

	if input.Key == nil {
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

	if v, ok := result.(dynamodb.DeleteItemOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")
}

func (mock DynamodbMock) BatchGetItem(ctx context.Context, input *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err,ok := result.(error); ok {
		return nil,err
	}

	if result == nil {
		return nil,nil
	}

	if v, ok := result.(dynamodb.BatchGetItemOutput); ok {
		return &v,nil
	}

	panic("This should not happen bro...")

}