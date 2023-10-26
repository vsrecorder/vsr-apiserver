package repositories

import (
	"context"

	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"gorm.io/gorm"
)

type BattleRepositoryInterface interface {
	FindById(
		ctx context.Context,
		id string,
	) (*daos.Battle, error)

	FindByGameId(
		ctx context.Context,
		gameId string,
	) ([]*daos.Battle, error)

	Save(
		ctx context.Context,
		dao *daos.Battle,
	) error

	Delete(
		ctx context.Context,
		id string,
		uid string,
	) error
}

type BattleRepository struct {
	db *gorm.DB
}

func NewBattleRepository(
	db *gorm.DB,
) BattleRepositoryInterface {
	return &BattleRepository{db}
}

func (r *BattleRepository) FindById(
	ctx context.Context,
	id string,
) (*daos.Battle, error) {
	dao := daos.Battle{}
	if tx := r.db.Where(&daos.Battle{ID: id}).First(&dao); tx.Error != nil {
		return nil, tx.Error
	}

	return &dao, nil
}

func (r *BattleRepository) FindByGameId(
	ctx context.Context,
	gameId string,
) ([]*daos.Battle, error) {
	var battles []*daos.Battle

	if tx := r.db.Where(&daos.Battle{GameId: gameId}).Find(&battles); tx.Error != nil {
		return nil, tx.Error
	}

	return battles, nil
}

func (r *BattleRepository) Save(
	ctx context.Context,
	dao *daos.Battle,
) error {
	if tx := r.db.Save(dao); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *BattleRepository) Delete(
	ctx context.Context,
	id string,
	uid string,
) error {
	if tx := r.db.Where(&daos.Record{ID: id, UserId: uid}).Delete(&daos.Battle{}); tx.Error != nil {
		return tx.Error
	}

	return nil
}
