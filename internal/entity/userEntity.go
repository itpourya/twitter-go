package entity

type User struct {
	ID         int64  `gorm:"primary_key:auto_increment" json:"-"`
	Firstname  string `gorm:"type:varchar(100)" json:"firstname"`
	Lastname   string `gorm:"type:varchar(100)" json:"lastname"`
	Username   string `gorm:"type:varchar(100);unique" json:"username"`
	Email      string `gorm:"type:varchar(100);unique" json:"email"`
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
