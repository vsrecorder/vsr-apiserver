package services

import (
	"context"
	"errors"

	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/dtos"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"github.com/vsrecorder/vsr-apiserver/pkg/services/models"
)

type RecordServiceInterface interface {
	Find(
		ctx context.Context,
		limit int,
		offset int,
	) ([]*models.Record, error)

	FindById(
		ctx context.Context,
		id string,
	) (*models.Record, error)

	FindByUID(
		ctx context.Context,
		uid string,
		limit int,
		offset int,
	) ([]*models.Record, error)

	FindGameById(
		ctx context.Context,
		id string,
	) ([]*models.Game, error)

	Create(
		ctx context.Context,
		uid string,
		dto *dtos.Record,
	) (*models.Record, error)

	Update(
		ctx context.Context,
		id string,
		uid string,
		dto *dtos.Record,
	) (*models.Record, error)

	Delete(
		ctx context.Context,
		id string,
		uid string,
	) error
}

type RecordService struct {
	recordRepository        repositories.RecordRepositoryInterface
	gameRepository          repositories.GameRepositoryInterface
	officialEventRepository repositories.OfficialEventRepositoryInterface
}

func NewRecordService(
	recordRepository repositories.RecordRepositoryInterface,
	gameRepository repositories.GameRepositoryInterface,
	officialEventRepository repositories.OfficialEventRepositoryInterface,
) RecordServiceInterface {
	return &RecordService{
		recordRepository,
		gameRepository,
		officialEventRepository,
	}
}

func createRecordModel(dao *daos.Record) *models.Record {
	record := models.Record{}

	record.ID = dao.ID
	record.CreatedAt = dao.CreatedAt
	record.UpdatedAt = dao.UpdatedAt
	record.OfficialEventId = dao.OfficialEventId
	record.UserId = dao.UserId
	record.DeckId = dao.DeckId

	return &record
}

func (s *RecordService) Find(
	ctx context.Context,
	limit int,
	offset int,
) ([]*models.Record, error) {
	daos, err := s.recordRepository.Find(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	records := []*models.Record{}
	for _, dao := range daos {
		record := createRecordModel(dao)
		records = append(records, record)
	}

	return records, nil
}

func (s *RecordService) FindById(
	ctx context.Context,
	id string,
) (*models.Record, error) {
	dao, err := s.recordRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	record := createRecordModel(dao)

	return record, nil
}

func (s *RecordService) FindByUID(
	ctx context.Context,
	uid string,
	limit int,
	offset int,
) ([]*models.Record, error) {
	daos, err := s.recordRepository.FindByUID(ctx, uid, limit, offset)

	if err != nil {
		return nil, err
	}

	records := []*models.Record{}
	for _, dao := range daos {
		record := createRecordModel(dao)
		records = append(records, record)
	}

	return records, nil
}

func (s *RecordService) FindGameById(
	ctx context.Context,
	id string,
) ([]*models.Game, error) {
	// 指定されたIdのRecordが存在するか確認
	if _, err := s.FindById(ctx, id); err != nil {
		return nil, err
	}

	daos, err := s.gameRepository.FindByRecordId(ctx, id)
	if err != nil {
		return nil, err
	}

	games := []*models.Game{}
	for _, dao := range daos {
		games = append(games, createGameModel(dao))
	}

	return games, nil

}

func (s *RecordService) Create(
	ctx context.Context,
	uid string,
	dto *dtos.Record,
) (*models.Record, error) {
	// 指定されたdto.OfficialEventIdのOfficialEventが存在するか確認
	if _, err := s.officialEventRepository.FindById(ctx, dto.OfficialEventId); err != nil {
		return nil, err
	}

	// TODO: 既に指定されたdto.OfficialEventIdでRecordが作成されているか確認

	id, err := generateId()
	if err != nil {
		return nil, err
	}

	dao := daos.Record{
		ID:              id,
		OfficialEventId: dto.OfficialEventId,
		UserId:          uid,
		DeckId:          dto.DeckId,
	}

	if err := s.recordRepository.Save(ctx, &dao); err != nil {
		return nil, err
	}

	record := createRecordModel(&dao)

	return record, nil
}

func (s *RecordService) Update(
	ctx context.Context,
	id string,
	uid string,
	dto *dtos.Record,
) (*models.Record, error) {
	// 指定されたdto.OfficialEventIdのOfficialEventが存在するか確認
	if _, err := s.officialEventRepository.FindById(ctx, dto.OfficialEventId); err != nil {
		return nil, err
	}

	// 指定されたidのRecordが存在するか確認
	dao, err := s.recordRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// 指定されたuidと取得したdaoのUserIdが一致しているか確認
	if dao.UserId != uid {
		return nil, errors.New("no authority")
	}

	dao.OfficialEventId = dto.OfficialEventId
	dao.DeckId = dto.DeckId

	if err := s.recordRepository.Save(ctx, dao); err != nil {
		return nil, err
	}

	record := createRecordModel(dao)

	return record, nil
}

func (s *RecordService) Delete(
	ctx context.Context,
	id string,
	uid string,
) error {
	// 指定されたIdのRecordが存在するか確認
	dao, err := s.recordRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	// 指定されたUIDと取得したdaoのUserIdが一致しているか確認
	if dao.UserId != uid {
		return errors.New("no authority")
	}

	return s.recordRepository.Delete(ctx, id, uid)
}
