package user

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Email        string        `bson:"email" json:"email"`
	PasswordHash string        `bson:"passwordHash" json:"-"`
	Role         string        `bson:"role" json:"role"`
	CreatedAt    time.Time     `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt,omitempty" json:"updatedAt"`
}

type PublicUser struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToPublicUser(u *User) PublicUser {
	return PublicUser{
		ID:        u.ID.Hex(),
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
