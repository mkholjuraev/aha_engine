package models

import "time"

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primary_key" query:"u.id"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:now()::timestamp" sql:"DEFAULT:now::timestamp"`
}

type User struct {
	BaseModel
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password" gorm:"not null"`
	Telephone   string `json:"telephone" gorm:"not null"`
	SocialLinks string `json:"social_links"`
	//TODO: add interests info
	PhotoID uint   `json:"photo_id"`
	Photo   Images `json:"photo"`
}

type Writer struct {
	BaseModel
	Biography       string           `json:"biography"`
	Profession      string           `json:"profession"`
	DistinctLikes   int              `json:"distinct_likes" gorm:"default:null"`
	DistinctViews   int              `json:"distinct_views" gorm:"default:null"`
	Specializations []Specialization `gorm:"many2many:writer_specializations;"`
	UserID          uint             `json:"user_id"`
	User            User             `gorm:"foreignKey:UserID"`
}

type Follower struct {
	BaseModel
	WriterID     uint `json:"writer_id"`
	WriterUser   User `gorm:"foreignKey:WriterID"`
	FollowerID   uint `json:"follower_id"`
	FollowerUser User `gorm:"foreignKey:FollowerID"`
}

type Post struct {
	BaseModel
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	WriterID    uint   `json:"writer_id"`
	Writer      Writer `gorm:"foreignKey:WriterID"`
	Content     string `json:"content"`
	Views       int    `json:"views" gorm:"default:null"`
	Likes       int    `json:"likes" gorm:"default:null"`
	Shares      int    `json:"shares" gorm:"default:null"`
	CoverImage  string `json:"cover_image"`
	ReadTime    int    `json:"read_time"`
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

type Specialization struct {
	BaseModel
	Name        string   `json:"name" gorm:"not null;unique"`
	Description string   `json:"description"`
	Writers     []Writer `gorm:"many2many:writer_specializations;"`
}

type Tags struct {
	ID   uint   `json:"id" gorm:"primary_key" query:"t.id"`
	Name string `json:"name" gorm:"not null;unique"`
}

type PostMetadata struct {
	TagIDJSON          []byte         `json:"tag_ids" gorm:"primary_key"`
	PostID             uint           `json:"post_id" gorm:"primary_key"`
	SpecializationID   uint           `json:"specialization_id" gorm:"primary_key"`
	PostFKEY           Post           `gorm:"foreignKey:PostID;constraint:OnDelete: CASCADE;"`
	SpecializationFKEY Specialization `gorm:"foreignKey:SpecializationID;constraint:OnDelete:SET NULL;"`
}
