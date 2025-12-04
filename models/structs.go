package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user model struct
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"createdAt" json:"created_at"`
	LastLogin time.Time          `bson:"lastLogin" json:"last_login"`
	Role      string             `bson:"role" json:"role"`
	UpdatedAt time.Time          `bson:"updated" json:"updated_at"`
}

// article struct
type Article struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title      string             `bson:"title" json:"title"`
	Slug       string             `bson:"slug" json:"slug"`
	Content    string             `bson:"content" json:"content"`
	AuthorID   primitive.ObjectID `bson:"author_id" json:"author_id"`
	CategoryID primitive.ObjectID `bson:"category_id" json:"category_id"`
	ImageURL   string             `bson:"image_url" json:"image_url"`
	Status     string             `bson:"status" json:"status"` // draft, published
	Tags       []string           `bson:"tags" json:"tags"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	Featured   bool               `bson:"featured" json:"featured"`
	Trending   bool               `bson:"trending" json:"trending"`
}

// Json reponse struct
type JSONresponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// User response struct
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
