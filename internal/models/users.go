package models

import "time"

type User struct {
	ID          int64      `json:"id" db:"user_id"`
	Email       string     `json:"email" db:"email"`
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    string     `json:"last_name" db:"last_name"`
	Birthday    time.Time  `json:"birthday" db:"birthday"`
	PhoneNumber *string    `json:"phone_number" db:"phone_number"`
	Subid       string     `json:"subid" db:"subid"`
	PhotoBucket *string    `json:"photo_bucket,omitempty" db:"photo_bucket"`
	PhotoKey    *string    `json:"photo_key,omitempty" db:"photo_key"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	ArchivedAt  *time.Time `json:"archived_at" db:"archived_at"`
}

type UsersCreateRequest struct {
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Birthday    time.Time `json:"birthday"`
	PhoneNumber *string   `json:"phone_number"`
	Subid       string    `json:"subid"`
}
