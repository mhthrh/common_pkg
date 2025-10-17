package auth

import "context"

type IAuth interface {
	Login(ctx context.Context, username, password, loginChallenge, host, method string) (string, string, error)
	Logout(ctx context.Context, idToken, refreshToken, host string) (string, error)
	TwoFactorVerificationCode(ctx context.Context, email, userKey, mobileNumber string) (string, error)
	VerifyTwoFactorAuth(ctx context.Context, loginChallenge, email, userKey, verificationCode, host string) (string, error)
	Consent(ctx context.Context, consentChallenge, host string) (string, error)
	//Redirect(ctx context.Context, code, state, host string) (model.AuthResponse, error)
	CheckTokenValidness(ctx context.Context, bearerToken, host string) (bool, string, error)
	ExchangeJWTTokenByAccessToken(ctx context.Context, token, apiPath, host string) (string, error)
	//RefreshTokenByRefreshToken(ctx context.Context, refreshToken, host string) (model.AuthResponse, error)
}
