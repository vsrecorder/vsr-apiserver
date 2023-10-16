package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/vsrecorder/import-officialevent-bat/model"
)

type OfficialEventRepositoryInterface interface {
	Find(
		ctx context.Context,
		limit int,
		offset int,
	) ([]*model.OfficialEvent, error)

	FindById(
		ctx context.Context,
		id uint,
	) (*model.OfficialEvent, error)

	FindByDate(
		ctx context.Context,
		startDate time.Time,
		endDate time.Time,
	) ([]*model.OfficialEvent, error)
}

type OfficialEventRepository struct {
	db *gorm.DB
}

func NewOfficialEventRepository(
	db *gorm.DB,
) OfficialEventRepositoryInterface {
	return &OfficialEventRepository{db}
}

func (r *OfficialEventRepository) Find(
	ctx context.Context,
	limit int,
	offset int,
) ([]*model.OfficialEvent, error) {
	var officialEvents []*model.OfficialEvent

	if tx := r.db.Limit(limit).Offset(offset).Find(&officialEvents); tx.Error != nil {
		return nil, tx.Error
	}

	return officialEvents, nil
}

func (r *OfficialEventRepository) FindById(
	ctx context.Context,
	id uint,
) (*model.OfficialEvent, error) {
	var officialEvent model.OfficialEvent

	if tx := r.db.Where(&model.OfficialEvent{Id: id}).First(&officialEvent); tx.Error != nil {
		return nil, tx.Error
	}

	return &officialEvent, nil
}

func (r *OfficialEventRepository) FindByDate(
	ctx context.Context,
	startDate time.Time,
	endDate time.Time,
) ([]*model.OfficialEvent, error) {
	var officialEvents []*model.OfficialEvent

	if tx := r.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&officialEvents); tx.Error != nil {
		return nil, tx.Error
	}

	return officialEvents, nil
}
