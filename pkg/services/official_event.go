package services

import (
	"context"
	"time"

	"github.com/vsrecorder/import-officialevent-bat/model"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
)

type OfficialEventServiceInterface interface {
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

type OfficialEventService struct {
	repository repositories.OfficialEventRepositoryInterface
}

func NewOfficialEventService(
	repository repositories.OfficialEventRepositoryInterface,
) OfficialEventServiceInterface {
	return &OfficialEventService{repository}
}

func (s *OfficialEventService) Find(
	ctx context.Context,
	limit int,
	offset int,
) ([]*model.OfficialEvent, error) {
	ret, err := s.repository.Find(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return ret, nil

}

func (s *OfficialEventService) FindById(
	ctx context.Context,
	id uint,
) (*model.OfficialEvent, error) {
	ret, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *OfficialEventService) FindByDate(
	ctx context.Context,
	startDate time.Time,
	endDate time.Time,
) ([]*model.OfficialEvent, error) {
	ret, err := s.repository.FindByDate(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
