package models

import (
	"time"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now();autoUpdateTime:milli" sql:"DEFAULT:'current_timestamp'"`
}

type User struct {
	BaseModel
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password" gorm:"not null"`
	Telephone   string `json:"telephone" gorm:"not null"`
	SocialLinks string `json:"social_links"`
}

type Writer struct {
	BaseModel
	Biograph      string `json:"biograph"`
	Description   string `json:"description"`
	DistinctLikes int    `json:"distinct_likes" gorm:"default:null"`
	DistinctViews int    `json:"distinct_views" gorm:"default:null"`
	UserID        uint   `json:"user_id"`
	User          User   `gorm:"foreignKey:UserID"`
}

type Follower struct {
	BaseModel
	WriterID     uint `json:"writer_id"`
	FollowerID   uint `json:"follower_id"`
	WriterUser   User `gorm:"foreignKey:WriterID"`
	FollowerUser User `gorm:"foreignKey:FollowerID"`
}

type Post struct {
	BaseModel
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	WriterID    uint   `json:"writer_id"`
	User        User   `gorm:"foreignKey:WriterID"`
	Content     string `json:"content"`
	Views       int    `json:"views" gorm:"default:null"`
	Likes       int    `json:"likes" gorm:"default:null"`
	Shares      int    `json:"shares" gorm:"default:null"`
}

type Notifications struct {
	BaseModel
	Title   string `json:"title" gorm:"not null"`
	Message string `json:"message" gorm:"not null"`
	PostID  uint   `json:"post_id"`
	Post    Post   `gorm:"foreignKey:PostID"`
	UserID  uint   `json:"user_id"`
	User    User
}

type Chat struct {
	BaseModel
	Title        string `json:"title" gorm:"not null"`
	Message      string `json:"message" gorm:"not null"`
	SenderID     uint   `json:"sender_id"`
	ReceiverID   uint   `json:"receiver_id"`
	SenderUser   User   `gorm:"foreignKey:SenderID"`
	ReceiverUser User   `gorm:"foreignKey:ReceiverID"`
}

type Images struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"not null"`
	Path      string `json:"path" gorm:"not null"`
	Type      int    `json:"type"`
	Extension string `json:"extension"`
}
