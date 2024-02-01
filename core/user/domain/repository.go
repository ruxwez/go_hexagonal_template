package userDomain

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) *User
	FindById(id uint) *User
}
