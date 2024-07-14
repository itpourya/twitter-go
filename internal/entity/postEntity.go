package entity

type Post struct {
	ID          int64     `gorm:"primary_key:auto_increment" json:"-"`
	Content     string    `gorm:"type:varchar(128)" json:"-"`
	AuthorEmail string    `gorm:"type:varchar(100)" json:"-"`
	User        *User     `gorm:"foreignKey:ID"`
	Bookmarks   []User    `gorm:"many2many:user_bookmarks;"`
	Likes       []User    `gorm:"many2many:user_likes;"`
	Comments    []Comment `gorm:"foreignKey:PostId;reference:ID;onDelete:CASCADE"`
}

type Comment struct {
	ID      int64  `gorm:"primary_key:auto_increment" json:"-"`
	Content string `gorm:"type:varchar(128)" json:"-"`
	User    *User  `gorm:"foreignKey:ID"`
	PostId  int    `gorm:"index"`
	Post    *Post  `gorm:"foreignKey:ID"`
}
