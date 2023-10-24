package services

import (
	"context"
	"errors"

	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/dtos"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"github.com/vsrecorder/vsr-apiserver/pkg/services/models"
)

type GameServiceInterface interface {
	FindById(
		ctx context.Context,
		id string,
	) (*models.Game, error)

	Create(
		ctx context.Context,
		uid string,
		dto *dtos.Game,
	) (*models.Game, error)

	Update(
		ctx context.Context,
		id string,
		uid string,
		dto *dtos.Game,
	) (*models.Game, error)

	Delete(
		ctx context.Context,
		id string,
		uid string,
	) error
}

type GameService struct {
	gameRepository   repositories.GameRepositoryInterface
	recordRepository repositories.RecordRepositoryInterface
}

func NewGameService(
	gameRepository repositories.GameRepositoryInterface,
	recordRepository repositories.RecordRepositoryInterface,
) GameServiceInterface {
	return &GameService{gameRepository, recordRepository}
}

func createGameModel(dao *daos.Game) *models.Game {
	model := &models.Game{}

	model.ID = dao.ID
	model.CreatedAt = dao.CreatedAt
	model.UpdatedAt = dao.UpdatedAt
	model.RecordId = dao.RecordId
	model.UserId = dao.UserId
	model.OpponentsUserId = dao.OpponentsUserId
	model.BO3Flg = dao.BO3Flg
	model.QualifyingRoundFlg = dao.QualifyingRoundFlg
	model.FinalTournamentFlg = dao.FinalTournamentFlg
	model.VictoryFlg = dao.VictoryFlg
	model.OpponentsDeckInfo = dao.OpponentsDeckInfo
	model.Memo = dao.Memo

	return model
}

func (s *GameService) FindById(
	ctx context.Context,
	id string,
) (*models.Game, error) {
	dao, err := s.gameRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	model := createGameModel(dao)

	return model, nil
}

func (s *GameService) Create(
	ctx context.Context,
	uid string,
	dto *dtos.Game,
) (*models.Game, error) {
	// 指定されたdto.RecordIdのRecordが存在するか確認
	if _, err := s.recordRepository.FindById(ctx, dto.RecordId); err != nil {
		return nil, err
	}

	id, err := generateId()
	if err != nil {
		return nil, err
	}

	dao := daos.Game{
		ID:                 id,
		RecordId:           dto.RecordId,
		UserId:             uid,
		OpponentsUserId:    dto.OpponentsUserId,
		BO3Flg:             dto.BO3Flg,
		QualifyingRoundFlg: dto.QualifyingRoundFlg,
		FinalTournamentFlg: dto.FinalTournamentFlg,
		VictoryFlg:         dto.VictoryFlg,
		OpponentsDeckInfo:  dto.OpponentsDeckInfo,
		Memo:               dto.Memo,
	}

	if err := s.gameRepository.Save(ctx, &dao); err != nil {
		return nil, err
	}

	model := createGameModel(&dao)

	return model, nil

}

func (s *GameService) Update(
	ctx context.Context,
	id string,
	uid string,
	dto *dtos.Game,
) (*models.Game, error) {
	// 指定されたdto.RecordIdのRecordが存在するか確認
	if _, err := s.recordRepository.FindById(ctx, dto.RecordId); err != nil {
		return nil, err
	}

	// 指定されたidのGameが存在するか確認
	dao, err := s.gameRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// 指定されたuidと取得したdaoのUserIdが一致しているか確認
	if dao.UserId != uid {
		return nil, errors.New("no authority")
	}

	dao.RecordId = dto.RecordId
	dao.OpponentsUserId = dto.OpponentsUserId
	dao.BO3Flg = dto.BO3Flg
	dao.QualifyingRoundFlg = dto.QualifyingRoundFlg
	dao.FinalTournamentFlg = dto.FinalTournamentFlg
	dao.VictoryFlg = dto.VictoryFlg
	dao.OpponentsDeckInfo = dto.OpponentsDeckInfo
	dao.Memo = dto.Memo

	if err := s.gameRepository.Save(ctx, dao); err != nil {
		return nil, err
	}

	model := createGameModel(dao)

	return model, nil
}

func (s *GameService) Delete(
	ctx context.Context,
	id string,
	uid string,
) error {
	// 指定されたIdのGameが存在するか確認
	dao, err := s.gameRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	// 指定されたUIDと取得したdaoのUserIdが一致しているか確認
	if dao.UserId != uid {
		return errors.New("no authority")
	}

	return s.gameRepository.Delete(ctx, id, uid)
}
