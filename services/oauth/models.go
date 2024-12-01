package oauth

import "time"

type Client struct {

	// The ID is the client id
	ID int `json:"id" gorm:"primaryKey"`

	// The ID is the client id
	Identifier string `json:"identifier" gorm:"type:text"`

	// The client secret
	Secret string `json:"secret" gorm:"type:text"`

	// The given name of the client e.g. "iOS", "Android", "Web", "Simon"...
	Name string `json:"name" gorm:"type:text"`

	// A flag indicating if the secret is a requirement
	SecretRequired bool `json:"secret_required"`

	// Unix timestamp of when the access token was created
	Created int64 `json:"created"`
}

func (a Client) TableName() string {
	return "oauth_clients"
}

func (a *Client) IsValidSecret(secret string) bool {
	if a.SecretRequired && a.Secret != secret {
		return false
	}
	return true
}

type ClientInput struct {

	// The given name of the client e.g. "iOS", "Android", "Web", "Simon"...
	Name string `json:"access_token,omitempty"`

	// A flag indicating if the secret is a requirement
	SecretRequired bool `json:"secret_required"`
}

type AccessToken struct {
	ID int `json:"id" gorm:"primaryKey"`

	// The ID is the actual token that can be used to
	// access the given resource with.
	AccessToken string `json:"access_token" gorm:"type:text"`

	// Unix timestamp of when the access token was created
	Created int64 `json:"created"`

	// Amount of seconds until the access token expires
	ExpiresIn int64 `json:"expires_in"`

	// The refresh token to be used to refresh the access
	// token. This can be optional.
	RefreshToken *string `json:"refresh_token" gorm:"type:text"`

	// The Identifiable is the unique key that should be
	// able to identify who this access token belongs to.
	Identifiable string `json:"identifiable"  gorm:"type:text"`
}

func (a AccessToken) TableName() string {
	return "oauth_access_tokens"
}

type AccessTokenInput struct {

	// The Identifiable is the unique key that should be
	// able to identify who this access token belongs to.
	Identifiable string `json:"identifiable"`

	GenerateRefreshToken bool
}

type RefreshToken struct {
	ID int `json:"id" gorm:"primaryKey"`

	// The refresh token to be used to refresh the access
	// token. This can be optional.
	RefreshToken string `json:"refresh_token" gorm:"type:text"`

	// Unix timestamp of when the access token was created
	Created int64 `json:"created"`

	// Amount of seconds until the access token expires. This can
	// be optional for the refresh token.
	ExpiresIn int64 `json:"expires_in"`

	// The Identifiable is the unique key that should be
	// able to identify who this access token belongs to.
	Identifiable string `json:"identifiable" gorm:"type:text"`
}

func (r RefreshToken) TableName() string {
	return "oauth_refresh_tokens"
}

type RefreshTokenInput struct {

	// The id of the access token that contained
	// the refresh token id
	AccessToken string `json:"access_token,omitempty"`

	// The Identifiable is the unique key that should be
	// able to identify who this access token belongs to.
	Identifiable string `json:"identifiable,omitempty"`
}

type TwoFactorToken struct {

	// The ID of the two factor token.
	ID string `json:"id" gorm:"primaryKey"`

	// The Identifiable is the unique key that should be
	// able to identify who this access token belongs to.
	Identifiable string `json:"identifiable" gorm:"type:text"`

	// The two factor token shown to the user which
	// is supposed to be used
	Token string `json:"token" gorm:"type:text"`

	// Unix timestamp of when the two factor token was created
	Created int64 `json:"created"`

	// Amount of seconds until the two factor token expires
	ExpiresIn int64 `json:"expires_in"`
}

func (t TwoFactorToken) TableName() string {
	return "oauth_two_factor_tokens"
}

type TwoFactorTokenInput struct {

	// The Identifiable is the unique key that should be
	// able to identify who this access token belongs to.
	Identifiable string `json:"identifiable,omitempty"`

	// The StaticTestToken
	DefinedToken *string `json:"defined_token,omitempty"`
}

/*
 * Below you'll find all of the methods from the above models.
 */

func (a AccessToken) IsValid() bool {
	return time.Unix(a.Created, 0).Add(time.Duration(a.ExpiresIn) * time.Second).After(time.Now())
}

func (r RefreshToken) IsValid() bool {
	return time.Unix(r.Created, 0).Add(time.Duration(r.ExpiresIn) * time.Second).After(time.Now())
}

func (t TwoFactorToken) IsValid() bool {
	return time.Unix(t.Created, 0).Add(time.Duration(t.ExpiresIn) * time.Second).After(time.Now())
}
