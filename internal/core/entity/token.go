package entity

type TokenManager interface {
	Sign(payload any) (string, error)
	Decode(tokenStr string) (any, error)
	Verify(tokenStr string) (bool, error)
}
