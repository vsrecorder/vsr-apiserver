package repositories

import (
	"context"

	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"gorm.io/gorm"
)

type RecordRepositoryInterface interface {
	TearDown() error

	Find(
		ctx context.Context,
		limit int,
		offset int,
	) ([]*daos.Record, error)

	FindById(
		ctx context.Context,
		id string,
	) (*daos.Record, error)

	FindByUID(
		ctx context.Context,
		uid string,
		limit int,
		offset int,
	) ([]*daos.Record, error)

	FindAllByUID(
		ctx context.Context,
		uid string,
	) ([]*daos.Record, error)

	FindByOfficialEventId(
		ctx context.Context,
		officialEventId uint,
	) ([]*daos.Record, error)

	FindByDeckId(
		ctx context.Context,
		deckId string,
	) ([]*daos.Record, error)

	Save(
		ctx context.Context,
		record *daos.Record,
	) error

	Delete(
		ctx context.Context,
		id string,
		uid string,
	) error
}

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(
	db *gorm.DB,
) RecordRepositoryInterface {
	return &RecordRepository{db}
}

func (r *RecordRepository) TearDown() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (r *RecordRepository) Find(
	ctx context.Context,
	limit int,
	offset int,
) ([]*daos.Record, error) {
	var records []*daos.Record

	if tx := r.db.Limit(limit).Offset(offset).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r *RecordRepository) FindById(
	ctx context.Context,
	id string,
) (*daos.Record, error) {
	var record daos.Record

	if tx := r.db.Where(&daos.Record{ID: id}).First(&record); tx.Error != nil {
		return nil, tx.Error
	}

	return &record, nil
}

func (r *RecordRepository) FindByUID(
	ctx context.Context,
	uid string,
	limit int,
	offset int,
) ([]*daos.Record, error) {
	var records []*daos.Record

	if tx := r.db.Where(&daos.Record{UserId: uid}).Limit(limit).Offset(offset).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r *RecordRepository) FindAllByUID(
	ctx context.Context,
	uid string,
) ([]*daos.Record, error) {
	var records []*daos.Record

	if tx := r.db.Where(&daos.Record{UserId: uid}).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r *RecordRepository) FindByOfficialEventId(
	ctx context.Context,
	officialEventId uint,
) ([]*daos.Record, error) {
	var records []*daos.Record

	if tx := r.db.Where(&daos.Record{OfficialEventId: officialEventId}).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r *RecordRepository) FindByDeckId(
	ctx context.Context,
	deckId string,
) ([]*daos.Record, error) {
	var records []*daos.Record

	if tx := r.db.Where(&daos.Record{DeckId: deckId}).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r *RecordRepository) Save(
	ctx context.Context,
	record *daos.Record,
) error {
	if tx := r.db.Save(record); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *RecordRepository) Delete(
	ctx context.Context,
	id string,
	uid string,
) error {
	if tx := r.db.Where(&daos.Record{ID: id, UserId: uid}).Delete(&daos.Record{}); tx.Error != nil {
		return tx.Error
	}

	return nil
}
