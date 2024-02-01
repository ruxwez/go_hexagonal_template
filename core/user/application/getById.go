package userApplication

import userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"

// GetById implements UserService.
func (s *userService) GetById(id int) *userDomain.User {
	return s.user.FindById(uint(id))
}
