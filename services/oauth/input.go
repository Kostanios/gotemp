package oauth

type NewInput struct {

	// DB info
	ClientsTableName        string
	AccessTokenTableName    string
	RefreshTokenTableName   string
	TwoFactorTokenTableName string

	// Limits
	AccessTokenExpiresAfter  int64
	RefreshTokenExpiresAfter int64
	TwoFactorExpiresAfter    int64
}
