package authInfrastructure

import authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"

// FindSession implements authDomain.AuthRepository.
func (a *authRepository) FindSession(token string) (*authDomain.Session, error) {
	var data *authDomain.Session

	if err := a.db.Model(&authDomain.Session{}).Where("token = ?", token).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
