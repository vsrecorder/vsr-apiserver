package services

import (
	"context"
	"errors"

	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/dtos"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"github.com/vsrecorder/vsr-apiserver/pkg/services/models"
)

type DeckServiceInterface interface {
	FindByIdWithUID(
		ctx context.Context,
		id string,
		uid string,
	) (*models.Deck, error)

	FindByUID(
		ctx context.Context,
		uid string,
		limit int,
		offset int,
	) ([]*models.Deck, error)

	FindRecordById(
		ctx context.Context,
		id string,
	) ([]*models.Record, error)

	Create(
		ctx context.Context,
		uid string,
		dto *dtos.Deck,
	) (*models.Deck, error)

	Update(
		ctx context.Context,
		id string,
		uid string,
		dto *dtos.Deck,
	) (*models.Deck, error)

	Delete(
		ctx context.Context,
		id string,
		uid string,
	) error
}

type DeckService struct {
	deckRepository   repositories.DeckRepositoryInterface
	recordRepository repositories.RecordRepositoryInterface
}

func NewDeckService(
	deckRepository repositories.DeckRepositoryInterface,
	recordRepository repositories.RecordRepositoryInterface,
) DeckServiceInterface {
	return &DeckService{
		deckRepository,
		recordRepository,
	}
}

func createDeckModel(dao *daos.Deck) *models.Deck {
	model := &models.Deck{}

	model.ID = dao.ID
	model.CreatedAt = dao.CreatedAt
	model.UpdatedAt = dao.UpdatedAt
	model.UserId = dao.UserId
	model.Name = dao.Name
	model.Code = dao.Code
	model.PrivateCodeFlg = dao.PrivateCodeFlg

	return model
}

func (s *DeckService) FindByIdWithUID(
	ctx context.Context,
	id string,
	uid string,
) (*models.Deck, error) {
	dao, err := s.deckRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	model := createDeckModel(dao)

	if model.PrivateCodeFlg && uid != model.UserId {
		model.Code = "ZZZZZZ-YYYYYY-ZZZZZZ"
	}

	return model, nil
}

func (s *DeckService) FindByUID(
	ctx context.Context,
	uid string,
	limit int,
	offset int,
) ([]*models.Deck, error) {
	daos, err := s.deckRepository.FindByUID(ctx, uid, limit, offset)

	if err != nil {
		return nil, err
	}

	decks := []*models.Deck{}
	for _, dao := range daos {
		deck := createDeckModel(dao)
		decks = append(decks, deck)
	}

	return decks, nil
}

func (s *DeckService) FindRecordById(
	ctx context.Context,
	id string,
) ([]*models.Record, error) {
	// 指定されたIdのDeckが存在するか確認
	if _, err := s.deckRepository.FindById(ctx, id); err != nil {
		return nil, err
	}

	daos, err := s.recordRepository.FindByDeckId(ctx, id)
	if err != nil {
		return nil, err
	}

	records := []*models.Record{}
	for _, dao := range daos {
		records = append(records, createRecordModel(dao))
	}

	return records, nil
}

func (s *DeckService) Create(
	ctx context.Context,
	uid string,
	dto *dtos.Deck,
) (*models.Deck, error) {
	id, err := generateId()
	if err != nil {
		return nil, err
	}

	dao := daos.Deck{
		ID:             id,
		UserId:         uid,
		Name:           dto.Name,
		Code:           dto.Code,
		PrivateCodeFlg: dto.PrivateCodeFlg,
	}

	if err := s.deckRepository.Save(ctx, &dao); err != nil {
		return nil, err
	}

	model := createDeckModel(&dao)

	return model, nil
}

func (s *DeckService) Update(
	ctx context.Context,
	id string,
	uid string,
	dto *dtos.Deck,
) (*models.Deck, error) {
	// TODO: 指定されたdto.Codeがトレーナーズウェブサイト上に存在するか確認したい(有効なデッキコードか否か)
	// https://www.pokemon-card.com/deck/result.html/deckID/{dto.Code}

	// 指定されたidのDeckが存在するか確認
	dao, err := s.deckRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// 指定されたuidと取得したdaoのUserIdが一致しているか確認
	if dao.UserId != uid {
		return nil, errors.New("no authority")
	}

	dao.Name = dto.Name
	dao.Code = dto.Code
	dao.PrivateCodeFlg = dto.PrivateCodeFlg

	if err := s.deckRepository.Save(ctx, dao); err != nil {
		return nil, err
	}

	model := createDeckModel(dao)

	return model, nil
}

func (s *DeckService) Delete(
	ctx context.Context,
	id string,
	uid string,
) error {
	// 指定されたIdのDeckが存在するか確認
	dao, err := s.deckRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	// 指定されたUIDと取得したdaoのUserIdが一致しているか確認
	if dao.UserId != uid {
		return errors.New("no authority")
	}

	return s.deckRepository.Delete(ctx, id, uid)
}
