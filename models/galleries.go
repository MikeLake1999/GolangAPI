package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Galleries struct {
	Id        int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	AccountId uint      `gorm:"NOT NULL" json:"account_id"`
	Name      string    `gorm:"NOT NULL" json:"name"`
	Brief     string    `gorm:"type:VARCHAR(256)" json:"brief"`
	Active    string    `gorm:"type:VARCHAR(256)" json:"active"`
	CreatedAt time.Time `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"NOT NULL" json:"updated_at"`
	Photos    []Photos  `gorm:"foreignkey:GalleryId"  json:"photos,omitempty" `
}
