package repositories

import (
	"context"

	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
)

type UserRepositoryInterface interface {
	FindById(
		ctx context.Context,
		id string,
	) (*daos.User, error)
}

type UserRepository struct {
	fbAuth *firebaseAuth.Client
}

func NewUserRepository(
	fbAuth *firebaseAuth.Client,
) UserRepositoryInterface {
	return &UserRepository{fbAuth}
}

func (r *UserRepository) FindById(
	ctx context.Context,
	id string,
) (*daos.User, error) {
	userRecord, err := r.fbAuth.GetUser(context.Background(), id)
	if err != nil {
		return nil, err
	}

	user := &daos.User{}
	user.UID = userRecord.UID
	user.DisplayName = userRecord.DisplayName
	user.PhotoURL = userRecord.PhotoURL

	return user, nil
}
