package userDomain

import "time"

// Model de Usuario en la base de datos
type User struct {
	Id       uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	Password string    `gorm:"notNull" json:"-"`
	Email    string    `gorm:"unique;notNull" json:"email"`
	Name     string    `gorm:"notNull" json:"name"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
}
