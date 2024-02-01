package authDomain

type AuthRepository interface {
	CreateSession(userId uint, ip string) (string, error)
	DeleteSession(token string) error
	FindSession(token string) (*Session, error)
}
