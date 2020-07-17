package repo

import (
	"context"
	"github.com/pkg/errors"
	"github.com/samuelmahr/cliqueup-service/internal/models"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	SetupTestDB()
	code := m.Run()
	os.Exit(code)
}

func TestUserRepository_CreateUser(t *testing.T) {
	now := time.Now().UTC()
	phone := "18001155188" // 1-800-CALLATT shoutout to carrot top
	type args struct {
		user models.UsersCreateRequest
	}

	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "happy path",
			want: models.User{
				ID:          1,
				Email:       "jake@statefarm.com",
				FirstName:   "JAke",
				LastName:    "From State Farm",
				Birthday:    now,
				PhoneNumber: &phone,
				Subid:       "cognito",
			},
			args: args{
				user: models.UsersCreateRequest{
					Email:       "jake@statefarm.com",
					FirstName:   "Jake",
					LastName:    "From State Farm",
					Birthday:    now,
					PhoneNumber: &phone,
					Subid:       "cognitoid",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PurgeTables()

			r := &UsersRepoType{
				db: DB,
			}

			got, err := r.CreateUser(context.Background(), tt.args.user)
			if err != nil && !tt.wantErr {
				t.Fatal(err)
			} else if err != nil && tt.wantErr {
				assert.True(t, errors.Is(err, err))
			} else {
				assert.Equal(t, tt.want.ID, got.ID)
				assert.Equal(t, tt.want.Email, got.Email)
				assert.Equal(t, tt.want.FirstName, got.FirstName)
				assert.Equal(t, tt.want.LastName, got.LastName)
				assert.Equal(t, tt.want.PhoneNumber, got.PhoneNumber)
				assert.Equal(t, tt.want.Birthday, got.Birthday)
				assert.Equal(t, tt.want.Subid, got.Subid)
			}
		})
	}
}
