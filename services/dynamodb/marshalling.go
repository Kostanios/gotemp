package dynamodbservice

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Unmarshal(av types.AttributeValue, out interface{}) error {
	return attributevalue.NewDecoder(func(options *attributevalue.DecoderOptions) {
		options.TagKey = "json"
	}).Decode(av, out)
}

func UnmarshalMap(m map[string]types.AttributeValue, out interface{}) error {
	return attributevalue.NewDecoder(func(options *attributevalue.DecoderOptions) {
		options.TagKey = "json"
	}).Decode(&types.AttributeValueMemberM{Value: m}, out)
}

func UnmarshalList(l []types.AttributeValue, out interface{}) error {
	return attributevalue.NewDecoder(func(options *attributevalue.DecoderOptions) {
		options.TagKey = "json"
	}).Decode(&types.AttributeValueMemberL{Value: l}, out)
}

func UnmarshalListOfMaps(l []map[string]types.AttributeValue, out interface{}) error {
	items := make([]types.AttributeValue, len(l))
	for i, m := range l {
		items[i] = &types.AttributeValueMemberM{Value: m}
	}

	return UnmarshalList(items, out)
}

func Marshal(in interface{}) (types.AttributeValue, error) {
	return attributevalue.NewEncoder(func(options *attributevalue.EncoderOptions) {
		options.TagKey = "json"
	}).Encode(in)
}

func MarshalMap(in interface{}) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.NewEncoder(func(options *attributevalue.EncoderOptions) {
		options.TagKey = "json"
	}).Encode(in)

	asMap, ok := av.(*types.AttributeValueMemberM)
	if err != nil || av == nil || !ok {
		return map[string]types.AttributeValue{}, err
	}

	return asMap.Value, nil
}

func MarshalList(in interface{}) ([]types.AttributeValue, error) {
	av, err := attributevalue.NewEncoder(func(options *attributevalue.EncoderOptions) {
		options.TagKey = "json"
	}).Encode(in)

	asList, ok := av.(*types.AttributeValueMemberL)
	if err != nil || av == nil || !ok {
		return []types.AttributeValue{}, err
	}

	return asList.Value, nil
}