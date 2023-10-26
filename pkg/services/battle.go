package services

import (
	"context"
	"errors"

	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/dtos"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"github.com/vsrecorder/vsr-apiserver/pkg/services/models"
)

type BattleServiceInterface interface {
	FindById(
		ctx context.Context,
		id string,
	) (*models.Battle, error)

	Create(
		ctx context.Context,
		uid string,
		dto *dtos.Battle,
	) (*models.Battle, error)

	Update(
		ctx context.Context,
		id string,
		uid string,
		dto *dtos.Battle,
	) (*models.Battle, error)

	Delete(
		ctx context.Context,
		id string,
		uid string,
	) error
}

type BattleService struct {
	battleRepository repositories.BattleRepositoryInterface
}

func NewBattleService(
	battleRepository repositories.BattleRepositoryInterface,
) BattleServiceInterface {
	return &BattleService{
		battleRepository,
	}
}

func createBattleModel(dao *daos.Battle) *models.Battle {
	model := models.Battle{}

	model.ID = dao.ID
	model.CreatedAt = dao.CreatedAt
	model.UpdatedAt = dao.UpdatedAt
	model.GameId = dao.GameId
	model.UserId = dao.UserId
	model.GoFirst = dao.GoFirst
	model.VictoryFlg = dao.VictoryFlg
	model.YourPrizeCards = dao.YourPrizeCards
	model.OpponentsPrizeCards = dao.OpponentsPrizeCards
	model.Memo = dao.Memo

	return &model
}

func (s *BattleService) FindById(
	ctx context.Context,
	id string,
) (*models.Battle, error) {
	dao, err := s.battleRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	model := createBattleModel(dao)

	return model, nil

}

func (s *BattleService) Create(
	ctx context.Context,
	uid string,
	dto *dtos.Battle,
) (*models.Battle, error) {
	id, err := generateId()
	if err != nil {
		return nil, err
	}

	dao := daos.Battle{
		ID:                  id,
		GameId:              dto.GameId,
		UserId:              uid,
		GoFirst:             dto.GoFirst,
		VictoryFlg:          dto.VictoryFlg,
		YourPrizeCards:      dto.YourPrizeCards,
		OpponentsPrizeCards: dto.OpponentsPrizeCards,
		Memo:                dto.Memo,
	}

	if err := s.battleRepository.Save(ctx, &dao); err != nil {
		return nil, err
	}

	model := createBattleModel(&dao)

	return model, nil
}

func (s *BattleService) Update(
	ctx context.Context,
	id string,
	uid string,
	dto *dtos.Battle,
) (*models.Battle, error) {
	// 指定されたIdのBattleが存在するか確認
	dao, err := s.battleRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// 指定されたuidと取得したdaoのUserIdが一致しているか確認
	if dao.UserId != uid {
		return nil, errors.New("no authority")
	}

	dao.GameId = dto.GameId
	dao.GoFirst = dto.GoFirst
	dao.VictoryFlg = dto.VictoryFlg
	dao.YourPrizeCards = dto.YourPrizeCards
	dao.OpponentsPrizeCards = dto.OpponentsPrizeCards
	dao.Memo = dto.Memo

	if err := s.battleRepository.Save(ctx, dao); err != nil {
		return nil, err
	}

	model := createBattleModel(dao)

	return model, nil
}

func (s *BattleService) Delete(
	ctx context.Context,
	id string,
	uid string,
) error {
	// 指定されたIdのBattleが存在するか確認
	dao, err := s.battleRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	// 指定されたUIDと取得したdaoのUserIdが一致しているか確認
	if dao.UserId != uid {
		return errors.New("no authority")
	}

	if err := s.battleRepository.Delete(ctx, id, uid); err != nil {
		return err
	}

	return nil
}
