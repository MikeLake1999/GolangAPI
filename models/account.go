package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	Id        uint        `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Name      string      `gorm:"type:VARCHAR(256);NOT NULL" json:"name"`
	Avatar    string      `gorm:"type:VARCHAR(256)" json:"avatar"`
	Address   string      `gorm:"type:VARCHAR(256);NOT NULL" json:"address"`
	Phone     string      `gorm:"type:VARCHAR(256);NOT NULL" json:"phone"`
	Email     string      `gorm:"type:VARCHAR(256);NOT NULL" json:"email"`
	Password  string      `gorm:"type:VARCHAR(64);NOT NULL" json:"-"`
	CreatedAt time.Time   `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt time.Time   `gorm:"NOT NULL" json:"updated_at"`
	Galleries []Galleries `json:"galleries,omitempty"`
	Photos    []Photos    `json:"photos,omitempty"`
	Reactions []Reactions `json:"reactions,omitempty"`
}
