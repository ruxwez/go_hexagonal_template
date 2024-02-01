package authInfrastructure

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"
)

// CreateSession implements authDomain.AuthRepository.
func (a *authRepository) CreateSession(userId uint, ip string) (string, error) {

	// Generamos un token con UUID de google
	token, _ := uuid.NewV7()

	// Creamos el modelo de sesion
	session := authDomain.Session{
		Token:    token.String(),
		UserID:   userId,
		LastIP:   ip,
		LastConn: time.Now(),
	}

	// Insertamos la sesion en la base de datos y comprobamos si hay errores
	if err := a.db.Create(&session).Error; err != nil {
		// Comprobamos si el error es de clave duplicada
		mysqlErr, isMySQLError := err.(*mysql.MySQLError)
		if isMySQLError && mysqlErr.Number == 1062 {
			// En el caso de que el error sea por duplicidad de token, volvemos a llamar a la funcion
			return a.CreateSession(userId, ip)
		}
		return "", authDomain.ErrSessionCreation
	}

	return token.String(), nil
}
