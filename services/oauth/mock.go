package oauth

var ResponseData []interface{}
var responseDataIndex int

type ServiceMock struct{}

func Mock() Interface {
	return ServiceMock{}
}

func Reset() {
	ResponseData = nil
	responseDataIndex = 0
}

func (mock ServiceMock) LoadClient(clientId string) (error, *Client) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	if v, ok := result.(*Client); ok {
		return nil, v
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) RemoveClient(clientId string) error {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err
	}

	if result == nil {
		return nil
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) GenerateClient(input ClientInput) (error, *Client) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	if v, ok := result.(*Client); ok {
		return nil, v
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) LoadAccessToken(token string) (error, *AccessToken) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	if v, ok := result.(*AccessToken); ok {
		return nil, v
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) RemoveAccessToken(token string) error {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err
	}

	if result == nil {
		return nil
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) GenerateAccessToken(input AccessTokenInput) (error, *AccessToken) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	if v, ok := result.(*AccessToken); ok {
		return nil, v
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) LoadRefreshToken(refreshToken string) (error, *RefreshToken) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	if v, ok := result.(*RefreshToken); ok {
		return nil, v
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) RemoveRefreshToken(refreshToken string) error {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err
	}

	if result == nil {
		return nil
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) GenerateRefreshToken(input RefreshTokenInput) (error, *RefreshToken) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	if v, ok := result.(*RefreshToken); ok {
		return nil, v
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) LoadTwoFactorToken(id string) (error, *TwoFactorToken) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	if v, ok := result.(*TwoFactorToken); ok {
		return nil, v
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) RemoveTwoFactorToken(id string) error {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err
	}

	if result == nil {
		return nil
	}

	panic("This should not happen bro...")
}

func (mock ServiceMock) GenerateTwoFactorToken(input TwoFactorTokenInput) (error, *TwoFactorToken) {

	result := ResponseData[responseDataIndex]
	responseDataIndex += 1

	if err, ok := result.(error); ok {
		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	if v, ok := result.(*TwoFactorToken); ok {
		return nil, v
	}

	panic("This should not happen bro...")
}
