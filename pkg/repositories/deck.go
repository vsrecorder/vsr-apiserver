package repositories

import (
	"context"

	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"gorm.io/gorm"
)

type DeckRepositoryInterface interface {
	FindById(
		ctx context.Context,
		id string,
	) (*daos.Deck, error)

	FindByUID(
		ctx context.Context,
		uid string,
		limit int,
		offset int,
	) ([]*daos.Deck, error)

	Save(
		ctx context.Context,
		dao *daos.Deck,
	) error

	Delete(
		ctx context.Context,
		id string,
		uid string,
	) error
}

type DeckRepository struct {
	db *gorm.DB
}

func NewDeckRepository(
	db *gorm.DB,
) DeckRepositoryInterface {
	return &DeckRepository{db}
}

func (r *DeckRepository) FindById(
	ctx context.Context,
	id string,
) (*daos.Deck, error) {
	dao := &daos.Deck{}

	if tx := r.db.Where(&daos.Deck{ID: id}).First(dao); tx.Error != nil {
		return nil, tx.Error
	}

	return dao, nil
}

func (r *DeckRepository) FindByUID(
	ctx context.Context,
	uid string,
	limit int,
	offset int,
) ([]*daos.Deck, error) {
	var decks []*daos.Deck

	if tx := r.db.Where(&daos.Deck{UserId: uid}).Limit(limit).Offset(offset).Find(&decks); tx.Error != nil {
		return nil, tx.Error
	}

	return decks, nil
}

func (r *DeckRepository) Save(
	ctx context.Context,
	dao *daos.Deck,
) error {
	if tx := r.db.Save(dao); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *DeckRepository) Delete(
	ctx context.Context,
	id string,
	uid string,
) error {
	if tx := r.db.Where(&daos.Deck{ID: id, UserId: uid}).Delete(&daos.Deck{}); tx.Error != nil {
		return tx.Error
	}

	return nil
}
