package repositories

import (
	"context"

	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"gorm.io/gorm"
)

type GameRepositoryInterface interface {
	FindById(
		ctx context.Context,
		id string,
	) (*daos.Game, error)

	Save(
		ctx context.Context,
		game *daos.Game,
	) error

	Delete(
		ctx context.Context,
		id string,
		uid string,
	) error
}

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(
	db *gorm.DB,
) GameRepositoryInterface {
	return &GameRepository{db}
}

func (r *GameRepository) FindById(
	ctx context.Context,
	id string,
) (*daos.Game, error) {
	game := &daos.Game{}

	if tx := r.db.Where(&daos.Game{ID: id}).First(game); tx.Error != nil {
		return nil, tx.Error
	}

	return game, nil
}

func (r *GameRepository) Save(
	ctx context.Context,
	game *daos.Game,
) error {
	if tx := r.db.Save(game); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *GameRepository) Delete(
	ctx context.Context,
	id string,
	uid string,
) error {
	if tx := r.db.Where(&daos.Game{ID: id, UserId: uid}).Delete(&daos.Game{}); tx.Error != nil {
		return tx.Error
	}

	return nil
}
