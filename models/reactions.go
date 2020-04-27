package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Reactions struct {
	Id        int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	AccountId uint      `gorm:"NOT NULL" json:"account_id"`
	PhotoId   int       `gorm:"NOT NULL" json:"photo_id"`
	Reaction  string    `gorm:"type:VARCHAR(256);NOT NULL" json:"reaction"`
	CreatedAt time.Time `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"NOT NULL" json:"updeted_at"`
}
