package authDomain

import "time"

type Session struct {
	Token    string `json:"token"`
	UserID   uint   `gorm:"notNull"`
	LastIP   string
	LastConn time.Time `json:"last_conn"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
}
