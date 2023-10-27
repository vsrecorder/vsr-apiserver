package services

import (
	"context"
	"time"

	oem "github.com/vsrecorder/import-officialevent-bat/pkg/models"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
	"github.com/vsrecorder/vsr-apiserver/pkg/services/models"
)

type OfficialEventServiceInterface interface {
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

	FindRecordById(
		ctx context.Context,
		id uint,
	) ([]*models.Record, error)
}

type OfficialEventService struct {
	officialEventRepository repositories.OfficialEventRepositoryInterface
	recordRepository        repositories.RecordRepositoryInterface
}

func NewOfficialEventService(
	officialEventRepository repositories.OfficialEventRepositoryInterface,
	recordRepository repositories.RecordRepositoryInterface,
) OfficialEventServiceInterface {
	return &OfficialEventService{
		officialEventRepository,
		recordRepository,
	}
}

func (s *OfficialEventService) Find(
	ctx context.Context,
	limit int,
	offset int,
) ([]*oem.OfficialEvent, error) {
	ret, err := s.officialEventRepository.Find(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return ret, nil

}

func (s *OfficialEventService) FindById(
	ctx context.Context,
	id uint,
) (*oem.OfficialEvent, error) {
	ret, err := s.officialEventRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *OfficialEventService) FindByDate(
	ctx context.Context,
	startDate time.Time,
	endDate time.Time,
) ([]*oem.OfficialEvent, error) {
	ret, err := s.officialEventRepository.FindByDate(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *OfficialEventService) FindRecordById(
	ctx context.Context,
	id uint,
) ([]*models.Record, error) {
	// 指定されたidのOfficialEventが存在するか確認
	if _, err := s.FindById(ctx, id); err != nil {
		return nil, err
	}

	daos, err := s.recordRepository.FindByOfficialEventId(ctx, id)
	if err != nil {
		return nil, err
	}

	records := []*models.Record{}
	for _, dao := range daos {
		records = append(records, createRecordModel(dao))
	}

	return records, nil
}
