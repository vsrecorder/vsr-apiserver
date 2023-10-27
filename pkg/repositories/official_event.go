package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"

	oem "github.com/vsrecorder/import-officialevent-bat/pkg/models"
)

type OfficialEventRepositoryInterface interface {
	Find(
		ctx context.Context,
		limit int,
		offset int,
	) ([]*oem.OfficialEvent, error)

	FindById(
		ctx context.Context,
		id uint,
	) (*oem.OfficialEvent, error)

	FindByDate(
		ctx context.Context,
		startDate time.Time,
		endDate time.Time,
	) ([]*oem.OfficialEvent, error)
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
) ([]*oem.OfficialEvent, error) {
	var officialEvents []*oem.OfficialEvent

	if tx := r.db.Limit(limit).Offset(offset).Find(&officialEvents); tx.Error != nil {
		return nil, tx.Error
	}

	return officialEvents, nil
}

func (r *OfficialEventRepository) FindById(
	ctx context.Context,
	id uint,
) (*oem.OfficialEvent, error) {
	var officialEvent oem.OfficialEvent

	if tx := r.db.Where(&oem.OfficialEvent{Id: id}).First(&officialEvent); tx.Error != nil {
		return nil, tx.Error
	}

	return &officialEvent, nil
}

func (r *OfficialEventRepository) FindByDate(
	ctx context.Context,
	startDate time.Time,
	endDate time.Time,
) ([]*oem.OfficialEvent, error) {
	var officialEvents []*oem.OfficialEvent

	if tx := r.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&officialEvents); tx.Error != nil {
		return nil, tx.Error
	}

	return officialEvents, nil
}
