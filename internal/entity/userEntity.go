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
	ID             int  `gorm:"primary_key:auto_increment" json:"-"`
	UserID         int  `gorm:"not null" json:"-"`
	FollowerUserID int  `gorm:"not null" json:"user"`
	User           User `gorm:"foreignKey:UserID" json:"-"`
	Follower       User `gorm:"foreignKey:FollowerUserID" json:"-"`
}

type Following struct {
	gorm.Model
	ID              int  `gorm:"primary_key:auto_increment" json:"-"`
	UserID          int  `gorm:"not null" json:"-"`
	FollowingUserID int  `gorm:"not null"`
	User            User `gorm:"foreignKey:UserID" json:"-"`
	Following       User `gorm:"foreignKey:FollowingUserID" json:"-"`
}
