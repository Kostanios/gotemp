package stripe

import (
	"bytes"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

func Mock() *stripe.Backends {
	return &stripe.Backends{
		API:StripeMock{},
	}
}

type StripeMock struct {}

var ResponseData []interface{}
var responseDataIndex int

func Reset() {
	ResponseData = nil
	responseDataIndex = 0
}

func (mock StripeMock) Call(method, path, key string, params stripe.ParamsContainer, v stripe.LastResponseSetter) error {
	if v,ok := v.(*stripe.Source); ok {
		if mockValue,ok := ResponseData[responseDataIndex].(stripe.Source); ok {
			*v = mockValue
			responseDataIndex += 1
			return nil
		}
	}

	if v,ok := v.(*stripe.PaymentIntent); ok {
		if mockValue,ok := ResponseData[responseDataIndex].(stripe.PaymentIntent); ok {
			*v = mockValue
			responseDataIndex += 1
			return nil
		}
	}

	if err,ok := ResponseData[responseDataIndex].(error); ok {
		responseDataIndex += 1
		return err
	}

	if  ResponseData[responseDataIndex] == nil {
		return nil
	}

	panic("This should not happen!")

	return nil
}

func (mock StripeMock) CallStreaming(method, path, key string, params stripe.ParamsContainer, v stripe.StreamingLastResponseSetter) error {
	panic("implement me")
}

func (mock StripeMock) CallRaw(method, path, key string, body *form.Values, params *stripe.Params, v stripe.LastResponseSetter) error {
	if v,ok := v.(*stripe.Source); ok {
		if mockValue,ok := ResponseData[responseDataIndex].(stripe.Source); ok {
			*v = mockValue
			responseDataIndex += 1
			return nil
		}
	}

	if v,ok := v.(*stripe.PaymentIntent); ok {
		if mockValue,ok := ResponseData[responseDataIndex].(stripe.PaymentIntent); ok {
			*v = mockValue
			responseDataIndex += 1
			return nil
		}
	}

	if err,ok := ResponseData[responseDataIndex].(error); ok {
		responseDataIndex += 1
		return err
	}

	if  ResponseData[responseDataIndex] == nil {
		return nil
	}

	panic("This should not happen!")

	return nil
}

func (mock StripeMock) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *stripe.Params, v stripe.LastResponseSetter) error {
	if v,ok := v.(*stripe.Source); ok {
		if mockValue,ok := ResponseData[responseDataIndex].(stripe.Source); ok {
			*v = mockValue
			responseDataIndex += 1
			return nil
		}
	}

	if v,ok := v.(*stripe.PaymentIntent); ok {
		if mockValue,ok := ResponseData[responseDataIndex].(stripe.PaymentIntent); ok {
			*v = mockValue
			responseDataIndex += 1
			return nil
		}
	}

	if err,ok := ResponseData[responseDataIndex].(error); ok {
		responseDataIndex += 1
		return err
	}

	if  ResponseData[responseDataIndex] == nil {
		return nil
	}

	panic("This should not happen!")

	return nil
}

func (mock StripeMock) SetMaxNetworkRetries(maxNetworkRetries int64) {
	panic("implement me")
}