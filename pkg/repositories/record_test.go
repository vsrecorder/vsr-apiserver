package repositories

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories/daos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	layout = "2006-01-02 15:04:05"
)

func setupMySQLMock() (*gorm.DB, sqlmock.Sqlmock, error) {
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
func TestRecordRepository(t *testing.T) {
	db, mock, err := setupMySQLMock()
	require.NoError(t, err)

	r := NewRecordRepository(db)

	for scenario, fn := range map[string]func(
		t *testing.T, r RecordRepositoryInterface, mock sqlmock.Sqlmock,
	){
		"Save": test_Save,
		//"Find": test_Find,
		//"Delete": test_Delete,
	} {
		fn := fn // ↑引数の順番通りに実行するための設定(https://github.com/golang/go/wiki/CommonMistakes)
		t.Run(scenario, func(t *testing.T) {
			fn(t, r, mock)
		})
	}

	/*
		t.Run("TearDown", func(t *testing.T) {
			test_TearDown(t, r)
		})
	*/
}

func test_Save(t *testing.T, r RecordRepositoryInterface, mock sqlmock.Sqlmock) {
	datetime, err := time.Parse(layout, "2023-10-21 09:53:30")
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{
		"id",
		"created_at",
		"update_at",
		"deleted_at",
		"official_event_id",
		"user_id",
		"deck_id",
	}).AddRow("01HD7Y3K8D6FDHMHTZ2GT41TN2", datetime, datetime, "NULL", "236790", "CeQ0Oa9g9uRThL11lj4l45VAg8p1", "")

	mock.ExpectBegin()

	mock.ExpectQuery(
		regexp.QuoteMeta(
			"INSERT INTO `records` (`id`,`created_at`,`updated_at`,`deleted_at`,`official_event_id`,`user_id`,`deck_id`) "+
				"VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "+
				"`updated_at`=?,`deleted_at`=VALUES(?),`official_event_id`=VALUES(?),`user_id`=VALUES(?),`deck_id`=VALUES(?)",
		)).
		WithArgs(
			"01HD7Y3K8D6FDHMHTZ2GT41TN2",
			"2023-10-21 09:53:30",
			"2023-10-21 09:53:30",
			"NULL",
			"236790",
			"CeQ0Oa9g9uRThL11lj4l45VAg8p1",
			"",
			"2023-10-21 09:53:30",
			"NULL",
			"236790",
			"CeQ0Oa9g9uRThL11lj4l45VAg8p1",
			"",
		).WillReturnRows(rows)
	mock.ExpectCommit()

	record := &daos.Record{
		ID:        "01HD7Y3K8D6FDHMHTZ2GT41TN2",
		CreatedAt: datetime,
		UpdatedAt: datetime,
		//DeletedAt:       ,
		OfficialEventId: 236790,
		UserId:          "CeQ0Oa9g9uRThL11lj4l45VAg8p1",
		DeckId:          "",
	}

	r.Save(context.Background(), record)

	{
		err := mock.ExpectationsWereMet()
		require.NoError(t, err)
		//require.Equal(t, true, false)
	}
}

func test_Find(t *testing.T, r RecordRepositoryInterface, mock sqlmock.Sqlmock) {
	datetime, err := time.Parse(layout, "2023-10-21 09:53:30")
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{
		"id",
		"created_at",
		"update_at",
		"deleted_at",
		"official_event_id",
		"user_id",
		"deck_id",
	}).AddRow("01HD7Y3K8D6FDHMHTZ2GT41TN2", datetime, datetime, nil, "236790", "CeQ0Oa9g9uRThL11lj4l45VAg8p1", "")
	//}).AddRow("", "", "", "", "", "", "")

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `records` WHERE `records`.`deleted_at` IS NULL LIMIT 20"),
	).WillReturnRows(rows)

	records, err := r.Find(context.Background(), 20, 0)

	fmt.Println(records)
	require.NoError(t, err)
	require.Equal(t, len(records), 1)
	require.Equal(t, records[0].ID, "01HD7Y3K8D6FDHMHTZ2GT41TN2")
	//require.Equal(t, true, false)
}

func test_Delete(t *testing.T, r RecordRepositoryInterface, mock sqlmock.Sqlmock) {
	require.Equal(t, true, false)
}

func test_TearDown(t *testing.T, r RecordRepositoryInterface) {
	require.Equal(t, true, false)
	//require.NoError(t, r.TearDown())
}
