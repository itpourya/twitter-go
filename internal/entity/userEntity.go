package entity

type User struct {
	ID         int64  `gorm:"primary_key:auto_increment" json:"-"`
	Firstname  string `gorm:"type:varchar(100)" json:"-"`
	Lastname   string `gorm:"type:varchar(100)" json:"-"`
	Username   string `gorm:"type:varchar(100);unique" json:"-"`
	Email      string `gorm:"type:varchar(100);unique" json:"-"`
	Password   string `gorm:"type:varchar(100)" json:"-"`
	IsActive   bool
	Posts      []Post     `gorm:"foreignKey:ID;onDelete:CASCADE"`
	Followers  []Follower `gorm:"foreignKey:ID"`
	Followings []Follower `gorm:"foreignKey:ID"`
}

type Follower struct {
	ID   int64 `gorm:"primary_key:auto_increment" json:"-"`
	User *User `gorm:"foreignKey:ID"`
}
