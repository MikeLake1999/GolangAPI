package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Photos struct {
	Id          int         `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	AccountId   uint        `gorm:"NOT NULL" json:"account_id"`
	GalleryId   int         `gorm:"NOT NULL" json:"gallery_id"`
	Name        string      `gorm:"type:VARCHAR(256);NOT NULL" json:"name"`
	Description string      `gorm:"type:TEXT;NOT NULL" json:"desciption"`
	Path        string      `gorm:"type:VARCHAR(256);NOT NULL" json:"path"`
	Size        int64       `gorm:"type:VARCHAR(256);NOT NULL" json:"size"`
	CreatedAt   time.Time   `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"NOT NULL" json:"updated_at"`
	Reactions   []Reactions `json:"reactions,omitempty"`
}
