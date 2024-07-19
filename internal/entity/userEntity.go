package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int    `gorm:"primary_key:auto_increment" json:"-"`
	Firstname string `gorm:"type:varchar(100)" json:"firstname"`
	Lastname  string `gorm:"type:varchar(100)" json:"lastname"`
	Username  string `gorm:"type:varchar(100);unique" json:"username"`
	Email     string `gorm:"type:varchar(100);unique" json:"email"`
	Password  string `gorm:"type:varchar(100)" json:"-"`
	IsActive  bool
}

type Follower struct {
	gorm.Model
	ID             int  `gorm:"primary_key:auto_increment" json:"id"`
	UserID         int  `gorm:"not null"`
	FollowerUserID int  `gorm:"not null"`
	User           User `gorm:"foreignKey:UserID"`
	Follower       User `gorm:"foreignKey:FollowerUserID"`
}

type Following struct {
	gorm.Model
	ID              int  `gorm:"primary_key:auto_increment" json:"id"`
	UserID          int  `gorm:"not null"`
	FollowingUserID int  `gorm:"not null"`
	User            User `gorm:"foreignKey:UserID"`
	Following       User `gorm:"foreignKey:FollowingUserID"`
}
