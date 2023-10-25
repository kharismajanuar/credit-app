package jwt

type JwtClaims struct {
	UserID uint
	Email  string
	Role   string
}
