package services

import (
	"context"

	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
	"github.com/vsrecorder/vsr-apiserver/pkg/services/models"
)

type UserServiceInterface interface {
	FindById(
		ctx context.Context,
		id string,
	) (*models.User, error)

	FindRecordsById(
		ctx context.Context,
		id string,
	) ([]*models.Record, error)

	FindGamesById(
		ctx context.Context,
		id string,
	) ([]*models.Game, error)

	FindDecksByIdWithUID(
		ctx context.Context,
		id string,
		uid string,
	) ([]*models.Deck, error)
}

type UserService struct {
	userRepository   repositories.UserRepositoryInterface
	recordRepository repositories.RecordRepositoryInterface
	gameRepository   repositories.GameRepositoryInterface
	deckRepository   repositories.DeckRepositoryInterface
}

func NewUserService(
	userRepository repositories.UserRepositoryInterface,
	recordRepository repositories.RecordRepositoryInterface,
	gameRepository repositories.GameRepositoryInterface,
	deckRepository repositories.DeckRepositoryInterface,
) UserServiceInterface {
	return &UserService{
		userRepository,
		recordRepository,
		gameRepository,
		deckRepository,
	}
}

func (s *UserService) FindById(
	ctx context.Context,
	id string,
) (*models.User, error) {
	dao, err := s.userRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	model := &models.User{}
	model.UID = dao.UID
	model.DisplayName = dao.DisplayName
	model.PhotoURL = dao.PhotoURL

	return model, nil
}

func (s *UserService) FindRecordsById(
	ctx context.Context,
	id string,
) ([]*models.Record, error) {
	daos, err := s.recordRepository.FindAllByUID(ctx, id)
	if err != nil {
		return nil, err
	}

	records := []*models.Record{}
	for _, dao := range daos {
		records = append(records, createRecordModel(dao))
	}

	return records, nil
}

func (s *UserService) FindGamesById(
	ctx context.Context,
	id string,
) ([]*models.Game, error) {
	daos, err := s.gameRepository.FindAllByUID(ctx, id)
	if err != nil {
		return nil, err
	}

	games := []*models.Game{}
	for _, dao := range daos {
		games = append(games, createGameModel(dao))
	}

	return games, nil
}

func (s *UserService) FindDecksByIdWithUID(
	ctx context.Context,
	id string,
	uid string,
) ([]*models.Deck, error) {
	daos, err := s.deckRepository.FindAllByUID(ctx, id)
	if err != nil {
		return nil, err
	}

	decks := []*models.Deck{}
	for _, dao := range daos {
		deck := createDeckModel(dao)

		if deck.PrivateCodeFlg && deck.UserId != uid {
			deck.Code = "ZZZZZZ-YYYYYY-ZZZZZZ"
		}

		decks = append(decks, deck)
	}

	return decks, nil
}
