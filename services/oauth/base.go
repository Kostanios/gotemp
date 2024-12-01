package oauth

import (
	"github.com/kostanios/gotemp/services/random"
	"gorm.io/gorm"
	"time"

	"github.com/twinj/uuid"
)

type Interface interface {
	LoadClient(clientId string) (error, *Client)
	RemoveClient(clientId string) error
	GenerateClient(input ClientInput) (error, *Client)

	LoadAccessToken(token string) (error, *AccessToken)
	RemoveAccessToken(token string) error
	GenerateAccessToken(input AccessTokenInput) (error, *AccessToken)

	LoadRefreshToken(refreshToken string) (error, *RefreshToken)
	RemoveRefreshToken(refreshToken string) error
	GenerateRefreshToken(input RefreshTokenInput) (error, *RefreshToken)

	LoadTwoFactorToken(id string) (error, *TwoFactorToken)
	RemoveTwoFactorToken(id string) error
	GenerateTwoFactorToken(input TwoFactorTokenInput) (error, *TwoFactorToken)
}

type Auth struct {
	AccessTokenExpiresAfter  int64
	RefreshTokenExpiresAfter int64
	TwoFactorExpiresAfter    int64

	// Services
	db gorm.DB
}

func New(db gorm.DB, input NewInput) Interface {
	return Auth{
		AccessTokenExpiresAfter:  input.AccessTokenExpiresAfter,
		RefreshTokenExpiresAfter: input.RefreshTokenExpiresAfter,
		TwoFactorExpiresAfter:    input.TwoFactorExpiresAfter,
		db:                       db,
	}
}

func (a Auth) LoadClient(clientId string) (error, *Client) {

	var result *Client

	if err := a.db.Table(Client{}.TableName()).Where("identifier = ?", clientId).First(&result).Error; err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}

		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	return nil, result
}

func (a Auth) RemoveClient(clientId string) error {

	if err := a.db.Table(Client{}.TableName()).Where("identifier = ?", clientId).Delete(&Client{}).Error; err != nil {
		return err
	}

	return nil
}

func (a Auth) GenerateClient(input ClientInput) (error, *Client) {

	client := Client{
		Name:           input.Name,
		Created:        time.Now().Unix(),
		SecretRequired: input.SecretRequired,
	}

	if input.SecretRequired {
		client.Secret = uuid.NewV1().String()
	}

	var result *Client

	if err := a.db.Table(Client{}.TableName()).Create(&client).Error; err != nil {
		return err, nil
	}

	return nil, result
}

func (a Auth) LoadAccessToken(token string) (error, *AccessToken) {

	var result *AccessToken

	if err := a.db.Table(AccessToken{}.TableName()).Where("access_token = ?", token).First(&result).Error; err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}

		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	return nil, result
}

func (a Auth) RemoveAccessToken(token string) error {

	if err := a.db.Table(AccessToken{}.TableName()).Where("access_token = ?", token).Delete(&AccessToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (a Auth) GenerateAccessToken(input AccessTokenInput) (error, *AccessToken) {

	accessToken := AccessToken{
		AccessToken:  uuid.NewV4().String(),
		Created:      time.Now().Unix(),
		ExpiresIn:    a.AccessTokenExpiresAfter,
		Identifiable: input.Identifiable,
	}

	if input.GenerateRefreshToken {

		err, refreshToken := a.GenerateRefreshToken(RefreshTokenInput{
			AccessToken:  accessToken.AccessToken,
			Identifiable: input.Identifiable,
		})

		if err != nil {
			return err, nil
		}

		accessToken.RefreshToken = &refreshToken.RefreshToken
	}

	if err := a.db.Table(AccessToken{}.TableName()).Create(&accessToken).Error; err != nil {
		return err, nil
	}

	return nil, &accessToken
}

func (a Auth) LoadRefreshToken(refreshToken string) (error, *RefreshToken) {

	var result *RefreshToken

	if err := a.db.Table(RefreshToken{}.TableName()).Where("refresh_token = ?", refreshToken).First(&result).Error; err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}

		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	return nil, result
}

func (a Auth) RemoveRefreshToken(refreshToken string) error {

	if err := a.db.Table(RefreshToken{}.TableName()).Where("refresh_token = ?", refreshToken).Delete(&RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (a Auth) GenerateRefreshToken(input RefreshTokenInput) (error, *RefreshToken) {

	refreshToken := RefreshToken{
		RefreshToken: uuid.NewV4().String(),
		Identifiable: input.Identifiable,
		Created:      time.Now().Unix(),
		ExpiresIn:    a.RefreshTokenExpiresAfter,
	}

	if err := a.db.Table(refreshToken.TableName()).Create(&refreshToken).Error; err != nil {
		return err, nil
	}

	return nil, &refreshToken
}

func (a Auth) LoadTwoFactorToken(id string) (error, *TwoFactorToken) {

	var result *TwoFactorToken

	if err := a.db.Table(TwoFactorToken{}.TableName()).Where("id = ?", id).First(&result).Error; err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}

		return err, nil
	}

	if result == nil {
		return nil, nil
	}

	return nil, result
}

func (a Auth) RemoveTwoFactorToken(id string) error {

	if err := a.db.Table(TwoFactorToken{}.TableName()).Where("id = ?", id).Delete(&TwoFactorToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (a Auth) GenerateTwoFactorToken(input TwoFactorTokenInput) (error, *TwoFactorToken) {

	twoFactorToken := TwoFactorToken{
		ID:           uuid.NewV4().String(),
		Identifiable: input.Identifiable,
		Token:        random.RandNumberString(6),
		Created:      time.Now().Unix(),
		ExpiresIn:    a.TwoFactorExpiresAfter,
	}

	if input.DefinedToken != nil {
		twoFactorToken.Token = *input.DefinedToken
	}

	if err := a.db.Table(twoFactorToken.TableName()).Create(&twoFactorToken).Error; err != nil {
		return err, nil
	}

	return nil, &twoFactorToken
}
