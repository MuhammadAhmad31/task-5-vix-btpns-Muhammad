package models

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserName string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;notnull" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
