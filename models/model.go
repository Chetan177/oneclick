package models

import (
	"time"
)

type Project struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type SystemData struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Token       string    `bson:"token" json:"token"`
	Name        string    `bson:"name" json:"name"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}



