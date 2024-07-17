package entity

type Post struct {
	ID             int64     `gorm:"primary_key:auto_increment" json:"id"`
	Content        string    `gorm:"type:varchar(128)" json:"content"`
	AuthorUsername string    `gorm:"type:varchar(100)" json:"author_username"`
	User           *User     `gorm:"foreignKey:ID" json:"-"`
	Bookmarks      []User    `gorm:"many2many:user_bookmarks;" json:"-"`
	Likes          []User    `gorm:"many2many:user_likes;" json:"-"`
	Comments       []Comment `gorm:"foreignKey:PostId;reference:ID;onDelete:CASCADE" json:"-"`
}

type Comment struct {
	ID      int64  `gorm:"primary_key:auto_increment" json:"-"`
	Content string `gorm:"type:varchar(128)" json:"-"`
	User    *User  `gorm:"foreignKey:ID"`
	PostId  int    `gorm:"index"`
	Post    *Post  `gorm:"foreignKey:ID"`
}
