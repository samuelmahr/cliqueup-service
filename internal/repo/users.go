package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samuelmahr/cliqueup-service/internal/models"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, newUser models.UsersCreateRequest) (models.User, error)
}

type UsersRepoType struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) UsersRepoType {
	return UsersRepoType{
		db: db,
	}
}

const createUserQuery = `
insert into users(email, first_name, last_name, birthday, phone_number, subid)
VALUES ($1, $2, $3, $4, $5, $6)
returning id as user_id, email, first_name, last_name, birthday, phone_number, subid, created_at, updated_at
`

func (ur *UsersRepoType) CreateUser(ctx context.Context, nUser models.UsersCreateRequest) (models.User, error) {
	var u models.User
	err := ur.db.QueryRowx(createUserQuery, nUser.Email, nUser.FirstName, nUser.LastName, nUser.Birthday, nUser.PhoneNumber, nUser.Subid).StructScan(&u)

	if err != nil {
		return models.User{}, errors.Wrap(err, "error creating user")
	}

	return u, nil
}
