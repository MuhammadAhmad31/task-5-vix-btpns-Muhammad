package app

import "time"

type PhotosUpdateApp struct {
	ID          uint64    `json:"id" form:"id" binding:"required"`
	Title       string    `json:"title" form:"title" binding:"required"`
	PhotoUrl    string    `gorm:"type:varchar(255)" json:"photo-url"`
	Description string    `json:"description" from:"description" binding:"required"`
	UserID      uint64    `json:"user_id,omitempty" from:"user_id,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
}

type PhotosCreateApp struct {
	Title       string    `json:"title" form:"title" binding:"required"`
	PhotoUrl    string    `gorm:"type:varchar(255)" json:"photo-url"`
	Description string    `json:"description" from:"description" binding:"required"`
	UserID      uint64    `json:"user_id,omitempty" from:"user_id,omitempty"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
}
