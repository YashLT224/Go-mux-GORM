package model

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(100)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(100)"`
	Age       int8   `json:"age" binding:"gte=1,lte=20"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(100)"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=108" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"min=2,max=108" gorm:"type:varchar(100)"`
	Url         string    `json:"url" binding:"required,url" gorm:"type:varchar(100);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersionID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
