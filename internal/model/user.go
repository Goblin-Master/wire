package model

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(32);not null"`
	Password string `gorm:"type:varchar(32);not null"`
}

func (u *User) TableName() string {
	return "user"
}
