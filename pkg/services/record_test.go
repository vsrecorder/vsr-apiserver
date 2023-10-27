package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQLMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn:                      mockDB,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)

	return db, mock, err

}

func TestRecordService(t *testing.T) {
	db, _, err := initMySQLMock()
	require.NoError(t, err)

	s := NewRecordService(
		repositories.NewRecordRepository(db),
		repositories.NewGameRepository(db),
		repositories.NewOfficialEventRepository(db),
	)

	for scenario, fn := range map[string]func(
		t *testing.T, s RecordServiceInterface,
	){
		"Create": test_Create,
		"Find":   test_Find,
		"Update": test_Update,
		"Delete": test_Delete,
	} {
		fn := fn // ↑引数の順番通りに実行するための設定(https://github.com/golang/go/wiki/CommonMistakes)
		t.Run(scenario, func(t *testing.T) {
			fn(t, s)
		})
	}
}

func test_Create(t *testing.T, s RecordServiceInterface) {
	require.Equal(t, true, false)
}

func test_Find(t *testing.T, s RecordServiceInterface) {
	require.Equal(t, true, false)
}

func test_Update(t *testing.T, s RecordServiceInterface) {
	require.Equal(t, true, false)
}

func test_Delete(t *testing.T, s RecordServiceInterface) {
	require.Equal(t, true, false)
}
